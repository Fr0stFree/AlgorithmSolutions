#include <stdio.h>

#define TOTAL 20

void insert(int array[], size_t size, int *position, int value);
void print_array(int array[], size_t size);

int main(void)
{
    int digs[TOTAL] = {0};
    size_t count = 0;
    size_t sz_ar = sizeof(digs) / sizeof(*digs);

    while (count < sz_ar && scanf("%d", &digs[count]) == 1)
        count++;

    for (int *ptr = digs; ptr != digs + count; ptr++)
    {
        if (*ptr == 5)
        {
            insert(digs, count, ptr+1, -5);
            if (count < TOTAL) {
                count++;
            }
            break;
        }
    }
    print_array(digs, count);
    return 0;
}

void insert(int array[], size_t size, int *position, int value)
{
    int *ptr;
    for (ptr = array + size; ptr > position; ptr--)
    {
        *ptr = *(ptr-1);
    }
    *position = value;
}

void print_array(int array[], size_t size)
{
    for (int *ptr = array; ptr != array + size; ptr++)
    {
        printf("%d ", *ptr);
    }
    putchar('\n');
}