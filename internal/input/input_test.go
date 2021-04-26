package input

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    Input
		wantErr bool
	}{
		{
			name: "Test case",
			args: args{
				input:
				`1
1 PT
1
1 1 10`,
			},
			want: Input{
				CrewMembers: map[int]string{1: "PT"},
				Parts: []Part{{
					ID:          1,
					CrewID:      1,
					TimeSpent:   10,
					CountryCode: "PT",
				},
				},
			},
			wantErr: false,
		},
		{
			name: "Error Test case",
			args: args{
				input:
				`1
BANANA PT
1
1 1 10`,
			},
			want: Input{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInput(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInput() got = %v, want %v", got, tt.want)
			}
		})
	}
}
