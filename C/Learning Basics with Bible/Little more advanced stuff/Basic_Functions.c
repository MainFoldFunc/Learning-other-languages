# include <stdio.h>

void Hello(char name[], int age){
    printf("Hello %s you are %d years old.\n", name, age);
}

int main(){
    char name[25];
    int age;
    printf("What is your name: ");
    scanf("%s", &name);
    printf("What is your age: ");
    scanf("%d", &age);
    Hello(name, age);

    return 0;
}