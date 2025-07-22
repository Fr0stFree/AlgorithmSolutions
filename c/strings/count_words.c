#include <stdio.h>
#include <string.h>

int main()
{
    char buffer[1000];
    int amount = 0;
    char *ptr = buffer;

    fgets(buffer, sizeof(buffer), stdin);
    while (strstr(ptr, "bomb") != NULL)
    {
        ptr = strstr(ptr, "bomb");
        ptr++;
        amount++;
    }
    printf("%d\n", amount);
    return 0;
}
