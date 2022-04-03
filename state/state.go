package state

import (
	"errors"
	"fmt"
	"github.com/cloudstruct/cargo/config"
	"github.com/cloudstruct/cargo/logging"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"sync"
)

type State struct {
	sync.Mutex
	gormDb   *gorm.DB
	Metadata *Metadata
}

// Singleton state instance
var globalState *State

func Load(cfg *config.Config) (*State, error) {
	var err error
	var gormDb *gorm.DB
	logger := logging.GetLogger()
	switch cfg.State.DatabaseDriver {
	case "sqlite":
		// Create parent directory if it doesn't exist
		parentDir := filepath.Dir(cfg.State.DatabaseDsn)
		if _, err := os.Stat(parentDir); errors.Is(err, os.ErrNotExist) {
			logger.Debugf("creating parent directory %s", parentDir)
			if err := os.MkdirAll(parentDir, os.ModePerm); err != nil {
				return nil, fmt.Errorf("failed creating parent directories for %s: %s", parentDir, err)
			}
		}
		gormDb, err = gorm.Open(sqlite.Open(cfg.State.DatabaseDsn), &gorm.Config{})
		if err == nil {
			logger.Infof("opened SQLite connection for state: %s", cfg.State.DatabaseDsn)
		} else {
			return nil, fmt.Errorf("failed creating SQLite connection for state: %s: %s", cfg.State.DatabaseDsn, err)
		}
	// TODO: add additional drivers
	default:
		return nil, fmt.Errorf("unsupported database driver: %s\n", cfg.State.DatabaseDriver)
	}
	globalState = &State{
		gormDb: gormDb,
	}
	if err := globalState.Load(); err != nil {
		return nil, fmt.Errorf("failed loading state data: %s", err)
	}
	return globalState, nil
}

func GetState() *State {
	return globalState
}

func (s *State) Load() error {
	s.Lock()
	defer s.Unlock()
	var err error
	s.Metadata, err = NewMetadata(s.gormDb)
	if err != nil {
		return err
	}
	// TODO: add additional state tables to load
	return nil
}

func (s *State) Save() error {
	s.Lock()
	defer s.Unlock()
	if err := s.Metadata.Save(); err != nil {
		return err
	}
	// TODO: add additional state tables to save
	return nil
}
