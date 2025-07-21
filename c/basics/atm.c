#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main() {
    int banknotes[] = {1, 2, 5, 10, 50, 100, 500, 1000, 5000};
    int needed_banknotes[sizeof(banknotes)/sizeof(banknotes[0])];

    memset(needed_banknotes, 0, sizeof(needed_banknotes));

    int asked_amount;
    scanf("%d", &asked_amount);

    for (size_t idx = sizeof(banknotes)/sizeof(banknotes[0])-1; idx != (size_t) - 1; idx--) {
        if (banknotes[idx] > asked_amount) {
            printf("%d %d\n", banknotes[idx], 0);
            continue;
        }
        needed_banknotes[idx] = asked_amount / banknotes[idx];
        asked_amount %= banknotes[idx];
        printf("%d %d\n", banknotes[idx], needed_banknotes[idx]);
    }

    return 0;
}