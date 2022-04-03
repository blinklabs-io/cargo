package main

import (
	"flag"
	"fmt"
	"github.com/cloudstruct/cargo/api"
	"github.com/cloudstruct/cargo/config"
	"github.com/cloudstruct/cargo/logging"
	"github.com/cloudstruct/cargo/state"
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

	// Handle -version
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
	// Sync logger on exit
	defer func() {
		if err := logger.Sync(); err != nil {
			// We don't actually care about the error here, but we have to do something
			// to appease the linter
			return
		}
	}()

	// Load state
	_, err = state.Load(cfg)
	if err != nil {
		logger.Fatalf("failed to load state: %s", err)
	}
	// TODO: remove me...this is mostly here to test state saving until we have
	// more actually using the state
	if err := state.GetState().Save(); err != nil {
		logger.Fatalf("failed to save state: %s", err)
	}

	// Start API listener
	logger.Infof("starting management API listener on %s:%d", cfg.Api.Address, cfg.Api.Port)
	if err := api.Start(cfg); err != nil {
		logger.Fatalf("failed to start API: %s", err)
	}
}
