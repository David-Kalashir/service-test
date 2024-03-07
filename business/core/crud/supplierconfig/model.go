package supplierconfig

import "github.com/google/uuid"

type ConfigType string

const (
	TURKISH_CONFIG_TYPE ConfigType = "TURKISH"
)

type Config struct {
	ID       uuid.UUID
	ConfigID uuid.UUID  // represent supplier config id depends on the type
	Type     ConfigType // ex: TURKISH
}

type TurkishConfig struct {
	ID       uuid.UUID
	Username string
	Password string
}
