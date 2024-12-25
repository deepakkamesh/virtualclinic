package main

import (
	"log"
	"net/rpc"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	log.SetFlags(log.LstdFlags)
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "192.168.68.119:2311",
			Usage: "Set host:Port",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "set_volume",
			Aliases: []string{"sv"},
			Usage:   "Set Volume",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:  "l",
					Usage: " Set Volume Level",
				},
			},
			Action: func(c *cli.Context) error {
				client, err := rpc.DialHTTP("tcp", c.String("host"))
				if err != nil {
					return err
				}
				var vol int
				err = client.Call("Server.SetVolume", c.Int("l"), &vol)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "get_volume",
			Aliases: []string{"v"},
			Usage:   "Get Volume Level",
			Action: func(c *cli.Context) error {
				client, err := rpc.DialHTTP("tcp", c.String("host"))
				if err != nil {
					return err
				}
				var vol int
				err = client.Call("Server.Volume", struct{}{}, &vol)
				if err != nil {
					return err
				}
				log.Printf("Volume: %d", vol)
				return nil
			},
		},
		{
			Name:    "check_printer",
			Aliases: []string{"p"},
			Usage:   "Check if printer is connected",
			Action: func(c *cli.Context) error {
				client, err := rpc.DialHTTP("tcp", c.String("host"))
				if err != nil {
					return err
				}
				err = client.Call("Server.CheckPrinter", struct{}{}, &struct{}{})
				if err != nil {
					return err
				}
				log.Printf("Printer Connected!")
				return nil
			},
		},
		{
			Name:    "check_otocam",
			Aliases: []string{"o"},
			Usage:   "Check if otocam is connected",
			Action: func(c *cli.Context) error {
				client, err := rpc.DialHTTP("tcp", c.String("host"))
				if err != nil {
					return err
				}
				err = client.Call("Server.CheckOtocam", struct{}{}, &struct{}{})
				if err != nil {
					return err
				}
				log.Printf("Otocam Connected!")
				return nil
			},
		},
		{
			Name:    "tunnel",
			Aliases: []string{"t"},
			Usage:   "Start/Stop Tunnel",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "id",
					Usage: " Tunnel id e.g -id ssh",
				},
				&cli.BoolFlag{
					Name:  "on",
					Usage: " Tunnel id e.g -id ssh --on",
				},
			},
			Action: func(c *cli.Context) error {
				client, err := rpc.DialHTTP("tcp", c.String("host"))
				if err != nil {
					return err
				}
				if c.Bool("on") {
					err = client.Call("Server.StartTunnel", c.String("id"), &struct{}{})
				} else {
					err = client.Call("Server.StopTunnel", c.String("id"), &struct{}{})
				}
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "gvc",
			Aliases: []string{"g"},
			Usage:   "GVC Controls",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "on",
					Usage: " Turn on/off GVC",
				},
			},
			Action: func(c *cli.Context) error {
				client, err := rpc.DialHTTP("tcp", c.String("host"))
				if err != nil {
					return err
				}
				if c.Bool("on") {
					err = client.Call("Server.StartRemoteGVC", struct{}{}, &struct{}{})
				} else {
					err = client.Call("Server.StopRemoteGVC", struct{}{}, &struct{}{})
				}
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "gvc_cam",
			Aliases: []string{"gc"},
			Usage:   "Switch GVC Camera",
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:  "c",
					Usage: " Camera Number e.g -c 1 or -c 2",
				},
			},
			Action: func(c *cli.Context) error {
				client, err := rpc.DialHTTP("tcp", c.String("host"))
				if err != nil {
					return err
				}
				err = client.Call("Server.SwitchGVCCamera", c.Int("c"), &struct{}{})
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "mute_gvc",
			Aliases: []string{"m"},
			Usage:   "Mute GVC",
			Action: func(c *cli.Context) error {
				client, err := rpc.DialHTTP("tcp", c.String("host"))
				if err != nil {
					return err
				}
				err = client.Call("Server.ToggleMuteGVC", struct{}{}, &struct{}{})
				if err != nil {
					return err
				}
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
