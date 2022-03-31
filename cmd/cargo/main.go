package main

import (
	"flag"
	"fmt"
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
	cmdFlags := cmdlineFlags{}
	flag.StringVar(&cmdFlags.configFile, "config", "", "path to config file")
	flag.BoolVar(&cmdFlags.version, "version", false, "display the version and exit")
	flag.Parse()

	if cmdFlags.version {
		fmt.Printf("cargo %s\n", version.GetVersionString())
		os.Exit(0)
	}

	cfg, err := config.Load(cmdFlags.configFile)
	if err != nil {
		fmt.Printf("Failed to load config: %s\n", err)
		os.Exit(1)
	}

	logging.Setup(&cfg.Logging)

	logger := logging.GetLogger()

	logger.Info("cargo started")

}
