package main

import (
	"os"
	"runtime"

	"event-log-collect/cmd"

	"github.com/urfave/cli"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "event-log-collect"
	app.Author = "hksanduo"
	app.Email = "hksanduo@gmail.com"
	app.Version = "V1.0"
	app.Usage = "collect operation system event log"
	app.Commands = []cli.Command{cmd.Scan}
	app.Flags = append(app.Flags, cmd.Scan.Flags...)

	app.Run(os.Args)
}
