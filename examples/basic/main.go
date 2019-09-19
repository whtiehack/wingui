package main

import (
	"github.com/lxn/win"
	"log"
	"os"

	"github.com/whtiehack/wingui"
)

func init() {
	log.SetOutput(os.Stdout)
}

var dlg *wingui.Dialog

// optional  genereate resource IDs
//go:generate go run github.com/whtiehack/wingui/tools/genids -filename ui/resource.h -packagename main
func main() {
	log.Printf("resource %v %#[1]v  \n", IDD_DIALOG)
	var err error
	dlg, err = wingui.NewDialog(IDD_DIALOG, 0, nil)
	if err != nil {
		log.Panic("main dialog create error", err)
	}
	dlg.SetIcon(IDI_ICON1)
	log.Println("dlg create end", dlg)
	btn := wingui.NewButton(IDB_OK)
	btn.OnClicked = modalBtnClicked
	closeBtn := wingui.NewButton(IDB_CANCEL)
	closeBtn.OnClicked = func() {
		dlg.Close()
	}
	dlg.BindWidgets(btn, closeBtn)

	normalBtn, _ := wingui.BindNewButton(IDB_NORMAL, dlg)
	normalBtn.OnClicked = normalBtnClicked

	bindWidgets(dlg)
	dlg.Show()
	wingui.SetCurrentDialog(dlg.Handle())
	wingui.MessageLoop()
	log.Println("stoped")
}

func normalBtnClicked() {
	d, _ := wingui.NewDialog(IDD_DIALOG_OK, dlg.Handle(), nil)
	d.Show()
}

func modalBtnClicked() {
	log.Println("btn clicked")
	wingui.NewModalDialog(IDD_DIALOG_OK, dlg.Handle(), nil, func(okdlg *wingui.Dialog) {
		okbtn := wingui.NewButton(IDB_OK)
		okbtn.OnClicked = func() {
			log.Println("modal btn click")
			okdlg.Close()
		}
		okdlg.BindWidgets(okbtn)
	})
}

func bindWidgets(dlg *wingui.Dialog) {
	image, _ := wingui.BindNewImage(IDP_BMP, dlg)
	// change bitmap from resource id
	btnChangeBmp, _ := wingui.BindNewButton(IDB_CHANGEBMP, dlg)
	btnChangeBmp.OnClicked = func() {
		bitmap, _ := wingui.NewBitmapFromResourceId(IDB_BITMAP1)
		org := image.LoadBitmap(bitmap.HBitmap())
		win.DeleteObject(win.HGDIOBJ(org))
		bitmap.Dispose()
	}
	// load bit map from file
	btnFileBmp, _ := wingui.BindNewButton(IDB_FILEBMP, dlg)
	btnFileBmp.OnClicked = func() {
		bitmap, err := wingui.NewBitmapFromFile("bitmap.jpg")
		if err != nil {
			log.Panic("!!en", err)
		}
		org := image.LoadBitmap(bitmap.HBitmap())
		win.DeleteObject(win.HGDIOBJ(org))
		bitmap.Dispose()
	}

	// create bitmap from hwnd
	hwndBmp, _ := wingui.BindNewButton(IDB_HWNDBMP, dlg)
	hwndBmp.OnClicked = func() {
		bitmap, err := wingui.NewBitmapFromWindow(dlg.Handle())
		if err != nil {
			log.Panic("!!en", err)
		}
		org := image.LoadBitmap(bitmap.HBitmap())
		win.DeleteObject(win.HGDIOBJ(org))
		bitmap.Dispose()
	}

	// combobox
	combobox, _ := wingui.BindNewComboBox(IDC_COMBOBOX, dlg)
	idx, _ := combobox.AddString("test1")
	combobox.SetItemData(idx, uintptr(101))
	idx, _ = combobox.AddString("test2")
	combobox.SetItemData(idx, uintptr(102))
	idx, _ = combobox.AddString("test3")
	combobox.SetItemData(idx, uintptr(104))
	combobox.DeleteString(1)
	d, _ := combobox.GetItemData(1)
	if d != 104 {
		log.Panic("!! error combo GetItemData")
	}
	str := combobox.GetLbText(1)
	if str != "test3" {
		log.Panic("!! error combo GetLbText")
	}
	log.Println("combo cur sel:", combobox.CurrentIndex())

	combobox.OnSelChange = func() {
		log.Println("! combo sel changed", combobox.CurrentIndex())
	}

	// listbox
	listbox, _ := wingui.BindNewListBox(IDL_LISTBOX, dlg)
	listbox.AddString("@1")
	listbox.AddString("@2")
	listbox.AddString("@3")
	listbox.AddString("@4")
	listbox.DeleteString(2)
	// other
}
