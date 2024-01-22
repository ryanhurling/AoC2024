package schematic

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestSchematic_GetCoordinatesAroundPoint(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		s    Schematic
		args args
		want []Coordinate
	}{
		{name: "simple 3x3 find", s: getSchematic([]string{"...", ".*.", "..."}), args: args{x: 1, y: 1}, want: []Coordinate{
			{
				x: 0,
				y: 0,
			},
			{
				x: 0,
				y: 1,
			},
			{
				x: 0,
				y: 2,
			},
			{
				x: 1,
				y: 0,
			},
			{
				x: 1,
				y: 2,
			},
			{
				x: 2,
				y: 0,
			},
			{
				x: 2,
				y: 1,
			},
			{
				x: 2,
				y: 2,
			},
		}},
		{name: "simple 3x2 find", s: getSchematic([]string{"...", ".*."}), args: args{x: 1, y: 1}, want: []Coordinate{
			{
				x: 0,
				y: 0,
			},
			{
				x: 0,
				y: 1,
			},
			{
				x: 0,
				y: 2,
			},
			{
				x: 1,
				y: 0,
			},
			{
				x: 1,
				y: 2,
			},
		}},
		{name: "simple 3x2 find corner", s: getSchematic([]string{"*..", "..."}), args: args{x: 0, y: 0}, want: []Coordinate{
			{
				x: 0,
				y: 1,
			},
			{
				x: 1,
				y: 0,
			},
			{
				x: 1,
				y: 1,
			},
		}},
		{name: "simple 3x2 find top edge", s: getSchematic([]string{".*.", "..."}), args: args{x: 0, y: 1}, want: []Coordinate{
			{
				x: 0,
				y: 0,
			},
			{
				x: 0,
				y: 2,
			},
			{
				x: 1,
				y: 0,
			},
			{
				x: 1,
				y: 1,
			},
			{
				x: 1,
				y: 2,
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.GetCoordinatesAroundPoint(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCoordinatesAroundPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getSchematic(s []string) Schematic {
	return Create2dArray(s)
}

func TestScan(t *testing.T) {
	type args struct {
		schematic Schematic
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "should return 1",
			args: args{schematic: getSchematic([]string{".1.", ".*.", "..."})},
			want: []int{1},
		},
		{
			name: "should return 1 and 3",
			args: args{schematic: getSchematic([]string{".1.", ".*.", "..3"})},
			want: []int{1, 3},
		},
		{
			name: "should return 1 without 3",
			args: args{schematic: getSchematic([]string{".1.", "*..", "..3"})},
			want: []int{1},
		},
		{
			name: "should return 12",
			args: args{schematic: getSchematic([]string{".12", "*..", "..3"})},
			want: []int{12},
		},
		{
			/*
				21..
				..*.
				....
			*/
			name: "should return 21",
			args: args{schematic: getSchematic([]string{"21..", "..*.", "...."})},
			want: []int{21},
		},
		{
			/*
				21..
				.*..
				....
			*/
			name: "should return 21 only",
			args: args{schematic: getSchematic([]string{"21..", ".*..", "...."})},
			want: []int{21},
		},
		{
			/*
				21..
				.*..
				....
			*/
			name: "should return 34 and 1",
			args: args{schematic: getSchematic([]string{"..34.", ".*.*.", "1...."})},
			want: []int{34, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Scan(tt.args.schematic); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPossibleGear(t *testing.T) {
	type args struct {
		char rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "should return true for *", args: args{char: '*'}, want: true},
		{name: "should return false for non *", args: args{char: '.'}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPossibleGear(tt.args.char); got != tt.want {
				t.Errorf("isPossibleGear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		file, err := os.ReadFile("../data.txt")
		if err != nil {
			panic(err)
		}
		fileContents := string(file)
		inputs := strings.Split(fileContents, "\r\n")
		b.StartTimer()
		s := Create2dArray(inputs)
		nums := Scan(s)
		Reduce(nums)
	}
	b.ReportAllocs()
}
