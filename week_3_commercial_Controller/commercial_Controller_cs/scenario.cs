using System;
namespace commercial_Controller
{
    public static class Scenario
    {
        
        public static void setup1(Battery battery)
        {
            battery.columnList[1].elevatorList[0].floorNumber = 20;
            battery.columnList[1].elevatorList[0].movement = "DOWN";
            battery.columnList[1].elevatorList[0].destinationFloor = 5;
            battery.columnList[1].elevatorList[0].requestList.Add(battery.columnList[1].elevatorList[0].destinationFloor);

            battery.columnList[1].elevatorList[1].floorNumber = 3;
            battery.columnList[1].elevatorList[1].movement = "UP";
            battery.columnList[1].elevatorList[1].destinationFloor = 15;
            battery.columnList[1].elevatorList[1].requestList.Add(battery.columnList[1].elevatorList[1].destinationFloor);


            battery.columnList[1].elevatorList[2].floorNumber =13;
            battery.columnList[1].elevatorList[2].movement = "DOWN";
            battery.columnList[1].elevatorList[2].destinationFloor = 1;
            battery.columnList[1].elevatorList[2].requestList.Add(battery.columnList[1].elevatorList[2].destinationFloor);


            battery.columnList[1].elevatorList[3].floorNumber = 15;
            battery.columnList[1].elevatorList[3].movement = "DOWN";
            battery.columnList[1].elevatorList[3].destinationFloor = 2;
            battery.columnList[1].elevatorList[3].requestList.Add(battery.columnList[1].elevatorList[3].destinationFloor);


            battery.columnList[1].elevatorList[4].floorNumber = 6;
            battery.columnList[1].elevatorList[4].movement = "DOWN";
            battery.columnList[1].elevatorList[4].destinationFloor = 1;
            battery.columnList[1].elevatorList[4].requestList.Add(battery.columnList[1].elevatorList[4].destinationFloor);


        }
        public static void setup2(Battery battery)
        {
            battery.columnList[2].elevatorList[0].floorNumber = 1;
            battery.columnList[2].elevatorList[0].movement = "IDLE";
            battery.columnList[2].elevatorList[0].destinationFloor = 21;
            battery.columnList[2].elevatorList[0].requestList.Add(battery.columnList[2].elevatorList[0].destinationFloor);


            battery.columnList[2].elevatorList[1].floorNumber = 23;
            battery.columnList[2].elevatorList[1].movement = "UP";
            battery.columnList[2].elevatorList[1].destinationFloor = 28;
            battery.columnList[2].elevatorList[1].requestList.Add(battery.columnList[2].elevatorList[1].destinationFloor);


            battery.columnList[2].elevatorList[2].floorNumber = 33;
            battery.columnList[2].elevatorList[2].movement = "DOWN";
            battery.columnList[2].elevatorList[2].destinationFloor = 1;
            battery.columnList[2].elevatorList[2].requestList.Add(battery.columnList[2].elevatorList[2].destinationFloor);


            battery.columnList[2].elevatorList[3].floorNumber = 40;
            battery.columnList[2].elevatorList[3].movement = "DOWN";
            battery.columnList[2].elevatorList[3].destinationFloor = 24;
            battery.columnList[2].elevatorList[3].requestList.Add(battery.columnList[2].elevatorList[3].destinationFloor);


            battery.columnList[2].elevatorList[4].floorNumber = 39;
            battery.columnList[2].elevatorList[4].movement = "DOWN";
            battery.columnList[2].elevatorList[4].destinationFloor = 1;
            battery.columnList[2].elevatorList[4].requestList.Add(battery.columnList[2].elevatorList[4].destinationFloor);



        }
        public static void setup3(Battery battery)
        {
            battery.columnList[3].elevatorList[0].floorNumber = 58;
            battery.columnList[3].elevatorList[0].movement = "DOWN";
            battery.columnList[3].elevatorList[0].destinationFloor = 1;
            battery.columnList[3].elevatorList[0].requestList.Add(battery.columnList[3].elevatorList[0].destinationFloor);


            battery.columnList[3].elevatorList[1].floorNumber = 50;
            battery.columnList[3].elevatorList[1].movement = "UP";
            battery.columnList[3].elevatorList[1].destinationFloor = 60;
            battery.columnList[3].elevatorList[1].requestList.Add(battery.columnList[3].elevatorList[1].destinationFloor);


            battery.columnList[3].elevatorList[2].floorNumber = 46;
            battery.columnList[3].elevatorList[2].movement = "UP";
            battery.columnList[3].elevatorList[2].destinationFloor = 58;
            battery.columnList[3].elevatorList[2].requestList.Add(battery.columnList[3].elevatorList[2].destinationFloor);


            battery.columnList[3].elevatorList[3].floorNumber = 1;
            battery.columnList[3].elevatorList[3].movement = "UP";
            battery.columnList[3].elevatorList[3].destinationFloor = 54;
            battery.columnList[3].elevatorList[3].requestList.Add(battery.columnList[3].elevatorList[3].destinationFloor);


            battery.columnList[3].elevatorList[4].floorNumber = 60;
            battery.columnList[3].elevatorList[4].movement = "DOWN";
            battery.columnList[3].elevatorList[4].destinationFloor = 1;
            battery.columnList[3].elevatorList[4].requestList.Add(battery.columnList[3].elevatorList[4].destinationFloor);



        }
        public static void setup4(Battery battery)
        {
            battery.columnList[0].elevatorList[0].floorNumber = -4;
            battery.columnList[0].elevatorList[0].movement = "IDLE";
            //battery.columnList[0].elevatorList[0].destinationFloor = null;

            battery.columnList[0].elevatorList[1].floorNumber = 1;
            battery.columnList[0].elevatorList[1].movement = "IDLE";
            //battery.columnList[0].elevatorList[1].destinationFloor = null;

            battery.columnList[0].elevatorList[2].floorNumber = -3;
            battery.columnList[0].elevatorList[2].movement = "DOWN";
            battery.columnList[0].elevatorList[2].destinationFloor = -5;
            battery.columnList[0].elevatorList[2].requestList.Add(battery.columnList[0].elevatorList[2].destinationFloor);


            battery.columnList[0].elevatorList[3].floorNumber = -6;
            battery.columnList[0].elevatorList[3].movement = "UP";
            battery.columnList[0].elevatorList[3].destinationFloor = 1;
            battery.columnList[0].elevatorList[3].requestList.Add(battery.columnList[0].elevatorList[3].destinationFloor);



            battery.columnList[0].elevatorList[3].load = 100000000000000000;


            battery.columnList[0].elevatorList[4].floorNumber = -1;
            battery.columnList[0].elevatorList[4].movement = "DOWN";
            battery.columnList[0].elevatorList[4].destinationFloor = -6;
            battery.columnList[0].elevatorList[4].requestList.Add(battery.columnList[0].elevatorList[4].destinationFloor);


        }

