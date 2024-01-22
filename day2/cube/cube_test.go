package cube

import (
	"reflect"
	"testing"
)

func TestGameString_ExtractTurns(t *testing.T) {
	tests := []struct {
		name string
		s    RoundString
		want Turns
	}{
		{name: "should extract turns", s: RoundString("3 red;"), want: Turns{Turn{
			colour: Red,
			number: 3,
		}}},
		{name: "should extract 2 turns", s: RoundString("3 red, 2 blue;"), want: Turns{Turn{
			colour: Red,
			number: 3,
		},
			Turn{
				colour: Blue,
				number: 2,
			}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ExtractTurns(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractTurns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTurn(t *testing.T) {
	type args struct {
		turnString string
	}
	tests := []struct {
		name string
		args args
		want *Turn
	}{
		{name: "should create struct", args: struct{ turnString string }{turnString: "3 red"}, want: &Turn{
			colour: Colour("red"),
			number: 3,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTurn(tt.args.turnString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTurn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRound(t *testing.T) {
	type args struct {
		roundString string
	}
	tests := []struct {
		name string
		args args
		want *Round
	}{
		{name: "should create simple round", args: args{roundString: "6 red;"}, want: &Round{Turns{Turn{
			colour: Red,
			number: 6,
		}}}},
		{name: "should create complex round", args: args{roundString: "6 green, 18 blue, 6 red;"}, want: &Round{Turns{Turn{
			colour: Green,
			number: 6,
		},
			Turn{
				colour: Blue,
				number: 18,
			},
			Turn{
				colour: Red,
				number: 6,
			}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRound(tt.args.roundString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGame(t *testing.T) {
	type args struct {
		gameString string
	}
	tests := []struct {
		name string
		args args
		want *Game
	}{
		{name: "should make simple game", args: args{gameString: "Game 1: 7 green"}, want: &Game{
			Id: 1,
			rounds: Rounds{Round{Turns{Turn{
				colour: Green,
				number: 7,
			}}}}},
		},
		{name: "should make complex game", args: args{gameString: "Game 1: 7 green; 5 red"}, want: &Game{
			Id: 1,
			rounds: Rounds{Round{Turns{Turn{
				colour: Green,
				number: 7,
			}}},
				Round{Turns{Turn{
					colour: Red,
					number: 5,
				}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGame(tt.args.gameString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_MinimumPossibleCubes(t *testing.T) {
	type fields struct {
		Id     int
		rounds Rounds
	}
	tests := []struct {
		name   string
		fields fields
		want   MinimumNeeded
	}{
		{name: "should return correct min", fields: fields{
			Id: 1,
			rounds: Rounds{Round{Turns{Turn{
				colour: Green,
				number: 3,
			}}}},
		}, want: MinimumNeeded{
			Red: Turn{
				colour: Red,
				number: 0,
			},
			Green: Turn{
				colour: Green,
				number: 3,
			},
			Blue: Turn{
				colour: Blue,
				number: 0,
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Game{
				Id:     tt.fields.Id,
				rounds: tt.fields.rounds,
			}
			if got := g.MinimumPossibleCubes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinimumPossibleCubes() = %v, want %v", got, tt.want)
			}
		})
	}
}
