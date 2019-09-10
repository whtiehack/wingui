#include <windows.h>

HWND dlg;

// Message Loop
void MessageLoop(){
    MSG msg;
    while (GetMessage(&msg, NULL, 0, 0))
    {
        if(dlg){
                if(!IsDialogMessage(dlg, &msg))
                {
                    TranslateMessage(&msg);
                    DispatchMessage(&msg);
                }
        }else{
            TranslateMessage(&msg);
            DispatchMessage(&msg);
        }
    }

}

void SetCurrentDialog(long long int h){
    dlg = (HWND)h;
}