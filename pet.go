package wizlib

import (
	"fmt"
	"math"
)

type PetAttributes struct {
	Strength     int64 `json:"strength"`
	Willpower    int64 `json:"will"`
	Intelligence int64 `json:"intelligence"`
	Power        int64 `json:"power"`
	Agility      int64 `json:"agility"`
	Happiness    int64 `json:"happiness"`
}

// PetCalculator provides methods for calculating various attributes based on pet attributes.
type PetCalculator struct{}

// NewPetCalculator creates a new instance of PetCalculator.
func NewPetCalculator() *PetCalculator {
	return &PetCalculator{}
}

// Calculate calculates various attributes based on pet attributes.
func (c *PetCalculator) Calculate(pa *PetAttributes) map[string]map[string]interface{} {
	result := make(map[string]map[string]interface{})
	result["damage"] = map[string]interface{}{
		"bringer": c.bitsetDiv(pa.Strength, pa.Willpower, pa.Power, 400.0),
		"giver":   c.bitsetDiv(pa.Strength, pa.Willpower, pa.Power, 200.0),
		"dealer":  c.bitsetTimes(pa.Strength, pa.Willpower, pa.Power, 0.0075),
	}
	result["resist"] = map[string]interface{}{
		"ward":  c.bitsetTimes(pa.Strength, pa.Agility, pa.Power, 0.012),
		"proof": c.bitsetDiv(pa.Strength, pa.Agility, pa.Power, 125.0),
	}
	result["critical"] = map[string]interface{}{
		"defender": c.bitsetTimes(pa.Intelligence, pa.Willpower, pa.Power, 0.024),
		"blocker":  c.bitsetTimes(pa.Intelligence, pa.Willpower, pa.Power, 0.02),
	}
	result["pierce"] = map[string]interface{}{
		"breaker": c.bitsetDiv(pa.Strength, pa.Agility, pa.Power, 400.0),
		"piercer": c.bitsetTimes(pa.Strength, pa.Agility, pa.Power, 0.0015),
	}
	result["stun"] = map[string]interface{}{
		"recal":  c.bitsetDiv(pa.Strength, pa.Intelligence, pa.Power, 125.0),
		"resist": c.bitsetDiv(pa.Strength, pa.Intelligence, pa.Power, 250.0),
	}
	result["healing"] = map[string]interface{}{
		"lively": c.bitsetTimes(pa.Strength, pa.Agility, pa.Power, 0.0065),
		"healer": c.bitsetTimes(pa.Strength, pa.Agility, pa.Power, 0.003),
		"medic":  c.bitsetTimes(pa.Strength, pa.Agility, pa.Power, 0.0065),
	}
	result["health"] = map[string]interface{}{
		"healthy": c.bitsetTimes(pa.Intelligence, pa.Agility, pa.Power, 0.003),
		"gift":    c.bitsetTimes(pa.Agility, pa.Willpower, pa.Power, 0.1),
		"add":     c.bitsetTimes(pa.Agility, pa.Willpower, pa.Power, 0.06),
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

// bitsetTimes calculates a result based on bitwise operations and a multiplier.
func (c *PetCalculator) bitsetTimes(att1, att2, att3 int64, offset float64) string {
	return fmt.Sprintf("%.1f%%", math.Round(float64((att1<<1)+(att2<<1)+(att3<<1))*offset*10)/10)
}

// bitsetDiv calculates a result based on bitwise operations and a divisor.
func (c *PetCalculator) bitsetDiv(att1, att2, att3 int64, offset float64) string {
	return fmt.Sprintf("%.1f%%", math.Round(float64((att1<<1)+(att2<<1)+(att3<<1))/offset*10)/10)
}

// CompareFactors compares two factors and returns a result based on a formula.
func CompareFactors(petOne, petTwo int64) int64 {
	return (100 * (11 - petOne) / (22 - (petOne + petTwo)))
}
