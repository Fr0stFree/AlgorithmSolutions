#include <stdio.h>

int main(void)
{
    short ar[10], marks[5];
    size_t count = 0;
    size_t sz_ar = sizeof(ar) / sizeof(*ar);

    while(count < sz_ar && scanf("%hd", &ar[count]) == 1)
        count++;

    short *ar_ptr = ar + count - 1;
    short *marks_ptr = marks;
    while (ar_ptr >= ar && marks_ptr - marks < 5) {
        *marks_ptr++ = *ar_ptr--;
    }
    
    for (short *ptr = marks; ptr < marks_ptr; ptr++) {
        printf("%hd ", *ptr);
    }
    putchar('\n');
    return 0;
}