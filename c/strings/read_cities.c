#include <stdio.h>
#include <string.h>
#define SIZE 50
#define AMOUNT 10

int main() {
    char cities[AMOUNT][SIZE];
    char result[(SIZE+1)*AMOUNT];
    short idx = 0;
    char current_city[SIZE];
    while (idx < AMOUNT) {
        if (fgets(current_city, SIZE, stdin) == NULL) {
            break;
        }
        char *last = strchr(current_city, '\n');
        if (last != NULL) {
            *last = '\0';
        }
        if (0 == strcmp("", current_city)) {
            continue;
        }
        strcpy(cities[idx], current_city);
        idx++;
    }


    for (short idx = 0; idx < AMOUNT; idx++) {
        if (strstr(cities[idx], "на")) {
            continue;
        }
        strcat(result, " ");
        strcat(result, cities[idx]);
    }
    puts(result);
    return 0;
}