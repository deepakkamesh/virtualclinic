// This file implements a server wrapper around sysagent to expose it as an RPC service.
package sysagent

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"

	vc "github.com/deepakkamesh/virtualclinic"
	"github.com/golang/glog"
)

type Server struct {
	checkURL    string    // http URL to validate if internet is connected.
	sysAgent    *SysAgent // New System Agent.
	rpcHostPort string    // RPC service information.
}

func NewServer(c vc.Config) *Server {
	tun, ok := c.Tunnels[vc.RPCTunID]
	if !ok {
		return nil
	}

	return &Server{
		checkURL:    c.CheckURL,
		sysAgent:    NewSysAgent(c),
		rpcHostPort: fmt.Sprintf("%v:%v", tun.LocalHost, tun.LocalPort),
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

func (s *Server) StartRPCTunnel() error {
	if err := s.sysAgent.StartTunnel(vc.RPCTunID); err != nil {
		return fmt.Errorf("failed to start rpc tunnel %v", err)
	}
	glog.Info("Started RPC Tunnel...[OK]")
	return nil
}

/************* RPC Service Functions *************/

func (s *Server) StartRemoteGVC(args struct{}, reply *struct{}) error {
	return s.sysAgent.StartRemoteGVC()
}

func (s *Server) StopRemoteGVC(args struct{}, reply *struct{}) error {
	return s.sysAgent.StopRemoteGVC()
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

func (s *Server) PrintScript(script string, reply *struct{}) error {
	return s.sysAgent.PrintScript(script)
}
