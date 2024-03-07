package turkishconfig

import "context"

type Core struct {
}

func (c *Core) GetConfig(ctx context.Context) (Config, error) {
	return Config{}, nil
}
