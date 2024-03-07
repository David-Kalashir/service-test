package supplierconfig

import (
	"context"

	"github.com/google/uuid"
)

type Core struct {
}

func (c Core) GetConfig(ctx context.Context, id uuid.UUID) (Config, error) {
	return Config{}, nil
}

func (c Core) GetTurkishConfig(ctx context.Context, id uuid.UUID) (TurkishConfig, error) {
	return TurkishConfig{}, nil
}
