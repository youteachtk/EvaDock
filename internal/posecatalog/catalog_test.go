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

func TestLoadIncludesAllApprovedPosesWithoutChangingTransparency(t *testing.T) {
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
	if len(poses) != 3 {
		t.Fatalf("expected sitting, curled and horizontal poses, got %d", len(poses))
	}
	if poses[0].File != "independientes/sentada.png" || poses[1].File != "independientes/pose_acurrucada_aprobada.png" || poses[2].File != "independientes/pose_horizontal_aprobada.png" {
		t.Fatalf("wrong approved image files: %#v", requested)
	}
	if len(requested) != 3 || requested[1] != poses[1].File || requested[2] != poses[2].File {
		t.Fatalf("approved images were not loaded: %#v", requested)
	}
	got := poses[2].Image.NRGBAAt(0, 0)
	if got != (color.NRGBA{R: 120, G: 55, B: 20, A: 75}) {
		t.Fatalf("transparency or colors changed: %#v", got)
	}
}

func TestApprovedCurledPNGMatchesOriginalBytes(t *testing.T) {
	data, err := os.ReadFile("../../assets/independientes/pose_acurrucada_aprobada.png")
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(data)
	const want = "083264e6997024c5c9837a17ea08af102f1503116dd5cb181abb24881f5e25b4"
	if got := hex.EncodeToString(sum[:]); got != want {
		t.Fatalf("approved curled PNG bytes changed: %s", got)
	}
	decoded, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if decoded.Bounds().Dx() != 1677 || decoded.Bounds().Dy() != 938 {
		t.Fatalf("unexpected curled image dimensions: %v", decoded.Bounds())
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
