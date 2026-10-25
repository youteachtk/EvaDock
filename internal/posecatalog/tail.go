package posecatalog

import (
	"image"
	"math"
	"time"
)

// TailShouldTwitch gives the reclined, closed-eye pose two brief tail-tip
// movements roughly every 25 seconds. All other poses stay completely still.
func TailShouldTwitch(elapsed time.Duration) bool {
	if elapsed < 0 {
		return false
	}
	phase := elapsed % (25 * time.Second)
	return (phase >= 6100*time.Millisecond && phase < 6450*time.Millisecond) ||
		(phase >= 6800*time.Millisecond && phase < 7150*time.Millisecond)
}

// TailTwitch creates a second render frame by subtly bending the lower-left
// tail-tip area. It leaves the source unchanged and does not edit stored PNGs.
// Coordinates are normalized to Eva's authorized independent reclined pose.
func TailTwitch(src *image.NRGBA, amplitude int) *image.NRGBA {
	if src == nil {
		return nil
	}
	dst := image.NewNRGBA(src.Bounds())
	copy(dst.Pix, src.Pix)
	if amplitude == 0 {
		return dst
	}
	w, h := src.Bounds().Dx(), src.Bounds().Dy()
	// Small oval traced around the dark terminal tail segment in the real
	// reclined asset. The former broad rectangle included Eva's flank.
	cx, cy := float64(w)*0.12, float64(h)*0.775
	rx, ry := float64(w)*0.075, float64(h)*0.042
	if rx < 1 || ry < 1 {
		return dst
	}
	for y := int(cy - ry); y < int(cy+ry); y++ {
		for x := int(cx - rx); x < int(cx+rx); x++ {
			nx, ny := (float64(x)-cx)/rx, (float64(y)-cy)/ry
			distance := nx*nx + ny*ny
			if distance >= 1 {
				continue
			}
			fade := 1 - distance
			delta := int(math.Round(float64(amplitude) * fade * fade))
			sourceY := y + delta
			if sourceY < 0 || sourceY >= h {
				continue
			}
			from := src.PixOffset(x+src.Bounds().Min.X, sourceY+src.Bounds().Min.Y)
			to := dst.PixOffset(x+dst.Bounds().Min.X, y+dst.Bounds().Min.Y)
			copy(dst.Pix[to:to+4], src.Pix[from:from+4])
		}
	}
	return dst
}
