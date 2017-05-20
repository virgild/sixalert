package main

import (
	"testing"
	"time"
)

func TestChecksum(t *testing.T) {
	t1, _ := time.Parse("01/02/2006 15:04PM", "01/01/2020 10:00AM")
	alert := NewTTCAlert("This is a test content", "None", t1)

	checksum := alert.Checksum()
	correctsum := "0616d9b1a0d672d159a96f2f39a87fc0"

	if checksum != correctsum {
		t.Errorf("Checksum expected \"%v\" but returned \"%v\"", correctsum, checksum)
	}
}
