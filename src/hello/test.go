package main

import "testing"

func TestDoSomething(t *testing.T) {
    str = "Hello World!"
    actual = AddYouSaid(str)
    expected = "You said: " + str
    if actual != expected {
        t.Errorf("Incorrect string returned. got: %s, expected: %s", actual, expected)
    }
}
