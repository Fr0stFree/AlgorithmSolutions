#include <stdio.h>
#include <stdlib.h>
#define N 100

typedef struct
{
    char a[N];
    unsigned int n;
} Decimal;

void elong_print(Decimal x);
void elong_set(Decimal *res, const char str[]);

int main()
{
    Decimal res;
    elong_set(&res, "12345678901234567890");
    elong_print(res);

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

void elong_set(Decimal *res, const char str[])
{
    int len = 0;
    while (str[len] != '\0') {
        len++;
    }

    res->n = len - 1;
    for (int idx = 0; str[idx] != '\0'; idx++)
    {
        res->a[idx] = str[len - 1 - idx] - '0';
    }
}