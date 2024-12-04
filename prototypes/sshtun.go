package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rgzr/sshtun"
)

func main() {
	// We want to connect to port 8080 on our machine to acces port 80 on my.super.host.com
	sshTun := sshtun.NewRemote(54950, "drguruswamyclinic.hyperlinkhome.com", 8080)
	sshTun.SetUser("dkg")

	/*// We print each tunneled state to see the connections status
	sshTun.SetTunneledConnState(func(tun *sshtun.SSHTun, state *sshtun.TunneledConnState) {
		log.Printf("%+v", state)
	})
	*/
	var i int

	// We set a callback to know when the tunnel is ready
	sshTun.SetConnState(func(tun *sshtun.SSHTun, state sshtun.ConnState) {
		switch state {
		case sshtun.StateStarting:
			log.Printf("STATE is Starting")
			i = 3
		case sshtun.StateStarted:
			log.Printf("STATE is Started")
		case sshtun.StateStopped:
			log.Printf("STATE is Stopped")
		}
	})

	// We start the tunnel (and restart it every time it is stopped)
	go func() {
		for {
			if err := sshTun.Start(context.Background()); err != nil {
				log.Printf("SSH tunnel error: %v", err)
				time.Sleep(time.Second) // don't flood if there's a start error :)
			}
		}
	}()
	/*
	   // We stop the tunnel every 20 seconds (just to see what happens)

	   	for {
	   		time.Sleep(time.Second * time.Duration(20))
	   		log.Println("Lets stop the SSH tunnel...")
	   		sshTun.Stop()
	   	}
	*/
	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	sshTun.Stop()
}
