package state

import (
	"github.com/cloudstruct/cargo/logging"
	"github.com/cloudstruct/cargo/version"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type metadataModel struct {
	ID        uuid.UUID `gorm:"type:uuid"`
	UpdatedAt time.Time
	Name      string `gorm:"unique"`
	Value     string
}

func (metadataModel) TableName() string {
	return "metadata"
}

func (m *metadataModel) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	m.ID = uuid
	return nil
}

type Metadata struct {
	gormDb *gorm.DB
}

func NewMetadata(gormDb *gorm.DB) (*Metadata, error) {
	m := &Metadata{
		gormDb: gormDb,
	}
	err := m.Load()
	return m, err
}

func (m *Metadata) Load() error {
	metadata := []metadataModel{}
	// Create table if it doesn't exist
	if !m.gormDb.Migrator().HasTable(&metadata) {
		logging.GetLogger().Debug("creating metadata table")
		if err := m.gormDb.AutoMigrate(&metadata); err != nil {
			return err
		}
	}
	result := m.gormDb.Find(&metadata)
	if len(metadata) > 0 {
		logging.GetLogger().Debugf("state was last written by Cargo %s at %s", metadata[0].Value, metadata[0].UpdatedAt.Format(time.RFC3339))
	}
	return result.Error
}

func (m *Metadata) Save() error {
	// Values to write into metadata table
	metadata := []metadataModel{
		{
			Name:  "CargoVersion",
			Value: version.GetVersionString(),
		},
	}
	// Automatically replace row if the 'name' column value matches
	onConflictClause := clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		UpdateAll: true,
	}
	// Create/update data in table
	result := m.gormDb.Clauses(onConflictClause).Create(&metadata)
	return result.Error
}
