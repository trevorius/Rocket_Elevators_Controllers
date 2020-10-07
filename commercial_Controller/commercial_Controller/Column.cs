using System;
using System.Collections.Generic;
using System.Linq;
using System.Linq.Expressions;
using System.Runtime.CompilerServices;
using System.Security.Cryptography.X509Certificates;
using System.Timers;
using System.Threading;

namespace commercial_Controller
{
    public class Column
    {
        public int name;
        public int floors;
        public int numberOfElevators;
        public bool online;
        public string toBase;
        public int Base;

        public List<CallButton> callbuttonList = new List<CallButton>();
        public List<Elevator> elevatorList = new List<Elevator>();
        public List<Elevator> onlineElevatorList = new List<Elevator>();

        public List<int> floorsServed = new List<int>();
        public List<int> callList = new List<int>();


        public Elevator selectedElevator;

        public char nameLetter;






        public Column(int aName, List<int> afloorsServed, int elevatorsPerCollumn, int aBase)
        {
            name = aName;
            floorsServed = afloorsServed;
            numberOfElevators = elevatorsPerCollumn;
            floors = floorsServed.Count;
            online = true;
            Base = aBase;

            nameLetter = (char)(name + 64);

            calculateToBase();
            createElevatorList();
            createCallButtonList();
            //createOnlineElevatorList();

        }
        public void RequestElevator(int FloorNumber)
        {
            createOnlineElevatorList();

            callList.Add(FloorNumber);

            if (!checkIfElevatorOnItsWay(FloorNumber))
            {
                sortElevatorByDistance(FloorNumber, onlineElevatorList);

                selectedElevator = null;


                foreach (Elevator elevator in onlineElevatorList)
                {
                    if (elevator.movement == "IDLE")
                    {
                        selectedElevator = elevator;
                    }
                }
                if (selectedElevator == null)
                {
                    selectedElevator = onlineElevatorList[onlineElevatorList.Count - 1];
                    selectedElevator.requestList.Add(FloorNumber);
                }

            }
            Console.WriteLine("xxxxxxxxxxxxxxxxxxxx         elevator {0} was selected       xxxxxxxxxxxxxxxxxxxxx", selectedElevator.name);
            selectedElevator.destinationFloor = Base;
            move(selectedElevator);

        }


        private void calculateToBase()
        {
            if (floorsServed[0] < Base)
            {
                toBase = "UP";
            }
            else
            {
                toBase = "DOWN";
            }
        }

        private void createCallButtonList()
        {
            foreach (int floor in floorsServed)
            {
                CallButton newCallButton = new CallButton(floor, toBase);
                callbuttonList.Add(newCallButton);
            }

        }

        private void createElevatorList()
        {


            int elevatorName = 1;
            while (elevatorName < numberOfElevators + 1)
            {
                Elevator newElevator = new Elevator(elevatorName, toBase);
                elevatorList.Add(newElevator);
                elevatorName++;

            }
        }

        public void createOnlineElevatorList()
        {
            foreach (Elevator elevator in elevatorList)
            {
                if (elevator.online)
                {
                    onlineElevatorList.Add(elevator);
                }
            }

        }

        public bool checkIfElevatorOnItsWay(int floor)
        {

            foreach (Elevator elevator in onlineElevatorList)
            {
                int absFloorNumber = Math.Abs(elevator.floorNumber);
                int distanceToCall = Math.Abs(Math.Abs(elevator.floorNumber) - Math.Abs(floor));
                int distanceToDestination = Math.Abs(Math.Abs(elevator.floorNumber) - Math.Abs(elevator.destinationFloor));

                //if elevators movement takes it through floor
                if (absFloorNumber >= Math.Abs(floor) && elevator.movement == elevator.toBase && distanceToCall <= distanceToDestination)
                {
                    selectedElevator = elevator;
                    return true;
                }
            }
            return false;


        }

        public void sortElevatorByDistance(int floor, List<Elevator> list)
        {

            foreach (Elevator elevator in list)
            {
                elevator.distance = Math.Abs(Math.Abs(elevator.floorNumber) - Math.Abs(floor));

            }
            list.Sort((x, y) => x.distance.CompareTo(y.distance));
        }

