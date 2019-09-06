windres -i ui/ui.rc -O coff -o ui.syso
go build -ldflags="-s -w -H windowsgui"

