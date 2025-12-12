# wingui (CLAUDE.md)

This file is project context for coding agents (and humans). Keep it concise and factual.

## TL;DR (Quick Start)

- Windows only (Win32 dialogs + controls).
- If you have a `.syso` file: use `go run .` (NOT `go run main.go`), otherwise resources won't load.
- For keyboard navigation (Tab/accelerators): call `wingui.SetCurrentDialog(dlg.Handle())` for the active dialog (single global).
- If you are building/using wingui as a DLL: call `wingui.InitHInstance("your.dll")` before creating dialogs.

## Repo Facts

- Module: `github.com/whtiehack/wingui`
- Go: `go 1.10` (see `go.mod`)
- Platform: Windows only (no build tags; non-Windows builds will fail)
- Main deps:
  - `github.com/lxn/win` (Win32 APIs/constants/types)
  - `golang.org/x/sys/windows` (low-level Windows syscall support)

## Project Layout (Key Files)

- `wingui.go` - `InitHInstance`, `MessageLoop`, `SetCurrentDialog`
- `dialog.go` - `Dialog`, message routing (`dialogWndProc`), widget binding
- `windows.base.go` - `WindowBase` (common operations + subclassing support)
- Controls:
  - `button.go` - Button/CheckBox/RadioButton
  - `edit.go` - Edit control
  - `static.go` - Static text
  - `combobox.go` - ComboBox
  - `listbox.go` - ListBox
  - `progressbar.go` - ProgressBar
  - `trackbar.go` - TrackBar/Slider
  - `tabcontrol.go` - TabControl (partial; see notes)
  - `image.go`, `bitmap.go` - image/bitmap helpers
- `winapi/` - extra Win32 wrappers/constants not in `lxn/win`
- `tools/genids/` - generate Go consts from `resource.h`
- `examples/` - runnable samples (many include prebuilt `.syso`)

## Runtime Model

### Core Abstractions

- `Widget` (`dialog.go`):
  - `WndProc(msg, wParam, lParam) uintptr`
  - `AsWindowBase() *WindowBase`
- `WindowBase` (`windows.base.go`): embedded by widgets; owns `hwnd`, `idd`, and subclassing state.
- `Dialog` (`dialog.go`):
  - Modeless: `CreateDialogParam`
  - Modal: `DialogBoxParam`
  - Routes messages to widgets via `dialogWndProc`

### Message Routing (Important)

`Dialog.dialogWndProc` routes to `Widget.WndProc` in three ways:

- Subclassed child controls: if `hwnd != dlg.hwnd` and `dlg.items[hwnd]` exists, route to that widget.
- Notifications sent to the dialog (e.g. `WM_COMMAND`): for messages where `lParam` is the child HWND, route to `dlg.items[HWND(lParam)]`.
- `WM_NOTIFY`: routes using `NMHDR.HwndFrom`.

Guideline for widget `WndProc`:
- Handle what you need, then call `w.AsWindowBase().WndProc(...)` to fall through to the previous WndProc when subclassing is active.

### Message Loop / Exit Behavior

`wingui.MessageLoop()`:
- Calls `runtime.LockOSThread()` (Windows UI should stay on one OS thread)
- Runs `GetMessage/TranslateMessage/DispatchMessage`
- If a current dialog is set (global), runs `IsDialogMessage` to support tab/keyboard navigation

`wingui.SetCurrentDialog(hwnd)`:
- Single global handle used by `IsDialogMessage`
- If multiple dialogs are shown, you must update the current one when focus changes (library does not manage this automatically)

Dialog lifetime:
- `Dialog` tracks a global dialog count; when the last dialog is destroyed, it calls `PostQuitMessage(0)`.

## Resources (.rc/.res/.syso)

### Generating `.syso`

Typical (MinGW/TDM-GCC) via `windres`:

```bash
windres -i ui/ui.rc -O coff -o ui.syso
# or
windres -i resource.res -O coff -o vsui.syso
```

