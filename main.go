// Eva Dock — Aplicacion nativa de Windows, con poses independientes.
// Controles HWND BUTTON (WM_COMMAND/BN_CLICKED), vista previa y mascota flotante.
// Recurso independiente de Eva; conserva los PNG anteriores solo como respaldo local.
// No usa red, cuentas, servicios de terceros ni modelos IA.
package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"time"
	"unsafe"

	"github.com/youteachtk/EvaDock/internal/posecatalog"
	"github.com/youteachtk/EvaDock/internal/win32alpha"
)

const (
	wmCreate           = 0x0001
	wmDestroy          = 0x0002
	wmSize             = 0x0005
	wmClose            = 0x0010
	wmPaint            = 0x000F
	wmCommand          = 0x0111
	wmTimer            = 0x0113
	wmLDown            = 0x0201
	wmLUp              = 0x0202
	wmMove             = 0x0200
	wmLDouble          = 0x0203
	wmRUp              = 0x0205
	wmWheel            = 0x020A
	wsOverlappedWindow = 0x00CF0000
	wsPopup            = 0x80000000
	wsChild            = 0x40000000
	wsVisible          = 0x10000000
	wsExLayered        = 0x00080000
	wsExTopmost        = 0x00000008
	wsExToolwindow     = 0x00000080
	ssLeft             = 0x00000000
	csDblClicks        = 0x0008
	swShow             = 5
	swHide             = 0
	swpNoSize          = 0x0001
	swpNoZOrder        = 0x0004
	swpNoActivate      = 0x0010
	dibRGBColors       = 0
	ulwAlpha           = 2
	srccopy            = 0x00CC0020
	idShow             = 201
	idSit              = 202
	idLie              = 203
	idSleep            = 204
	idAuto             = 205
	idSmaller          = 206
	idBigger           = 207
	idBring            = 208
	idExit             = 209
	panelW             = 745
	panelH             = 490
	previewW           = 322
	previewH           = 300
)

type point struct{ X, Y int32 }
type siz struct{ CX, CY int32 }
type rect struct{ Left, Top, Right, Bottom int32 }
type classEx struct {
	CbSize     uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   uintptr
	Icon       uintptr
	Cursor     uintptr
	Background uintptr
	MenuName   uintptr
	ClassName  uintptr
	IconSmall  uintptr
}
type msgT struct {
	Hwnd     uintptr
	Message  uint32
	_        uint32
	WParam   uintptr
	LParam   uintptr
	Time     uint32
	Pt       point
	LPrivate uint32
}
type paintStruct struct {
	Hdc       uintptr
	Erase     int32
	PaintRect rect
	Restore   int32
	IncUpdate int32
	Reserved  [32]byte
}
type bitmapHeader struct {
	Size          uint32
	Width         int32
	Height        int32
	Planes        uint16
	BitCount      uint16
	Compression   uint32
	SizeImage     uint32
	XPelsPerMeter int32
	YPelsPerMeter int32
	ClrUsed       uint32
	ClrImportant  uint32
}
type bitmapInfo struct {
	H      bitmapHeader
	Colors [1]uint32
}

