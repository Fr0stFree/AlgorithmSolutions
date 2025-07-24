#include <stdio.h>
#include <string.h>

void delete(char *symbol);

int main(void)
{
    char str[100] = {0};
    fgets(str, sizeof(str)-1, stdin);
    char* ptr_n = strrchr(str, '\n');
    if(ptr_n != NULL)
        *ptr_n = '\0';

    for (char *next = strchr(str, '-'); next != NULL; next = strchr(next+1, '-')) {
        while (*(next+1) == '-'){
            delete(next+1);
        }
    }
    puts(str);
    return 0;
}


void delete(char *symbol) {
    char *ptr = symbol;
    while (*ptr != '\0') {
        *ptr = *(ptr+1);
        ptr++;
    }
}