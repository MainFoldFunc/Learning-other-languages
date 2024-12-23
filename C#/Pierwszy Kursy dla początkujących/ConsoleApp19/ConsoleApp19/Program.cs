using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp19
{
    class Program
    {
        static void Main(string[] args)
        {
            int run = 1;
            while (run > 0)
            {
                try
                {
                    Console.Write("Enter an int: ");
                    int num1 = Convert.ToInt32(Console.ReadLine());
                    Console.Write("Enter an int: ");
                    int num2 = Convert.ToInt32(Console.ReadLine());
                    Console.Write("Enter a symbole (+, -, *, /): ");
                    string symbole = Console.ReadLine();
                    if (symbole == "+")
                    {
                        Console.WriteLine(plus(num1, num2));
                        Console.Write("Do you want to run program again?");
                        string answer = Console.ReadLine();
                        run--;
                        if (answer == "yes")
                        {
                            run = run + 1;
                        }
                        else
                        {
                            run = 0;
                        }
                    }
                    else if (symbole == "-")
                    {
                        Console.WriteLine(minus(num1, num2));
                        Console.Write("Do you want to run program again?");
                        string answer = Console.ReadLine();
                        run--;
                        if (answer == "yes")
                        {
                            run = run + 1;
                        }
                        else
                        {
                            run = 0;
                        }
                    }
                    else if (symbole == "*")
                    {
                        Console.WriteLine(multi(num1, num2));
                        Console.Write("Do you want to run program again?");
                        string answer = Console.ReadLine();
                        run--;
                        if (answer == "yes")
                        {
                            run = run + 1;
                        }
                        else
                        {
                            run = 0;
                        }
                    }
                    else if (symbole == "/")
                    {
                        Console.WriteLine(division(num1, num2));
                        Console.Write("Do you want to run program again?");
                        string answer = Console.ReadLine();
                        run--;
                        if (answer == "yes")
                        {
                            run = run + 1;
                        }
                        else
                        {
                            run = 0;
                        }
                    }
                    else
                    {
                        Console.WriteLine("There is no such a sign!");
                        Console.Write("Do you want to run program again?");
                        string answer = Console.ReadLine();
                        run--;
                        if (answer == "yes")
                        {
                            run = run + 1;
                        }
                        else
                        {
                            run = 0;
                        }
                    }
                }

                catch (DivideByZeroException)
                {
                    Console.WriteLine("You cannot divied by 0");

                }
                catch (FormatException)
                {
                    Console.WriteLine("You cannot enter a string here");
                }
                catch
                {
                    Console.WriteLine("Unspecified Error");
                }
            }
            
            


        }
        static int plus(int i, int j)
        {
            return i + j;
        }
        static int minus(int i, int j)
        {
            return i - j;
        }
        static int multi(int i, int j)
        {
            return i * j;
        }
        static int division(int i, int j)
        {
            return i / j;
        }

    }
}