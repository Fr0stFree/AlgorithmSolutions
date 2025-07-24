#include <stdio.h>
#include <string.h>
#define SIZE 50

int main() {
    char cities[6][SIZE];
    size_t amount = sizeof(cities)/sizeof(*cities);
    char first[SIZE] = "";
    char second[SIZE] = "";

    for (size_t idx = 0; idx < amount; idx++) {
        scanf("%49s", cities[idx]);
        size_t length = strlen(cities[idx]);
        if (length > strlen(first)) {
            strcpy(second, first);
            strcpy(first, cities[idx]);
            continue;
        } 
        if (length > strlen(second)) {
            strcpy(second, cities[idx]);
            continue;
        }
    }
    for (size_t idx = 0; idx < amount; idx++) {
        if (0 == strcmp(first, cities[idx]) || 0 == strcmp(second, cities[idx])) {
            printf("%s ", cities[idx]);
        }
    }
    putchar('\n');
}