# include <stdio.h>

int main()
{
    int x = 1;
    int y = 2;
    int temporary;
    temporary= x;
    x = y;
    y = temporary;
    printf("%d\n", x);
    printf("%d\n", y);
    return 0;
}