#include <stdio.h>

#define TOTAL 10

int main(void)
{
    int digs[TOTAL] = {0};
    size_t count = 0;
    size_t sz_ar = sizeof(digs) / sizeof(*digs);

    while (count < sz_ar && scanf("%d", &digs[count]) == 1)
        count++;

    int out[TOTAL];
    size_t out_count = 0;
    int inserts = 0;

    for (size_t i = 0; i < count && out_count < TOTAL; i++) {
        out[out_count++] = digs[i];
        if (digs[i] == 5 && out_count < TOTAL) {
            inserts++;
            out[out_count++] = -inserts;
        }
    }

    for (size_t i = 0; i < out_count; i++) {
        printf("%d ", out[i]);
    }
    return 0;
}