var (
	u32                                         = syscall.NewLazyDLL("user32.dll")
	gdi                                         = syscall.NewLazyDLL("gdi32.dll")
	kern                                        = syscall.NewLazyDLL("kernel32.dll")
	pReg                                        = u32.NewProc("RegisterClassExW")
	pCreate                                     = u32.NewProc("CreateWindowExW")
	pDef                                        = u32.NewProc("DefWindowProcW")
	pShow                                       = u32.NewProc("ShowWindow")
	pEnable                                     = u32.NewProc("EnableWindow")
	pDestroy                                    = u32.NewProc("DestroyWindow")
	pGetMsg                                     = u32.NewProc("GetMessageW")
	pTranslate                                  = u32.NewProc("TranslateMessage")
	pDispatch                                   = u32.NewProc("DispatchMessageW")
	pQuit                                       = u32.NewProc("PostQuitMessage")
	pSetText                                    = u32.NewProc("SetWindowTextW")
	pSetTimer                                   = u32.NewProc("SetTimer")
	pGetSystem                                  = u32.NewProc("GetSystemMetrics")
	pSetPos                                     = u32.NewProc("SetWindowPos")
	pGetCursor                                  = u32.NewProc("GetCursorPos")
	pCapture                                    = u32.NewProc("SetCapture")
	pReleaseCapture                             = u32.NewProc("ReleaseCapture")
	pBeginPaint                                 = u32.NewProc("BeginPaint")
	pEndPaint                                   = u32.NewProc("EndPaint")
	pInvalidate                                 = u32.NewProc("InvalidateRect")
	pGetDC                                      = u32.NewProc("GetDC")
	pReleaseDC                                  = u32.NewProc("ReleaseDC")
	pUpdate                                     = u32.NewProc("UpdateLayeredWindow")
	pCreateMem                                  = gdi.NewProc("CreateCompatibleDC")
	pDeleteDC                                   = gdi.NewProc("DeleteDC")
	pCreateDIB                                  = gdi.NewProc("CreateDIBSection")
	pSelect                                     = gdi.NewProc("SelectObject")
	pDelete                                     = gdi.NewProc("DeleteObject")
	pStretch                                    = gdi.NewProc("StretchDIBits")
	pModule                                     = kern.NewProc("GetModuleHandleW")
	pMessageBox                                 = u32.NewProc("MessageBoxW")
	pPostMessage                                = u32.NewProc("PostMessageW")
	pIsWindow                                   = u32.NewProc("IsWindow")
	pForeground                                 = u32.NewProc("SetForegroundWindow")
	poseList                                    []posecatalog.Pose
	active                                      int
	recostadaTailUp                             *image.NRGBA
	tailRaised                                  bool
	poseSince                                   = time.Now()
	panel, pet, statusH, showButton, autoButton uintptr
	petX, petY                                  int32
	petW                                        = 340
	petH                                        = 310
	visible                                     = true
	automatic                                   = false
	petFallback                                 = false
	dragging                                    bool
	downPt                                      point
	downX, downY                                int32
	dragged                                     bool
	actionCount                                 int
	lastAuto                                    = time.Now()
	lastWalk                                    = time.Now()
	logFile                                     string
	instance                                    uintptr
)

func u(s string) *uint16   { return syscall.StringToUTF16Ptr(s) }
func ptr(s string) uintptr { return uintptr(unsafe.Pointer(u(s))) }
func logErr(s string) {
	if logFile != "" {
		f, _ := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if f != nil {
			_, _ = f.WriteString(time.Now().Format("2006-01-02 15:04:05") + " " + s + "\n")
			_ = f.Close()
		}
	}
}
func n32(x int32) uintptr { return uintptr(int64(x)) }
func posByCursor() point  { p := point{}; pGetCursor.Call(uintptr(unsafe.Pointer(&p))); return p }
func decode() error {
	poses, err := posecatalog.Load(bundledAsset)
	if err != nil {
		return err
	}
	poseList = poses
	recostadaTailUp = posecatalog.TailTwitch(poseList[1].Image, 24)
	poseSince = time.Now()
	return nil
}

func poseImage() *image.NRGBA {
	if active == 1 && tailRaised && recostadaTailUp != nil {
		return recostadaTailUp
	}
	return poseList[active].Image
}

func class(name string, proc uintptr, style uint32) error {
	c := classEx{CbSize: uint32(unsafe.Sizeof(classEx{})), Style: style, WndProc: proc, Instance: instance, ClassName: ptr(name)}
	r, _, err := pReg.Call(uintptr(unsafe.Pointer(&c)))
	if r == 0 {
		return fmt.Errorf("RegisterClassEx %s: %v", name, err)
	}
	return nil
}
func child(typ, caption string, x, y, w, h int, id int) uintptr {
	hWnd, _, err := pCreate.Call(0, ptr(typ), ptr(caption), wsChild|wsVisible, uintptr(x), uintptr(y), uintptr(w), uintptr(h), panel, uintptr(id), instance, 0)
	if hWnd == 0 {
		logErr(fmt.Sprintf("No se creo boton %d: %v", id, err))
	}
	return hWnd
}
func setupPanel() {
	child("STATIC", "EVA DOCK v"+version+"  |  Mascota de escritorio", 16, 15, 700, 30, 0)
	child("STATIC", "Eva debe verse en una ventana flotante fuera de este panel.", 16, 45, 700, 24, 0)
	showButton = child("BUTTON", "Ocultar a Eva", 16, 80, 171, 34, idShow)
	child("BUTTON", "Traer Eva al frente", 196, 80, 171, 34, idBring)
	child("BUTTON", "Sentada", 16, 122, 171, 34, idSit)
	child("BUTTON", "Recostada", 196, 122, 171, 34, idLie)
	child("BUTTON", "Dormida", 16, 164, 171, 34, idSleep)
	autoButton = child("BUTTON", "Rutina automatica: NO", 196, 164, 171, 34, idAuto)
	child("BUTTON", "Tamano -", 16, 206, 171, 34, idSmaller)
	child("BUTTON", "Tamano +", 196, 206, 171, 34, idBigger)
	child("STATIC", "De pie, caminando y salto: pendientes de nuevas fotos.", 16, 261, 365, 47, 0)
	child("BUTTON", "Cerrar Eva Dock", 16, 326, 351, 38, idExit)
	child("STATIC", "VISTA PREVIA - imagen independiente", 405, 80, 325, 23, 0)
	statusH = child("STATIC", "Estado: iniciando...", 16, 385, 711, 55, 0)
	child("STATIC", "Arrastra a Eva para moverla. Las tres poses son independientes.", 16, 446, 707, 20, 0)
}

