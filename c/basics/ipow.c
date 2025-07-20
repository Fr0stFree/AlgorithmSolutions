#include <stdio.h>

long long int ipow(long long int a, unsigned int n, unsigned int * depth);

int main() {
    int depth = 0;
    int result = ipow(3, 12, &depth);
    printf("%d %d\n", result, depth);
    return 0;   
}

long long int ipow(long long int a, unsigned int n, unsigned int * depth) {
    static int IS_FIRST_CALL = 1;
    if (IS_FIRST_CALL == 1) {
        *depth = 1;
        IS_FIRST_CALL = 0;
    }
    if (n == 0) {
        IS_FIRST_CALL = 1;
        return 1;
    }
    if (n == 1) {
        IS_FIRST_CALL = 1;
        return a;
    }
    (*depth)++;
    long long int result;
    if (n % 2 == 0) {
        result = ipow(a, n / 2, depth);
        return result * result;
    }
    result = ipow(a, n - 1, depth);
    return result * a;
}