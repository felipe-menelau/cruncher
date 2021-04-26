package output

import (
	"fmt"
	"sort"
)

type Output struct {
	RocketeersAverageTime map[int]float64
	CountriesAverageTime  map[string]float64
	PartsAverageTime      map[int]float64
}

func SerializeOutput(output Output) string {
	result := ""

	rocketeerKeys := make([]int, 0, len(output.RocketeersAverageTime))
	for k, _ := range output.RocketeersAverageTime {
		rocketeerKeys = append(rocketeerKeys, k)
	}
	sort.Ints(rocketeerKeys)
	for _, key := range rocketeerKeys {
		result = result + fmt.Sprintf("%d %.2f\n", key, output.RocketeersAverageTime[key])
	}

	result = result + "---\n"
	countryKeys := make([]string, 0, len(output.CountriesAverageTime))
	for k, _ := range output.CountriesAverageTime {
		countryKeys = append(countryKeys, k)
	}
	sort.Strings(countryKeys)
	for _, key := range countryKeys {
		result = result + fmt.Sprintf("%s %.2f\n", key, output.CountriesAverageTime[key])
	}

	result = result + "---\n"
	for key, value := range output.PartsAverageTime {
		result = result + fmt.Sprintf("%d %.2f\n", key, value)
	}

	return result
}
