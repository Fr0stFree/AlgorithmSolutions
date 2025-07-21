#include <stdio.h>
#include <stdlib.h>


int find_index(int customers[], int size, int looking_id);

int main() {
    int purchase_amount;
    scanf("%d", &purchase_amount);

    int *customers = malloc(purchase_amount * sizeof(int));
    int unique_customers = 0;
    for (int idx = 0; idx < purchase_amount; idx++) {
        customers[idx] = -1;
    }
    
    for (int idx = 0; idx != purchase_amount; idx++) {
        double timestamp, receipt_id;
        int customer_id;
        scanf("%lf %d %lf", &timestamp, &customer_id, &receipt_id);
        if (find_index(customers, purchase_amount, customer_id) == -1) {
            customers[idx] = customer_id;
            unique_customers++;
        }
    }
    free(customers);
    printf("%d", unique_customers);
    return 0;
}

int find_index(int customers[], int size, int looking_id) {
    for (int idx = 0; idx != size; idx++) {
        if (looking_id == customers[idx]) {
            return idx;
        }
    }
    return -1;
}