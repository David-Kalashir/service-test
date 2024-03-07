package turkish

import (
	"github.com/ardanlabs/service/business/core/crud/flight"
	"github.com/ardanlabs/service/business/core/crud/turkishairlines"
)

func toCoreFlights(turkishFlts []turkishairlines.Flight) ([]flight.Flight, error) {

	return []flight.Flight{}, nil
}
