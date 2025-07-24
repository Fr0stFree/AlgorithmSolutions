#include <stdio.h>

int main() {
    char data[6][30];
    size_t len = sizeof(data) / sizeof(*data);
    for (size_t idx = 0; idx < len; idx++) {
        scanf("%29s", data[idx]);
    }

    for (size_t idx = 0; idx < len; idx++) {
        char letter = data[idx][0];
        if (letter == 'G') {
            printf("%s ", data[idx]);
        }
    }
    return 0;
}