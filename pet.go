package wizlib

import (
	"fmt"

	"github.com/icza/gox/mathx"
)

// CompareFactors compares two factors and returns a result based on a formula.
func CompareFactors(petOne, petTwo int64) int64 {
	return (100 * (11 - petOne) / (22 - (petOne + petTwo)))
}

// bitsetTimes calculates a result based on bitwise operations and a multiplier.
func bitsetTimes(att1, att2, att3 int64, offset float64) string {
	return fmt.Sprintf("%.1f%%", mathx.Round(float64(att1<<1+att2<<1+att3)*offset, 0.1))
}

// bitsetDiv calculates a result based on bitwise operations and a divisor.
func bitsetDiv(att1, att2, att3 int64, offset float64) string {
	return fmt.Sprintf("%.1f%%", mathx.Round(float64(att1<<1+att2<<1+att3)/offset, 0.1))
}

type PetAttributes struct {
	Strength     int64 `json:"strength"`
	Willpower    int64 `json:"will"`
	Intelligence int64 `json:"intelligence"`
	Power        int64 `json:"power"`
	Agility      int64 `json:"agility"`
	Happiness    int64 `json:"happiness"`
}

// Calculate calculates various attributes based on pet attributes.
func (pa *PetAttributes) Calculate() map[string]map[string]interface{} {
	result := make(map[string]map[string]interface{})
	result["damage"] = map[string]interface{}{
		"bringer": bitsetDiv(pa.Strength, pa.Willpower, pa.Power, 400.0),
		"giver":   bitsetDiv(pa.Strength, pa.Willpower, pa.Power, 200.0),
		"dealer":  bitsetTimes(pa.Strength, pa.Willpower, pa.Power, 0.0075),
	}
	result["resist"] = map[string]interface{}{
		"ward":  bitsetTimes(pa.Strength, pa.Agility, pa.Power, 0.012),
		"proof": bitsetDiv(pa.Strength, pa.Agility, pa.Power, 125.0),
	}
	result["critical"] = map[string]interface{}{
		"defender": bitsetTimes(pa.Intelligence, pa.Willpower, pa.Power, 0.024),
		"blocker":  bitsetTimes(pa.Intelligence, pa.Willpower, pa.Power, 0.02),
	}
	result["pierce"] = map[string]interface{}{
		"breaker": bitsetDiv(pa.Strength, pa.Agility, pa.Power, 400.0),
		"piercer": bitsetTimes(pa.Strength, pa.Agility, pa.Power, 0.0015),
	}
	result["stun"] = map[string]interface{}{
		"recal":  bitsetDiv(pa.Strength, pa.Intelligence, pa.Power, 125.0),
		"resist": bitsetDiv(pa.Strength, pa.Intelligence, pa.Power, 250.0),
	}
	result["healing"] = map[string]interface{}{
		"lively": bitsetTimes(pa.Strength, pa.Agility, pa.Power, 0.0065),
		"healer": bitsetTimes(pa.Strength, pa.Agility, pa.Power, 0.003),
		"medic":  bitsetTimes(pa.Strength, pa.Agility, pa.Power, 0.0065),
	}
	result["health"] = map[string]interface{}{
		"healthy": bitsetTimes(pa.Intelligence, pa.Agility, pa.Power, 0.003),
		"gift":    bitsetTimes(pa.Agility, pa.Willpower, pa.Power, 0.1),
		"add":     bitsetTimes(pa.Agility, pa.Willpower, pa.Power, 0.06),
	}
	result["attributes"] = map[string]interface{}{
		"strength":  pa.Strength,
		"intellect": pa.Intelligence,
		"agility":   pa.Agility,
		"willpower": pa.Willpower,
		"power":     pa.Power,
		"happiness": pa.Happiness,
	}
	return result
}
