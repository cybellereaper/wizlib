package wizlib

import (
	"regexp"
	"strconv"
	"strings"
)

type PetInfo struct {
	School         string   `json:"school,omitempty"`
	Pedigree       int      `json:"pedigree,omitempty"`
	Egg            string   `json:"egg,omitempty"`
	HatchMinutes   int      `json:"hatchMinutes,omitempty"`
	HatchCost      string   `json:"hatchCost,omitempty"`
	Kiosk          string   `json:"kiosk,omitempty"`
	Description    string   `json:"description,omitempty"`
	Strength       int      `json:"strength,omitempty"`
	Intellect      int      `json:"intellect,omitempty"`
	Agility        int      `json:"agility,omitempty"`
	Will           int      `json:"will,omitempty"`
	Power          int      `json:"power,omitempty"`
	Talents        []string `json:"talents,omitempty"`
	DerbyAbilities []string `json:"derbyAbilities,omitempty"`
	LikeClass      string   `json:"likeClass,omitempty"`
	LikedSnacks    []string `json:"likedSnacks,omitempty"`
	LovedSnacks    []string `json:"lovedSnacks,omitempty"`
	LooksLike      []string `json:"looksLike,omitempty"`
	FormedHybrids  []string `json:"formedHybrids,omitempty"`
	FishChestLoc   []string `json:"fishChestLocations,omitempty"`
}

func ParseItemInfobox(template string) map[string]string {
	result := make(map[string]string)

	lines := strings.Split(template, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "|") {
			parts := strings.SplitN(line[1:], "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])

				if value != "" {
					result[key] = value
				}
			}
		}
	}

	return result
}

func ParsePetInfo(infoBox string) PetInfo {
	petInfo := PetInfo{}
	r := regexp.MustCompile(`\| (.*?) = (.*?)(?:\n|$)`)
	matches := r.FindAllStringSubmatch(infoBox, -1)

	for _, match := range matches {
		field := strings.TrimSpace(match[1])
		value := strings.TrimSpace(match[2])

		switch field {
		case "school":
			petInfo.School = value
		case "pedigree":
			pedigree, _ := strconv.Atoi(value)
			petInfo.Pedigree = pedigree
		case "egg":
			petInfo.Egg = value
		case "hatchm":
			hatchMinutes, _ := strconv.Atoi(value)
			petInfo.HatchMinutes = hatchMinutes
		case "hatchcost":
			petInfo.HatchCost = value
		case "kiosk":
			petInfo.Kiosk = value
		case "descrip":
			petInfo.Description = value
		case "strength":
			strength, _ := strconv.Atoi(value)
			petInfo.Strength = strength
		case "intellect":
			intellect, _ := strconv.Atoi(value)
			petInfo.Intellect = intellect
		case "agility":
			agility, _ := strconv.Atoi(value)
			petInfo.Agility = agility
		case "will":
			will, _ := strconv.Atoi(value)
			petInfo.Will = will
		case "power":
			power, _ := strconv.Atoi(value)
			petInfo.Power = power
		case "talent1", "talent2", "talent3", "talent4", "talent5", "talent6", "talent7", "talent8", "talent9", "talent10":
			petInfo.Talents = append(petInfo.Talents, value)
		case "derbyability1", "derbyability2", "derbyability3", "derbyability4", "derbyability5", "derbyability6", "derbyability7", "derbyability8", "derbyability9", "derbyability10":
			petInfo.DerbyAbilities = append(petInfo.DerbyAbilities, value)
		case "likeclass":
			petInfo.LikeClass = value
		case "likedsnacks":
			snacks := strings.Split(value, ";\n")
			petInfo.LikedSnacks = append(petInfo.LikedSnacks, snacks...)
		case "lovedsnacks":
			snacks := strings.Split(value, ";\n")
			petInfo.LovedSnacks = append(petInfo.LovedSnacks, snacks...)
		case "lookslike":
			looks := strings.Split(value, ";\n")
			petInfo.LooksLike = append(petInfo.LooksLike, looks...)
		case "formedhybrids":
			hybrids := strings.Split(value, "::")
			hybridList := make([]string, len(hybrids))
			for i, hybrid := range hybrids {
				hybridList[i] = strings.TrimSpace(hybrid)
			}
			petInfo.FormedHybrids = hybridList
		case "fishchestlocations":
			locations := strings.Split(value, ";\n")
			petInfo.FishChestLoc = append(petInfo.FishChestLoc, locations...)
		}
	}

	return petInfo
}
