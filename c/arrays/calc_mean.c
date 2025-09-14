#include <stdio.h>

double mean_ar(const int *ar, size_t len_ar, int(*predicate)(int))
{   
    int sum = 0;
    short amount = 0;
    for (short idx = 0; idx < len_ar; idx++) {
        if (predicate(ar[idx])) {
            sum += ar[idx];
            amount++;
        }
    }
    double mean = (double)sum / amount;
    return mean;
}

int is_valid_mark(int num) {
    return num >= 1 && num <= 5;
}

int main(void)
{
    int marks[20] = {0};
    int x;
    int count = 0;
    while(scanf("%d", &x) == 1) {
        marks[count] = x;
        count++;
    }
    double mean = mean_ar(marks, count, is_valid_mark);
    printf("%.1lf", mean);
    return 0;
}