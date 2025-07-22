#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <string.h>

typedef struct
{
    unsigned int length;
    char course[6];
} Direction;

Direction retrieve_direction(char str[]);

int main()
{
    int x = 0;
    int y = 0;

    char buffer[10];
    while (fgets(buffer, 10, stdin) != NULL)
    {
        buffer[strcspn(buffer, "\n")] = '\0';
        if (strcmp(buffer, "Treasure!") == 0)
        {
            break;
        }
        Direction dir = retrieve_direction(buffer);
        if (strcmp(dir.course, "North") == 0) {
            y += dir.length;
        } else if (strcmp(dir.course, "South") == 0) {
            y -= dir.length;
        } else if (strcmp(dir.course, "East") == 0) {
            x += dir.length;
        } else if (strcmp(dir.course, "West") == 0) {
            x -= dir.length;
        }
    }
    printf("%d %d", x, y);
    return 0;
}

Direction retrieve_direction(char str[])
{
    Direction direction;
    int breakpoint = 0;
    for (int idx = 0; str[idx] != ' ' && str[idx] != '\0'; idx++)
    {
        breakpoint++;
    }
    strncpy(direction.course, str, breakpoint);
    direction.course[breakpoint] = '\0';
    direction.length = (unsigned int)atoi(&str[breakpoint+1]);
    return direction;
}

// North 5
// East 3
// South 1
// Treasure!