package main

import (
	"log"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/lxn/win"
	"github.com/whtiehack/wingui"
)

func init() {
	log.SetOutput(os.Stdout)
}

var dlg *wingui.Dialog

// optional  genereate resource IDs
//
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
	log.Println("combo cur sel:", combobox.GetCurSel())

	combobox.OnSelChange = func() {
		log.Println("! combo sel changed", combobox.GetCurSel())
	}

	// listbox
	listbox, _ := wingui.BindNewListBox(IDL_LISTBOX, dlg)
	listbox.AddString("@1")
	listbox.AddString("@2")
	listbox.AddString("@3")
	listbox.AddString("@4")
	listbox.DeleteString(2)

	listbox.OnDoubleClick = func() {
		log.Println(" list box double click getcursel:", listbox.GetCurSel(),
			" GetSelectedIndexes:", listbox.GetSelectedIndexes())
	}
	listbox.OnSelChange = func() {
		log.Println("list box on sel change:", listbox.GetCurSel())
	}

	// checkbox
	checkbox, _ := wingui.BindNewButton(IDB_CHECK, dlg)
	checkbox.SetCheck(1)
	// radios
	radios, _ := wingui.BindNewButtons([]uintptr{IDB_RADIO1, IDB_RADIO2}, dlg)
	radios[1].Click()
	log.Println("radiosstatus", radios[0].GetCheck(), radios[1].GetCheck())

	// slider bar
	sliderBar, _ := wingui.BindNewTrackBar(IDS_SLIDER, dlg)
	sliderBar.ClearSel(false)

	go func() {
		time.Sleep(3 * time.Second)
		checkbox.Click()
		radios[0].SetCheck(1)
		radios[1].SetCheck(0)
		log.Println("radiosstatus", radios[0].GetCheck(), radios[1].GetCheck())

		linesize := sliderBar.GetLineSize()
		numTics := sliderBar.GetNumTics()
		pagesize := sliderBar.GetPageSize()
		pos := sliderBar.GetPos()
		rangeMax := sliderBar.GetRangeMax()
		rangeMin := sliderBar.GetRangeMin()
		log.Println("sliderBar linesize:", linesize, " numtics:", numTics,
			" pagesize:", pagesize, " pos:", pos, " rangeMin:", rangeMin, " rangeMax:", rangeMax)

	}()

	// progress bar
	progressBar, _ := wingui.BindNewProgressBar(IDP_PROGRESSBAR, dlg)

	go func() {
		pos := 0
		for pos < 100 {
			pos = progressBar.DeltaPos(1)
			time.Sleep(50 * time.Millisecond)
		}
		progressBar.DeltaPos(-50)
	}()

	// tab control
	tabcontrol, _ := wingui.BindTabControl(IDC_TAB1, dlg)
	tabcontrol.OnSelChange = func(newIndex int) {
		dlg.SetText("Dialog Hello - tab " + strconv.Itoa(newIndex))
	}

	tab1, err := wingui.NewDialog(IDD_DIALOG_TAB1, tabcontrol.Handle(), nil)
	if err != nil {
		log.Println("tab1 error", err)
		return
	}
	lv1, _ := wingui.BindNewListView(IDC_TAB_LIST1, tab1)
	lv1.SetExtendedStyle(win.LVS_EX_FULLROWSELECT | win.LVS_EX_GRIDLINES | win.LVS_EX_DOUBLEBUFFER)
	lv1.InsertColumn(0, "Name", 120, win.LVCFMT_LEFT)
	lv1.InsertColumn(1, "Value", 80, win.LVCFMT_LEFT)
	for i := 0; i < 5; i++ {
		idx := lv1.InsertItem(-1, "Item "+strconv.Itoa(i))
		lv1.SetItemText(idx, 1, "V"+strconv.Itoa(i))
	}
	lv1.OnItemActivate = func(index int) {
		if index < 0 {
			return
		}
		tab1.MessageBox("Activate: "+lv1.GetItemText(index, 0), "ListView", win.MB_OK|win.MB_ICONINFORMATION)
	}

	tabBtn, _ := wingui.BindNewButton(IDC_TAB_BUTTON1, tab1)
	tabBtn.SetText("Add Item")
	var addSeq int32 = 5
	tabBtn.OnClicked = func() {
		n := int(atomic.AddInt32(&addSeq, 1) - 1)
		idx := lv1.InsertItem(-1, "Item "+strconv.Itoa(n))
		lv1.SetItemText(idx, 1, "V"+strconv.Itoa(n))
	}
	tabcontrol.AddDialogPage("Page1", tab1)

	tab2, err := wingui.NewDialog(IDD_DIALOG_TAB1, tabcontrol.Handle(), nil)
	if err != nil {
		log.Println("tab2 error", err)
		return
	}
	lv2, _ := wingui.BindNewListView(IDC_TAB_LIST1, tab2)
	lv2.SetExtendedStyle(win.LVS_EX_FULLROWSELECT | win.LVS_EX_GRIDLINES | win.LVS_EX_DOUBLEBUFFER)
	lv2.InsertColumn(0, "Name", 120, win.LVCFMT_LEFT)
	lv2.InsertColumn(1, "Value", 80, win.LVCFMT_LEFT)
	for i := 0; i < 3; i++ {
		idx := lv2.InsertItem(-1, "Row "+strconv.Itoa(i))
		lv2.SetItemText(idx, 1, time.Now().Format("15:04:05"))
	}

	tabBtn2, _ := wingui.BindNewButton(IDC_TAB_BUTTON1, tab2)
	tabBtn2.SetText("Show Selected")
	tabBtn2.OnClicked = func() {
		sel := lv2.SelectedIndex()
		if sel < 0 {
			tabBtn2.MessageBox("No selection", "ListView", win.MB_OK|win.MB_ICONWARNING)
			return
		}
		tabBtn2.MessageBox("Selected: "+lv2.GetItemText(sel, 0)+" / "+lv2.GetItemText(sel, 1), "ListView", win.MB_OK|win.MB_ICONINFORMATION)
	}
	tabcontrol.AddDialogPage("Page2", tab2)
	tabcontrol.Select(0)
	// other

}
