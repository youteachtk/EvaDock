package idleframe

import (
	"bytes"
	"testing"
)

func TestTailTwitchOnlyWarpsTipOfHorizontalPose(t *testing.T) {
	w, h := 100, 100
	src := make([]byte, w*h*4)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			i := (y*w + x) * 4
			src[i], src[i+3] = byte(x*2), 255
		}
	}
	original := bytes.Clone(src)
	warped := TailTwitch(src, w, h, 0.125)
	if len(warped) != len(src) {
		t.Fatal("warping changed frame dimensions")
	}
	if !bytes.Equal(src, original) {
		t.Fatal("tail movement must never modify an approved source frame")
	}
	if !bytes.Equal(warped[:4], src[:4]) {
		t.Fatal("top-left outside the tail region was changed")
	}
	center := (82*w + 17) * 4
	if warped[center] == src[center] {
		t.Fatal("tail tip was not animated")
	}
}

func TestTailTwitchStartsAndEndsOnOriginalFrame(t *testing.T) {
	src := make([]byte, 40*40*4)
	for i := range src {
		src[i] = byte(i % 251)
	}
	if !bytes.Equal(TailTwitch(src, 40, 40, 0), src) || !bytes.Equal(TailTwitch(src, 40, 40, 1), src) {
		t.Fatal("tip movement must return exactly to approved pose")
	}
}
