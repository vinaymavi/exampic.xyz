#include <iostream>
#include <emscripten.h>

extern "C" {
    char*  get_string(int index){
    char* s = (char*)malloc(100);
    char str1[] = "Hello from C++!";
    char str2[] = "Hello World";
    char str3[] = "Hello there :)";
    char* strArr[] = {str1, str2, str3};


    sprintf(s,"%s", strArr[index]);
    return s;
}
}


int main(){
    std::cout << "Hello, webassembly!" << std::endl;
    char* s = get_string(1);
    std::cout << s << std::endl;
    return 0;
}