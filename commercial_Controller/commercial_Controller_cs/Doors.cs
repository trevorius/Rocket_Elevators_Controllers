using System;
using System.Threading;
namespace commercial_Controller
{
    public class Doors
    {
        public bool open ;
        public int openTime ;
        public bool safeToClose ;
        public bool passengerDetector;

        public Doors()
        {
            open = false;
            openTime = 2;
            safeToClose = true;
            passengerDetector = false;
        }
        public void action()
        {
            open = true;
            Console.WriteLine("open Doors");
            Thread.Sleep(openTime * 1000);
            checkIfSafeToClose();
            Console.WriteLine("Close Doors");
            open = false;

        }
        public void checkIfSafeToClose()
        {
            while (passengerDetector)
            {
                safeToClose = false;
                Thread.Sleep(openTime * 1000);
            }
           
                safeToClose = true;
            


        }
    }
}

