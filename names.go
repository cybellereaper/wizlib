package wizlib

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Name struct {
	First  string
	Middle string
	Last   string
}

type AcceptedNames struct {
	Names []string `json:"names"`
}

func CreateName(input string, acceptedNames AcceptedNames) (string, error) {
	// Define a regular expression pattern to match the input name
	pattern := fmt.Sprintf(`(?i)^(%s)( (%s))?((%s))?$`, strings.Join(acceptedNames.Names, "|"), strings.Join(acceptedNames.Names, "|"), strings.Join(acceptedNames.Names, "|"))
	nameRegex := regexp.MustCompile(pattern)

	// Parse the input name into a Name struct
	nameParts := strings.Split(input, " ")
	var name Name
	switch len(nameParts) {
	case 1:
		name.First = nameParts[0]
	case 2:
		name.First = nameParts[0]
		name.Last = nameParts[1]
	default:
		name.First = nameParts[0]
		name.Middle = strings.Join(nameParts[1:len(nameParts)-1], " ")
		name.Last = nameParts[len(nameParts)-1]
	}

	if nameRegex.MatchString(input) {
		return fmt.Sprintf("%s %s%s", name.First, name.Middle, name.Last), nil
	}

	return "", errors.New("sorry, the name is not accepted")
}

func IsNameAccepted(name Name, acceptedNames []string) bool {
	// Launch a goroutine to compare each name in the list of accepted names
	for _, acceptedName := range acceptedNames {
		if acceptedName == fmt.Sprintf("%s %s %s", name.First, name.Middle, name.Last) {
			return true
		}
	}
	return false
}
