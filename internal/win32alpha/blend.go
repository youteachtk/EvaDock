package win32alpha

// BlendFunction matches the Win32 BLENDFUNCTION structure (four BYTE fields).
// UpdateLayeredWindow takes a pointer to this structure, not its packed value.
type BlendFunction struct {
	BlendOp             uint8
	BlendFlags          uint8
	SourceConstantAlpha uint8
	AlphaFormat         uint8
}

// PremultipliedAlphaBlend enables per-pixel alpha with full global opacity.
func PremultipliedAlphaBlend() BlendFunction {
	return BlendFunction{SourceConstantAlpha: 255, AlphaFormat: 1}
}
