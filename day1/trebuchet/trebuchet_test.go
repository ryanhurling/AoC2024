package trebuchet

import (
	"os"
	"strings"
	"testing"
)

func TestCalculateCalibration(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{name: "empty slice returns 0", args: args{values: make([]string, 0)}, want: 0},
		{name: "t1est returns 11", args: args{values: []string{"t1est"}}, want: 11},
		{name: "t1es2t returns 12", args: args{values: []string{"t1es2t"}}, want: 12},
		{name: "t1e4s2t returns 12", args: args{values: []string{"t1e4s2t"}}, want: 12},
		{name: "t1e4s2t, T2e3st returns 35", args: args{values: []string{"t1e4s2t", "t2e3st"}}, want: 35},
		{name: "one4s2t returns 12", args: args{values: []string{"one4s2t"}}, want: 12},
		{name: "4nineeightseven2 returns 42", args: args{values: []string{"4nineeightseven2"}}, want: 42},
		{name: "7pqrstsixteen returns 76", args: args{values: []string{"7pqrstsixteen"}}, want: 76},
		{name: "zoneight234 returns 14", args: args{values: []string{"zoneight234"}}, want: 14},
		{name: "array should return 281", args: args{values: []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}}, want: 281},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := CalculateCalibration(tt.args.values); got != tt.want || err != nil {
				t.Errorf("CalculateCalibration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test(t *testing.T) {
	file, err := os.ReadFile("../data.txt")
	if err != nil {
		panic(err)
	}
	fileContents := string(file)
	inputs := strings.Split(fileContents, "\r\n")
	if calibration, err := CalculateCalibration(inputs); calibration != 53515 || err != nil {
		t.Errorf("CalculateCalibration() = %v, want %v", calibration, 53515)
	}
}

func TestConvertWordToDigit(t *testing.T) {
	type args struct {
		word          string
		startingIndex int32
	}
	tests := []struct {
		name     string
		args     args
		want     int32
		wantSkip int32
	}{
		{name: "should return 1", args: args{
			word:          "one",
			startingIndex: 0,
		}, want: 1, wantSkip: 2},
		{name: "should return 3", args: args{
			word:          "three",
			startingIndex: 0,
		}, want: 3, wantSkip: 4},
		{name: "should return 2", args: args{
			word:          "two",
			startingIndex: 0,
		}, want: 2, wantSkip: 2},
		{name: "should return 2", args: args{
			word:          "testtwo",
			startingIndex: 4,
		}, want: 2, wantSkip: 2},
		{name: "should return error", args: args{
			word:          "testtwo",
			startingIndex: 0,
		}, want: -1, wantSkip: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, skip := ConvertWordToDigit(tt.args.word, tt.args.startingIndex)
			if skip != tt.wantSkip {
				t.Errorf("ConvertWordToDigit() error = %v, wantSkip %v", skip, tt.wantSkip)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertWordToDigit() got = %v, want %v", got, tt.want)
			}
		})
	}
}
