package main

import (
	"flag"
	"fmt"
	"github.com/cloudstruct/cargo/api"
	"github.com/cloudstruct/cargo/config"
	"github.com/cloudstruct/cargo/logging"
	"github.com/cloudstruct/cargo/version"
	"os"
)

type cmdlineFlags struct {
	configFile string
	version    bool
}

func main() {
	// Parse commandline
	cmdFlags := cmdlineFlags{}
	flag.StringVar(&cmdFlags.configFile, "config", "", "path to config file")
	flag.BoolVar(&cmdFlags.version, "version", false, "display the version and exit")
	flag.Parse()

	if cmdFlags.version {
		fmt.Printf("cargo %s\n", version.GetVersionString())
		os.Exit(0)
	}

	// Load config
	cfg, err := config.Load(cmdFlags.configFile)
	if err != nil {
		fmt.Printf("Failed to load config: %s\n", err)
		os.Exit(1)
	}

	// Configure logging
	logging.Setup(&cfg.Logging)
	logger := logging.GetLogger()

	// Start API listener
	logger.Infof("starting management API listener on %s:%d", cfg.Api.Address, cfg.Api.Port)
	if err := api.Start(cfg); err != nil {
		logger.Fatalf("failed to start API: %s", err)
		os.Exit(1)
	}
}