Notes:
- Use a `windres` that matches your target arch (32-bit vs 64-bit).
- The `.syso` must be in the package directory you build/run (commonly the `main` package).

### Resource IDs (`resource.h`)

Controls are referenced by numeric IDs. Optional helper to generate Go consts:

```bash
go run tools/genids/genids.go -filename=ui/resource.h -packagename=main
```

`tools/genids` behavior:
- Output file name: `path.Base(filename) + ".go"` (e.g. `resource.h.go`)
- Output location: current working directory (where you run the command)
- Regex only captures positive integer IDs; things like `#define IDC_STATIC (-1)` are ignored.

## Usage Patterns

### Modeless Dialog (Typical App)

```go
dlg, _ := wingui.NewDialog(IDD_DIALOG, 0, nil)
dlg.SetIcon(IDI_ICON)

btn, _ := wingui.BindNewButton(IDC_BUTTON, dlg)
btn.OnClicked = func() { dlg.Close() }

dlg.Show()

// Optional but recommended for tab/keyboard navigation:
wingui.SetCurrentDialog(dlg.Handle())

wingui.MessageLoop()
```

### Modal Dialog

```go
ret := wingui.NewModalDialog(IDD_DIALOG, 0, nil, func(dlg *wingui.Dialog) {
	// Bind widgets here (WM_INITDIALOG time)
})
_ = ret
```

Notes:
- Modal dialogs run their own message loop internally; do not call `wingui.MessageLoop()` just for a modal dialog.

### Subclassing (Intercept Messages on the Control Itself)

```go
edit := wingui.NewEdit(IDC_EDIT)
edit.Subclassing = true // must set BEFORE binding
dlg.BindWidgets(edit)
```

## Widget Helpers (Create/Bind)

- `NewEdit` / `BindNewEdit`
- `NewButton` / `BindNewButton` / `BindNewButtons`
- `NewStatic` / `BindNewStatic`
- `NewImage` / `BindNewImage`
- `NewComboBox` / `BindNewComboBox`
- `NewListBox` / `BindNewListBox`
- `NewProgressBar` / `BindNewProgressBar`
- `NewTrackBar` / `BindNewTrackBar`
- `NewTabControl` / `BindTabControl` (partial implementation)

## Adding a New Widget (Checklist)

1. Create `xxx.go` with `type X struct { WindowBase ... }`
2. Implement `WndProc`:
   - handle notifications/messages you care about
   - call `x.AsWindowBase().WndProc(msg, wParam, lParam)` as fallback
3. Provide `NewX(idd)` and `BindNewX(idd, dlg)` helpers
4. Keep behavior consistent with existing widgets (mostly `WM_COMMAND`/`WM_NOTIFY`-based callbacks)
5. Update docs/examples if it is user-facing

## Troubleshooting

- `Create Dialog error:<idd>`:
  - resource not linked (missing `.syso`, wrong package dir, used `go run main.go`)
  - wrong `hInstance` (DLL scenario: call `InitHInstance("your.dll")`)
- `GetDlgItem Error` while binding:
  - wrong control ID
  - binding before dialog `WM_INITDIALOG` (for modal dialogs, bind inside callback)
- Tab key/accelerators not working:
  - call `wingui.SetCurrentDialog(dlg.Handle())`
  - ensure controls have tab-stop styles in the dialog resource
- Widget callback not firing:
  - ensure widget is bound (`dlg.BindWidgets(...)` succeeded)
  - for `WM_COMMAND`-based handlers, many widgets check `lParam == uintptr(widget.hwnd)`

## Status (Controls)

Implemented:
- Edit, Static, Image
- ComboBox, ListBox
- Button (Push/Check/Radio)
- TrackBar, ProgressBar

Partial / TODO:
- TabControl (partial)
- ListView, TreeView, DateTimePicker, Menu, etc.

