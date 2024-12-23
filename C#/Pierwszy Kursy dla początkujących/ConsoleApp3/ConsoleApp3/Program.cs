using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp3
{
    class Program
    {
        static void Main(string[] args)
        {
            string color_1, color_2, name_1;
            Console.Write("What color should roses be: ");
            color_1 = Console.ReadLine();
            Console.Write("What coloure should violets be: ");
            color_2 = Console.ReadLine();
            Console.Write("Who do you love: ");
            name_1 = Console.ReadLine();

            Console.WriteLine("Roses are " + color_1);
            Console.WriteLine("Violets are " + color_2);
            Console.WriteLine("I love " + name_1);

            Console.ReadLine();
        }
    }

}