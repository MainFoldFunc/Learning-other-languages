using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp11
{
    class Program
    {
        static void Main(string[] args)
        {
            while (!false)
            {
                int num = 1;
                while (num < 5)
                {
                    Console.WriteLine(num);
                    num++;
                }
                if (num < 5)
                {
                    num = 1;
                }
            }
            

            
        }
    }
}