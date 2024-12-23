#include <stdio.h>

int main() {
    int rows, cols;
    char symbols;

    printf("Number of rows: ");
    scanf("%d", &rows);

    printf("Number of columns: ");
    scanf("%d", &cols);

    printf("Sign: ");
    getchar();
    scanf("%c", &symbols);

    for (int i = 1; i <= rows; i++) {
        for (int j = 1; j <= cols; j++) {
            printf("%c\t", symbols);
        }
        printf("\n");
    }

    return 0;
}
