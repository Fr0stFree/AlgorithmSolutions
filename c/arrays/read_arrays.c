#include <stdio.h>
#define TOTAL 5

int main()
{
    int ar_1[TOTAL];
    int ar_2[TOTAL*2];

    for (short idx = 0; idx < TOTAL; idx++)
    {
        scanf("%d", ar_1 + idx);
    }
    for (short idx = 0; idx < TOTAL*2; idx++)
    {
        ar_2[idx] = (idx < TOTAL) ? ar_1[idx] : -1;
    }
    return 0;
}