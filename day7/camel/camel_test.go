package camel

import (
	"reflect"
	"testing"
)

func TestCreateHand(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want Hand
	}{
		{name: "should return correct hand", args: args{input: "32T3K 765"}, want: Hand{
			cards: [5]Card{'3', '2', 'T', '3', 'K'},
			bid:   765,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateHand(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateHand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHand_GetHandType(t *testing.T) {
	type fields struct {
		cards [5]Card
		bid   int
	}
	tests := []struct {
		name   string
		fields fields
		want   HandType
	}{
		{name: "should return one pair", fields: struct {
			cards [5]Card
			bid   int
		}{cards: [5]Card{'2', '2', '3', '4', '5'}, bid: 100}, want: ONE_PAIR},
		{name: "should return two pair", fields: struct {
			cards [5]Card
			bid   int
		}{cards: [5]Card{'2', '2', '3', '3', '5'}, bid: 100}, want: TWO_PAIR},
		{name: "should return three of a kind", fields: struct {
			cards [5]Card
			bid   int
		}{cards: [5]Card{'2', '2', '2', '4', '5'}, bid: 100}, want: THREE_OF_A_KIND},
		{name: "should return full house", fields: struct {
			cards [5]Card
			bid   int
		}{cards: [5]Card{'2', '2', '2', '4', '4'}, bid: 100}, want: FULL_HOUSE},
		{name: "should return four of a kind", fields: struct {
			cards [5]Card
			bid   int
		}{cards: [5]Card{'2', '2', '2', '2', '5'}, bid: 100}, want: FOUR_OF_A_KIND},
		{name: "should return one pair", fields: struct {
			cards [5]Card
			bid   int
		}{cards: [5]Card{'2', '2', '2', '2', '2'}, bid: 100}, want: FIVE_OF_A_KIND},
		{name: "should return high card", fields: struct {
			cards [5]Card
			bid   int
		}{cards: [5]Card{'2', '6', '3', '4', '5'}, bid: 100}, want: HIGH_CARD},
		{name: "should return handle joker ", fields: struct {
			cards [5]Card
			bid   int
		}{cards: [5]Card{'K', 'T', 'J', 'J', 'T'}, bid: 100}, want: FOUR_OF_A_KIND},

		{name: "should return handle joker ", fields: struct {
			cards [5]Card
			bid   int
		}{cards: [5]Card{'3', 'K', '6', '6', 'J'}, bid: 100}, want: THREE_OF_A_KIND},
		{name: "should return handle joker ", fields: struct {
			cards [5]Card
			bid   int
		}{cards: [5]Card{'3', 'J', 'J', 'J', 'J'}, bid: 100}, want: FIVE_OF_A_KIND},
		{name: "should return handle joker ", fields: struct {
			cards [5]Card
			bid   int
		}{cards: [5]Card{'3', '3', '3', '3', 'J'}, bid: 100}, want: FIVE_OF_A_KIND},
		{name: "should return handle joker ", fields: struct {
			cards [5]Card
			bid   int
		}{cards: [5]Card{'3', '3', '3', '4', 'J'}, bid: 100}, want: FOUR_OF_A_KIND},
		{name: "should return handle joker ", fields: struct {
			cards [5]Card
			bid   int
		}{cards: [5]Card{'3', '4', 'J', 'J', 'J'}, bid: 100}, want: FOUR_OF_A_KIND},
		{name: "should return handle joker ", fields: struct {
			cards [5]Card
			bid   int
		}{cards: [5]Card{'3', '3', 'J', 'J', 'J'}, bid: 100}, want: FIVE_OF_A_KIND},
		{name: "should return handle joker ", fields: struct {
			cards [5]Card
			bid   int
		}{cards: [5]Card{'3', '3', '3', 'J', 'J'}, bid: 100}, want: FIVE_OF_A_KIND},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hand{
				cards: tt.fields.cards,
				bid:   tt.fields.bid,
			}
			if got := h.GetHandType(); got != tt.want {
				t.Errorf("GetHandType() = %v, want %v", got, tt.want)
			}
		})
	}
}
