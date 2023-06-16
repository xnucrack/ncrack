#include <stdio.h>

void hello_from_the_library(void)
{
    printf("Hello from dylib\n");
}

void to_replace(void)
{
    printf("I am not replaced\n");
}
