#include <stdio.h>

int main() {
    int round_length;
    int distance_ran;
    scanf("%d %d", &round_length, &distance_ran);

    int rounds_ran = distance_ran / round_length;
    int finish_distance_diff = distance_ran % round_length;

    printf("%d %d", rounds_ran, finish_distance_diff);
    return 0;
}