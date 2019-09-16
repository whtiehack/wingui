package wingui

import (
	"errors"
	"fmt"
	"github.com/lxn/win"
	"syscall"
	"unsafe"
)

type Bitmap struct {
	hBmp       win.HBITMAP
	hPackedDIB win.HGLOBAL
	size       Size
	dpi        int
}

func (b *Bitmap) Dispose() {
	if b.hBmp != 0 {
		win.DeleteObject(win.HGDIOBJ(b.hBmp))
		win.GlobalUnlock(b.hPackedDIB)
		win.GlobalFree(b.hPackedDIB)
		b.hPackedDIB = 0
		b.hBmp = 0
	}
}

func (b *Bitmap) GetHBitmap() win.HBITMAP {
	return b.hBmp
}

func (b *Bitmap) GetSize() Size {
	return b.size
}

func NewBitmapFromResource(name string) (bm *Bitmap, err error) {
	return newBitmapFromResource(syscall.StringToUTF16Ptr(name))
}

func NewBitmapFromResourceId(idd uintptr) (bm *Bitmap, err error) {
	return newBitmapFromResource(win.MAKEINTRESOURCE(idd))
}

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
