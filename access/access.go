package access

import (
	"github.com/fuks-kit/doorman/fuks"
	"sync"
	"time"
)

type Config struct {
	UpdateInterval time.Duration
	FallbackPath   string
	RecoveryPath   string
	SheetId        string
}

type accessList = map[uint32]fuks.AuthorisedUser

type Validator struct {
	mu             sync.RWMutex
	SheetId        string     `json:"-"`
	FallbackAccess accessList `json:"-"`
	FuksAccess     accessList `json:"fuks-access"`
	SheetAccess    accessList `json:"sheet-access"`
}

func NewValidator(config Config) (validator Validator) {
	validator = Validator{
		SheetId: config.SheetId,
	}

	if config.FallbackPath != "" {
		validator.readFallbackFrom(config.FallbackPath)
	}

	if config.RecoveryPath != "" {
		validator.readRecoveryFrom(config.RecoveryPath)
	}

	if config.UpdateInterval > 0 {
		validator.startUpdater(config.UpdateInterval, config.RecoveryPath)
	}

	return
}