func setStatus(action string) {
	actionCount++
	msg := fmt.Sprintf("Accion %d: %s. Pose: %s. En escritorio: %v.", actionCount, action, poseList[active].Name, visible)
	if statusH != 0 {
		pSetText.Call(statusH, ptr(msg))
	}
	logErr(msg)
}
func repaintPreview() {
	if panel != 0 {
		r := rect{Left: 401, Top: 105, Right: 730, Bottom: 376}
		pInvalidate.Call(panel, uintptr(unsafe.Pointer(&r)), 1)
	}
}
func changePose(to int, action string) {
	if to < 0 || to >= len(poseList) {
		return
	}
	active = to
	tailRaised = false
	poseSince = time.Now()
	lastAuto = time.Now()
	drawPet()
	repaintPreview()
	setStatus(action)
}
func movePet(x, y int32) {
	petX = x
	petY = y
	if pet != 0 {
		pSetPos.Call(pet, ^uintptr(0), n32(x), n32(y), 0, 0, swpNoSize|swpNoActivate)
	}
}
func resizePet(delta int) {
	n := petW + delta
	if n < 190 {
		n = 190
	}
	if n > 650 {
		n = 650
	}
	petH = int(math.Round(float64(n) * 310.0 / 340.0))
	petW = n
	if pet != 0 {
		pSetPos.Call(pet, ^uintptr(0), n32(petX), n32(petY), uintptr(petW), uintptr(petH), swpNoActivate)
		drawPet()
	}
	setStatus(fmt.Sprintf("Tamano %d px", petW))
}
func showPet(v bool) {
	visible = v
	if pet != 0 {
		if visible {
			pShow.Call(pet, swShow)
			drawPet()
		} else {
			pShow.Call(pet, swHide)
		}
	}
	if showButton != 0 {
		if visible {
			pSetText.Call(showButton, ptr("Ocultar a Eva"))
		} else {
			pSetText.Call(showButton, ptr("Mostrar a Eva"))
		}
	}
	setStatus("Visibilidad cambiada")
}
func act(id int) {
	switch id {
	case idShow:
		showPet(!visible)
	case idBring:
		showPet(true)
		sw, _, _ := pGetSystem.Call(0)
		x := int32(845)
		if sw < 1250 {
			x = int32(sw) - int32(petW) - 30
		}
		if x < 25 {
			x = 25
		}
		movePet(x, 125)
		pSetPos.Call(pet, ^uintptr(0), n32(x), n32(125), uintptr(petW), uintptr(petH), swpNoActivate)
		drawPet()
		setStatus("Eva colocada a la vista, al frente")
	case idSit:
		changePose(0, "Eva sentada")
	case idLie:
		changePose(1, "Eva recostada")
	case idSleep:
		changePose(2, "Eva dormida")
	case idAuto:
		automatic = !automatic
		lastAuto = time.Now()
		if automatic {
			pSetText.Call(autoButton, ptr("Rutina automatica: SI"))
		} else {
			pSetText.Call(autoButton, ptr("Rutina automatica: NO"))
		}
		setStatus("Rutina automatica cambiada")
	case idSmaller:
		resizePet(-42)
	case idBigger:
		resizePet(42)
	case idExit:
		if panel != 0 {
			pDestroy.Call(panel)
		}
	}
}

