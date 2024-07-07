package main

import (
	"bytes"
	"os/exec"
	"testing"
)

func Test_CalculateSubnets(t *testing.T) {
	const cmd = "go"
	args := []string{
		"run",
		"main.go",
		"-cidr",
		"10.1.1.0/28",
		"-size",
		"29",
	}
	expected := []byte{
		27, 91, 48, 109, 49, 48, 46, 49, 46, 49, 46, 48, 47, 50,
		57, 10, 49, 48, 46, 49, 46, 49, 46, 56, 47, 50, 57, 10,
	}
	actual, err := exec.Command(cmd, args...).Output()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(actual, expected) {
		t.Fatalf("mismatch\n"+
			"  actual: %s (%v)\n"+
			"expected: %s (%v)",
			string(actual), actual, string(expected), expected)
	}

}
