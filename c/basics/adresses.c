#include <stdio.h>

int main(void)
{
    double value;

    scanf("%lf", &value);

    char *ptr = (char*)&value;
    for (short idx = 0; idx < sizeof(double); idx++) {
        printf("%d ", (signed char)*(ptr+idx));
    }

    return 0;
}