package main

import (
    "testing"
    "time"
)

func TestCleanInput(t *testing.T) {

    cases := []struct {
        input   string
        expected []string

    }{
        {
            input:    "  ",
            expected: []string{},
        },
        {
            input:    "  hello  ",
            expected: []string{"hello"},
        },
        {
            input:    "  hello  world  ",
            expected: []string{"hello", "world"},
        },
        {
            input:    "  HellO  World  ",
            expected: []string{"hello", "world"},
        },
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

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
