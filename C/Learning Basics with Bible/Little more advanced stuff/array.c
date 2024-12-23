#include <stdio.h>
#include <stdlib.h>

int main() {
    int elements_of_array;
    printf("How many elements in the list: ");
    scanf("%d", &elements_of_array);

    double *prices = (double *)malloc(elements_of_array * sizeof(double));
    if (prices == NULL) {
        printf("Memory allocation failed.\n");
        return 1;
    }

    // Input loop
    for (int j = 0; j < elements_of_array; j++) {
        printf("What is the %d element of the array: ", j + 1);
        if (scanf("%lf", &prices[j]) != 1) {
            printf("Invalid input. Exiting.\n");
            free(prices);
            return 1;
        }
    }

    for (int i = 0; i < elements_of_array; i++) {
        printf("The %dth element of the list is %.1lf\n", i + 1, prices[i]);
    }
    // printf("\nArray has %lld bytes", sizeof(prices)); HELL NAH //

    free(prices);

    return 0;
}
