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
    char str[100], res_str[100] = "I love language C ";
    fgets(str, sizeof(str), stdin);
    strip_string(str, sizeof(str));

    strncat(res_str, str, sizeof(res_str) - strlen(res_str) - 1);
    puts(res_str);

    return 0;
}