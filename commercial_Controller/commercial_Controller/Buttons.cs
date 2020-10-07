using System;
using System.Reflection;

namespace commercial_Controller
{

    public class Button
    {
        public string name;
        public bool isPressed = false;

        public Button(string aName)
        {
            name = aName;
        }

    }

    public class FloorRequestButton : Button
    {
        public int nameint;
        public FloorRequestButton(int aName):base (aName.ToString())
        {
            nameint = aName;
        }
    }

    public class CallButton : Button        
    {
        public int nameint;
        public string direction;
        public CallButton(int aName, string aDirection):base(aName.ToString())
        {
            nameint = aName;
            direction = aDirection;

            
        }

    }
 }

