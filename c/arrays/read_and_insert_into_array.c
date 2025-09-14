#include <stdio.h>

#define NUMBERS_AMOUNT 20

int read_array(double *, int);

int main() {
    double numbers[NUMBERS_AMOUNT];

    int number_of_elements = read_array(numbers, NUMBERS_AMOUNT);
    
    double sum = 0;
    for (int i = 0; i < number_of_elements; i++) {
        sum += numbers[i];
    }
    double mean = sum / number_of_elements;
    
    printf("%.2lf\n", mean);
    return 0;
}

int read_array(double *array_ptr, int size) {
    double current_number;
    double *ptr = array_ptr;
    int count = 0;

    while (count < size && scanf("%lf", &current_number) == 1)
    {
        *ptr = current_number;
        ptr++;
        count++;
    }
    return count;
}
