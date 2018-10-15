package main

import(
    "testing"
    "regexp"
)

var validGoVersion = regexp.MustCompile(`^go[0-9]\.[0-9].+$`)

func TestMyGoVersion(t *testing.T) {
	if validGoVersion.MatchString(goversion()) != true  {
		t.Fatal("Test fail")
	}
}
