#include <stdio.h>

#define TOTAL     9


void delete(short array[], size_t size, short *to_delete);
void print_array(short array[], size_t size);

int main(void)
{
    short pows[TOTAL] = {0};
    size_t count = 0;
    size_t sz_ar = sizeof(pows) / sizeof(*pows);

    while(count < sz_ar && scanf("%hd", &pows[count]) == 1)
        count++;

    for (short *ptr = pows; ptr != pows + count; ptr++) {
        if (*ptr % 2 == 0) {
            delete(pows, count, ptr);
            count--;
            break;
        }
    }
    print_array(pows, count);
    return 0;
}

void delete(short array[], size_t size, short *to_delete) {
    for (short *ptr = to_delete; ptr != array + size - 1; ptr++) {
        *ptr = *(ptr + 1);
    }
}

void print_array(short array[], size_t size) {
    for (short *ptr = array; ptr != array + size; ptr++) {
        printf("%d ", *ptr);
    }
    putchar('\n');
}