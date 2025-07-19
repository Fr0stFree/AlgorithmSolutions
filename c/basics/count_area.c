#include <stdio.h>

int main() {
    float height;
    float width;

    scanf("%f %f", &height, &width);

    float area = height * width;

    printf("%f", area);
    return 0;
}