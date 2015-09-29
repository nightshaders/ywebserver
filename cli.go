package main

import (
	cmd "github.com/codegangsta/cli"
	"github.com/nightshaders/ywebserver/config"
	"github.com/nightshaders/ywebserver/server"
)

func triggerServer(start StartServer) func(*cmd.Context) {
	return func(c *cmd.Context) {
		conf := &config.WebConf{
			Port:                 c.Int("port"),
			SiteRoot:             c.String("root"),
			DefaultFile:          c.String("default-html"),
			ServeEmbedddedAssets: c.Bool("serve-embedded-assets"),
		}
		newServer := server.NewServer(conf)
		newServer.Conf.ApplyDefaults()
		start(newServer)
	}
}

type StartServer func(*server.Server)

func NewCli(s StartServer) *cmd.App {
	app := cmd.NewApp()
	app.Name = "sites"
	app.Version = "0.0.1"
	app.Usage = "Sites is an executable for both server and client."
	app.Commands = []cmd.Command{
		{
			Name:    "web",
			Aliases: []string{"s"},
			Usage:   "Start the web server.",
			Action:  triggerServer(s),
			Flags: []cmd.Flag{
				cmd.BoolFlag{
					Name:  "serve-embedded-assets",
					Usage: "Serves embedded assets. (ignores root flag)",
				},
				cmd.IntFlag{
					Name:  "port",
					Value: 1111,
					Usage: "Port for web-server to bind.",
				},
				cmd.StringFlag{
					Name:  "root",
					Value: "src/web_pair/.www/dest/_site",
					Usage: "Path to static site assets.",
				},
				cmd.StringFlag{
					Name:  "default-html",
					Value: "index.html",
					Usage: "File to serve for empty root '/' path.",
				},
			},
		},
	}
	return app
}
