package turkishairlines

import "context"

type Core struct {
}

func (c *Core) AirLowFairSearch(ctx context.Context, req AirLowFairSearchRequest) ([]Flight, error) {
	return []Flight{}, nil
}
