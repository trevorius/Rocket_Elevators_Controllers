using System;

namespace commercial_Controller
{
    class Program
    {
        static void Main(string[] args)
        {
            Battery testBattery = new Battery(66, 21, 5, 1, 1, 8, 18);

            Console.WriteLine(testBattery.stories);
            //Console.WriteLine(testBattery.interfaceDisplay.goTo);
            //Console.WriteLine(testBattery.columnList[0]);
            Console.WriteLine("column : {0}",testBattery.columnList[0].name.ToString());
            foreach (Column element in testBattery.columnList )
            {
                foreach (int i in element.floorsServed)
                {
                    Console.WriteLine(" column{1} serves : {0}", i.ToString(), element.name);
                }
            }
        }

    }
}
