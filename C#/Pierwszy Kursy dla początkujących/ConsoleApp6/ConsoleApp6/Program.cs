using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp6
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine(method1(3));
            Console.ReadLine();

        }

        static int method1(int num)
        {
            int num_1 = num * num * num;
            return num_1;
        }
    }
}