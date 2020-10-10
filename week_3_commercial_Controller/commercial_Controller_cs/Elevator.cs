using System;
using System.Collections.Generic;

namespace commercial_Controller
{
    public class Elevator
    {
        public const int MAXLOAD = 10000;


        public int name;
        public int floorNumber;
        public int destinationFloor;
        public int distance;
        public int timer;

        public float load; 

        public bool online;

        public Doors Doors;
        public FloorDisplay FloorDisplay;

        public string movement;
        public string toBase;

        public List<int> requestList = new List<int>();





        public Elevator(int aName, string aToBase)
        {
            name = aName;
            toBase = aToBase;

            floorNumber = 0;            
            timer = 0;
            load = 100;

            online= true;

            Doors = new Doors();
            FloorDisplay = new FloorDisplay();

            movement = "IDLE";

            
        }
        public bool isElevatorFull()
        {
            if(load <= Elevator.MAXLOAD )
            {
                return false;
            }
            else
            {
                return true;
            }
        }

        public void FloorDisplayUpdate(int Number, string direction)
        {
            FloorDisplay.number = Number;
            FloorDisplay.direction = direction;
            FloorDisplay.noProblem();
        }

        
       
        
    }
}
