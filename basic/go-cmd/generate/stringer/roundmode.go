package stringer

//go:generate stringer -type=RoundingMode

type RoundingMode byte

const (
	ToNearestEven RoundingMode = iota
	ToNearestAway
	ToZero
	AwayFromZero
	ToNegativeInf
	ToPositiveInf
)
