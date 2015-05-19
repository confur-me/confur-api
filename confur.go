package main

import (
	"flag"
	"fmt"
	"github.com/confur-me/confur-api/api/migrator"
	"github.com/confur-me/confur-api/lib/config"
	"github.com/confur-me/confur-api/server"
	//"github.com/confur-me/confur-api/test/fixtures"
	"os"
)

var configFile *string

func init() {
	const (
		help_description   = "print this help"
		config_description = "path to config file"
		usage              = `
Usage:
  confur -config=/path/to/config.yml start : start server
  confur -config=/path/to/config.yml db:drop : Drop database
  confur -config=/path/to/config.yml db:migrate : Migrate database
  confur -config=/path/to/config.yml fixtures:seed : Seed test data
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
		app := server.Application{}
		app.Run()
	case "db:migrate":
		migrator.DbMigrate()
	case "db:drop":
		migrator.DbDrop()
	case "fixtures:seed":
		//fixtures.Seed()
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
