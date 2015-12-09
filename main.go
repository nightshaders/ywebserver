package main

import (
	"github.com/nightshaders/ywebserver/server"
	"os"
)

func main() {
	server.NewCli(ExampleServe).Run(os.Args)
}
