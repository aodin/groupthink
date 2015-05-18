package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/aodin/volta/config"
	"github.com/codegangsta/cli"
	"github.com/codegangsta/envy/lib"

	"github.com/aodin/groupthink/server"
)

func main() {
	// Bootstrap the environment
	envy.Bootstrap()

	app := cli.NewApp()
	app.Name = "groupthink"
	app.Usage = "Start the groupthink server"
	app.Action = startServer
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log, l",
			Value: "",
			Usage: "Sets the log output file path",
		},
		cli.StringFlag{
			Name:  "config, c",
			Value: "./settings.json",
			Usage: "Sets the configuration file",
		},
	}
	app.Run(os.Args)
}

func startServer(c *cli.Context) {
	logF := c.String("log")
	file := c.String("config")
	// Set the log output - if no path given, use stdout
	// TODO log rotation?
	if logF != "" {
		if err := os.MkdirAll(filepath.Dir(logF), 0776); err != nil {
			log.Panic(err)
		}
		l, err := os.OpenFile(logF, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Panic(err)
		}
		defer l.Close()
		log.SetOutput(l)
	}
	conf := connect(file)
	log.Panic(server.New(conf).ListenAndServe())
}

func connect(file string) config.Config {
	// Parse the given configuration file
	conf, err := config.ParseFile(file)
	if err != nil {
		log.Panicf("groupthink: could not parse configuration: %s", err)
	}
	return conf
}
