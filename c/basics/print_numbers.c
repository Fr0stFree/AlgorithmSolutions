#include <stdio.h>
#define N 1000

int main()
{
    int a[N]; // массив для N чисел
    int i;    // номер текущей ячейки
    int n;    // сколько чисел дано

    scanf("%d", &n);

    for (i = 0; i < n; i++) {
        scanf("%d", &a[i]);
    }

    for (i = 0; i < n; i++) {
        a[i] % 2 == 0 && printf("%d ", a[i]);
    }
    printf("\n");

    for (i = 0; i < n; i++) {
        a[i] % 2 == 1 && printf("%d ", a[i]);
    }
    printf("\n");
    return 0;
}
