// +build cgo

package round

import "math"
import "testing"

import "github.com/mhale/cround"
import "github.com/mhale/round"

const RangeMin = -1000.0
const RangeMax = +1000.0
const RangeStep = +0.05

func TestRoundVersusCRound(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	t.Run("ToNearestEven", func(t *testing.T) {
		for i := RangeMin; i <= RangeMax; i += RangeStep {
			// Prevent the accumulated inaccuracy in i from influencing the rounding
			i = math.Trunc(i*100) / 100
			result := round.ToNearestEven(i)
			resultInC := cround.ToNearestEven(i)
			if result != resultInC {
				t.Errorf("ToNearestEven(%.2f) = %.2f, expected %.2f", i, result, resultInC)
			}
		}
	})

	t.Run("ToNearestAway", func(t *testing.T) {
		for i := RangeMin; i <= RangeMax; i += RangeStep {
			// Prevent the accumulated inaccuracy in i from influencing the rounding
			i = math.Trunc(i*100) / 100
			result := round.ToNearestAway(i)
			resultInC := cround.ToNearestAway(i)
			if result != resultInC {
				t.Errorf("ToNearestAway(%.2f) = %.2f, expected %.2f", i, result, resultInC)
			}
		}
	})

	t.Run("ToZero", func(t *testing.T) {
		for i := RangeMin; i <= RangeMax; i += RangeStep {
			// Prevent the accumulated inaccuracy in i from influencing the rounding
			i = math.Trunc(i*100) / 100
			result := round.ToZero(i)
			resultInC := cround.ToZero(i)
			if result != resultInC {
				t.Errorf("ToZero(%.2f) = %.2f, expected %.2f", i, result, resultInC)
			}
		}
	})

	// C does not have an "away from zero" rounding mode, so we cannot test it.

	t.Run("ToPositiveInf", func(t *testing.T) {
		for i := RangeMin; i <= RangeMax; i += RangeStep {
			// Prevent the accumulated inaccuracy in i from influencing the rounding
			i = math.Trunc(i*100) / 100
			result := round.ToPositiveInf(i)
			resultInC := cround.ToPositiveInf(i)
			if result != resultInC {
				t.Errorf("ToPositiveInf(%.2f) = %.2f, expected %.2f", i, result, resultInC)
			}
		}
	})

	t.Run("ToNegativeInf", func(t *testing.T) {
		for i := RangeMin; i <= RangeMax; i += RangeStep {
			// Prevent the accumulated inaccuracy in i from influencing the rounding
			i = math.Trunc(i*100) / 100
			result := round.ToNegativeInf(i)
			resultInC := cround.ToNegativeInf(i)
			if result != resultInC {
				t.Errorf("ToNegativeInf(%.2f) = %.2f, expected %.2f", i, result, resultInC)
			}
		}
	})

}
