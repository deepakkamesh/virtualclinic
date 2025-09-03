// This file implements a server wrapper around sysagent to expose it as an RPC & HTTP service.
package sysagent

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
	"strings"
	"time"

	vc "github.com/deepakkamesh/virtualclinic"
	"github.com/golang/glog"
)

type Server struct {
	checkURL     string    // http URL to validate if internet is connected.
	sysAgent     *SysAgent // New System Agent.
	rpcHostPort  string    // RPC service information.
	httpHostPort string    // HTTP service information.
}

// response Struct to return JSON.
type response struct {
	Msg        string
	IsMsgError bool
	Data       interface{}
}

type pingResponse struct {
	IsSystemReady      bool `json:"isSystemReady"`
	IsGVCReady         bool `json:"isGVCReady"`
	IsPrinterConnected bool `json:"isPrinterConnected"`
	IsOtocamConnected  bool `json:"isOtocamConnected"`
	Volume             int  `json:"volume"`
}

func NewServer(c vc.Config) *Server {
	tun, ok := c.Tunnels[vc.RPCTunID]
	if !ok {
		return nil
	}

	httpTun, ok := c.Tunnels[vc.HTTPTunID]
	if !ok {
		return nil
	}

	return &Server{
		checkURL:     c.CheckURL,
		sysAgent:     NewSysAgent(c),
		rpcHostPort:  fmt.Sprintf("%v:%v", tun.LocalHost, tun.LocalPort),
		httpHostPort: fmt.Sprintf("%v:%v", httpTun.LocalHost, httpTun.LocalPort),
	}
}

func (s *Server) StopServer() error {
	return nil
}

func (s *Server) ValidateInternetConnection(timeout time.Duration) error {
	client := http.Client{
		Timeout: timeout,
	}
	_, err := client.Get(s.checkURL)
	return err
}

func (s *Server) StartRPCService() error {
	if err := rpc.Register(s); err != nil {
		return fmt.Errorf("failed to register rpc service %v", err)
	}
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", s.rpcHostPort)
	if e != nil {
		return fmt.Errorf("failed to start rpc service %v", e)
	}
	go http.Serve(l, nil)
	return nil
}

func (s *Server) StartHTTPService() error {

	// Http routers.
	http.HandleFunc("/api/ping", s.ping)
	http.HandleFunc("/api/gvcready", s.gvcready)
	http.HandleFunc("/api/startremotegvc", s.startRemoteGVC)
	http.HandleFunc("/api/stopremotegvc", s.stopRemoteGVC)
	http.HandleFunc("/api/switchgvccamera", s.switchGVCCamera)
	http.HandleFunc("/api/togglegvcmute", s.toggleGVCMute)
	http.HandleFunc("/api/checkprinter", s.checkPrinter)
	http.HandleFunc("/api/checkotocam", s.checkOtoCam)
	http.HandleFunc("/api/setvolume", s.setVolume)
	http.HandleFunc("/api/volume", s.volume)
	http.HandleFunc("/api/printscript", s.printScript)

	// TODO: Setup SSL.
	go http.ListenAndServe(s.httpHostPort, nil)
	return nil
}

func (s *Server) StartRPCTunnel() error {
	if err := s.sysAgent.StartTunnel(vc.RPCTunID); err != nil {
		return fmt.Errorf("failed to start rpc tunnel %v", err)
	}
	glog.Info("Started RPC Tunnel...[OK]")
	return nil
}

func (s *Server) StartHTTPTunnel() error {
	if err := s.sysAgent.StartTunnel(vc.HTTPTunID); err != nil {
		return fmt.Errorf("failed to start http tunnel %v", err)
	}
	glog.Info("Started HTTP Tunnel...[OK]")
	return nil
}

/************* RPC Wrapper Functions *************/

func (s *Server) StartRemoteGVC(args struct{}, reply *struct{}) error {
	return s.sysAgent.StartRemoteGVC()
}

func (s *Server) StopRemoteGVC(args struct{}, reply *struct{}) error {
	return s.sysAgent.StopRemoteGVC()
}

func (s *Server) IsGVCOpen(args struct{}, reply *bool) error {
	*reply = s.sysAgent.IsGVCOpen()
	return nil
}

func (s *Server) SwitchGVCCamera(camera int, reply *struct{}) error {
	return s.sysAgent.SwitchGVCCamera(camera)
}

func (s *Server) ToggleMuteGVC(args struct{}, reply *struct{}) error {
	return s.sysAgent.ToggleMuteGVC()
}

