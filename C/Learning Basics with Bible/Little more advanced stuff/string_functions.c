# include <stdio.h>
# include <string.h>

int main(){
    char name[] = "Jon";
    char Second_name[] = "Malinkowicz";
    int len;
    strlwr(name);
    printf("%s\n", name);

    strupr(name);
    printf("%s\n", name);

    strcat(name, Second_name);
    printf("%s\n", name);
    printf("%s\n", Second_name);

    strncat(name, Second_name, 2);
    printf("%s\n", name);
    printf("%s\n", Second_name);

    strcpy(name, Second_name);
    printf("%s\n", name);
    printf("%s\n", Second_name);

    strncpy(name, Second_name, 4);
    printf("%s\n", name);
    printf("%s\n", Second_name);
    len = strlen(name);
    print("%d", len);

    return 0;
}