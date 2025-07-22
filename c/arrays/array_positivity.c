#include <stdio.h>
#define SIZE 10

int main(void)
{
    int numbers[SIZE] = {0};
    int current_number;
    for (short idx = 0; idx < SIZE; idx++)
    {
        int amount = scanf("%d", &current_number);
        if (!amount)
            break;

        numbers[idx] = current_number;
    }

    short result = 1;
    size_t len = sizeof(numbers) / sizeof (*numbers);
    for (size_t idx = 0; idx < len; idx++) {
        if (numbers[idx] < 0 || numbers[idx] % 2 == 0) {
            result = 0;
            break;
        }
    }
    printf("%d", result);
    return 0;
}