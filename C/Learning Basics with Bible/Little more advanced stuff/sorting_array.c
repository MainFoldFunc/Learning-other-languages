#include <stdio.h>
#include <stdlib.h>

void Sort(double list[], int size) {
    for (int i = 0; i < size - 1; i++) {
        for (int j = 0; j < size - i - 1; j++) {
            if (list[j] > list[j + 1]) {
                double temp = list[j];
                list[j] = list[j + 1];
                list[j + 1] = temp;
            }
        }
    }
}

int main() {
    int elements_of_array;
    printf("How many elements in the list: ");
    scanf("%d", &elements_of_array);

    double *prices = (double *)malloc(elements_of_array * sizeof(double));
    if (prices == NULL) {
        printf("Memory allocation failed.\n");
        return 1;
    }

    for (int j = 0; j < elements_of_array; j++) {
        printf("What is the %d element of the array: ", j + 1);
        if (scanf("%lf", &prices[j]) != 1) {
            printf("Invalid input. Exiting.\n");
            free(prices);
            return 1;
        }
    }

    Sort(prices, elements_of_array);

    printf("Sorted array:\n");
    for (int i = 0; i < elements_of_array; i++) {
        printf("%.2lf ", prices[i]);
    }
    printf("\n");

    free(prices);

    return 0;
}
