package posecatalog

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
)

// Pose is a complete, individually authored image rather than a crop of a pose sheet.
type Pose struct {
	Name  string
	File  string
	Image *image.NRGBA
}

var definitions = []struct{ name, file string }{
	{"Sentada", "independientes/sentada.png"},
	{"Recostada", "independientes/recostada.png"},
	{"Dormida", "independientes/dormida.png"},
}

// Load always validates the whole catalog before returning any pose.
func Load(readAsset func(string) ([]byte, error)) ([]Pose, error) {
	poses := make([]Pose, 0, len(definitions))
	for _, d := range definitions {
		data, err := readAsset(d.file)
		if err != nil {
			return nil, fmt.Errorf("recurso %s: %w", d.file, err)
		}
		decoded, err := png.Decode(bytes.NewReader(data))
		if err != nil {
			return nil, fmt.Errorf("PNG %s: %w", d.file, err)
		}
		b := decoded.Bounds()
		if b.Dx() < 200 || b.Dy() < 200 {
			return nil, fmt.Errorf("imagen %s demasiado pequeña", d.file)
		}
		img := image.NewNRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		draw.Draw(img, img.Bounds(), decoded, b.Min, draw.Src)
		transparent, visible := false, false
		for i := 3; i < len(img.Pix); i += 4 {
			if img.Pix[i] < 16 {
				transparent = true
			}
			if img.Pix[i] > 128 {
				visible = true
			}
			if transparent && visible {
				break
			}
		}
		if !transparent || !visible {
			return nil, fmt.Errorf("imagen %s sin transparencia útil", d.file)
		}
		poses = append(poses, Pose{Name: d.name, File: d.file, Image: img})
	}
	return poses, nil
}

// AutomaticIndex chooses only the three authorized, non-standing poses.
func AutomaticIndex(tick int) int {
	n := len(definitions)
	return ((tick % n) + n) % n
}