// convert one original RGBA pose to composited BGRA bytes, alpha preserved when transparent.
func composeCanvas(width, height int, im *image.NRGBA, transparent bool) []byte {
	out := make([]byte, width*height*4)
	cr, cg, cb := float64(249), float64(244), float64(238)
	if !transparent {
		for j := 0; j < len(out); j += 4 {
			out[j] = byte(cb)
			out[j+1] = byte(cg)
			out[j+2] = byte(cr)
			out[j+3] = 255
		}
	}
	b := im.Bounds()
	s := math.Min(float64(width-16)/float64(b.Dx()), float64(height-16)/float64(b.Dy()))
	if transparent {
		s = math.Min(float64(width-10)/float64(b.Dx()), float64(height-10)/float64(b.Dy()))
	}
	tw := int(float64(b.Dx()) * s)
	th := int(float64(b.Dy()) * s)
	if tw <= 0 || th <= 0 {
		return out
	}
	ox := (width - tw) / 2
	oy := height - th - 7
	for y := 0; y < th; y++ {
		sy := (float64(y)+0.5)*float64(b.Dy())/float64(th) - 0.5
		y0 := int(math.Floor(sy))
		fy := sy - float64(y0)
		if y0 < 0 {
			y0 = 0
			fy = 0
		}
		if y0 >= b.Dy()-1 {
			y0 = b.Dy() - 1
			fy = 0
		}
		y1 := y0 + 1
		if y1 >= b.Dy() {
			y1 = y0
		}
		for x := 0; x < tw; x++ {
			sx := (float64(x)+0.5)*float64(b.Dx())/float64(tw) - 0.5
			x0 := int(math.Floor(sx))
			fx := sx - float64(x0)
			if x0 < 0 {
				x0 = 0
				fx = 0
			}
			if x0 >= b.Dx()-1 {
				x0 = b.Dx() - 1
				fx = 0
			}
			x1 := x0 + 1
			if x1 >= b.Dx() {
				x1 = x0
			}
			c0 := y0*im.Stride + x0*4
			c1 := y0*im.Stride + x1*4
			c2 := y1*im.Stride + x0*4
			c3 := y1*im.Stride + x1*4
			ix := ((y+oy)*width + x + ox) * 4
			if ix < 0 || ix+3 >= len(out) {
				continue
			}
			var rgba [4]float64
			for k := 0; k < 4; k++ {
				rgba[k] = (1-fy)*((1-fx)*float64(im.Pix[c0+k])+fx*float64(im.Pix[c1+k])) + fy*((1-fx)*float64(im.Pix[c2+k])+fx*float64(im.Pix[c3+k]))
			}
			a := rgba[3] / 255
			if transparent {
				out[ix] = byte(math.Min(255, rgba[2]*a))
				out[ix+1] = byte(math.Min(255, rgba[1]*a))
				out[ix+2] = byte(math.Min(255, rgba[0]*a))
				out[ix+3] = byte(rgba[3])
			} else {
				out[ix] = byte(math.Min(255, rgba[2]*a+cb*(1-a)))
				out[ix+1] = byte(math.Min(255, rgba[1]*a+cg*(1-a)))
				out[ix+2] = byte(math.Min(255, rgba[0]*a+cr*(1-a)))
				out[ix+3] = 255
			}
		}
	}
	return out
}
func drawPet() bool {
	if pet == 0 || len(poseList) == 0 {
		return false
	}
	if petFallback {
		pInvalidate.Call(pet, 0, 1)
		return true
	}
	buf := composeCanvas(petW, petH, poseImage(), true)
	info := bitmapInfo{H: bitmapHeader{Size: 40, Width: int32(petW), Height: -int32(petH), Planes: 1, BitCount: 32}}
	screenDC, _, _ := pGetDC.Call(0)
	if screenDC == 0 {
		logErr("GetDC error")
		return false
	}
	defer pReleaseDC.Call(0, screenDC)
	mem, _, _ := pCreateMem.Call(screenDC)
	if mem == 0 {
		logErr("CreateCompatibleDC error")
		return false
	}
	defer pDeleteDC.Call(mem)
	var bits uintptr
	bmp, _, err := pCreateDIB.Call(screenDC, uintptr(unsafe.Pointer(&info)), dibRGBColors, uintptr(unsafe.Pointer(&bits)), 0, 0)
	if bmp == 0 || bits == 0 {
		logErr("CreateDIBSection error: " + err.Error())
		return false
	}
	old, _, _ := pSelect.Call(mem, bmp)
	dest := unsafe.Slice((*byte)(unsafe.Pointer(bits)), len(buf))
	copy(dest, buf)
	p := point{petX, petY}
	size := siz{int32(petW), int32(petH)}
	orig := point{}
	// UpdateLayeredWindow expects a pointer to BLENDFUNCTION, not packed bytes.
	blend := win32alpha.PremultipliedAlphaBlend()
	ok, _, e := pUpdate.Call(pet, screenDC, uintptr(unsafe.Pointer(&p)), uintptr(unsafe.Pointer(&size)), mem, uintptr(unsafe.Pointer(&orig)), 0, uintptr(unsafe.Pointer(&blend)), ulwAlpha)
	runtime.KeepAlive(&blend)
	if ok == 0 {
		logErr("UpdateLayeredWindow error: " + e.Error())
	}
	pSelect.Call(mem, old)
	pDelete.Call(bmp)
	runtime.KeepAlive(buf)
	return ok != 0
}

