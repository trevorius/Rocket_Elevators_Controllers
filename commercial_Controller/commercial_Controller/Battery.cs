using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;

namespace commercial_Controller
{
    public class Battery
    {
        public const int FloorsPerColumn = 20;

        public int stories;
        public int floors;
        public int basements;
        public int baseFloor;
        public int rc;
        public int numberOfCollumns;
        public bool alarm;
        public int numberOfBasementCollumns;
        public int elevatorsPerCollumn;

        public int leavingTime;
        public int arrivingTime;

        public List<Column> columnList = new List<Column>();
        public List<FloorRequestButton> FloorRequestButtonList = new List<FloorRequestButton>();

        public InterfaceDisplay interfaceDisplay;


        //battery constructor
        public Battery(int aFloors, int aBasements, int elevators, int BaseFloorNumber, int RCnumber, int _leavingTime, int _arrivingTime)
        {


            stories = aFloors;
            floors = aFloors - aBasements;
            basements = aBasements;
            baseFloor = BaseFloorNumber;
            rc = RCnumber;
            alarm = false;
            elevatorsPerCollumn = elevators;

            numberOfCollumns = (int)Math.Ceiling((double)floors / (double)Battery.FloorsPerColumn);
            numberOfBasementCollumns = (int)Math.Ceiling((double)basements / (double)Battery.FloorsPerColumn);


            leavingTime = _leavingTime;
            arrivingTime = _arrivingTime;

            interfaceDisplay = new InterfaceDisplay();

            Console.WriteLine("basementcollumns : ");
            Console.WriteLine(numberOfBasementCollumns);

            createColumnList();
            createFloorRequestButtonList();

        }

        private void createColumnList()
        {
            if (numberOfBasementCollumns != 0)
            {
                createBasementColumnList(numberOfBasementCollumns, elevatorsPerCollumn, basements, baseFloor);
            }
            createAboveGroundColumnList(numberOfCollumns, elevatorsPerCollumn, floors, baseFloor);
        }

        private void createBasementColumnList(int basementColumns, int elevatorsPerCollumn, int basements, int baseFloor)
        {
            //calculate number of floors per coloumn for this battery
            int floorsPerColumn = (int)(basements / basementColumns);
            Console.WriteLine("basements Percolumn : {0}", floorsPerColumn);
            int remFloorsPerColumn = basements % basementColumns;
            int columnCounter = 0;
            while (columnCounter < basementColumns)
            {
                List<int> basementsServed = new List<int>();

                if (basements <= Battery.FloorsPerColumn)
                {
                    for (int i = -1; i >= (basements * -1); i--)
                    {
                        basementsServed.Add(i);
                    }

                }
                else
                {   //if there are more basements to serve after this column
                    if ((columnCounter + 1) * floorsPerColumn < basements - remFloorsPerColumn)
                    {
                        for (int i = columnCounter * floorsPerColumn * -1; i >= (columnCounter + 1) * floorsPerColumn * -1; i--)
                        {
                            basementsServed.Add(i);
                        }
                    }
                    else
                    {   //last column of basements
                        for (int i = ((columnCounter * floorsPerColumn) + 1) * -1; i >= (basements * -1); i--)
                        {
                            basementsServed.Add(i);
                        }
                    }
                }
                if (rc != 0)
                {
                    basementsServed.Remove(0);
                }
                basementsServed.Add(baseFloor);
                Column newColumn = new Column((columnCounter + 1 ), basementsServed, elevatorsPerCollumn, baseFloor);
                columnList.Add(newColumn);
                columnCounter++;
            }
        }

        private void createAboveGroundColumnList(int numberOfColumns, int elevatorsPerColumn, int floors, int baseFloor)
        {
            //calculate number of floors per coloumn for this battery
            int floorsPerColumn = (int)(floors / numberOfColumns);
            Console.WriteLine("floorsPercolumn : {0}", floorsPerColumn);
            int remFloorsPerColumn = floors % numberOfColumns;
            int columnCounter = 0;
            while (columnCounter < numberOfColumns)
            {
                List<int> floorsServed = new List<int>();

                if (floors <= Battery.FloorsPerColumn)
                {
                    for (int i = 1; i <= basements; i++)
                    {
                        floorsServed.Add(i);
                    }

                }
                else
                {   //if there are more floors to serve after this column
                    if ((columnCounter + 1) * floorsPerColumn < floors - remFloorsPerColumn)
                    {
                        for (int i = columnCounter * floorsPerColumn +1; i <= (columnCounter + 1) * floorsPerColumn; i++)
                        {
                            floorsServed.Add(i);
                        }
                    }
                    else
                    {   //last column of floors
                        for (int i = (columnCounter * floorsPerColumn) + 1; i <= floors; i++)
                        {
                            floorsServed.Add(i);
                        }
                    }
                }
                Console.WriteLine("RC is : {0}", rc);
                if (rc != 0)
                {
                    floorsServed.Remove(0);
                }
                else
                {
                    if (columnCounter == numberOfCollumns - 1)
                    {
                        floorsServed.RemoveAt(floorsPerColumn-1);
                    }
                }
                floorsServed.Remove(baseFloor);
                floorsServed.Add(baseFloor);
                Column newColumn = new Column((columnCounter + 1 + numberOfBasementCollumns), floorsServed, elevatorsPerCollumn, baseFloor);
                columnList.Add(newColumn);
                columnCounter++;
                               
            }
        }

