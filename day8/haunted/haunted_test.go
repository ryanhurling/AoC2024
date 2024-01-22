package haunted

import (
	"reflect"
	"testing"
)

func TestCreateGraph(t *testing.T) {
	type args struct {
		entries []string
	}
	tests := []struct {
		name string
		args args
		want Graph
	}{
		{name: "should create graph with single Node", args: args{entries: []string{"AAA = (BBB, CCC)"}}, want: map[string]Node{
			"AAA": {
				id:       "AAA",
				leftKey:  "BBB",
				rightKey: "CCC",
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateGraph(tt.args.entries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateInstructions(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want Instructions
	}{
		{name: "should return LR", args: args{str: "LR"}, want: Instructions{LEFT, RIGHT}},
		{name: "should return LRRL", args: args{str: "LRRL"}, want: Instructions{LEFT, RIGHT, RIGHT, LEFT}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateInstructions(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateInstructions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createNode(t *testing.T) {
	type args struct {
		entry string
	}
	tests := []struct {
		name string
		args args
		want Node
	}{
		{name: "should creat Node", args: args{entry: "AAA = (BBB, CCC)"}, want: Node{
			id:       "AAA",
			leftKey:  "BBB",
			rightKey: "CCC",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := createNode(tt.args.entry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountSteps(t *testing.T) {
	type args struct {
		graph        Graph
		instructions Instructions
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "should return 2", args: args{
			graph: Graph{
				"AAA": Node{
					id:       "AAA",
					leftKey:  "BBB",
					rightKey: "CCC",
				},
				"BBB": Node{
					id:       "BBB",
					leftKey:  "DDD",
					rightKey: "EEE",
				},
				"CCC": Node{
					id:       "CCC",
					leftKey:  "ZZZ",
					rightKey: "GGG",
				},
				"DDD": Node{
					id:       "DDD",
					leftKey:  "DDD",
					rightKey: "DDD",
				},
				"EEE": Node{
					id:       "EEE",
					leftKey:  "EEE",
					rightKey: "EEE",
				},
				"GGG": Node{
					id:       "GGG",
					leftKey:  "GGG",
					rightKey: "GGG",
				},
				"ZZZ": Node{
					id:       "ZZZ",
					leftKey:  "BBB",
					rightKey: "CCC",
				},
				"_ROOT_": Node{
					id:       "AAA",
					leftKey:  "BBB",
					rightKey: "CCC",
				},
			},
			instructions: Instructions{RIGHT, LEFT},
		}, want: 2},
		{name: "should return 4", args: args{
			graph: Graph{
				"AAA": Node{
					id:       "AAA",
					leftKey:  "BBB",
					rightKey: "CCC",
				},
				"BBB": Node{
					id:       "BBB",
					leftKey:  "DDD",
					rightKey: "EEE",
				},
				"CCC": Node{
					id:       "CCC",
					leftKey:  "GGG",
					rightKey: "ZZZ",
				},
				"DDD": Node{
					id:       "DDD",
					leftKey:  "DDD",
					rightKey: "DDD",
				},
				"EEE": Node{
					id:       "EEE",
					leftKey:  "EEE",
					rightKey: "EEE",
				},
				"GGG": Node{
					id:       "GGG",
					leftKey:  "GGG",
					rightKey: "CCC",
				},
				"ZZZ": Node{
					id:       "ZZZ",
					leftKey:  "BBB",
					rightKey: "CCC",
				},
				"_ROOT_": Node{
					id:       "AAA",
					leftKey:  "BBB",
					rightKey: "CCC",
				},
			},
			instructions: Instructions{RIGHT, LEFT, RIGHT},
		}, want: 4},
		{name: "should return 6", args: args{
			graph: Graph{
				"AAA": Node{
					id:       "AAA",
					leftKey:  "BBB",
					rightKey: "BBB",
				},
				"BBB": Node{
					id:       "BBB",
					leftKey:  "AAA",
					rightKey: "ZZZ",
				},
				"ZZZ": Node{
					id:       "ZZZ",
					leftKey:  "ZZZ",
					rightKey: "ZZZ",
				},
				"_ROOT_": Node{
					id:       "AAA",
					leftKey:  "BBB",
					rightKey: "BBB",
				},
			},
			instructions: Instructions{LEFT, LEFT, RIGHT},
		}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSteps(tt.args.graph, tt.args.instructions); got != tt.want {
				t.Errorf("CountSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
