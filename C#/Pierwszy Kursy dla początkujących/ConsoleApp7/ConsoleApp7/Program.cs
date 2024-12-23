using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp7
{
    class Program
    {
        static void Main(string[] args)
        {
            while (!false)
            {
                bool ismale = true;
                Console.Write("Are you a male(yes or no ): ");
                string question = Console.ReadLine();
                if (question == "yes")
                {
                    ismale = true;
                }
                else if (question == "no")
                {
                    ismale = false;
                }
                else
                {
                    ismale = false;
                }


                if (ismale)
                {
                    Console.WriteLine("You are male");
                }
                else
                {
                    Console.WriteLine("You are a female");
                }
                Console.ReadLine();
            }

            
        }
    }
}