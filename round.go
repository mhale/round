package round

import "math"

const Epsilon = 0.0000001

// Round returns the nearest integer value of x, with ties rounded to even integers.
//
// Special cases are:
//	Round(±0) = ±0
//	Round(±Inf) = ±Inf
//	Round(NaN) = NaN
func Round(x float64) float64 {
	return ToNearestEven(x)
}

// RoundTo returns the nearest integer value of x, to dp decimal places, with ties rounded to even integers.
//
// Special cases are:
//	Round(±0, dp) = ±0
//	Round(±Inf, dp) = ±Inf
//	Round(NaN, dp) = NaN
func RoundTo(x float64, dp float64) float64 {
	x = x * math.Pow(10, dp)
	result := ToNearestEven(x)
	return result / math.Pow(10, dp)
}

// ToNearestEven returns the nearest integer value of x, with ties rounded with ties rounded to even integers.
//
// Special cases are:
//	ToNearestEven(±0) = ±0
//	ToNearestEven(±Inf) = ±Inf
//	ToNearestEven(NaN) = NaN
func ToNearestEven(x float64) float64 {
	return toNearest(x, true)
}

// ToNearestAway returns the nearest integer value of x, with ties rounded away from zero.
//
// Special cases are:
//	ToNearestAway(±0) = ±0
//	ToNearestAway(±Inf) = ±Inf
//	ToNearestAway(NaN) = NaN
func ToNearestAway(x float64) float64 {
	return toNearest(x, false)
}

// ToZero returns the nearest integer value toward zero. If x is negative it is rounded up, if it is positive it is rounded down.
//
// Special cases are:
//	ToZero(±0) = ±0
//	ToZero(±Inf) = ±Inf
//	ToZero(NaN) = NaN
func ToZero(x float64) float64 {
	return math.Trunc(x)
}

// AwayFromZero returns the nearest integer value away from zero. If x is negative it is rounded down, if it is positive it is rounded up.
//
// Special cases are:
//	AwayFromZero(±0) = ±0
//	AwayFromZero(±Inf) = ±Inf
//	AwayFromZero(NaN) = NaN
func AwayFromZero(x float64) float64 {
	if x >= 0 {
		return math.Ceil(x)
	} else {
		return math.Floor(x)
	}
}

// ToPositiveInf returns the nearest integer value greater than x.
//
// Special cases are:
//	ToPositiveInf(±0) = ±0
//	ToPositiveInf(±Inf) = ±Inf
//	ToPositiveInf(NaN) = NaN
func ToPositiveInf(x float64) float64 {
	return math.Ceil(x)
}

// ToNegativeInf returns the nearest integer value less than x.
//
// Special cases are:
//	ToNegativeInf(±0) = ±0
//	ToNegativeInf(±Inf) = ±Inf
//	ToNegativeInf(NaN) = NaN
func ToNegativeInf(x float64) float64 {
	return math.Floor(x)
}

// Internal function not intended to be called from outside the package.
func toNearest(x float64, tiesToEven bool) float64 {
	if x == 0 || math.IsNaN(x) || math.IsInf(x, 0) {
		return x
	}

	if x < 0.0 {
		return -toNearest(-x, tiesToEven)
	}

	// Split x into integer and fractional parts
	intPart, fracPart := math.Modf(x)

	// If x is halfway between two integers
	if math.Abs(fracPart-0.5) < Epsilon {
		// If rounding to nearest even and integer part is even, return it
		if tiesToEven {
			if math.Mod(intPart, 2.0) < Epsilon {
				return intPart
			}
		}

		// Else round to furthest integer
		return math.Ceil(intPart + 0.5)
	}

	// Else round to closest integer
	return math.Floor(x + 0.5)
}
