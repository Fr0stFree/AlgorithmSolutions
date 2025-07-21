#include <stdio.h>
#include <stdlib.h>
#define N 100

typedef struct
{
    char a[N];
    unsigned int n;
} Decimal;

void elong_print(Decimal x);

int main()
{
    Decimal x = {{7, 4, 1}, 2};
    Decimal zero = {{0}, 0};

    elong_print(x);
    elong_print(zero);

    return 0;
}

void elong_print(Decimal x)
{
    for (int idx = x.n; idx >= 0; idx--)
    {
        printf("%d", x.a[idx]);
    }
    printf("\n");
}