#include <windows.h>


// Message Loop
void MessageLoop(){
    MSG msg;
    while (GetMessage(&msg, NULL, 0, 0))
    {
            TranslateMessage(&msg);
            DispatchMessage(&msg);
    }

}