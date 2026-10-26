package posecatalog

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestLoadIncludesApprovedHorizontalPoseWithoutChangingTransparency(t *testing.T) {
	img := image.NewNRGBA(image.Rect(0, 0, 2, 1))
	img.SetNRGBA(0, 0, color.NRGBA{R: 120, G: 55, B: 20, A: 75})
	img.SetNRGBA(1, 0, color.NRGBA{R: 30, G: 40, B: 50, A: 255})
	var pngBytes bytes.Buffer
	if err := png.Encode(&pngBytes, img); err != nil {
		t.Fatal(err)
	}

	var requested []string
	poses, err := Load(func(path string) ([]byte, error) {
		requested = append(requested, path)
		return pngBytes.Bytes(), nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(poses) != 2 {
		t.Fatalf("expected sitting and approved horizontal poses, got %d", len(poses))
	}
	if poses[0].File != "independientes/sentada.png" || poses[1].File != "independientes/pose_horizontal_aprobada.png" {
		t.Fatalf("wrong image files: %q, %q", poses[0].File, poses[1].File)
	}
	if len(requested) != 2 || requested[1] != poses[1].File {
		t.Fatalf("horizontal image was not loaded from approved file: %#v", requested)
	}
	got := poses[1].Image.NRGBAAt(0, 0)
	if got != (color.NRGBA{R: 120, G: 55, B: 20, A: 75}) {
		t.Fatalf("transparency or colors changed: %#v", got)
	}
}

func TestApprovedPNGMatchesOriginalBytes(t *testing.T) {
	data, err := os.ReadFile("../../assets/independientes/pose_horizontal_aprobada.png")
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(data)
	got := hex.EncodeToString(sum[:])
	const want = "72391725920009d6234574fbbbcde9a37fc38fb79b4cbd6d03af51a5a503d044"
	if got != want {
		t.Fatalf("approved PNG bytes changed: %s", got)
	}
	decoded, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if decoded.Bounds().Dx() != 1536 || decoded.Bounds().Dy() != 1024 {
		t.Fatalf("unexpected approved image dimensions: %v", decoded.Bounds())
	}
}
