#include <stdio.h>

#define SIZE_BUFFER     11

void print_array(int array[], size_t size);

int main(void)
{
    int buffer[SIZE_BUFFER];
    size_t count = 0;
    size_t sz_ar = sizeof(buffer) / sizeof(*buffer);

    while(count < sz_ar && scanf("%d", &buffer[count]) == 1)
        count++;

    int temp;
    int *left_ptr = buffer;
    int *right_ptr = (count % 2 == 0) ? buffer + (count / 2) : buffer + (count / 2) + 1;
    while (right_ptr != buffer + count)
    {
        temp = *left_ptr;
        *left_ptr = *right_ptr;
        *right_ptr = temp;
        right_ptr++;
        left_ptr++;
    }
    print_array(buffer, count);
    return 0;
}

void print_array(int array[], size_t size) {
    for (int *ptr = array; ptr != array + size; ptr++) {
        printf("%d ", *ptr);
    }
    putchar('\n');
}