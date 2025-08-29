#include <stdio.h>
#include <string.h>

void strip_string(char* str, int max_len)
{
    int count = 0;
    while(*str++ != '\0' && count++ < max_len);

    if(count > 1) {
        str -= 2;
        if(*str == '\n')
            *str = '\0';
    }
}

int main(void)
{
    char str[100];
    fgets(str, sizeof(str), stdin);
    strip_string(str, sizeof(str));
    char* str_ptr = str;
    for (char* ptr = strchr(str_ptr, 'e'); ptr != NULL; ptr = strchr(str_ptr, 'e')) {
        printf("%.*s", (int)(ptr - str_ptr), str_ptr);
        str_ptr = ptr + 1;
    }
    printf("%s\n", str_ptr);
    return 0;
}