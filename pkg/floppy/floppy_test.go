package floppy

import (
	"reflect"
	"testing"
)

func TestWriteRead(t *testing.T) {
	floppy := NewFloppy(80, 18, 512)
	writeData := []byte("Hello")
	floppy.Write(MagneticHead0, 0, 2, writeData)
	readData := floppy.Read(MagneticHead0, 0, 2)
	if reflect.DeepEqual(readData, writeData) {
		t.Errorf("Expect: %v, got %v", writeData, readData)
	}
}
