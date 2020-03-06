package main

import (
	"log"
	"os"

	"github.com/fabiodcorreia/ozone/server"
	cli "github.com/urfave/cli/v2"
)

var version = "development"

func main() {
	var hostname string
	var port string
	var defaultResource string
	var rootDir string

	app := &cli.App{
		Version:   version,
		Name:      "Ozone Server",
		Usage:     "Ozone is a HTTP Server for in-memory resource serving",
		UsageText: "ozone [global options]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "hostname",
				Aliases:     []string{"n"},
				Value:       "127.0.0.1",
				Usage:       "Set the hostname",
				Destination: &hostname,
			},
			&cli.StringFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Value:       "8080",
				Usage:       "Set the listening port",
				Destination: &port,
			},
			&cli.StringFlag{
				Name:        "default",
				Aliases:     []string{"d"},
				Value:       "/index.html",
				Usage:       "Set the default resource to use on 404",
				Destination: &defaultResource,
			},
			&cli.StringFlag{
				Name:        "root",
				Aliases:     []string{"r"},
				Value:       ".",
				Usage:       "Set root directory to serve",
				Destination: &rootDir,
			},
		},
		Action: func(c *cli.Context) error {
			server.StartServer(server.Configuration{
				Hostname:        hostname,
				Port:            port,
				DefaultResource: defaultResource,
				RootPath:        rootDir,
			})
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
