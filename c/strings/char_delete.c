#include <stdio.h>

void delete(char string[], char *symbol);

int main(void)
{
    char str[100] = "Best string!";
    char *ptr = str;
    while (*ptr != '\0') {
        if (*ptr != 's') {
            ptr++;
            continue;
        }
        delete(str, ptr);
    }
    return 0;
}

void delete(char string[], char *symbol) {
    char *ptr = symbol;
    while (*ptr != '\0') {
        *ptr = *(ptr+1);
        ptr++;
    }
}