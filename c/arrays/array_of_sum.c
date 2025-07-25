#include <stdio.h>

#define ROWS 3
#define COLS 4

void print_array(int array[ROWS][COLS]);

int main(void)
{
    int a[ROWS][COLS] = {0};
    int b[ROWS][COLS] = {0};
    int *ptr_a = &a[0][0];
    int *ptr_b = &b[0][0];
    int res[ROWS][COLS] = {0};

    int x;
    for (int i = 0; i < ROWS * COLS && scanf("%d", &x) == 1; ++i)
        *ptr_a++ = x;
    for (int i = 0; i < ROWS * COLS && scanf("%d", &x) == 1; ++i)
        *ptr_b++ = x;

    for (short row_idx = 0; row_idx < ROWS; row_idx++)
        for (short col_idx = 0; col_idx < COLS; col_idx++)
            res[row_idx][col_idx] = a[row_idx][col_idx] + b[row_idx][col_idx];

    print_array(res);
    return 0;
}

void print_array(int array[ROWS][COLS])
{
    for (short row_idx = 0; row_idx < ROWS; row_idx++)
    {
        for (short col_idx = 0; col_idx < COLS; col_idx++)
        {
            printf("%d", array[row_idx][col_idx]);
            if (col_idx != COLS - 1)
                printf(" ");
        }
        putchar('\n');
    }
}