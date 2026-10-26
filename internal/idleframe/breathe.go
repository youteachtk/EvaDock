// Package idleframe renders a very slight, temporary idle expansion.
// The original approved sprite bytes are never modified.
package idleframe

import "math"

// Pulse expands a pre-multiplied BGRA canvas by at most 1.2% around the
// center, with the feet/bottom anchored. It keeps transparent edges intact.
func Pulse(src []byte, width, height int, expansion float64) []byte {
	if width <= 0 || height <= 0 || len(src) != width*height*4 {
		return nil
	}
	dst := make([]byte, len(src))
	if expansion <= 0 {
		copy(dst, src)
		return dst
	}
	if expansion > 0.012 {
		expansion = 0.012
	}
	scale := 1 + expansion
	centerX := float64(width-1) / 2
	bottom := float64(height - 1)
	for y := 0; y < height; y++ {
		sy := bottom + (float64(y)-bottom)/scale
		for x := 0; x < width; x++ {
			sx := centerX + (float64(x)-centerX)/scale
			x0, y0 := int(math.Floor(sx)), int(math.Floor(sy))
			fx, fy := sx-float64(x0), sy-float64(y0)
			at := (y*width + x) * 4
			var channels [4]float64
			for yy := 0; yy <= 1; yy++ {
				iy := y0 + yy
				if iy < 0 || iy >= height {
					continue
				}
				wy := 1 - fy
				if yy == 1 {
					wy = fy
				}
				for xx := 0; xx <= 1; xx++ {
					ix := x0 + xx
					if ix < 0 || ix >= width {
						continue
					}
					wx := 1 - fx
					if xx == 1 {
						wx = fx
					}
					weight := wx * wy
					from := (iy*width + ix) * 4
					for channel := 0; channel < 4; channel++ {
						channels[channel] += float64(src[from+channel]) * weight
					}
				}
			}
			for channel := 0; channel < 4; channel++ {
				dst[at+channel] = byte(math.Round(math.Min(255, channels[channel])))
			}
		}
	}
	return dst
}
