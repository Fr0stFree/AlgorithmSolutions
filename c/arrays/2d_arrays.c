#include <stdio.h>
#define ROWS 3
#define COLUMNS 3

void print_array(int array[COLUMNS][ROWS]);

int main(void)
{
    int ar_2D[ROWS][COLUMNS];
    for (short col_idx = 0; col_idx < COLUMNS; col_idx++)
    {
        for (short row_idx = 0; row_idx < ROWS; row_idx++)
        {
            scanf("%d", &ar_2D[col_idx][row_idx]);
        }
    }

    print_array(ar_2D);
    return 0;
}

void print_array(int array[COLUMNS][ROWS])
{
    for (short col_idx = 0; col_idx < COLUMNS; col_idx++)
    {
        for (short row_idx = 0; row_idx < ROWS; row_idx++)
        {
            printf("%d ", array[col_idx][row_idx]);
        }
        putchar('\n');
    }
}