package access

import (
	"github.com/fuks-kit/doorman/workspace"
	"sync"
	"time"
)

type Config struct {
	UpdateInterval time.Duration
	FallbackPath   string
	RecoveryPath   string
}

type accessList = map[uint32]workspace.AuthorisedUser

type Validator struct {
	mu             sync.RWMutex
	FallbackAccess accessList `json:"-"`
	FuksAccess     accessList `json:"fuks-access"`
	SheetAccess    accessList `json:"sheet-access"`
}

func NewValidator(config Config) (validator *Validator) {
	// Always export the Validators pointer,
	// because otherwise updates are not persisted properly!
	validator = &Validator{}

	if config.RecoveryPath != "" {
		validator.readRecoveryFrom(config.RecoveryPath)
	}

	if config.FallbackPath != "" {
		validator.readFallbackFrom(config.FallbackPath)
	}

	if config.UpdateInterval > 0 {
		validator.startUpdater(config.UpdateInterval, config.RecoveryPath)
	}

	return
}
