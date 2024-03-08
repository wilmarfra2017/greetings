package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Wilmar"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Wilmar")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Juan") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHellowEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`hello("") = %q, %v, quiere "", error`, msg, err)
	} 
}
