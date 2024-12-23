# include <stdio.h>
# include <string.h>
int main(){
    char cars[][10] = {"Lambo", "Suzuki", "Tesla"};
    // strcpy(cars[0], "Ferrari"); //
    for (int i = 0; i < sizeof(cars) / sizeof(cars[0]); i++){
        printf("%s\n", cars[i]);
    }
    return 0;
}