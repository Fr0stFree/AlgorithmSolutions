#include <stdio.h>
#include <string.h>

int main(void)
{
    char str[100] = {0};
    fgets(str, sizeof(str) - 1, stdin);
    char *ptr_n = strrchr(str, '\n');
    if (ptr_n != NULL)
        *ptr_n = '\0';

    const char looking_char = '-';
    char new_str[100] = {0};

    char *left_border = str;
    for (char *right_border = strchr(left_border, looking_char); right_border != NULL; right_border = strchr(left_border, looking_char))
    {
        strncat(new_str, left_border, right_border - left_border);
        strcat(new_str, "-+-");
        left_border = right_border + 1;
    }
    strcat(new_str, left_border);
    puts(new_str);
    return 0;
}