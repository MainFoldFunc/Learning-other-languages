using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp14
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.Write("What is going to be powerd (only positive numbers): ");
            int powered = Convert.ToInt32(Console.ReadLine());
            Console.Write("What is going to be powering by (only positive numbers): ");
            int poweringby = Convert.ToInt32(Console.ReadLine());
            Console.WriteLine("Your powered numbers are eaqule to " + power(powered,poweringby));
            Console.ReadLine();
        }

        static int power(int x, int y)
        {
            int result = 1;
            for (int i = 0; i < y; i++)
            {
                result = result * x; 
            }

            return result;
        }
    }
}