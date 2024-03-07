package turkish

import (
	"context"

	"github.com/ardanlabs/service/business/core/crud/flight"
	"github.com/ardanlabs/service/business/core/crud/supplierconfig"
	"github.com/ardanlabs/service/business/core/crud/turkishairlines"
)

// Ask: is it correct to use suppplier config?!
type Supplier struct {
	Turkishairlines *turkishairlines.Core
	suppliercfg     *supplierconfig.Core
}

func NewSupplier() *Supplier {
	return &Supplier{}
}

// Create inserts a new user into the database.
func (s *Supplier) SearchFlight(ctx context.Context, req flight.SupplierSearchFlight) ([]flight.Flight, error) {
	cfg, err := s.suppliercfg.GetTurkishConfig(ctx, req.ConfigID)
	if err != nil {
		return nil, err
	}

	flts, err := s.Turkishairlines.AirLowFairSearch(ctx, turkishairlines.AirLowFairSearchRequest{
		Username:    cfg.Username,
		Password:    cfg.Password,
		Origin:      req.Origin,
		Destination: req.Destination,
	})
	if err != nil {
		return nil, err
	}

	return toCoreFlights(flts)
}
