# include <stdio.h>
void hello(char[], int); // function prototype


int main(){
    char name[] = "Jons";
    int age = 19;
    hello(name, age);
    return 0;
}
void hello(char name[], int age){
    printf("Hello %s and you are %d years old", name);
}