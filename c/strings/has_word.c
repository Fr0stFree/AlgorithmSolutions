#include <stdio.h>
#include <string.h>
#include <stdbool.h>

int main()
{
    char words[1000];
    bool has_bomb = false;

    while (scanf("%999s", words) == 1)
    {
        if (0 == strcmp(words, "bomb"))
        {
            has_bomb = true;
        }
    }
    char *result = has_bomb ? "YES" : "NO";
    printf("%s\n", result);
    return 0;
}
