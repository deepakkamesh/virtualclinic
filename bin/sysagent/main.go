// Binary for sysagent.
package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/deepakkamesh/virtualclinic"
	"github.com/deepakkamesh/virtualclinic/sysagent"
	"github.com/golang/glog"
)

func main() {
	var (
		cfgPath = flag.String("config_file", "./virtualclinic.conf.toml", "config file for virtual clinic")
	)
	flag.Parse()

	var conf virtualclinic.Config
	_, err := toml.DecodeFile(*cfgPath, &conf)
	if err != nil {
		glog.Errorf("Failed to open config file: %v", err)
		return
	}

	server := sysagent.NewServer(conf)
	if server == nil {
		glog.Errorf("Failed to start sysagent. Bad Config?")
		return
	}

	// Wait till we have internet connection.
	for i := 0; ; i++ {
		err = server.ValidateInternetConnection(10 * time.Second)
		if err == nil {
			break
		}
		time.Sleep(1 * time.Minute)
		glog.Errorf("Network error %v. Retrying attempt %v", err, i)
	}
	glog.Info("Validating Network...[OK]")

	// Start RPC Service.
	if err := server.StartRPCService(); err != nil {
		glog.Errorf("Failed to start RPC service %v", err)
		return
	}
	glog.Info("Started RPC Service...[OK]")

	// Start RPC Tunnel.
	if err := server.StartRPCTunnel(); err != nil {
		glog.Errorf("Failed to start RPC tunnel %v", err)
		return
	}

	glog.Info("Virtual Clinic Started")

	// Wait for signal to stop.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	server.StopServer()

}
