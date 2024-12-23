using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp2
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.Write("Write a number: ");
            double num_1 = Convert.ToDouble(Console.ReadLine());
            Console.Write("Write a second number: ");
            double num_2 = Convert.ToDouble(Console.ReadLine());

            Console.WriteLine("Answer is: " + (num_1 + num_2));


            Console.ReadLine();
        }
    }
}