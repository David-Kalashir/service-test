package flight

import (
	"context"

	"github.com/ardanlabs/service/business/core/crud/delegate"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/google/uuid"
)

type FlightStorer interface {
	Create(ctx context.Context, flt Flight) error
	Update(ctx context.Context, flt Flight) error
	Delete(ctx context.Context, flt Flight) error
}

type TicketStorer interface {
	Create(ctx context.Context, tkt Ticket) error
	Update(ctx context.Context, tkt Ticket) error
	Delete(ctx context.Context, tkt Ticket) error
}

type Supplier interface {
	SearchFlight(ctx context.Context, req SupplierSearchFlight) ([]Flight, error)
}

// Core manages the set of APIs for flight access.
type Core struct {
	log          *logger.Logger
	flightStorer FlightStorer
	ticketStorer FlightStorer
	delegate     *delegate.Delegate
	suppliers    map[string]Supplier
}

// NewCore constructs a flight core API for use.
func NewCore(
	log *logger.Logger,
	delegate *delegate.Delegate,
	flightStorer FlightStorer,
	ticketStorer FlightStorer,
	suppliers map[string]Supplier,
) *Core {
	return &Core{
		log:          log,
		delegate:     delegate,
		flightStorer: flightStorer,
		ticketStorer: ticketStorer,
		suppliers:    suppliers,
	}
}

func (c *Core) Search(ctx context.Context, req Search) ([]Flight, error) {
	// get flight configs base on user information
	cfgID := uuid.New()

	flts := make([]Flight, 0)
	for _, supplier := range c.suppliers {
		supplierFlts, err := supplier.SearchFlight(ctx, SupplierSearchFlight{
			ConfigID:    cfgID,
			Origin:      req.Origin,
			Destination: req.Destination,
		})
		if err != nil {
			return nil, err
		}
		flts = append(flts, supplierFlts...)
	}

	return flts, nil
}
