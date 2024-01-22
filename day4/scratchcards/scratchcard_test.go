package scratchcards

import (
	"reflect"
	"testing"
)

func TestScratchCard_GetScore(t *testing.T) {
	type fields struct {
		winningNumbers []int
		givenNumbers   []int
	}
	tests := []struct {
		name   string
		fields fields
		want   uint
	}{
		{name: "should return 0 when no matching numbers", fields: struct {
			winningNumbers []int
			givenNumbers   []int
		}{winningNumbers: []int{1, 2, 3}, givenNumbers: []int{4, 5, 6}}, want: 0},
		{name: "should return 1 when 1 matching number", fields: struct {
			winningNumbers []int
			givenNumbers   []int
		}{winningNumbers: []int{1, 2, 3}, givenNumbers: []int{1, 5, 6}}, want: 1},
		{name: "should return 2 when 2 matching number", fields: struct {
			winningNumbers []int
			givenNumbers   []int
		}{winningNumbers: []int{1, 2, 3}, givenNumbers: []int{1, 2, 6}}, want: 2},
		{name: "should return 4 when 3 matching number", fields: struct {
			winningNumbers []int
			givenNumbers   []int
		}{winningNumbers: []int{1, 2, 3}, givenNumbers: []int{1, 2, 3}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ScratchCard{
				winningNumbers: tt.fields.winningNumbers,
				givenNumbers:   tt.fields.givenNumbers,
			}
			if got := c.GetScore(); got != tt.want {
				t.Errorf("GetScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateScratchCard(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want *ScratchCard
	}{
		{name: "should return created card", args: struct{ input string }{input: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"}, want: &ScratchCard{
			id:             "Card 1",
			winningNumbers: []int{41, 48, 83, 86, 17},
			givenNumbers:   []int{83, 86, 6, 31, 17, 9, 48, 53},
			count:          1,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createScratchCard(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createScratchCard() = %v, want %v", got, tt.want)
			}
		})
	}
}
