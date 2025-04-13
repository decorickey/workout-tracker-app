package aggregate

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Exercise struct {
	name string
	memo string
}

type Record struct {
	exerciseID uuid.UUID
	memo       string
}

type RecordSet struct {
	recordID uuid.UUID
	weight   decimal.Decimal
	reps     uint
}
