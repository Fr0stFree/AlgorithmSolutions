#include <stdio.h>

int main(void)
{
    int marks[100];
    int current_number = 0;
    int *marks_ptr = marks;

    while (1)
    {
        scanf("%d", &current_number);
        if (current_number == 78)
            break;

        *marks_ptr++ = current_number;
    }

    for (int *ptr = marks; ptr < marks_ptr; ptr++)
    {
        printf("%d ", *ptr);
    }
    return 0;
}