func paintPetRegular(h uintptr) {
	ps := paintStruct{}
	hdc, _, _ := pBeginPaint.Call(h, uintptr(unsafe.Pointer(&ps)))
	if hdc == 0 {
		return
	}
	defer pEndPaint.Call(h, uintptr(unsafe.Pointer(&ps)))
	if len(poseList) == 0 {
		return
	}
	buf := composeCanvas(petW, petH, poseImage(), false)
	info := bitmapInfo{H: bitmapHeader{Size: 40, Width: int32(petW), Height: -int32(petH), Planes: 1, BitCount: 32}}
	pStretch.Call(hdc, 0, 0, uintptr(petW), uintptr(petH), 0, 0, uintptr(petW), uintptr(petH), uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&info)), 0, srccopy)
	runtime.KeepAlive(buf)
}
func paintPanel(h uintptr) {
	ps := paintStruct{}
	hdc, _, _ := pBeginPaint.Call(h, uintptr(unsafe.Pointer(&ps)))
	if hdc == 0 {
		return
	}
	defer pEndPaint.Call(h, uintptr(unsafe.Pointer(&ps)))
	if len(poseList) == 0 {
		return
	}
	buf := composeCanvas(previewW, previewH, poseImage(), false)
	info := bitmapInfo{H: bitmapHeader{Size: 40, Width: previewW, Height: -previewH, Planes: 1, BitCount: 32}}
	pStretch.Call(hdc, 405, 105, previewW, previewH, 0, 0, previewW, previewH, uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&info)), 0, srccopy)
	runtime.KeepAlive(buf)
}
func panelProc(h uintptr, m uint32, w, l uintptr) uintptr {
	switch m {
	case wmUpdateReady:
		if updateStaged != "" {
			startReplacement(updateStaged)
			if panel != 0 {
				pDestroy.Call(panel)
			}
		}
		return 0
	case wmCommand:
		id := int(w & 0xffff)
		// BN_CLICKED = HIWORD(wParam)==0. Do not lose messages to child controls.
		notify := (w >> 16) & 0xffff
		if notify == 0 && id >= idShow && id <= idExit {
			act(id)
			return 0
		}
	case wmPaint:
		paintPanel(h)
		return 0
	case wmTimer:
		if w == 1 {
			if active == 1 && visible {
				next := posecatalog.TailShouldTwitch(time.Since(poseSince))
				if next != tailRaised {
					tailRaised = next
					drawPet()
					repaintPreview()
				}
			}
			if automatic && len(poseList) > 0 && time.Since(lastAuto) > 12*time.Second {
				changePose(posecatalog.AutomaticIndex(active+1), "Rutina automatica")
			}
		}
		return 0
	case wmClose:
		pDestroy.Call(h)
		return 0
	case wmDestroy:
		if pet != 0 {
			pDestroy.Call(pet)
			pet = 0
		}
		pQuit.Call(0)
		return 0
	}
	r, _, _ := pDef.Call(h, uintptr(m), w, l)
	return r
}
func petProc(h uintptr, m uint32, w, l uintptr) uintptr {
	switch m {
	case wmPaint:
		if petFallback {
			paintPetRegular(h)
			return 0
		}
	case wmLDown:
		dragging = true
		dragged = false
		downPt = posByCursor()
		downX = petX
		downY = petY
		pCapture.Call(h)
		return 0
	case wmMove:
		if dragging {
			p := posByCursor()
			dx, dy := p.X-downPt.X, p.Y-downPt.Y
			if dx > 2 || dx < -2 || dy > 2 || dy < -2 {
				dragged = true
			}
			movePet(downX+dx, downY+dy)
		}
		return 0
	case wmLUp:
		if dragging {
			dragging = false
			pReleaseCapture.Call()
			if !dragged {
				setStatus("Eva seleccionada")
			}
		}
		return 0
	case wmLDouble:
		dragging = false
		pReleaseCapture.Call()
		setStatus("Eva seleccionada")
		return 0
	case wmWheel:
		if int16((w>>16)&0xffff) > 0 {
			resizePet(24)
		} else {
			resizePet(-24)
		}
		return 0
	case wmRUp:
		if panel != 0 {
			pShow.Call(panel, swShow)
			pForeground.Call(panel)
		}
		return 0
	case wmClose:
		showPet(false)
		return 0
	case wmDestroy:
		pet = 0
		return 0
	}
	r, _, _ := pDef.Call(h, uintptr(m), w, l)
	return r
}
func main1() error {
	runtime.LockOSThread()
	logFile = filepath.Join(configDir(), "eva.log")
	_ = os.MkdirAll(configDir(), 0700)
	logErr("Inicio Eva Dock " + version)
	if err := decode(); err != nil {
		return fmt.Errorf("Imagenes de Eva: %w", err)
	}
	instance, _, _ = pModule.Call(0)
	var err error
	if err := class("EvaDockPanelV08", syscall.NewCallback(panelProc), 0); err != nil {
		return err
	}
	if err := class("EvaDockPetV08", syscall.NewCallback(petProc), csDblClicks); err != nil {
		return err
	}
	panel, _, err = pCreate.Call(0, ptr("EvaDockPanelV08"), ptr("Eva Dock "+version+" - Controles"), wsOverlappedWindow|wsVisible, 65, 75, panelW, panelH, 0, 0, instance, 0)
	if panel == 0 {
		return fmt.Errorf("Crear panel: %v", err)
	}
	setupPanel()
	sw, _, _ := pGetSystem.Call(0)
	sh, _, _ := pGetSystem.Call(1)
	petX = int32(sw) - int32(petW) - 30
	if petX > 835 {
		petX = 835
	}
	if petX < 20 {
		petX = 20
	}
	petY = 120
	_ = sh
	pet, _, err = pCreate.Call(wsExLayered|wsExTopmost|wsExToolwindow, ptr("EvaDockPetV08"), ptr("Eva - mascota"), wsPopup, n32(petX), n32(petY), uintptr(petW), uintptr(petH), 0, 0, instance, 0)
	if pet == 0 {
		logErr("Crear mascota flotante: " + err.Error())
	} else {
		if !drawPet() {
			// Win32 can reject UpdateLayeredWindow. Show a normal always-on-top
			// window as a verifiable fallback rather than leaving Eva invisible.
			logErr("Layered window error, starting visible fallback window")
			pDestroy.Call(pet)
			pet = 0
		}
	}
	if pet == 0 {
		petFallback = true
		pet, _, err = pCreate.Call(wsExTopmost|wsExToolwindow, ptr("EvaDockPetV08"), ptr("Eva - modo visible"), wsPopup|wsVisible, n32(petX), n32(petY), uintptr(petW), uintptr(petH), 0, 0, instance, 0)
		if pet == 0 {
			return fmt.Errorf("No se pudo abrir la ventana de Eva: %v", err)
		}
		pInvalidate.Call(pet, 0, 1)
		setStatus("Eva visible en modo alternativo (con fondo)")
	} else {
		setStatus("Eva visible en modo transparente")
	}
	pShow.Call(pet, swShow)
	pSetPos.Call(pet, ^uintptr(0), n32(petX), n32(petY), uintptr(petW), uintptr(petH), swpNoActivate)

	pShow.Call(panel, swShow)
	pSetTimer.Call(panel, 1, 150, 0)
	repaintPreview()
	go autoUpdateCheck()
	var ms msgT
	for {
		res, _, e := pGetMsg.Call(uintptr(unsafe.Pointer(&ms)), 0, 0, 0)
		if int32(res) == -1 {
			return fmt.Errorf("Mensajes Windows: %v", e)
		}
		if res == 0 {
			break
		}
		pTranslate.Call(uintptr(unsafe.Pointer(&ms)))
		pDispatch.Call(uintptr(unsafe.Pointer(&ms)))
	}
	return nil
}
func main() {
	if handleUpdaterArgs() {
		return
	}
	if firstInstallIfNeeded() {
		return
	}
	if e := main1(); e != nil {
		logErr(e.Error())
		pMessageBox.Call(0, ptr(e.Error()), ptr("Eva Dock - Error"), 0x10)
	}
}
