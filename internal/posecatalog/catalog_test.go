package posecatalog

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func diskAsset(path string) ([]byte, error) {
	return os.ReadFile(filepath.Join("..", "..", "assets", filepath.FromSlash(path)))
}

func TestThreeIndependentTransparentPoses(t *testing.T) {
	got, err := Load(diskAsset)
	if err != nil {
		t.Fatalf("cargar recursos reales: %v", err)
	}
	wantNames := []string{"Sentada", "Recostada", "Dormida"}
	wantPaths := []string{
		"independientes/sentada.png", "independientes/recostada.png", "independientes/dormida.png",
	}
	if len(got) != len(wantNames) {
		t.Fatalf("se esperaban 3 poses, se cargaron %d", len(got))
	}
	for i, p := range got {
		if p.Name != wantNames[i] || p.File != wantPaths[i] {
			t.Errorf("pose %d = %q, %q", i, p.Name, p.File)
		}
		if p.Image == nil || p.Image.Bounds().Dx() < 200 || p.Image.Bounds().Dy() < 200 {
			t.Errorf("pose %d no tiene imagen válida", i)
		}
	}
}

func TestLoadRejectsMissingAsset(t *testing.T) {
	_, err := Load(func(string) ([]byte, error) { return nil, os.ErrNotExist })
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("se esperaba error identificable por recurso ausente, resultado: %v", err)
	}
}

func TestLoadRejectsCorruptPNG(t *testing.T) {
	_, err := Load(func(string) ([]byte, error) { return []byte("esto no es png"), nil })
	if err == nil {
		t.Fatal("PNG dañado debe producir error")
	}
}

func TestLoadRejectsOpaqueRectangle(t *testing.T) {
	img := image.NewNRGBA(image.Rect(0, 0, 400, 400))
	for i := 0; i < len(img.Pix); i += 4 {
		img.Pix[i], img.Pix[i+1], img.Pix[i+2], img.Pix[i+3] = 255, 255, 255, 255
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		t.Fatal(err)
	}
	_, err := Load(func(string) ([]byte, error) { return buf.Bytes(), nil })
	if err == nil {
		t.Fatal("una imagen completamente opaca crearía un fondo rectangular")
	}
}

func TestAutomaticIndexUsesOnlyThreePoses(t *testing.T) {
	want := []int{0, 1, 2, 0, 1, 2, 0, 1}
	for tick, expected := range want {
		if got := AutomaticIndex(tick); got != expected {
			t.Fatalf("tick %d: se esperaba pose %d, se obtuvo %d", tick, expected, got)
		}
	}
}

func TestTailTwitchOnlyChangesTailTipRegion(t *testing.T) {
	img := image.NewNRGBA(image.Rect(0, 0, 200, 200))
	for y := 150; y < 175; y++ {
		for x := 12; x < 40; x++ {
			img.SetNRGBA(x, y, color.NRGBA{R: 80, G: 40, B: 20, A: 255})
		}
	}
	original := append([]byte(nil), img.Pix...)
	out := TailTwitch(img, 12)
	if !bytes.Equal(original, img.Pix) {
		t.Fatal("la imagen original debe permanecer intacta")
	}
	if bytes.Equal(original, out.Pix) {
		t.Fatal("una oscilación no puede producir el mismo fotograma")
	}
	for y := 0; y < 200; y++ {
		for x := 0; x < 200; x++ {
			if x < 7 || x >= 42 || y < 140 || y >= 175 {
				if img.NRGBAAt(x, y) != out.NRGBAAt(x, y) {
					t.Fatalf("la animación alteró un píxel fuera de la cola: (%d,%d)", x, y)
				}
			}
		}
	}
}

func TestTailTwitchIsOccasional(t *testing.T) {
	for _, tm := range []time.Duration{0, 5 * time.Second, 8 * time.Second, 17 * time.Second} {
		if TailShouldTwitch(tm) {
			t.Errorf("no debe moverse la cola a los %s", tm)
		}
	}
	for _, tm := range []time.Duration{6200 * time.Millisecond, 6900 * time.Millisecond, 31200 * time.Millisecond} {
		if !TailShouldTwitch(tm) {
			t.Errorf("debe existir una oscilación corta a los %s", tm)
		}
	}
}

func TestRealRecumbentBodyStaysPixelIdentical(t *testing.T) {
	poses, err := Load(diskAsset)
	if err != nil {
		t.Fatal(err)
	}
	base := poses[1].Image
	wag := TailTwitch(base, 24)
	// These are hand-inspected coordinates in the 1448x1086 reclined image:
	// flank, shoulder, cheek, white front paw, and eye/ear respectively.
	bodyPoints := []image.Point{{695, 760}, {780, 745}, {970, 645}, {1200, 920}, {1000, 320}}
	for _, pt := range bodyPoints {
		if base.NRGBAAt(pt.X, pt.Y) != wag.NRGBAAt(pt.X, pt.Y) {
			t.Fatalf("cuerpo alterado en (%d,%d): %v -> %v", pt.X, pt.Y,
				base.NRGBAAt(pt.X, pt.Y), wag.NRGBAAt(pt.X, pt.Y))
		}
	}
	changed := 0
	for y := 0; y < base.Bounds().Dy(); y++ {
		for x := 0; x < base.Bounds().Dx(); x++ {
			if base.NRGBAAt(x, y) == wag.NRGBAAt(x, y) {
				continue
			}
			// The visible tail tip curls along Eva's far left, not under her belly.
			if x < 60 || x > 295 || y < 790 || y > 895 {
				t.Fatalf("píxel ajeno a la punta de cola alterado: (%d,%d)", x, y)
			}
			changed++
		}
	}
	if changed < 10 {
		t.Fatalf("la punta de cola no cambió: sólo %d píxeles", changed)
	}
}
