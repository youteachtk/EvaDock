package win32alpha

import (
	"bytes"
	"testing"
	"unsafe"
)

// UpdateLayeredWindow expects a pointer to this four-byte Win32 ABI struct.
// Passing the numeric value 0x01ff0000 as a pointer makes the call fail.
func TestPremultipliedAlphaBlendIsValidWin32Argument(t *testing.T) {
	blend := PremultipliedAlphaBlend()
	if got := unsafe.Sizeof(blend); got != 4 {
		t.Fatalf("BLENDFUNCTION must use four Win32 BYTEs, got %d", got)
	}
	got := unsafe.Slice((*byte)(unsafe.Pointer(&blend)), 4)
	want := []byte{0, 0, 255, 1}
	if !bytes.Equal(got, want) {
		t.Fatalf("invalid BLENDFUNCTION bytes: %v; want %v", got, want)
	}
}
