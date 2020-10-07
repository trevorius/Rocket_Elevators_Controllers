using System;
namespace commercial_Controller
{
    public class InterfaceDisplay
    {
        public string goTo;
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
