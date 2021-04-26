package input

import (
	"errors"
	"strconv"
	"strings"
)

type Input struct {
	CrewMembers map[int]string
	Parts       []Part
}

type Part struct {
	ID          int
	CrewID      int
	TimeSpent   float64
	CountryCode string
}

func ParseInput(input string) (Input, error) {
	splittedInput := strings.Split(input, "\n")
	result := Input{
		CrewMembers: map[int]string{},
		Parts:       []Part{},
	}

	crewCount, err := strconv.Atoi(splittedInput[0])
	if err != nil {
		return Input{}, errors.New("invalid input, non-numeric crew count")
	}

	partCount, err := strconv.Atoi(splittedInput[crewCount+1])
	if err != nil {
		return Input{}, errors.New("invalid input, non-numeric part count")
	}

	crewStart := 1
	partStart := crewCount + 2

	for i := 0; i < crewCount; i++ {
		element := splittedInput[crewStart+i]
		crewLine := strings.Split(element, " ")
		crewID, err := strconv.Atoi(crewLine[0])
		if err != nil {
			return Input{}, errors.New("invalid input, non-numeric crew ID")
		}
		result.CrewMembers[crewID] = crewLine[1]
	}

	for i := 0; i < partCount; i++ {
		element := splittedInput[partStart+i]
		partLine := strings.Split(element, " ")
		partID, err := strconv.Atoi(partLine[0])
		if err != nil {
			return Input{}, errors.New("invalid input, non-numeric part ID")
		}
		crewID, err := strconv.Atoi(partLine[1])
		if err != nil {
			return Input{}, errors.New("invalid input, non-numeric crew ID")
		}
		minutesWorked, err := strconv.ParseFloat(partLine[2], 64)
		if err != nil {
			return Input{}, errors.New("invalid input, non-numeric minutes worked")
		}

		result.Parts = append(result.Parts, Part{
			ID:          partID,
			CrewID:      crewID,
			TimeSpent:   minutesWorked,
			CountryCode: result.CrewMembers[crewID],
		})
	}
	return result, nil
}