        private void createFloorRequestButtonList()
        {
            for (int i = basements * -1 ; i <= floors ; i++ )
            {
                FloorRequestButton newFloorRequestButton = new FloorRequestButton(i);
                FloorRequestButtonList.Add(newFloorRequestButton);
                
            }
            if (rc != 0)
            {
                var RC = FloorRequestButtonList.Find(x => x.nameint == 0);
                FloorRequestButtonList.Remove(RC);
            }else
            {
                FloorRequestButtonList.RemoveAt(stories);
            }

        }

        public void pullAlarm()
        {
            alarm = !alarm;
            if (alarm)
            {
                Console.WriteLine("WARNING! WARNING! WARNING! WARNING! ALARMS ARE RINGING WARNING! WARNING! WARNING! WARNING! WARNING! ");
            }
            foreach (Column column in columnList)
            {
                column.online = !alarm;
                foreach (Elevator elevator in column.elevatorList)
                {
                    elevator.online = !alarm;
                }
            }
                
                
        }

        public Column selectColumn(int floor)
        {
            foreach(Column column in columnList)
            {
                if (column.floorsServed.Contains(floor))
                {
                    return column;
                }
            }
            return null;
        }

        public void AssignElevator(int RequestedFloor)

        {
            Elevator selectedElevator = null;
            int indexOfRequestedFloor = FloorRequestButtonList.IndexOf(FloorRequestButtonList.Where(button => button.nameint == RequestedFloor).FirstOrDefault());
            FloorRequestButtonList[indexOfRequestedFloor].isPressed = true;

            Column selectedColumn = null;
            selectedColumn = selectColumn(RequestedFloor);

            if (selectedColumn.online)
            {

                selectedColumn.createOnlineElevatorList();

                selectedColumn.sortElevatorByDistance(baseFloor, selectedColumn.onlineElevatorList);

                foreach (Elevator elevator in selectedColumn.onlineElevatorList)
                {
                    if (elevator.distance == 0 && elevator.movement == "IDLE")
                    {
                        selectedElevator = elevator;
                    }
                }
                if (selectedColumn.checkIfElevatorOnItsWay(baseFloor)) 
                {
                    selectedElevator = selectedColumn.selectedElevator;
                }
                foreach (Elevator elevator in selectedColumn.onlineElevatorList)
                {
                    if (elevator.movement == "IDLE")
                    {
                        selectedElevator = elevator;
                    }
                }

                if (selectedElevator == null)//all elevators are moving away from basefloor we will selectthe one whose last request is the smallest value.
                { 
                    selectedColumn.onlineElevatorList.Sort((x, y) => x.requestList[x.requestList.Count-1].CompareTo(y.requestList[y.requestList.Count-1]));
                    selectedElevator = selectedColumn.onlineElevatorList[0];
                }
                selectedElevator.requestList.Add(RequestedFloor);

                if (selectedElevator.toBase == "UP")
                {
                    selectedElevator.requestList.Sort((a, b) => b.CompareTo(a));// sortdescending
                }
                else
                {
                    selectedElevator.requestList.Sort((a, b) => a.CompareTo(b));//sort ascending
                }
                selectedElevator.destinationFloor = selectedElevator.requestList[selectedElevator.requestList.Count - 1];
                selectedColumn.move(selectedElevator);

                interfaceDisplay.goTo = ("xxxxxxxxxxxxxxxxxxxxxx Go to column {0}, elevator {1} to go to floor : {2} xxxxxxxxxxxxxxxxxxxxxxxxxx", selectedColumn.name, selectedElevator.name, RequestedFloor).ToString();
                interfaceDisplay.display(interfaceDisplay.goTo);
            }
            else
            {
                interfaceDisplay.display("SORRY the column to your desired floor is offline.");
            }
                    

        }


    }
}


    
   