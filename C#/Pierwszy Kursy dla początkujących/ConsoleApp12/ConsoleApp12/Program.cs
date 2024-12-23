using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp12
{
    class Program
    {
        static void Main(string[] args)
        {
            string login = "gsliv";
            string guessit = "Hello";
            string playerpassguess = "";
            string loginguest = "";
            int howmanyguessespass = 10;
            int howmanyguessesguess = 10;
            bool outoffpass = false;
            bool outoffguess = false;
            while (loginguest != login && howmanyguessesguess >= 1)
            {
                Console.Write("What is the login?: ");
                loginguest = Console.ReadLine();
                Console.WriteLine("Ok i will help you the length of login is " + login.Length);
                howmanyguessesguess--;
                Console.WriteLine("You have " + howmanyguessesguess + " Guesses");
                if (howmanyguessesguess < 1)
                {
                    outoffguess = true;
                }
                else if (howmanyguessesguess > 1)
                {
                    outoffguess = false;
                }
                if (outoffguess == true)
                {
                    Console.WriteLine("You run out off guesses");
                    break;
                }
            }
            if (outoffguess == true)
            {
                Console.WriteLine("You are looser!");
                
            }
            else
            {
                Console.WriteLine("This is correct login");
            }


            if (outoffguess == false)
            {
                while (guessit != playerpassguess && howmanyguessespass <= 10)
                {
                    Console.Write("What is the password: ");
                    playerpassguess = Console.ReadLine();
                    Console.WriteLine("Ok i will help you the length of password is " + guessit.Length);
                    Console.WriteLine("You have " + howmanyguessespass + " Guesses");
                    howmanyguessespass--;
                    if (howmanyguessespass < 1)
                    {
                        outoffpass = true;
                    }
                    else if (howmanyguessespass > 1)
                    {
                        outoffpass = false;
                    }
                    if (outoffpass == true)
                    {
                        Console.WriteLine("You run out off guesses");
                        break;
                    }
                }
                if (outoffpass == true)
                {
                    Console.WriteLine("You are looser!");
                }
                else
                {
                    Console.WriteLine("This is correct password");
                }
            }
        }
            


    }
}
