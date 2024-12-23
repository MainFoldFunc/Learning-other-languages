using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp5
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.Write("What is your age: ");
            string age = Console.ReadLine();
            Console.Write("What is your name: ");
            string name = Console.ReadLine();
            HelloToUser(name, age); // Calling method with one argument - name //
            Console.ReadLine();
        }

        static void HelloToUser(string name, string age) // Method, its like function in python //
        {
            Console.WriteLine("Hello " + name + " you are " + age);
            Console.ReadLine();
        }
    }
}