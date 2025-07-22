#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char *my_strcat(char *dest, const char *src)
{
    char *result = dest;
    dest += strlen(dest);
    while (*src != '\0') {
        *dest = *src;
        dest++;
        src++;
    }
    *dest = '\0';
    return result;
}

int main()
{
    char first[100] = "poly";
    char *second = "ndrome";
    char *concatenated = my_strcat(first, second);
    printf("%s\n", concatenated);
    return 0;
}