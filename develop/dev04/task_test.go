package main

import (
	"reflect"
	"testing"
)

func TestSearchAnagramsInDict(t *testing.T) {
	tests := []struct {
		input  []string
		output map[string][]string
	}{
		{
			input: []string{"слиток", "пятак", "листок", "тяпка", "пятка", "столик"},
			output: map[string][]string{
				"слиток": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
		},
		{
			input: []string{"яблоко", "окОлбя", "бАНан", "НАНАБ", "вИноград", "грУШа", "апельсин", "иновдагр", "градвино", "боб", "обб"},
			output: map[string][]string{
				"банан":    {"банан", "нанаб"},
				"виноград": {"виноград", "градвино", "иновдагр"},
				"боб":      {"боб", "обб"},
				"яблоко":   {"околбя", "яблоко"},
			},
		},
		{
			input:  []string{"один"},
			output: map[string][]string{},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := searchAnagramsInDict(test.input)
			if !reflect.DeepEqual(result, test.output) {
				t.Errorf("Expected %v, but got %v", test.output, result)
			}
		})
	}
}
