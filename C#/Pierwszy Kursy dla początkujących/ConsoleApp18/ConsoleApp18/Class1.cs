using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp18
{
    internal class Class1
    {
        public string title;
        public string autor;
        public int strony;
        public bool book;

        public Class1(string atitle, string aautor, int astrony, bool abook)
        {
            atitle = title;
            aautor = autor;
            astrony = strony;
            abook = book;
            Console.WriteLine("Crating book. the name of the book is: " + title + "\n The author of the book is: " + autor + "\n there are " + strony + "int this book \n");
        }
    }
}
