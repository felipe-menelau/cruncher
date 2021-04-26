package cruncher

import (
	"cruncher/internal/input"
	"cruncher/internal/output"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func Test_calculateCountriesAverage(t *testing.T) {
	type args struct {
		parsedInput input.Input
		result      *output.Output
		group       *sync.WaitGroup
	}

	happyResult := output.Output{
		CountriesAverageTime: map[string]float64{},
	}

	var wg sync.WaitGroup

	tests := []struct {
		name string
		args args
	}{
		{name: "Test Case",
			args: args{
				parsedInput: input.Input{
					CrewMembers: map[int]string{1: "BR", 2: "BR", 3: "US"},
					Parts: []input.Part{
						{ID: 25, CrewID: 1, TimeSpent: 10, CountryCode: "BR"},
						{ID: 42, CrewID: 2, TimeSpent: 20, CountryCode: "BR"},
						{ID: 22, CrewID: 3, TimeSpent: 10, CountryCode: "US"}},
				},
				result: &happyResult,
				group:  &wg,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg.Add(1)
			calculateCountriesAverage(tt.args.parsedInput, tt.args.result, tt.args.group)
			assert.Equal(t, 15.00, happyResult.CountriesAverageTime["BR"])
			assert.Equal(t, 10.00, happyResult.CountriesAverageTime["US"])
		})
	}
}

func Test_calculateRocketeerAverage(t *testing.T) {
	type args struct {
		parsedInput input.Input
		result      *output.Output
		group       *sync.WaitGroup
	}

	happyResult := output.Output{
		RocketeersAverageTime: map[int]float64{},
	}

	var wg sync.WaitGroup

	tests := []struct {
		name string
		args args
	}{
		{name: "Test Case",
			args: args{
				parsedInput: input.Input{
					CrewMembers: map[int]string{1: "BR", 3: "US"},
					Parts: []input.Part{
						{ID: 25, CrewID: 1, TimeSpent: 10, CountryCode: "BR"},
						{ID: 42, CrewID: 1, TimeSpent: 20, CountryCode: "BR"},
						{ID: 52, CrewID: 3, TimeSpent: 55, CountryCode: "US"},
						{ID: 22, CrewID: 3, TimeSpent: 10, CountryCode: "US"}},
				},
				result: &happyResult,
				group:  &wg,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg.Add(1)
			calculateRocketeerAverage(tt.args.parsedInput, tt.args.result, tt.args.group)
			assert.Equal(t, 15.00, happyResult.RocketeersAverageTime[1])
			assert.Equal(t, 32.50, happyResult.RocketeersAverageTime[3])
		})
	}
}

func Test_calculatePartsAverage(t *testing.T) {
	type args struct {
		parsedInput input.Input
		result      *output.Output
		group       *sync.WaitGroup
	}

	happyResult := output.Output{
		PartsAverageTime: map[int]float64{},
	}

	var wg sync.WaitGroup

	tests := []struct {
		name string
		args args
	}{
		{name: "Test Case",
			args: args{
				parsedInput: input.Input{
					CrewMembers: map[int]string{1: "BR", 3: "US"},
					Parts: []input.Part{
						{ID: 25, CrewID: 1, TimeSpent: 22, CountryCode: "BR"},
						{ID: 25, CrewID: 3, TimeSpent: 24, CountryCode: "BR"},
						{ID: 22, CrewID: 1, TimeSpent: 33, CountryCode: "US"},
						{ID: 22, CrewID: 3, TimeSpent: 14, CountryCode: "US"}},
				},
				result: &happyResult,
				group:  &wg,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg.Add(1)
			calculatePartsAverage(tt.args.parsedInput, tt.args.result, tt.args.group)
			assert.Equal(t, 23.00, happyResult.PartsAverageTime[25])
			assert.Equal(t, 23.50, happyResult.PartsAverageTime[22])
		})
	}
}
