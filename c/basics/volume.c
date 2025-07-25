#include <stdio.h>

int get_volume(int h, int v, int d) {
    return h * v * d;
}

int main(void)
{
    int height, width, depth;
    scanf("%d %d %d", &height, &width, &depth);
    int result = get_volume(height, width, depth);
    printf("%d\n", result);
    return 0;
}