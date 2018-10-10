package modtest

import (
	"testing"
)

func TestGetVersion(t *testing.T) {
	actual := GetVersion()
	expected := "v0.0.1"

	if actual != expected {
		t.Errorf("Got %s, Want %s", actual, expected)
	}
}
