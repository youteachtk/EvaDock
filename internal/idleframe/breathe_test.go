package idleframe

import (
	"bytes"
	"testing"
)

func TestPulseLeavesApprovedSourceIntact(t *testing.T) {
	w, h := 20, 20
	src := make([]byte, w*h*4)
	for y := 4; y < 16; y++ {
		for x := 4; x < 16; x++ {
			pos := (y*w + x) * 4
			src[pos], src[pos+1], src[pos+2], src[pos+3] = 40, 80, 120, 255
		}
	}
	before := bytes.Clone(src)
	still := Pulse(src, w, h, 0)
	if !bytes.Equal(still, src) {
		t.Fatal("zero pulse must render exactly the original approved pixels")
	}
	expanded := Pulse(src, w, h, 0.01)
	if len(expanded) != len(src) {
		t.Fatalf("changed bitmap size: %d", len(expanded))
	}
	if !bytes.Equal(src, before) {
		t.Fatal("breathing animation modified the approved image in memory")
	}
	if expanded[(0*w+0)*4+3] != 0 {
		t.Fatal("transparent background became opaque")
	}
}

func TestPulseDoesNotFadeFullyOpaquePixels(t *testing.T) {
	w, h := 31, 31
	src := make([]byte, w*h*4)
	for i := 0; i < len(src); i += 4 {
		src[i], src[i+1], src[i+2], src[i+3] = 20, 80, 170, 255
	}
	frame := Pulse(src, w, h, 0.009)
	center := (15*w + 15) * 4
	if frame[center+3] != 255 || frame[center] != 20 || frame[center+1] != 80 || frame[center+2] != 170 {
		t.Fatalf("fully opaque internal pixel faded: %v", frame[center:center+4])
	}
}
