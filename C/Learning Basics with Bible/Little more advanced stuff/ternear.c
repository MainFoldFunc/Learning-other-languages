# include <stdio.h>

int Find_Max(int a,int b){
    return (a > b) ? a : b;
}

int main(){
    int a, b;
    printf("Enter a: ");
    scanf("%d", &a);
    printf("Enter b: ");
    scanf("%d", &b);
    int max = Find_Max(a, b);
    printf("Max of those two values is: %d", max);
    return 0;
}