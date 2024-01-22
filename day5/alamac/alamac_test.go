package alamac

import (
	"reflect"
	"testing"
)

func TestCreateConversionMap(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want ConversionMap
	}{
		{name: "should create conversion map with one entry", args: struct{ input []string }{input: []string{"50 89 2"}}, want: ConversionMap{entries: []ConversionMapEntry{{
			destinationRangeStart: 50,
			sourceRangeStart:      89,
			rangeLength:           2,
		}}}},
		{name: "should create conversion map with two entries", args: struct{ input []string }{input: []string{"50 89 2", "52 50 48"}}, want: ConversionMap{entries: []ConversionMapEntry{{
			destinationRangeStart: 50,
			sourceRangeStart:      89,
			rangeLength:           2,
		},
			{
				destinationRangeStart: 52,
				sourceRangeStart:      50,
				rangeLength:           48,
			}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateConversionMap(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateConversionMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createConversionEntry(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want ConversionMapEntry
	}{
		{name: "should create entry", args: struct{ input string }{input: "50 89 2"}, want: ConversionMapEntry{
			destinationRangeStart: 50,
			sourceRangeStart:      89,
			rangeLength:           2,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createConversionEntry(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createConversionEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConversionMap_IsMapped(t *testing.T) {
	type fields struct {
		entries []ConversionMapEntry
	}
	type args struct {
		sourceNo int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "should return true if the source is mapped", fields: struct{ entries []ConversionMapEntry }{entries: []ConversionMapEntry{
			{
				destinationRangeStart: 2,
				sourceRangeStart:      2,
				rangeLength:           2,
			},
		}}, args: struct{ sourceNo int }{sourceNo: 3}, want: true},
		{name: "should return true if the source is mapped, source number is start", fields: struct{ entries []ConversionMapEntry }{entries: []ConversionMapEntry{
			{
				destinationRangeStart: 2,
				sourceRangeStart:      2,
				rangeLength:           2,
			},
		}}, args: struct{ sourceNo int }{sourceNo: 2}, want: true},
		{name: "should return false if the source is not mapped, source is less", fields: struct{ entries []ConversionMapEntry }{entries: []ConversionMapEntry{
			{
				destinationRangeStart: 2,
				sourceRangeStart:      2,
				rangeLength:           2,
			},
		}}, args: struct{ sourceNo int }{sourceNo: 1}, want: false},
		{name: "should return false if the source is not mapped, source is higher", fields: struct{ entries []ConversionMapEntry }{entries: []ConversionMapEntry{
			{
				destinationRangeStart: 2,
				sourceRangeStart:      2,
				rangeLength:           2,
			},
		}}, args: struct{ sourceNo int }{sourceNo: 4}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := ConversionMap{
				entries: tt.fields.entries,
			}
			if got := m.IsMapped(tt.args.sourceNo); got != tt.want {
				t.Errorf("IsMapped() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConversionMap_GetConvertedIndex(t *testing.T) {
	type fields struct {
		entries []ConversionMapEntry
	}
	type args struct {
		sourceNo int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{name: "should return given number if the source is not mapped", fields: struct{ entries []ConversionMapEntry }{entries: []ConversionMapEntry{
			{
				destinationRangeStart: 2,
				sourceRangeStart:      2,
				rangeLength:           2,
			},
		}}, args: struct{ sourceNo int }{sourceNo: 4}, want: 4},
		{name: "should return converted number if the source is mapped", fields: struct{ entries []ConversionMapEntry }{entries: []ConversionMapEntry{
			{
				destinationRangeStart: 2,
				sourceRangeStart:      2,
				rangeLength:           2,
			},
		}}, args: struct{ sourceNo int }{sourceNo: 3}, want: 3},
		{name: "should return converted number if the source is mapped", fields: struct{ entries []ConversionMapEntry }{entries: []ConversionMapEntry{
			{
				destinationRangeStart: 2,
				sourceRangeStart:      2,
				rangeLength:           2,
			},
		}}, args: struct{ sourceNo int }{sourceNo: 2}, want: 2},
		{name: "should return converted number if the source is mapped", fields: struct{ entries []ConversionMapEntry }{entries: []ConversionMapEntry{
			{
				destinationRangeStart: 1,
				sourceRangeStart:      2,
				rangeLength:           2,
			},
		}}, args: struct{ sourceNo int }{sourceNo: 3}, want: 2},
		{name: "should return converted number if the source is mapped - large", fields: struct{ entries []ConversionMapEntry }{entries: []ConversionMapEntry{
			{
				destinationRangeStart: 52,
				sourceRangeStart:      50,
				rangeLength:           48,
			},
		}}, args: struct{ sourceNo int }{sourceNo: 50}, want: 52},
		{name: "should return converted number if the source is mapped - large", fields: struct{ entries []ConversionMapEntry }{entries: []ConversionMapEntry{
			{
				destinationRangeStart: 50,
				sourceRangeStart:      98,
				rangeLength:           2,
			},
			{
				destinationRangeStart: 52,
				sourceRangeStart:      50,
				rangeLength:           48,
			},
		}}, args: struct{ sourceNo int }{sourceNo: 99}, want: 51},
		{name: "should return converted number if the source is mapped - large", fields: struct{ entries []ConversionMapEntry }{entries: []ConversionMapEntry{
			{
				destinationRangeStart: 50,
				sourceRangeStart:      98,
				rangeLength:           2,
			},
			{
				destinationRangeStart: 52,
				sourceRangeStart:      50,
				rangeLength:           48,
			},
		}}, args: struct{ sourceNo int }{sourceNo: 97}, want: 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := ConversionMap{
				entries: tt.fields.entries,
			}
			if got := m.GetConvertedIndex(tt.args.sourceNo); got != tt.want {
				t.Errorf("GetConvertedIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
