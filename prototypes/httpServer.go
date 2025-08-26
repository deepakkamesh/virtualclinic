package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Server struct {
	hostPort string
}

// response Struct to return JSON.
type response struct {
	Msg        string
	IsMsgError bool
	Data       interface{}
}

type pingResponse struct {
	IsGVCReady         bool `json:"isGVCReady"`
	IsPrinterConnected bool `json:"isPrinterConnected"`
	IsOtocamConnected  bool `json:"isOtocamConnected"`
	Volume             int  `json:"volume"`
}

// NewServer returns an initialized server.
func NewServer(hostPort string) *Server {
	return &Server{
		hostPort: hostPort,
	}
}

// Start starts the http server.
func (s *Server) Start() error {

	// Http routers.
	http.HandleFunc("/api/ping", s.ping)
	http.HandleFunc("/api/gvcready", s.gvcready)
	http.HandleFunc("/api/startremotegvc", s.startRemoteGVC)
	http.HandleFunc("/api/stopremotegvc", s.stopRemoteGVC)
	http.HandleFunc("/api/switchgvccamera", s.SwitchGVCCamera)
	http.HandleFunc("/api/togglegvcmute", s.toggleGVCMute)
	http.HandleFunc("/api/checkprinter", s.checkPrinter)
	http.HandleFunc("/api/checkotocam", s.checkOtoCam)
	http.HandleFunc("/api/setvolume", s.setVolume)
	http.HandleFunc("/api/volume", s.volume)
	http.HandleFunc("/api/printscript", s.printScript)

	// TODO: Setup SSL.
	return http.ListenAndServe(s.hostPort, nil)
}

// ping handles the ping request to check if the server is running.
func (s *Server) ping(w http.ResponseWriter, r *http.Request) {
	writePingResponse(w, &pingResponse{
		IsGVCReady:         true,
		IsPrinterConnected: false,
		IsOtocamConnected:  false,
		Volume:             50,
	})
}

// gvcready handles the request to check if GVC is ready.
func (s *Server) gvcready(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, &response{
		Msg:        "GVC is ready",
		IsMsgError: false,
	})
}

func (s *Server) startRemoteGVC(w http.ResponseWriter, r *http.Request) {
	// Placeholder for starting remote GVC.
	writeResponse(w, &response{
		Msg:        "Requested Remote GVC Start",
		IsMsgError: false,
	})
}

func (s *Server) stopRemoteGVC(w http.ResponseWriter, r *http.Request) {
	// Placeholder for stopping remote GVC.
	writeResponse(w, &response{
		Msg:        "Stopped Remote GVC failed",
		IsMsgError: true,
	})
}
func (s *Server) SwitchGVCCamera(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	camera := strings.TrimSpace(r.Form.Get("camera"))

	// Start Video Stream.
	if camera != "1" && camera != "2" {
		writeResponse(w, &response{
			Msg:        "Only camera 1(gvc) or 2(otocam) required",
			IsMsgError: true,
		})
		return
	}
	// Placeholder for switching camera logic.

	writeResponse(w, &response{
		Msg:        fmt.Sprintf("Switching to Camera %s", camera),
		IsMsgError: false,
	})
}

func (s *Server) toggleGVCMute(w http.ResponseWriter, r *http.Request) {
	// Placeholder for toggling GVC mute.
	writeResponse(w, &response{
		Msg:        "GVC Mute Toggled",
		IsMsgError: false,
	})
}

var c = 0

func (s *Server) checkPrinter(w http.ResponseWriter, r *http.Request) {
	// Placeholder for checking printer status.
	fmt.Println("request received")
	c++

	if c%2 == 0 {
		writeResponse(w, &response{
			Msg:        "Printer is connected",
			IsMsgError: false,
		})
		return
	}
	writeResponse(w, &response{
		Msg:        "Printer is not connected",
		IsMsgError: true,
	})

}

func (s *Server) checkOtoCam(w http.ResponseWriter, r *http.Request) {
	// Placeholder for checking oto camera status.
	writeResponse(w, &response{
		Msg:        "OtoCam is not connected",
		IsMsgError: true,
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
	// Placeholder for setting volume logic.

	writeResponse(w, &response{
		Msg:        fmt.Sprintf("Remote Volume set to %d%%", vol),
		IsMsgError: false,
	})
}

func (s *Server) volume(w http.ResponseWriter, r *http.Request) {
	// Placeholder for getting volume level.
	writeResponse(w, &response{
		Msg:        "Volume level retrieved",
		IsMsgError: false,
		Data:       50, // Example volume level
	})
}

// print prints the pdf file that was generated.
func (s *Server) printScript(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	/*	if err != nil {
		writeResponse(w, &response{
			Err: fmt.Sprintf("Print Error: %v", err),
		})
		log.Printf("Print Error %v : %v", fname, err)
		return
	}*/

	// If Exec is successful, send back command output.
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

func main() {
	server := NewServer(":58080") // Change to your desired host and port

	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	log.Println("Server started on :8080")
	fmt.Print("Server started on :8080\n")
}
