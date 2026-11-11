package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
	{
		input: "   hEllo World ",
		expected: []string{"hello", "world"},
	},
	{
		input: "i  GoT yA",
		expected: []string{"i", "got", "ya"},
	},
	{
		input: "AboBa",
		expected: []string{"aboba"},
	},
}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("length error")
			continue
		}
		for i := range actual {
		if actual[i] != c.expected[i]{
			t.Errorf("word %v doesn't match with %v", actual[i], c.expected[i])
		} 
	}
	}
	
}