package posecatalog

import (
	"bytes"
	"fmt"
	"image"
	"image/draw"
	"image/png"
)

type Pose struct {
	Name  string
	File  string
	Image *image.NRGBA
}

func Load(asset func(string) ([]byte, error)) ([]Pose, error) {
	files := []struct{ name, path string }{
		{"Sentada", "independientes/sentada.png"},
		{"Horizontal", "independientes/pose_horizontal_aprobada.png"},
	}
	poses := make([]Pose, 0, len(files))
	for _, entry := range files {
		data, err := asset(entry.path)
		if err != nil {
			return nil, fmt.Errorf("cargar %s: %w", entry.path, err)
		}
		img, err := png.Decode(bytes.NewReader(data))
		if err != nil {
			return nil, fmt.Errorf("decodificar %s: %w", entry.path, err)
		}
		bounds := img.Bounds()
		if bounds.Empty() {
			return nil, fmt.Errorf("imagen vacia: %s", entry.path)
		}
		n := image.NewNRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
		draw.Draw(n, n.Bounds(), img, bounds.Min, draw.Src)
		poses = append(poses, Pose{Name: entry.name, File: entry.path, Image: n})
	}
	return poses, nil
}
