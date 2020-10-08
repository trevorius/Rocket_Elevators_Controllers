using System;
namespace commercial_Controller
{
    public class InterfaceDisplay
    {
        public string goTo;
        public char gotoColumn;
        public int gotoElevator;
        public int gotoRequest;
        public string floorServedBy;
        public string message;

        public InterfaceDisplay()
        {
            goTo = "selected elevator";
            floorServedBy = "selected column";
        }
        public void display(string msg)
        {
            Console.WriteLine("interface Displays : {0}", msg);
        }
        public void displaygoto()
        {
            Console.WriteLine("interface displays : ");
            Console.WriteLine("xxxxxxxxxxxxxxxxxxxxxx Go to column {0}, elevator {1} to go to floor : {2} xxxxxxxxxxxxxxxxxxxxxxxxxx", gotoColumn, gotoElevator, gotoRequest);
        }

    }

    public class FloorDisplay
    {
        public int number;
        public string direction;
        public string message;

        public FloorDisplay()
        {
            number = 0;
            direction = "";
            message = "";
        }

        public void noProblem()
        {
            Console.WriteLine("shaft floorDisplays show : {0} going {1}", number, direction);
        }
        public void messageDisplay()
        {
            Console.WriteLine("floor display shows : ");
            Console.WriteLine(message);
        }
    }
}
