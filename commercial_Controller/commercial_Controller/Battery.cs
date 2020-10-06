using System;
using System.Collections;
using System.Collections.Generic;

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
            //public FloorRequestButton[] FloorRequestButtonList;

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
                Column newColumn = new Column((columnCounter + 1 ), basementsServed, elevatorsPerCollumn);
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
                    {   //last column of basements
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
                Column newColumn = new Column((columnCounter + 1 + numberOfBasementCollumns), floorsServed, elevatorsPerCollumn);
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
                var RC = FloorRequestButtonList.Find(x => x.name == 0);
                FloorRequestButtonList.Remove(RC);
            }else
            {
                FloorRequestButtonList.RemoveAt(stories);
            }

        }
    }
}


    /*
    

    SEQUENCE create_FloorRequestButtonList USING Floors
       SET counter to NEGATIVE basements
        WHILE counter LESS THAN Floors
            SET FloorRequestButton to INSTANTIATE FloorRequestButton WITH counter
            ADD FloorRequestButton to FloorRequestButtonList
        ENDWHILE
    ENDSEQUENCE
    
    

 "    SEQUENCE timeCheck USING leavingTime AND arrivingTime AND time
        CASE time IS GRETATER THAN (arrivingTime - 1hour) OR SMALLER THAN (arrivingTime + 1hour)   
            RETURN 'ARRIVING'
        CASE time IS GRETATER THAN (leavingTime - 1hour) OR SMALLER THAN (leavingTime + 1hour)  
            RETURN 'LEAVING'

        ENDCASE
    ENDSEQUENCE
 "
ENDDEFINE
             
            */
   