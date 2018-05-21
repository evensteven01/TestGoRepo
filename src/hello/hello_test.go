package main

import "testing"

func TestDoSomething(t *testing.T) {
    var str string = "Hello World!"
    var actual string = AddYouSaid(str)
    var expected string = "You said: " + str
    if actual != expected {
        t.Errorf("Incorrect string returned. got: %s, expected: %s", actual, expected)
    }
}
