#include <stdio.h>
#include <string.h>

int main(void)
{
    char str[9] = {0};
    fgets(str, sizeof(str), stdin);
    char *ptr_n = strrchr(str, '\n');
    if (ptr_n != NULL)
        *ptr_n = '\0';

    char hours[3] = {str[0], str[1], '\0'};
    char mins[3] = {str[3], str[4], '\0'};
    char sec[3] = {str[6], str[7], '\0'};
    char new_str[9] = {};
    sprintf(new_str, "%s:%s:%s", sec, mins, hours);
    puts(new_str);
    return 0;
}