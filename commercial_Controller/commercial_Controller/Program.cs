using System;
using System.ComponentModel.Design;
using System.Threading;

namespace commercial_Controller
{
    class Program
    {
        static void Main(string[] args)
        {
            Battery testBattery = new Battery(66, 6, 5, 1, 1, 8, 18);

            Console.WriteLine("Building configuration for the test scenarios");
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
                Console.WriteLine("\n.\n.\n.");    
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
                    testBattery.pullAlarm();
                    Scenario.scenario1(testBattery);
                }
            }





            /*
            Console.WriteLine(testBattery.stories);
            //Console.WriteLine(testBattery.interfaceDisplay.goTo);
            //Console.WriteLine(testBattery.columnList[0]);
            //Console.WriteLine("column : {0}",testBattery.columnList[0].name.ToString());
            
             foreach (FloorRequestButton button in testBattery.FloorRequestButtonList)
            {
                Console.Write(" {0} ", button.name);
            }

            Console.WriteLine("  .  ");

            foreach (Column element in testBattery.columnList )
            {
                foreach (int i in element.floorsServed)
                {
                    Console.WriteLine(" column{1} serves : {0}", i.ToString(), element.name);
                }
            }
            //scenario 3
            testBattery.columnList[3].elevatorList[0].floorNumber = 58;
            testBattery.columnList[3].elevatorList[0].movement = "DOWN";
            testBattery.columnList[3].elevatorList[0].destinationFloor = 1;
            testBattery.columnList[3].elevatorList[1].floorNumber = 50;
            testBattery.columnList[3].elevatorList[1].movement = "UP";
            testBattery.columnList[3].elevatorList[1].destinationFloor = 60;
            testBattery.columnList[3].elevatorList[2].floorNumber = 46;
            testBattery.columnList[3].elevatorList[2].movement = "UP";
            testBattery.columnList[3].elevatorList[2].destinationFloor = 58;
            testBattery.columnList[3].elevatorList[3].floorNumber = 1;
            testBattery.columnList[3].elevatorList[3].movement = "UP";
            testBattery.columnList[3].elevatorList[3].destinationFloor = 54;
            testBattery.columnList[3].elevatorList[4].floorNumber = 60;
            testBattery.columnList[3].elevatorList[4].movement = "DOWN";
            testBattery.columnList[3].elevatorList[4].destinationFloor = 1;

            testBattery.columnList[3].RequestElevator(54);

            //scenario 4

            testBattery.columnList[0].elevatorList[0].floorNumber = -4;
            testBattery.columnList[0].elevatorList[0].movement = "IDLE";
            //testBattery.columnList[0].elevatorList[0].destinationFloor = null;
            testBattery.columnList[0].elevatorList[1].floorNumber = 1;
            testBattery.columnList[0].elevatorList[1].movement = "IDLE";
            //testBattery.columnList[0].elevatorList[1].destinationFloor = null;
            testBattery.columnList[0].elevatorList[2].floorNumber = -3;
            testBattery.columnList[0].elevatorList[2].movement = "DOWN";
            testBattery.columnList[0].elevatorList[2].destinationFloor = -5;
            testBattery.columnList[0].elevatorList[3].floorNumber = -6;
            testBattery.columnList[0].elevatorList[3].movement = "UP";
            testBattery.columnList[0].elevatorList[3].destinationFloor = 1;
            testBattery.columnList[0].elevatorList[4].floorNumber = -1;
            testBattery.columnList[0].elevatorList[4].movement = "DOWN";
            testBattery.columnList[0].elevatorList[4].destinationFloor = -6;

            testBattery.columnList[0].RequestElevator(-3);

            */


        }

    }
}
