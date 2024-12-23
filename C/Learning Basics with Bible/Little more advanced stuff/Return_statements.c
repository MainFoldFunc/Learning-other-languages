#include <stdio.h>
float sqare(float a){
    return a * a;
}

int main(){
    int side_of_sqare;
    float area_of_sqare;
    printf("Enter a side of the sqare: ");
    scanf("%d", &side_of_sqare);
    area_of_sqare = sqare(side_of_sqare);
    printf("The area of that sqare is: %.2f", area_of_sqare);

    return 0;
}