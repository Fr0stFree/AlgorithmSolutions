#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int main(void)
{
    char str[200] = {0};
    fgets(str, sizeof(str) - 1, stdin);
    char *ptr_n = strrchr(str, '\n');
    if (ptr_n != NULL)
        *ptr_n = '\0';

    double csv[50] = {0};
    char *ptr = strchr(str, ':') + strlen(": ");

    for (short idx = 0; idx < 200 && ptr != NULL; idx++)
    {
        sscanf(ptr, "%lf;", csv+idx);
        ptr = strchr(ptr, ';');
        ptr != NULL && (ptr++);
    }
    return 0;
}