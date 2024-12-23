using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp17
{
    class Program
    {
        static void Main(string[] args)
        {
            try
            {
                Console.Write("Enter a number: ");
                int i = Convert.ToInt32(Console.ReadLine());
                Console.Write("Enter a number: ");
                int j = Convert.ToInt32(Console.ReadLine());

                Console.WriteLine(i / j);
            }
            catch (DivideByZeroException)
            {
                Console.WriteLine("You cannot divide by zero");
            }
            catch (FormatException)
            {
                Console.WriteLine("You need to enter en int.");
            }
            
            

            Console.ReadLine();
        }
    }
}