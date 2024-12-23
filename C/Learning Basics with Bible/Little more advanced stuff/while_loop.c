#include <stdio.h>
#include <string.h>

int main() {
    char name[25];

    while (1) { // Keep looping until valid input
        printf("What is your name: ");
        fgets(name, 25, stdin);

        // Remove the trailing newline character if present
        size_t len = strlen(name);
        if (len > 0 && name[len - 1] == '\n') {
            name[len - 1] = '\0';
        }

        // Check if the input is empty
        if (strlen(name) > 0) {
            break; // Exit the loop if name is valid
        }

        printf("You did not enter your name. Please try again.\n");
    }

    printf("Hello %s!\n", name);
    return 0;
}
