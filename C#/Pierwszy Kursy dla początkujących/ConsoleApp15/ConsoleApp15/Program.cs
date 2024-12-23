using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp15
{
    class Program
    {
        static void Main(string[] args)
        {
            int[,] a = {
            {1, 2, 3, },
            {4, 5, 6, },
            {6, 8, 9, },
            {10, 11, 12, } };
            Console.WriteLine(a[0,0]);
            Console.ReadLine();
        }
    }
}