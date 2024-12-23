using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp1
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.Write("Whats your name?: ");
            string Name = Console.ReadLine();
            int Age = 30;
            double Best_Number = 8.52;
            char Favorite_Letter = 'A';
            bool true_false = true;
            Console.WriteLine("Hi my name is " + Name + " \nand i am " + Age + "\nMy favorite number is " + Best_Number + " \nAnd my Favorite letter is "+ Favorite_Letter);
            Console.WriteLine("And my name is " +Name.Length + " Letters long");
            Console.WriteLine("And this is my name all capital letters " + Name.ToUpper());
            Console.WriteLine("And this is my name all small letters " + Name.ToLower());
            Console.WriteLine("The first letter of my name is " + Name[0]);
            Console.WriteLine(Name.Contains("John"));
            Console.WriteLine("My age plus my favourite number is " + (Age + Best_Number));
            Console.WriteLine("My age to the power of my favourite number is: ");
            Console.WriteLine(Math.Pow(Age, Best_Number));
            Console.WriteLine("Square root of my age is: ");
            Console.WriteLine(Math.Sqrt(Age) );
            Console.WriteLine("My Favorite number rounded is: ");
            Console.WriteLine(Math.Round(Best_Number));




            Console.ReadLine();

        }
    }

}