func (s *Server) StartTunnel(args string, reply *struct{}) error {
	return s.sysAgent.StartTunnel(args)
}

func (s *Server) StopTunnel(args string, reply *struct{}) error {
	return s.sysAgent.StopTunnel(args)
}

func (s *Server) CheckPrinter(args struct{}, reply *struct{}) error {
	return s.sysAgent.CheckPrinter()
}

func (s *Server) CheckOtocam(args struct{}, reply *struct{}) error {
	return s.sysAgent.CheckOtocam()
}

func (s *Server) Volume(args struct{}, reply *int) error {
	vol, err := s.sysAgent.Volume()
	if err != nil {
		return err
	}
	*reply = vol
	return nil
}

func (s *Server) SetVolume(vol int, reply *struct{}) error {
	return s.sysAgent.SetVolume(vol)
}

func (s *Server) PrintScript(script []*FormattedLine, reply *struct{}) error {
	return s.sysAgent.PrintScript(script)
}

/************* HTTP Wrapper Functions *************/
// ping handles the ping request to check if the server is running.
func (s *Server) ping(w http.ResponseWriter, r *http.Request) {

	writePingResponse(w, &pingResponse{
		IsSystemReady:      true,
		IsGVCReady:         s.sysAgent.IsGVCOpen(),
		IsPrinterConnected: s.sysAgent.CheckPrinter() == nil,
		IsOtocamConnected:  s.sysAgent.CheckOtocam() == nil,
		Volume: func() int {
			vol, err := s.sysAgent.Volume()

			if err != nil {
				return -1
			}
			return vol
		}(),
	})
}

// TODO: gvcready handles the request to check if GVC is ready.
func (s *Server) gvcready(w http.ResponseWriter, r *http.Request) {
	msg := "GVC is not open"
	isMsgErr := true

	if s.sysAgent.IsGVCOpen() {
		msg = "GVC is open"
		isMsgErr = false
	}

	writeResponse(w, &response{
		Msg:        msg,
		IsMsgError: isMsgErr,
	})
}

func (s *Server) startRemoteGVC(w http.ResponseWriter, r *http.Request) {
	if err := s.sysAgent.StartRemoteGVC(); err != nil {
		writeResponse(w, &response{
			Msg:        fmt.Sprintf("GVC Start failed: %v", err),
			IsMsgError: true,
		})
		return
	}
	writeResponse(w, &response{
		Msg:        "Remote Clinic Started",
		IsMsgError: false,
	})
}

func (s *Server) stopRemoteGVC(w http.ResponseWriter, r *http.Request) {

	if err := s.sysAgent.StopRemoteGVC(); err != nil {
		writeResponse(w, &response{
			Msg:        fmt.Sprintf("GVC Stop failed: %v", err),
			IsMsgError: true,
		})
		return
	}
	writeResponse(w, &response{
		Msg:        "Remote Clinic Stopped",
		IsMsgError: false,
	})
}

func (s *Server) switchGVCCamera(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	camera := strings.TrimSpace(r.Form.Get("camera"))

	if camera != "1" && camera != "2" {
		writeResponse(w, &response{
			Msg:        "Only camera 1(gvc) or 2(otocam) required",
			IsMsgError: true,
		})
		return
	}

	cam, _ := strconv.Atoi(camera)
	// TODO: This probably should be be done under a goroutine to avoid blocking the HTTP request.
	if err := s.sysAgent.SwitchGVCCamera(cam); err != nil {
		writeResponse(w, &response{
			Msg:        fmt.Sprintf("Switching to camera %s failed %v", camera, err),
			IsMsgError: true,
		})
		return
	}

	writeResponse(w, &response{
		Msg:        fmt.Sprintf("Switching to Camera %s", camera),
		IsMsgError: false,
	})
}

func (s *Server) toggleGVCMute(w http.ResponseWriter, r *http.Request) {

	if err := s.sysAgent.ToggleMuteGVC(); err != nil {
		writeResponse(w, &response{
			Msg:        fmt.Sprintf("Muting failed: %v", err),
			IsMsgError: true,
		})
		return
	}
	writeResponse(w, &response{
		Msg:        "GVC Mute Toggled",
		IsMsgError: false,
	})
}

func (s *Server) checkPrinter(w http.ResponseWriter, r *http.Request) {
	if err := s.sysAgent.CheckPrinter(); err != nil {
		writeResponse(w, &response{
			Msg:        fmt.Sprintf("Printer is not connected: %v", err),
			IsMsgError: true,
		})
		return
	}
	writeResponse(w, &response{
		Msg:        "Printer is connected",
		IsMsgError: false,
	})
}

