package output

import "fmt"

type Output struct {
	RocketeersAverageTime map[int]float64
	CountriesAverageTime  map[string]float64
	PartsAverageTime      map[int]float64
}

func SerializeOutput(output Output) string {
	result := ""
	for key, value := range output.RocketeersAverageTime {
		result = result + fmt.Sprintf("%d %.2f\n", key, value)
	}

	result = result + "---\n"
	for key, value := range output.CountriesAverageTime {
		result = result + fmt.Sprintf("%s %.2f\n", key, value)
	}

	result = result + "---\n"
	for key, value := range output.PartsAverageTime {
		result = result + fmt.Sprintf("%d %.2f\n", key, value)
	}

	return result
}
