package oslinfra

import (
	"testing"
)

func TestSSHRefresh(t *testing.T) {
	actual := SSHRefresh()
	expected := "helloworld"
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
