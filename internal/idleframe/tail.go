package idleframe

import "math"

// TailTwitch applies a brief, tiny, smooth warp only to the visible tail-tip
// region of Eva's approved horizontal sprite. The source PNG is never edited.
// The original head, paws and torso are outside this deformation field.
// progress runs from 0 to 1, returning exactly to the original in both ends.
func TailTwitch(src []byte, width, height int, progress float64) []byte {
	if width <= 0 || height <= 0 || len(src) != width*height*4 {
		return nil
	}
	dst := make([]byte, len(src))
	copy(dst, src)
	if progress <= 0 || progress >= 1 {
		return dst
	}
	envelope := math.Pow(math.Sin(math.Pi*progress), 2)
	amount := 0.04 * float64(width) * envelope * math.Sin(4*math.Pi*progress)
	if math.Abs(amount) < 0.01 {
		return dst
	}
	cx, cy := 0.17*float64(width), 0.82*float64(height)
	rx, ry := 0.16*float64(width), 0.22*float64(height)
	for y := 0; y < height; y++ {
		dy := (float64(y) - cy) / ry
		if math.Abs(dy) >= 1 {
			continue
		}
		for x := 0; x < width; x++ {
			dx := (float64(x) - cx) / rx
			radius2 := dx*dx + dy*dy
			if radius2 >= 1 {
				continue
			}
			falloff := 1 - radius2
			sx := float64(x) - amount*falloff*falloff
			if sx < 0 || sx >= float64(width-1) {
				continue
			}
			x0 := int(math.Floor(sx))
			fx := sx - float64(x0)
			pos := (y*width + x) * 4
			from := (y*width + x0) * 4
			for c := 0; c < 4; c++ {
				dst[pos+c] = byte(math.Round((1-fx)*float64(src[from+c]) + fx*float64(src[from+c+4])))
			}
		}
	}
	return dst
}
