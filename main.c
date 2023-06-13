#include <stdio.h>

extern void hello_from_the_library(void);
extern void to_replace(void);

int main(void)
{
    hello_from_the_library();
    to_replace();
    return 0;
}
