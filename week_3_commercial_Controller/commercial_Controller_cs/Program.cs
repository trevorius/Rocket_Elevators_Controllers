using System;
using System.ComponentModel.Design;
using System.Threading;

namespace commercial_Controller
{
    class Program
    {
        static void Main(string[] args)
        {

            Console.WriteLine("Building configuration for the test scenarios");

            Battery testBattery = new Battery(66, 6, 5, 1, 1, 8, 18);

            Thread.Sleep(500);

            Console.WriteLine("The building floors are as Follows : ");
            foreach (FloorRequestButton button in testBattery.FloorRequestButtonList)
            {
                Console.Write(" {0} ", button.name);
            }

            Scenario.setup1(testBattery);
            Scenario.setup2(testBattery);
            Scenario.setup3(testBattery);
            Scenario.setup4(testBattery);

            string selection = "";

            while (selection != "exit")
            {
                Console.WriteLine("\n\n\n");    
                Console.WriteLine("Select scenario to run 1 - 2 - 3 - 4 - Drill - exit");
                selection = Console.ReadLine();

            
                if(selection == "1")
                {
                    Scenario.scenario1(testBattery);

                }
                else if(selection == "2")
                {
                    Scenario.scenario2(testBattery);
                }
                else if (selection == "3")
                {
                    Scenario.scenario3(testBattery);
                }
                else if (selection == "4")
                {
                    Scenario.scenario4(testBattery);
                }
                else if (selection == "Drill")
                {
                    
                    Scenario.Drill(testBattery);
                }
            }

            Console.WriteLine("\n");
            Console.WriteLine("To demonstrate reusability of the project i will build a new building with the following stats : ");
            Console.WriteLine(" total Floors  : 109, basements : 23, elevators per column : 6, Lobby number is 0  ");



            Battery reusability = new Battery(109, 23, 6, 0, 0, 8, 18);

            Console.WriteLine("The  new building floors are as Follows : ");
            foreach (FloorRequestButton button in reusability.FloorRequestButtonList)
            {
                Console.Write(" {0} ", button.name);
            }

            Console.WriteLine("\n");

            foreach (Column column in reusability.columnList)
            {

                Console.WriteLine("Column {0} serves floors from :{1} to {2} and has {3} elevators.", column.nameLetter, column.floorsServed[0], column.floorsServed[column.floorsServed.Count - 2], column.elevatorList.Count);
            }

            Console.WriteLine("press any key to exit");
            Console.ReadLine();
        }
        

    }
}
