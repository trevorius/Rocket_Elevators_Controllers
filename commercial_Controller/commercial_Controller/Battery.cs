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

        private void createBasementColumnList(int basementColumns,int elevatorsPerCollumn,int basements, int baseFloor)
        {
            int floorsPerColumn = (int)(basements / basementColumns);
            Console.WriteLine("floorsPercolumn : {0}", floorsPerColumn);
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
                    if ((columnCounter + 1) * Battery.FloorsPerColumn < basements) 
                    {
                        for (int i = columnCounter * Battery.FloorsPerColumn * -1; i >= (columnCounter + 1) * Battery.FloorsPerColumn * -1; i--)
                        {
                            basementsServed.Add(i);
                        }
                    }
                    else
                    {   //last column of basements
                        for (int i = ((columnCounter * Battery.FloorsPerColumn) + 1) * -1; i >= (basements * -1); i--) 
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
                Column newColumn = new Column((columnCounter+1), basementsServed, elevatorsPerCollumn);
                columnList.Add(newColumn);
                columnCounter++;
            }
        }

        private void createAboveGroundColumnList(int numberOfCollumns, int elevatorsPerCollumn, int floors, int baseFloor)
        {
            /*
            int columnCounter = 0;
            while (columnCounter < numberOfCollumns)
            {
                if (floors <= )
            }
            
    SEQUENCE create_Collumns USING NumberOfCollumns AND ElevatorsPerCollumn AND Floors And FloorsPerCollumn AND Base
        SET CollumnCounter to 1
        WHILE CollumnCounter LESS THAN collumnlist length

            IF Floors LESS THAN FloorsPerCollumn + 1
                Then SET floorsServed to a list of numbers from 1 to Floors
                    ADD Base to floorsServed
            ELSE SET floorsServed to a list of numbers from CollumnCounter MULTIPLIES FloorsPerCollumn to(CollumnCounter PLUS 1) MULTIPLIES FloorsPerCollumn
                    IF RC NOT 0 THEN REMOVE 0 from floorsServed
                    ADD Base to floorsServed
            ENDIF

            SET CollumnID to NumberOfBasementCollumns + CollumnCounter
            SET Collumn to INSTANTIATE Collumn WITH CollumnID AND floorsServed AND ElevatorsPerCollumn
            ADD Collumn to CollumnList
            INCREMENT CollumnCounter by +1
        ENDWHILE
    ENDSEQUENCE
            */
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
    
    CALL create_BCollumns WITH NumberOfBasementCollumns AND ElevatorsPerCollumn AND Basements  AND FloorsPerCollumn
    CALL create_Collumns WITH NumberOfCollumns AND ElevatorsPerCollumn AND Floors And FloorsPerCollumn AND Base
    CALL create_FloorRequestButtonList WITH Floors

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
        /*
            int columnCounter = 1;
            while (columnCounter <= numberOfCollumns)
            {

                Column newColumn = new Column(columnCounter);
                columnList.Add(newColumn);


                columnCounter ++;
            }

                

            
            */

   