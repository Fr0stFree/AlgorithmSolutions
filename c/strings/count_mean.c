#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int main(void)
{
    char str[255];
    fgets(str, sizeof(str)-1, stdin);
    char* ptr_n = strrchr(str, '\n');
    if(ptr_n != NULL)
        *ptr_n = '\0';

    int amount = 0, sum = 0;
    char *ptr = strchr(str, ':') + strlen(": ");
    int mark;
    while (ptr != NULL) {
        amount++;
        sum += atoi(ptr);
        
        ptr = strchr(ptr, ',');
        if (ptr != NULL) {
            ptr += strlen(", ");
        }
    }
    float mean = (float)sum / amount;
    printf("%.3f\n", mean);
    return 0;
}