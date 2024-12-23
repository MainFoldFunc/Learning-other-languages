using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp10
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.Write("What is your favorite day (int): ");
            int num_tyg = Convert.ToInt32(Console.ReadLine());
            Console.WriteLine(method(num_tyg));
            Console.ReadLine();
        }
        static string method(int daynum)
        {
            string dayname;

            switch (daynum)
            {
                case 1:
                    dayname = "Monday";
                    break;
                case 2:
                    dayname = "Tuesday";
                    break;
                case 3:
                    dayname = "Wendsday";
                    break;
                case 4:
                    dayname = "Thursday";
                    break;
                case 5:
                    dayname = "Friday";
                    break;
                case 6:
                    dayname = "Sturday";
                    break;
                case 7:
                    dayname = "Sunday";
                    break;
                default:
                    dayname = "Invalid number of the day";
                    break;
                
            }

            return dayname;
        }
    }
}