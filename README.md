

# Windows GUI Library  - wingui
Golang GUI library for windows, UI created by ResEdit or Visual Studio Resource Editor

Lightweight, high performance and small executable file size Windows GUI library.

# UI Design tool

> ResEdit

> Visual Studio Resource Editor

> Other dialog box editor




# Requires

### GCC

MinGW
http://www.mingw.org

or 

TDM-GCC
http://tdm-gcc.tdragon.net/


#### windres.exe

This tool in `TDM-GCC-64/bin/windres.exe`,don't know where in MinGW.

# Usage

### Simple usage:

Generate x.syso file from rc or res file use `windres.exe` tool.

genereate syso:
`windres -i emptyProject/Debug/resource.res -O coff -o vsui.syso`

or

`windres -i ui/ui.rc -O coff -o ui.syso`

main.go
```go
package main

import "github.com/whtiehack/wingui"

func main() {
	dlg, _ := wingui.NewDialog(101, 0, nil)
	dlg.SetIcon(105)
	btnok, _ := wingui.BindNewButton(1002, dlg)
	btncancel, _ := wingui.BindNewButton(1003, dlg)
	btnok.OnClicked = func() {
		dlg.Close()
	}
	btncancel.OnClicked = btnok.OnClicked
	dlg.Show()
	// This invoke is optional.
	wingui.SetCurrentDialog(dlg.Handle())
	wingui.MessageLoop()
}


```


run:
`go run .`

Don't use `go run main.go`, because golang can't load x.syso files.




[More examples](https://github.com/whtiehack/wingui/tree/master/examples)

# Examples

see https://github.com/whtiehack/wingui/tree/master/examples
Welcome PRs.


# References 

https://github.com/lxn/walk

https://github.com/sumorf/gowi



# Screenshot

<details><summary><b> UI Screenshot details </b> </summary><br>


### ResEdit
![resedit](res/resedit.png)

![wowjump](res/wowjumpres.png)
[ResEdit Download](http://www.resedit.net/)


### Visual Studio Resource Editor

![vsreseditor](res/vsreseditor.png)


### Effect
![resedit](res/resedit_show.png)
![wowjump](res/wowjump.png)

### File size
![size](res/size.png)


</details>
