#include <stdio.h>


void insert(char *position, char value);

int main(void)
{
    char str[50] = "best string!";
    insert(str, 't');
    insert(str+1, 'h');
    insert(str+2, 'e');
    insert(str+3, ' ');

    return 0;
}

void insert(char *position, char value)
{
    char *end = position;
    while (*end != '\0')
    {
        end++;
    }
    for (char *ptr = end + 1; ptr > position; ptr--) {
        *ptr = *(ptr - 1);
    }
    *position = value;
}