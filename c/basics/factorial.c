#include <stdio.h>


int main() {
	int number;
	int result = 1;

	scanf("%d", &number);

	for (int index = 1; index <= number; index++) {
		result *= index;
	}
	printf("%d", result);
	return 0;
}
