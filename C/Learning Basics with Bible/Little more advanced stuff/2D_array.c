#include <stdio.h>

int main() {
    int size_rows;
    int size_cols;
    printf("How many rows: ");
    scanf("%d", &size_rows);
    printf("How many cols: ");
    scanf("%d", &size_cols);
    int nums[size_rows][size_cols];
    int rows = sizeof(nums) / sizeof(nums[0]);
    int cols = sizeof(nums[0]) / sizeof(nums[0][0]);

    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            printf("Add element to array at position [%d][%d]: ", i, j);
            scanf("%d", &nums[i][j]);
        }
    }

    printf("Array elements are:\n");
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            printf("%d ", nums[i][j]);
        }
        printf("\n");
    }

    return 0;
}

