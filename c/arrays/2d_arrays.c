#include <stdio.h>
#define ROWS 4
#define COLUMNS 3

void print_array(int array[ROWS][COLUMNS]);

int main(void)
{
    int ar_2D[ROWS][COLUMNS];
    for (short row_idx = 0; row_idx < ROWS; row_idx++)
    {
        for (short col_idx = 0; col_idx < COLUMNS; col_idx++)
        {
            scanf("%d", &ar_2D[row_idx][col_idx]);
        }
    }

    print_array(ar_2D);
    return 0;
}

void print_array(int array[ROWS][COLUMNS])
{
    for (short row_idx = 0; row_idx < COLUMNS; row_idx++)
    {
        for (short col_idx = 0; col_idx < ROWS; col_idx++)
        {
            (col_idx < ROWS - 1) ? printf("%d ", array[col_idx][row_idx]) : printf("%d", array[col_idx][row_idx]);
        }
        putchar('\n');
    }
}