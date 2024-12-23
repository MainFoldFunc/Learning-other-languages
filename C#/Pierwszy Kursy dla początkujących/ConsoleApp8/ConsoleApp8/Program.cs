using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp8
{
    class Program
    {
        static void Main(string[] args)
        {

            Method_1(3, 5);
            
            Console.ReadLine();
        }
        static void Method_1(int num_1, int num_2)
        {
            if (num_1 == num_2)
            {
                Console.WriteLine("Num_1 is eaqule to Num_2");
            }
            else if (num_1 > num_2)
            {
                Console.WriteLine("Num_1 is bigger than the Num_2");
            }
            else if (num_1 < num_2)
            {
                Console.WriteLine("Num_2 is bigger than the Num_1");
            }
            
        }
    }

}