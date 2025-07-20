#include <stdio.h>
#include <stdbool.h>

typedef struct {
    float x;
    float y;
} Point;

bool is_belong(Point);

int main() {
    Point point;
    scanf("%f %f", &point.x, &point.y);
    if (is_belong(point)) {
        printf("YES");
    } else {
        printf("NO");
    }
    return 0;
}

bool is_belong(Point point) {
    if (point.y > 3 + point.x) {
        return false;
    }
    if (point.y > 3 - point.x) {
        return false;
    }
    if (point.y < -2) {
        return false;
    }
    return true;
}