package sysagent

import (
	"context"
	"fmt"
	"os"
	"time"

	vc "github.com/deepakkamesh/virtualclinic"
	"github.com/golang/glog"
	volume "github.com/itchyny/volume-go"
	"github.com/kenshaw/escpos"
	"github.com/rgzr/sshtun"
)

type SysAgent struct {
	browser         *Browser                  // Browser struct.
	optionsXPath    string                    // XPath for the GVC options.
	joinNowSelector string                    // CSS Selector for GVC JoinNow.
	gvcID           string                    // GVC ID.
	checkURL        string                    // http URL to validate if internet is connected.
	tunnels         map[string]vc.Tunnel      // SSH tunnels list.
	tunnelConn      map[string]*sshtun.SSHTun // List of SSH tunnels.
	OtoCamDevice    string                    // OtoCam Device
	PrinterDevice   string                    // Printer Device
}

// NewSysAgent starts a new SysAgent Service
// BrowserWindowStates are proto.(BrowserWindowStateNormal|BrowserWindowStateMinimized|BrowserWindowStateMaximized|BrowserWindowStateFullscreen)
func NewSysAgent(c vc.Config) *SysAgent {
	sys := &SysAgent{
		browser:         NewBrowser(c.ChromeBin, c.ChromeUserDir, c.BrowserWindowState),
		optionsXPath:    c.GVCOptionsXPath,
		joinNowSelector: c.GVCJoinNowSelector,
		gvcID:           c.GVCID,
		checkURL:        c.CheckURL,
		OtoCamDevice:    c.OtoCamDevice,
		PrinterDevice:   c.PrinterDevice,
		tunnels:         make(map[string]vc.Tunnel),
		tunnelConn:      make(map[string]*sshtun.SSHTun),
	}
	sys.tunnels = c.Tunnels
	return sys
}

// StartTunnel starts a new ssh tunnel for tunID from config file.
func (s *SysAgent) StartTunnel(tunID string) error {
	tun, ok := s.tunnels[tunID]
	if !ok {
		return fmt.Errorf("tunnel information not found for %v", tunID)
	}
	_, ok = s.tunnelConn[tunID]
	if ok {
		return fmt.Errorf("tunnel already started for %v", tunID)
	}

	tunnelConn := sshtun.NewRemote(tun.LocalPort, tun.RemoteHost, tun.RemotePort)
	tunnelConn.SetLocalHost(tun.LocalHost)
	tunnelConn.SetUser(tun.User)

	// We set a callback to know when the tunnel is ready
	tunnelConn.SetConnState(func(tun *sshtun.SSHTun, state sshtun.ConnState) {
		switch state {
		case sshtun.StateStarting:
			glog.Infof("Tunnel for %v is Starting", tunID)
		case sshtun.StateStarted:
			glog.Infof("Tunnel for %v  is Started", tunID)
		case sshtun.StateStopped:
			glog.Infof("Tunnel for %v  is Stopped", tunID)
		}
	})

	go func() {
		for {
			if err := tunnelConn.Start(context.Background()); err != nil {
				glog.Warningf("SSH tunnel error: %v. Will auto retry forever.", err)
				time.Sleep(10 * time.Second)
				continue
			}
			return // Tunnel was stopped.
		}
	}()
	s.tunnelConn[tunID] = tunnelConn
	return nil
}

// StartTunnel stops new ssh tunnel for tunID from config file.
func (s *SysAgent) StopTunnel(tunID string) error {
	tun, ok := s.tunnelConn[tunID]
	if !ok {
		return fmt.Errorf("tunnel not started for %v", tunID)
	}
	tun.Stop()
	delete(s.tunnelConn, tunID)
	return nil
}

func (s *SysAgent) StartRemoteGVC() error {
	return s.browser.GVC(s.gvcID, s.joinNowSelector)
}

func (s *SysAgent) StopRemoteGVC() error {
	return s.browser.Close()
}

func (s *SysAgent) ToggleMuteGVC() error {
	return s.browser.ToggleMuteGVC()
}

// SwitchGVCCamera changes the camera on GVC session between OtoCam/WebCam
func (s *SysAgent) SwitchGVCCamera(camera int) error {
	return s.browser.SwitchGVCCamera(camera, s.optionsXPath)
}

// SetVolume sets the audio output volume as a %.
func (s *SysAgent) SetVolume(vol int) error {
	return volume.SetVolume(vol)
}

// CheckPrinter checks the status of the printer.
func (s *SysAgent) CheckPrinter() error {
	_, err := os.Stat(s.PrinterDevice)
	if os.IsNotExist(err) {
		return fmt.Errorf("printer not connected")
	} else if err != nil {
		return fmt.Errorf("failed to open printer: %v", err)
	}
	return nil
}

// CheckOtocam checks the status of the otocam.
func (s *SysAgent) CheckOtocam() error {
	_, err := os.Stat(s.OtoCamDevice)
	if os.IsNotExist(err) {
		return fmt.Errorf("otocam not connected")
	} else if err != nil {
		return fmt.Errorf("failed to open otocam: %v", err)
	}
	return nil
}

// Volume returns the current volume level as a %.
func (s *SysAgent) Volume() (int, error) {
	vol, err := volume.GetVolume()
	if err != nil {
		return 0, fmt.Errorf("failed to get volume: %+v", err)
	}
	return vol, nil
}

// PrintScript prints the script to the printer device.
func (s *SysAgent) PrintScript(lines []FormattedLine) error {
	f, err := os.OpenFile(s.PrinterDevice, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	defer f.Close()
	p := escpos.New(f)
	p.Init()

	for _, line := range lines {
		p.SetFont(line.Font)
		p.SetFontSize(line.FontSize[0], line.FontSize[1])
		p.SetAlign(line.Align)
		p.SetEmphasize(line.Emphasize)
		p.SetSmooth(line.Smooth)
		p.SetUnderline(line.Underline)
		p.Write(line.Text)
		p.FormfeedN(line.FormFeed)
	}
	p.Cut()
	p.End()
	return nil
}
