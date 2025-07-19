#include <stdio.h>

float fahr(int cel);

int main()
{
    int cel;
    float f;

    scanf("%d", &cel);
    f = fahr(cel);
    printf("%.2f\n", f);

    return 0;
}

float fahr(int cel) {
    float result = cel * 9 / 5 + 32;
    return result;
}