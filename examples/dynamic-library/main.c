

#include <stdio.h>
#include <Windows.h>

typedef int(*showui)();
int main(){
    HMODULE h;
    FARPROC proc;
    printf("main start\n");
    h = LoadLibraryA("dll/test.dll");
    //h = LoadLibraryA("F:/goproj/checkdll_log/dll/test.dll");
    printf("main load library %X\n",h);
    proc = GetProcAddress(h, "Showui");
    printf("main GetProcAddress %X\n",proc);
    (showui)(proc)();
    printf("main end\n");
    return 0;
}