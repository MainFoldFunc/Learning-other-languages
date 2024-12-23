using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp21
{
    internal class Student
    {
        public string name;
        public int age;
        public string Class;

        public Student(string aname, int aage, string awhatclass)
        {
            aname = name;
            aage = age;
            awhatclass = Class;
        }
        public bool agee()
        {
            if (age >= 16)
            {
                return true;
            }
            return false;
        }
    }
}
