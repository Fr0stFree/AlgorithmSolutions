#include <stdio.h>

#define ROWS 3
#define COLS 5

int main(void)
{
    short vls[ROWS][COLS] = {0};
    short *ptr_vls = &vls[0][0];
    double mean[ROWS] = {0};

    short x;
    for (int i = 0; i < ROWS * COLS && scanf("%hd", &x) == 1; ++i)
        *ptr_vls++ = x;

    short row_idx = 0;
    while (row_idx < ROWS) {
        int sum = 0, count = 0;
        while (count < COLS) {
            sum += vls[row_idx][count];
            count++;
        }
        mean[row_idx] = (float)sum / count;
        row_idx++;
    }

    for (double *ptr = mean; ptr < mean + ROWS; ptr++) {
        printf("%.2f ", *ptr);
    }
    return 0;
}
