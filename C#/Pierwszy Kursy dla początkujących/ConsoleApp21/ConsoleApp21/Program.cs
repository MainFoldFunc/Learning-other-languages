using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp21
{
    class Program
    {
        static void Main(string[] args)
        {
            Student student1 = new Student("Jake", 20, "11C");

            Student student2 = new Student("Micheal", 15, "8A");

            Console.WriteLine(student1.agee());

            Console.WriteLine(student2.agee());

            Console.ReadLine();
        }

        
    }
}