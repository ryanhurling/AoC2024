package race

import "testing"

func TestRace_GetLowestPossibleHeldTime(t *testing.T) {
	type fields struct {
		time           int
		recordDistance int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "should return 6 for lowest", fields: struct {
			time           int
			recordDistance int
		}{time: 53, recordDistance: 250}, want: 6},
		{name: "should return 2 for lowest", fields: struct {
			time           int
			recordDistance int
		}{time: 7, recordDistance: 9}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Race{
				time:           tt.fields.time,
				recordDistance: tt.fields.recordDistance,
			}
			if got := r.GetLowestPossibleHeldTime(); got != tt.want {
				t.Errorf("GetLowestPossibleHeldTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRace_GetHighestPossibleHeldTime(t *testing.T) {
	type fields struct {
		time           int
		recordDistance int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "should return 47 for highest", fields: struct {
			time           int
			recordDistance int
		}{time: 53, recordDistance: 250}, want: 47},
		{name: "should return 5 for highest", fields: struct {
			time           int
			recordDistance int
		}{time: 7, recordDistance: 9}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Race{
				time:           tt.fields.time,
				recordDistance: tt.fields.recordDistance,
			}
			if got := r.GetHighestPossibleHeldTime(); got != tt.want {
				t.Errorf("GetHighestPossibleHeldTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRace_TotalPossibleWinCases(t *testing.T) {
	type fields struct {
		time           int
		recordDistance int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "should return 4", fields: struct {
			time           int
			recordDistance int
		}{time: 7, recordDistance: 9}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Race{
				time:           tt.fields.time,
				recordDistance: tt.fields.recordDistance,
			}
			if got := r.TotalPossibleWinCases(); got != tt.want {
				t.Errorf("TotalPossibleWinCases() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRaces_TotalPossibleNoOfCases(t *testing.T) {
	tests := []struct {
		name string
		r    Races
		want int
	}{
		{name: "should return 288", r: Races{
			Race{
				time:           7,
				recordDistance: 9,
			},
			Race{
				time:           15,
				recordDistance: 40,
			},
			Race{
				time:           30,
				recordDistance: 200,
			},
		}, want: 288},
		{name: "should return 71503", r: Races{
			Race{
				time:           71530,
				recordDistance: 940200,
			},
		}, want: 71503},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.TotalPossibleNoOfCases(); got != tt.want {
				t.Errorf("TotalPossibleNoOfCases() = %v, want %v", got, tt.want)
			}
		})
	}
}
