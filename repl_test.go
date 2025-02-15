package main

import (
    "testing"
)

func TestCleanInput(t *testing.T) {

    cases := []struct {
        input   string
        expected []string

    }{
        {
            input: " hello  world ",
            expected: []string{"hello", "world"},
        },
        // add more cases here...
    }

    for _, c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("Actual and Expected Slices do not match")
            t.Errorf("Actual: %v\nExpected: %v", actual, c.expected)
            t.Fail()
        }
        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord {
                t.Errorf("Word and ExpectedWord do not mach")
                t.Errorf("Word: %v\nExpectedWord: %v", word, expectedWord)
                t.Fail()
            }
        }
    }
}
