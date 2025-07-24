#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(void)
{
    const short MAX_AMOUNT = 20;
    double digits[20];

    short count = 0;
    while (count != MAX_AMOUNT) {
        short has_scanned = scanf("%lf", (digits + count));
        if (!has_scanned) break;
        count++;
    }

    double *ptr_d = malloc(sizeof(double) * count);
    memcpy(ptr_d, digits, count*sizeof(double));

    free(ptr_d);
    return 0;
}