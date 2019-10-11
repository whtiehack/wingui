package wingui

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"

	"github.com/lxn/win"
)

// Bitmap struct
type Bitmap struct {
	hBmp       win.HBITMAP
	hPackedDIB win.HGLOBAL
	size       Size
	dpi        int
}

// Dispose release resource.
func (b *Bitmap) Dispose() {
	if b.hBmp != 0 {
		win.DeleteObject(win.HGDIOBJ(b.hBmp))
		win.GlobalUnlock(b.hPackedDIB)
		win.GlobalFree(b.hPackedDIB)
		b.hPackedDIB = 0
		b.hBmp = 0
	}
}

// HBitmap get handle HBITMAP
func (b *Bitmap) HBitmap() win.HBITMAP {
	return b.hBmp
}

// Size get bitmap size
func (bmp *Bitmap) Size() Size {
	return bmp.size
}

// NewBitmapFromResource create a new Bitmap from resource.
func NewBitmapFromResource(name string) (bm *Bitmap, err error) {
	return newBitmapFromResource(syscall.StringToUTF16Ptr(name))
}

// NewBitmapFromResourceId create a new Bitmap from resource ID.
func NewBitmapFromResourceId(idd uintptr) (bm *Bitmap, err error) {
	return newBitmapFromResource(win.MAKEINTRESOURCE(idd))
}

// NewBitmapFromFile create a new Bitmap from  file. Could use jpg file :)
func NewBitmapFromFile(filePath string) (*Bitmap, error) {
	var si win.GdiplusStartupInput
	si.GdiplusVersion = 1
	if status := win.GdiplusStartup(&si, nil); status != win.Ok {
		return nil, errors.New(fmt.Sprintf("GdiplusStartup failed with status '%s'", status))
	}
	defer win.GdiplusShutdown()

	var gpBmp *win.GpBitmap
	if status := win.GdipCreateBitmapFromFile(syscall.StringToUTF16Ptr(filePath), &gpBmp); status != win.Ok {
		return nil, errors.New(fmt.Sprintf("GdipCreateBitmapFromFile failed with status '%s' for file '%s'", status, filePath))
	}
	defer win.GdipDisposeImage((*win.GpImage)(gpBmp))

	var hBmp win.HBITMAP
	if status := win.GdipCreateHBITMAPFromBitmap(gpBmp, &hBmp, 0); status != win.Ok {
		return nil, errors.New(fmt.Sprintf("GdipCreateHBITMAPFromBitmap failed with status '%s' for file '%s'", status, filePath))
	}

	return newBitmapFromHBITMAP(hBmp)
}

func NewBitmapFromWindow(hwnd win.HWND) (*Bitmap, error) {
	hdcMem := win.CreateCompatibleDC(0)
	if hdcMem == 0 {
		return nil, errors.New("CreateCompatibleDC failed")
	}
	defer win.DeleteDC(hdcMem)

	var r win.RECT
	if !win.GetWindowRect(hwnd, &r) {
		return nil, errors.New("GetWindowRect failed")
	}

	hdc := win.GetDC(hwnd)
	width, height := r.Right-r.Left, r.Bottom-r.Top
	hBmp := win.CreateCompatibleBitmap(hdc, width, height)
	win.ReleaseDC(hwnd, hdc)

	hOld := win.SelectObject(hdcMem, win.HGDIOBJ(hBmp))
	flags := win.PRF_CHILDREN | win.PRF_CLIENT | win.PRF_ERASEBKGND | win.PRF_NONCLIENT | win.PRF_OWNED
	win.SendMessage(hwnd, win.WM_PRINT, uintptr(hdcMem), uintptr(flags))

	win.SelectObject(hdcMem, hOld)
	return newBitmapFromHBITMAP(hBmp)
}

func newBitmapFromResource(res *uint16) (bm *Bitmap, err error) {
	if hBmp := win.LoadImage(hInstance, res, win.IMAGE_BITMAP, 129, 57, win.LR_CREATEDIBSECTION); hBmp == 0 {
		err = lastError("LoadImage")
	} else {
		bm, err = newBitmapFromHBITMAP(win.HBITMAP(hBmp))
	}
	return
}

func newBitmapFromHBITMAP(hBmp win.HBITMAP) (bmp *Bitmap, err error) {
	var dib win.DIBSECTION
	if win.GetObject(win.HGDIOBJ(hBmp), unsafe.Sizeof(dib), unsafe.Pointer(&dib)) == 0 {
		return nil, errors.New("GetObject failed")
	}

	bmih := &dib.DsBmih

	bmihSize := uintptr(unsafe.Sizeof(*bmih))
	pixelsSize := uintptr(int32(bmih.BiBitCount)*bmih.BiWidth*bmih.BiHeight) / 8

	totalSize := uintptr(bmihSize + pixelsSize)

	hPackedDIB := win.GlobalAlloc(win.GHND, totalSize)
	dest := win.GlobalLock(hPackedDIB)
	defer win.GlobalUnlock(hPackedDIB)

	src := unsafe.Pointer(&dib.DsBmih)

	win.MoveMemory(dest, src, bmihSize)

	dest = unsafe.Pointer(uintptr(dest) + bmihSize)
	src = dib.DsBm.BmBits

	win.MoveMemory(dest, src, pixelsSize)

	return &Bitmap{
		hBmp:       hBmp,
		hPackedDIB: hPackedDIB,
		size: Size{
			int(bmih.BiWidth),
			int(bmih.BiHeight),
		},
	}, nil
}
