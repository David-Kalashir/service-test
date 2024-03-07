package flight

import "github.com/google/uuid"

type Flight struct {
}

type Ticket struct {
}

type Search struct {
	UserID      string
	Origin      string
	Destination string
}

type SupplierSearchFlight struct {
	ConfigID    uuid.UUID
	Origin      string
	Destination string
}