        public static void scenario1(Battery battery)
        {
            setup1(battery);
            battery.AssignElevator(20);
            Console.WriteLine("\n\n\n");
            Console.WriteLine("column {0} elevator {1} was selected", battery.SelectedColumn.nameLetter, battery.SelectedElevator.name);
            Console.WriteLine("elevator B5 was expected to be sent...");
            Console.WriteLine("\n\n\n");
        }
        public static void scenario2(Battery battery)
        {
            setup2(battery);
            battery.AssignElevator(36);
            Console.WriteLine("\n\n\n");
            Console.WriteLine("column {0} elevator {1} was selected", battery.SelectedColumn.nameLetter, battery.SelectedElevator.name);
            Console.WriteLine("elevator C1 was expected to be sent...");
            Console.WriteLine("\n\n\n");
        }
        public static void scenario3(Battery battery)
        {
            setup3(battery);
            battery.columnList[3].RequestElevator(54);
            Console.WriteLine("\n\n\n");
            Console.WriteLine("column {0} elevator {1} was selected", battery.columnList[3].nameLetter, battery.columnList[3].selectedElevator.name);
            Console.WriteLine("elevator D1 was expected to be sent...");
            Console.WriteLine("\n\n\n");
        }
        public static void scenario4(Battery battery)
        {

            setup4(battery);
            battery.columnList[0].RequestElevator(-3);

            Console.WriteLine("\n\n\n");
            Console.WriteLine("column {0} elevator {1} was selected", battery.columnList[0].nameLetter, battery.columnList[0].selectedElevator.name);
            Console.WriteLine("elevator A4 was expected to be sent...");
            Console.WriteLine("\n\n\n");
        }
        public static void Drill(Battery battery)
        {
            battery.pullAlarm();
            setupDrill(battery);
            battery.columnList[0].RequestElevator(-3);


            Console.WriteLine("\n\n\n");
            Console.WriteLine("Drill calls the alarm method thereby toggleing alarm status for the building, which in turn will toggle online status.");
            Console.WriteLine(" I tried to run Request 3rd basement.. if alarm's were on you would have got an error message.\n If not an elevator should have taken you to floor -3...");
            Console.WriteLine("\n\n\n");

        }

        public static void setupDrill(Battery battery)
        {
            battery.columnList[0].elevatorList[0].floorNumber = -4;
            battery.columnList[0].elevatorList[0].movement = "DOWN";
            battery.columnList[0].elevatorList[0].destinationFloor = -5;
            battery.columnList[0].elevatorList[0].requestList.Add(battery.columnList[0].elevatorList[0].destinationFloor );

            battery.columnList[0].elevatorList[1].floorNumber = -5;
            battery.columnList[0].elevatorList[1].movement = "DOWN";
            battery.columnList[0].elevatorList[1].destinationFloor = -6;
            battery.columnList[0].elevatorList[1].requestList.Add(battery.columnList[0].elevatorList[1].destinationFloor);


            battery.columnList[0].elevatorList[2].floorNumber = -4;
            battery.columnList[0].elevatorList[2].movement = "DOWN";
            battery.columnList[0].elevatorList[2].destinationFloor = -6;
            battery.columnList[0].elevatorList[2].requestList.Add(battery.columnList[0].elevatorList[2].destinationFloor);


            battery.columnList[0].elevatorList[3].floorNumber = -1;
            battery.columnList[0].elevatorList[3].movement = "DOWN";
            battery.columnList[0].elevatorList[3].destinationFloor = -6;
            battery.columnList[0].elevatorList[3].requestList.Add(battery.columnList[0].elevatorList[3].destinationFloor);


            battery.columnList[0].elevatorList[4].floorNumber = -2;
            battery.columnList[0].elevatorList[4].movement = "DOWN";
            battery.columnList[0].elevatorList[4].destinationFloor = -5;
            battery.columnList[0].elevatorList[4].requestList.Add(battery.columnList[0].elevatorList[4].destinationFloor);



        }
    }
}