        public void move(Elevator elevator)
        {

            Console.WriteLine("column {0} elevator {1} is on floor : {2}", name, elevator.name, elevator.floorNumber);
            if (elevator.floorNumber == elevator.destinationFloor)
            //elevator is at it's current destination, needs to check if it has other destinations lined up.
            {
                //if all calls or requests have not been fullfilled this will catch them
                if (elevator.requestList.Count != 0)
                {
                    elevator.destinationFloor = elevator.requestList[0];
                }
                else if (callList.Count != 0)
                {
                    elevator.destinationFloor = callList[0];
                }

                //must open doors for potential passengers
                elevator.Doors.action();
            }

            while (elevator.floorNumber != elevator.destinationFloor)

            {

                if (!elevator.isElevatorFull())
                {
                    //sets movement directionof the elevator
                    if (elevator.destinationFloor > elevator.floorNumber)
                    {
                        elevator.movement = "UP";

                        elevator.floorNumber++;
                        if (!floorsServed.Contains(0) && elevator.floorNumber == 0)
                        {
                            elevator.floorNumber++;
                        }


                        //Console.WriteLine("column {0} elevator {1} moved to {2}", name, elevator.name, elevator.floorNumber);

                    }
                    else
                    {
                        elevator.movement = "DOWN";

                        elevator.floorNumber--;
                        if (!floorsServed.Contains(0) && elevator.floorNumber == 0)
                        {
                            elevator.floorNumber--;
                        }

                        //Console.WriteLine("column {0} elevator {1} moved to {2}", name, elevator.name, elevator.floorNumber);
                    }

                    foreach(int floor in floorsServed)
                    {
                        if (floor == elevator.floorNumber)
                        {
                            Thread.Sleep(500);
                        }
                    }
                    elevator.FloorDisplayUpdate(elevator.floorNumber, elevator.movement);


                    bool floorIsInCallList = callList.Contains(elevator.floorNumber);
                    bool floorIsInRequestList = elevator.requestList.Contains(elevator.floorNumber);

                    //if elevator crosses a floor it should stop at (on a call or request list)
                    if ((floorIsInCallList && elevator.movement == elevator.toBase) ||
                        (floorIsInRequestList && elevator.movement != elevator.toBase))
                    {
                        elevator.Doors.action();
                        int indexOfCall = callbuttonList.IndexOf(callbuttonList.Where(call => call.nameint == elevator.floorNumber).FirstOrDefault());
                        callbuttonList[indexOfCall].isPressed = false;
                        callList.Remove(elevator.floorNumber);
                        elevator.requestList.Remove(elevator.floorNumber);
                    }



                    if (elevator.floorNumber == elevator.destinationFloor)                       
                        //elevator has reached it's destination
                    {
                        elevator.Doors.action();
                        //if all calls or requests have not been fullfilled this will catch them
                        if (callList.Count != 0)
                        {
                            elevator.destinationFloor = callList[0];
                        }
                        else if (elevator.requestList.Count != 0)
                        {
                            elevator.destinationFloor = elevator.requestList[0];
                        }
                    }

                }
                else
                {
                    elevator.FloorDisplay.message = "FULL!";
                    elevator.FloorDisplay.messageDisplay();
                    Console.WriteLine("enter a value lower than {0}", Elevator.MAXLOAD);
                    string newLoad = Console.ReadLine();
                    elevator.load = float.Parse(newLoad);
                    elevator.FloorDisplayUpdate(elevator.floorNumber, elevator.movement);


                }

            }
            if (elevator.floorNumber == elevator.destinationFloor)
            {
                elevator.movement = "IDLE";
                //elevator.Doors.action();
                int indexOfCall = callbuttonList.IndexOf(callbuttonList.Where(call => call.nameint == elevator.floorNumber).FirstOrDefault());
                callbuttonList[indexOfCall].isPressed = false;
                goToIdle();

            }
        }  

        private void goToIdle()
        {
            Console.WriteLine("going to Idle on column {0}", name);

            /*
            SET counter to 0
    FOR every elevator of collumn
        IF movement IS false
            THEN
                INCREMENT counter by +1
        ENDIF
    ENDFOR
    IF counter IS 0
        RETURN 'avoid a division by 0 on next line'
    ENDIF SET idleFloor to(floors DIVIDED BY (counter) ROUNDED DOWN)
        SET counter to 1
        FOR every elevator of collumn
            IF movement IS false
                IF counter is 1
                    THEN
                        SET  destinationFloor of elevator of collumn TO Base of collumn of collumnlist of   Battery
                ELSE
                    SET destinationFloor of elevator of collumn  TO(idleFloor MULTIPLIES counter)
                    INCREMENT counter by + 1
                ENDIF
                CALL moveElevator WITH elevator of collumn

            ENDIF
        ENDFOR
    ENDIF
            */

        }
    }
}