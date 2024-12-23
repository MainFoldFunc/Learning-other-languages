using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp4
{
    class Program
    {
        static void Main(string[] args)
        {
            int[] num_1 = {7, 13, 20, 21}; // Array, its like list in python  //
            string[] maine_foinds = {"Dejmian", "Oljek", "Bżil"};

            num_1[1] = 1; // updating the second element on the array //

            Console.WriteLine(num_1[1]); // Printing the first element in Array //
            Console.ReadLine();

        }
    }
}