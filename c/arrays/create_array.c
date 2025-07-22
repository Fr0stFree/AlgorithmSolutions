#include <stdio.h>
#include <stdlib.h>

int *new_array(size_t size) {
    int *array = (int*)malloc(size * sizeof(int));
    return array;
}

int main() {
    size_t size;
    scanf("%zu", &size);
    int *array = new_array(size);
    for (size_t idx = 0; idx < size; idx++) {
        printf("array[%zu] = %d\n", idx, array[idx]);
    }
    free(array);
    return 0;
}