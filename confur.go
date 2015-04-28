package main

import (
	"flag"
	"fmt"
	"github.com/confur-me/confur-api/app"
	"github.com/confur-me/confur-api/lib/config"
	"os"
)

var configFile *string

func init() {
	const (
		help_description   = "print this help"
		config_description = "path to config file"
		usage              = `
Usage:
  confur -config=/path/to/config.yml start
  confur -config=/path/to/config.yml migrate
`
	)
	flag.Bool("h", false, help_description)
	configFile = flag.String("c", "confur.yml", config_description)

	flag.Usage = func() {
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Print(usage)
		os.Exit(0)
	}
}

func main() {
	args := parseFlags()
	if err := config.Read(*configFile); err != nil {
		os.Exit(0)
	}

	switch args[0] {
	case "start":
		app := app.Application{}
		app.Run()
	case "migrate":
		//migrations.DbMigrate()
		panic("Not implemented yet")
	default:
		flag.Usage()
	}
}

func parseFlags() []string {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
	}
	return args
}
