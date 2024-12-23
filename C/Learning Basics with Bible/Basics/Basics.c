# include <stdio.h>
# include <stdbool.h>
# include <string.h>
# include <math.h>
# include <ctype.h>
int main(){
    // This is a comment //
    /*
    This is multi line comment
    */
    // Print statements //

    // printf("Hello world!\n"); // \n for new line //
    // printf("I am learning C\n");
    // printf("1\t2\t3\t\n"); // \t for tab the text //
    // printf("\"Hello\"\n\n");

    // Variables //

    // int xy = 4; // Intiger uses 4 byers of memory //
    // int age = 14;
    // float gpa = 4.555555; // Floating point number can store 32 bits of precision //
    // char grade = 'C'; // Stores a letter //
    // char namex[] = "Hello"; // A list of chars. In C there is no string type //
    // double precise = 3.141592653589793; // Doubles can store about 64 bits //
    // bool True_or_False = true; // Boolians stores one byte can only store true or false values(0 or 1) //
    // char int_no_way = 127; // In chars you can store intigers from -128 to 127, 1 byte //
    // unsigned char more_than_127 = 255; // Unsigned means that we don't allow any negatiwe numbers. Stil 1 byte
    // short int shorter = 32000; // Short int can place numbers from -32768 to 32767 2 bytes //
    // unsigned short int longer_but_short = 62000; // No negatives. Still 2 bytes of memory. //
    // unsigned int more = 2147345; // Still for bytes of memory but no negatives //
    // unsigned long long what_is_this_syntax = 43245932759348; // Ok guys who did this 8 bytes //

    // printf("%d\n", xy); // %d is for decimal //
    // printf("%d\n", age);
    // printf("%f\n", gpa); // %f is for float //
    // printf("%c\n", grade); // %c is for character //
    // printf("%s\n", namex); // %s is for string //
    // printf("%lf\n", precise); // %lf for double //
    // printf("%d\n", True_or_False); // %d for Boolians to // 
    // printf("%d, %d\n", int_no_way, more_than_127); // for intiger chars you can use %d or %c //
    // printf("%u", more); // %u for unsigned intigers //
    // printf("%llu\n\n", what_is_this_syntax); // %llu for long long intigers //

    // More format specifiers //

    // float score_1 = 4.56;
    // float score_2 = 3.73; 
    // float score_3 = 3.28; 
    // float score_4 = 4.56; 
    // printf("Score one is: %-8.2f\n", score_1); // %-8 is for aligning text left or right //
    // printf("Score one is: %-8.2f\n", score_2); // .2 is for seting how many decimal points would number have //
    // printf("Score one is: %-8.2f\n", score_3); // f is a type //
    // printf("Score one is: %-8.2f\n\n", score_4);

    // Constants //

    // const float pi = 3.14; // Float that canot be changed //
    // pi = 3 pi cannot be changed //
    // printf("%.2f\n", pi);

    // Arithemtic operators //

    // int x = 5;
    // int y = 2;

    // int z = x + y; // z is eaqule to x + y //
    // int a = x - y; // a is eaqule to x - y //
    // float b = x / y; // b is eaqule to x / y //
    // int c = x * y; // c is eaqule to x * y //
    // int d = y % x; // d is eaqule to rest of the division of x and y //
    // int e = x++; // e is eaqule to x + 1 //
    // int f = y--; // f is eaqule to y - 1 //
    // printf("%d\n", z);
    // printf("%d\n", a);
    // printf("%.2f\n", b);
    // printf("%d\n", c);
    // printf("%d\n", d);
    // printf("%d\n", e);
    // printf("%d\n", f);

    // User input //

    // char name[25] = "";
    // int agex = 34;
    // printf("What is your name?: ");
    //scanf("%s", &name); Scans previous line for input to the space and asigns it to variable //
    // fgets(name, 25, stdin); // Writes a max 25 bytes character to variable name //
    //name[strlen(name) - 1] = "\0";  removes new line character in string(list of chars) //
    // printf("Your name is: %s\nYour age is: %d\n\n", name, agex);

    // Maths //

    // double A = sqrt(9); // Takes the sqare root of 9 //
    // double B = pow(2, 4); // Takes 2 to the power of 4 //
    // int C = round(3.14); // rounds up 3.14 //
    // int D = ceil(3.6751);
    // int E = floor(3.99);
    // int F = fabs(-100); // Tels the absolute value of number //
    // double G = log2(4);
    // double H = sin(45);
    // double I = cos(45);
    // double J = tan(45);
    // printf("%.2lf\n", A);
    // printf("%.2lf\n", B);
    // printf("%d\n", C);
    // printf("%d\n", D);
    // printf("%d\n", E);
    // printf("%d\n", F);
    // printf("%.2lf\n", G);
    // printf("%.2lf\n", H);
    // printf("%.2lf\n", I);
    // printf("%.2lf\n", J);

    // Circle cirfurmance //
    // const double pix = 3.1419;
    // double radius;
    // double circumframe;
    // double area;

    // printf("What is a radius of the circle?: ");
    // scanf("%lf", &radius);
    // circumframe = 2 * pix * radius;
    // area = pow(pix * radius, 2);
    // printf("The cirumfernence is: %lf\n", circumframe);
    // printf("The area of this circle is: %lf", area);

    // Pitagoras //
    // double a;
    // double b;
    // double c;
    // printf("What is the first side of the triangle?: ");
    // scanf("%lf", &a);

    // printf("What is the second side of the triangle?: ");
    // scanf("%lf", &b);
    // c = sqrt(pow(a, 2) + pow(b, 2));
    // printf("The length of the fird side of the triangle is: %.2lf", c);
    // return 0;

    // If statements //

    // int age;
    // printf("What is your age?: ");
    // scanf("%d", &age);

    // if(age > 18){
    //     printf("You are an adult!");
    // }

    // else if (age<18 && age > 0)
    // {
    //     printf("You are a child!");
    // }

    // else{
    //     printf("Your age is not true!");
    // }

    // Swich statement //

    // char grade;
    // printf("What is your grade(A, B, C, etc.)?:");
    // scanf("%c", &grade);
    // switch (grade)
    // {
    //     case 'A':
    //         printf("Your grade is excelent.\n");
    //         break;
    //     case 'B':
    //         printf("Your grade is very good.\n");
    //         break;
    //     case 'C':
    //         printf("Your grade is good.\n");
    //         break;
    //     case 'D':
    //         printf("Your grade is not so good.\n");
    //         break;
    //     case 'E':
    //         printf("Your grade is bad.\n");
    //         break;
    //     case 'F':
    //         printf("Your grade is very bad.\n");
    //         break;
    //     default:
    //         printf("Please enter valid grade.");
    // }

    // Temperature converter //

    // char unit;
    // float temp;

    // printf("What is the temperature: \n");
    // scanf("%f", &temp);

    // printf("Is the temperature in Celsius (C) or Fahrenheit (F): \n");
    // scanf(" %c", &unit);

    // unit = toupper(unit);

    // if (unit == 'C') {
    //     temp = (temp * 9 / 5) + 32;
    //     printf("Temperature in F is: %.2f\n", temp);
    // } else if (unit == 'F') {
    //     temp = ((temp - 32) * 5) / 9;
    //     printf("The temperature in C is: %.1f\n", temp);
    // } else {
    //     printf("It is not a valid unit of temperature.\n");
    // }

    // return 0;

    // Calculator //

    // char sign;
    // double num1;
    // double num2;
    // double resoult;
    // int end;
    // printf("What is a sign of the equation: ");
    // scanf("%c", &sign);
    // printf("What is a num1: ");
    // scanf("%lf", &num1);
    // printf("What is a num2: ");
    // scanf("%lf", &num2);
    // switch (sign)
    // {
    // case '+':
    //     resoult = num1 + num2;
    //     printf("Resoult is: %.2lf\n", resoult);
    //     break;
    // case '-':
    //     resoult = num1 - num2;
    //     printf("Resoult is: %.2lf\n", resoult);
    //     break;
    // case '*':
    //     resoult = num1 * num2;
    //     printf("Resoult is: %.2lf\n", resoult);
    //     break;
    // case '/':
    //     resoult = num1 / num2;
    //     printf("Resoult is: %.2lf\n", resoult);
    //     break;
    
    // default:
    //     printf("There is no such thing.");
    //     break;
    // }
    // printf("Press enter to shut down...");
    // getchar();
    // getchar();

    // return 0;

    // Logical operators //

    // int age1;
    // int age2;

    // printf("\nWhat is your age: ");
    // scanf("%d", &age1);
    // printf("\nWhat is your friends age: ");
    // scanf("%d", &age2);

    // if(age1 >= 18 && age2 >= 18){
    //     printf("\nYou are both adults.");
    // }
    // else if (age1 < 18 && age2 < 18){
    //     printf("\nYou are both children.");
    // }
    // else{
    //     printf("One of you is adult and the second one is child.");
    // }

    // float temp = 10;
    // if (temp <= 10 || temp >= 20){
    //     printf("It is not okey.");
    // }
    // else{
    //     print("It is to cold or to hot.");
    // }
    // if (!true){
    //     printf("How did you done that?");
    // }
    // else{
    //     printf("Normal");
    // }
}