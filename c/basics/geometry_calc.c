#include <stdio.h>
#include <math.h>

typedef struct {
    int x;
    int y;
} Point;

typedef struct {
    Point a;    // начало отрезка
    Point b;    // конец отрезка
    float len;  // длина отрезка
} Line;

float distance(Point a, Point b) {
    int dx = b.x - a.x;
    int dy = b.y - a.y;
    double result = sqrt(pow(dx, 2) + pow(dy, 2));
    return result;
}

void scanLine(Line * t) {
    scanf("%d %d %d %d", &t -> a.x, &t -> a.y, &t -> b.x, &t -> b.y);
};

void printLine(Line t) {
    double d = distance(t.a, t.b);
    printf("%d %d %d %d %.3f", t.a.x, t.a.y, t.b.x, t.b.y, d);
}

void rotRLine(Line * t) {
    Point a = t->a;
    Point b = t->b;

    t->a.x = a.y;
    t->a.y = -a.x;
    t->b.x = b.y;
    t->b.y = -b.x;
}

int main() {
    Line t;

    scanLine(&t);
    rotRLine(&t);
    printLine(t);

    return 0;
}