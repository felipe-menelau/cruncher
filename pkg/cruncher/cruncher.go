package cruncher

import (
	"cruncher/internal/input"
	"cruncher/internal/output"
	"sync"
)

func Crunch(inputString string) (string, error) {
	var wg sync.WaitGroup
	parsedInput, err := input.ParseInput(inputString)

	if err != nil {
		return "", err
	}

	result := output.Output{
		RocketeersAverageTime: map[int]float64{},
		CountriesAverageTime:  map[string]float64{},
		PartsAverageTime:      map[int]float64{},
	}

	wg.Add(3)
	go calculateCountriesAverage(parsedInput, &result, &wg)
	go calculateRocketeerAverage(parsedInput, &result, &wg)
	go calculatePartsAverage(parsedInput, &result, &wg)

	wg.Wait()
	return output.SerializeOutput(result), nil
}

func calculateCountriesAverage(parsedInput input.Input, result *output.Output, group *sync.WaitGroup) {
	defer group.Done()
	countryCounter := make(map[string]int)

	for _, element := range parsedInput.Parts {
		result.CountriesAverageTime[element.CountryCode] += element.TimeSpent
		countryCounter[element.CountryCode] += 1
	}

	for key := range result.CountriesAverageTime {
		result.CountriesAverageTime[key] = result.CountriesAverageTime[key] / float64(countryCounter[key])
	}
}

func calculateRocketeerAverage(parsedInput input.Input, result *output.Output, group *sync.WaitGroup) {
	defer group.Done()
	rocketeerCounter := make(map[int]int)

	for _, element := range parsedInput.Parts {
		result.RocketeersAverageTime[element.CrewID] += element.TimeSpent
		rocketeerCounter[element.CrewID] += 1
	}

	for key := range result.RocketeersAverageTime {
		result.RocketeersAverageTime[key] = result.RocketeersAverageTime[key] / float64(rocketeerCounter[key])
	}
}

func calculatePartsAverage(parsedInput input.Input, result *output.Output, group *sync.WaitGroup) {
	defer group.Done()
	partsCounter := make(map[int]int)

	for _, element := range parsedInput.Parts {
		result.PartsAverageTime[element.ID] += element.TimeSpent
		partsCounter[element.ID] += 1
	}

	for key := range result.PartsAverageTime {
		result.PartsAverageTime[key] = result.PartsAverageTime[key] / float64(partsCounter[key])
	}
}
