using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp9
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.Write("Podaj liczbę: ");
            double  num_1 = Convert.ToDouble(Console.ReadLine());
            Console.Write("Podaj liczbę: ");
            double num_2 = Convert.ToDouble(Console.ReadLine());
            Console.Write("Podaj znak matematyczny(+, -, *, /):");
            string znak = Console.ReadLine();

            if (znak == "+")
            {
                Console.WriteLine("Twój wynik to: " + (num_1 + num_2));
                Console.ReadLine();
            }
            else if (znak == "-")
            {
                Console.WriteLine("Twój wynik to: " + (num_1 - num_2));
                Console.ReadLine();
            }
            else if(znak == "*")
            {
                Console.WriteLine("Twój wynik to: " + (num_1 * num_2));
                Console.ReadLine();
            }
            else if (znak == "/")
            {
                Console.WriteLine("Twój wynik to: " + (num_1 / num_2));
                Console.ReadLine();
            }
            else
            {
                Console.WriteLine("Podałeś zły znak!");
                Console.ReadLine();
            }
        }
    }
}