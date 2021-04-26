package output

import "testing"

func TestSerializeOutput(t *testing.T) {
	type args struct {
		output Output
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case",
			args: args{output: Output{
				RocketeersAverageTime: map[int]float64{1: 2.5},
				CountriesAverageTime:  map[string]float64{"BR": 10.0},
				PartsAverageTime:      map[int]float64{2: 20.0},
			}},
			want: "1 2.50\n---\nBR 10.00\n---\n2 20.00\n",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SerializeOutput(tt.args.output); got != tt.want {
				t.Errorf("SerializeOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}
