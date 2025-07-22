#include <stdio.h>

int find_node(int, int);

int main() {
    int a, b;
    scanf("%d %d", &a, &b);
    int nod = find_node(a, b);
    printf("%d", nod);
    return 0;
}

int find_node(int a, int b) {
    if (a == b) {
        return b;
    }
    if (b > a) {
        int buf = a;
        a = b;
        b = buf;
    }
    int remainder = a % b;
    if (remainder == 0) {
        return b;
    }
    return find_node(b, remainder);
}