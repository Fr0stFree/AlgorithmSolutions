#include <stdio.h>
#include <limits.h>

int main() {
	int amount, current_number;
    scanf("%d", &amount);

    int lowest = INT_MAX;
    while (amount) {
        scanf("%d", &current_number);
        lowest = (current_number < lowest) ? current_number : lowest;
        amount--;
    }
    printf("%d", lowest);
    
	return 0;
}
