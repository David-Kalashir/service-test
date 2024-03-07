package turkishconfig

import "github.com/google/uuid"

type Config struct {
	ID       uuid.UUID
	Username string
	Password string
}
