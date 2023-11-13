// Copyright 2023 Blink Labs Software
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"github.com/blinklabs-io/cargo/api"
	"github.com/blinklabs-io/cargo/config"
	"github.com/blinklabs-io/cargo/logging"
	"github.com/blinklabs-io/cargo/state"
	"github.com/blinklabs-io/cargo/version"
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
	flag.BoolVar(
		&cmdFlags.version,
		"version",
		false,
		"display the version and exit",
	)
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
	logger.Infof(
		"starting management API listener on %s:%d",
		cfg.Api.Address,
		cfg.Api.Port,
	)
	if err := api.Start(cfg); err != nil {
		logger.Fatalf("failed to start API: %s", err)
	}
}
