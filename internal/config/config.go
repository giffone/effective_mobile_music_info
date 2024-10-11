package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Debug        bool
	AppPort      uint64
	DatabaseAddr string
}

func New() *Config {
	debug := false
	debugStr := strings.ToLower(os.Getenv("DEBUG"))
	if debugStr == "true" {
		debug = true
	}

	portsEnv := must("APP_PORTS")
	ports := strings.Split(portsEnv, ":")
	if len(ports) != 2 {
		log.Fatalf("[config]: app_ports format is wrong - %s: must be like \"8080:8080\" in env file", portsEnv)
	}

	port, err := strconv.ParseUint(ports[1], 10, 64)
	if err != nil || port == 0 {
		log.Fatalf("[config]: parse port: %d: %s", port, err)
	}

	c := Config{
		Debug:        debug,
		AppPort:      port,
		DatabaseAddr: must("DATABASE_URL"),
	}

	if debug {
		c.Print()
	}

	return &c
}

func must(envName string) string {
	val := os.Getenv(envName)
	if val == "" {
		s := fmt.Sprintf("you can set by example: -%s '...'", envName)
		f := flag.String(envName, "", s)
		flag.Parse()

		if f == nil || *f == "" {
			log.Fatalf("[config]: %s is undefined, use -help for examples", envName)
		}
		val = *f
	}

	return val
}

func (c *Config) Print() {
	cfgStr := fmt.Sprintf(`
App Config: {
	Debug:			%t
	AppPort:		%d
	DatabaseAddr:	%s
}`, c.Debug, c.AppPort, c.DatabaseAddr)

	log.Println(cfgStr)
}