func (s *Server) checkOtoCam(w http.ResponseWriter, r *http.Request) {
	if err := s.sysAgent.CheckOtocam(); err != nil {
		writeResponse(w, &response{
			Msg:        fmt.Sprintf("OtoCam is not connected: %v", err),
			IsMsgError: true,
		})
		return
	}
	writeResponse(w, &response{
		Msg:        "OtoCam is connected",
		IsMsgError: false,
	})
}

func (s *Server) setVolume(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	volume := strings.TrimSpace(r.Form.Get("volume"))

	vol, err := strconv.Atoi(volume) // Atoi returns an int and an error
	if err != nil {
		writeResponse(w, &response{
			Msg:        "Number between 0 to 100 required",
			IsMsgError: true,
		})
		return
	}
	if vol < 0 || vol > 100 {
		writeResponse(w, &response{
			Msg:        "Volume must be between 0 and 100",
			IsMsgError: true,
		})
		return
	}
	// Set the volume using sysagent.
	if err := s.sysAgent.SetVolume(vol); err != nil {
		writeResponse(w, &response{
			Msg:        fmt.Sprintf("Failed to set volume  %v", err),
			IsMsgError: true,
		})
		return
	}
	writeResponse(w, &response{
		Msg:        fmt.Sprintf("Remote Volume set to %d%%", vol),
		IsMsgError: false,
	})
}

func (s *Server) volume(w http.ResponseWriter, r *http.Request) {
	vol, err := s.sysAgent.Volume()
	if err != nil {
		writeResponse(w, &response{
			Msg:        fmt.Sprintf("Failed to get volume: %v", err),
			IsMsgError: true,
		})
		return
	}
	writeResponse(w, &response{
		Msg:        "Volume level retrieved",
		IsMsgError: false,
		Data:       vol,
	})
}

// print prints the pdf file that was generated.
func (s *Server) printScript(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseMultipartForm(10); err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	lines := strings.Split(r.FormValue("script"), "\n")

	script := func() []*FormattedLine {
		loc, _ := time.LoadLocation("Asia/Kolkata") // Always print date/time in India time.
		now := time.Now().In(loc)
		date := now.Format("2 Jan 2006  3:04 pm")

		lines := []*FormattedLine{}
		line := Line("Dr. R Guruswamy", FontSize([2]uint8{2, 2}), Smooth(1), Align("center"), Underline(6), Emphasize(3), FormFeed(2))
		lines = append(lines, line)
		line = Line("Ph:+91-9840084500 / Email:dr.guruswamy@gmail.com", FontSize([2]uint8{1, 1}), Smooth(1), Align("left"))
		lines = append(lines, line)
		line = Line("_______________________________________________", FontSize([2]uint8{1, 1}), Smooth(1), Align("center"), FormFeed(2))
		lines = append(lines, line)
		line = Line(date, FontSize([2]uint8{1, 1}), Smooth(1), Align("right"), FormFeed(2))
		lines = append(lines, line)

		return lines
	}()

	var line *FormattedLine

	for i := 0; i < len(lines); i++ {
		// First line contains name and should be bold.
		if i == 0 {
			line = Line(lines[0], FontSize([2]uint8{1, 1}), Emphasize(1), Smooth(1), Align("center"), FormFeed(2))
		} else {
			line = Line(lines[i], FontSize([2]uint8{1, 1}), Smooth(1), Align("left"), FormFeed(1))
		}
		script = append(script, line)
	}

	if err := s.sysAgent.PrintScript(script); err != nil {
		writeResponse(w, &response{
			Msg:        fmt.Sprintf("Prescription Print Error: %v", err),
			IsMsgError: true,
		})
		return
	}

	writeResponse(w, &response{
		Msg:        "Prescription Printed",
		IsMsgError: false,
	})

}

// writeResponse writes the response json object to w. If unable to marshal
// it writes a http 500.
func writeResponse(w http.ResponseWriter, resp *response) {
	w.Header().Set("Content-Type", "application/json")
	js, e := json.Marshal(resp)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	//	log.Printf("Writing json response %s", js)
	w.Write(js)
}

func writePingResponse(w http.ResponseWriter, resp *pingResponse) {
	w.Header().Set("Content-Type", "application/json")
	js, e := json.Marshal(resp)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
