using System;
using System.Collections.Generic;

namespace commercial_Controller
{
    public class Column
    {
        public int name;
        public List<int> floorsServed = new List<int>();
        
        public Column(int aName, List<int> afloorsServed, int elevatorsPerCollumn)
        {
            name = aName;
            floorsServed = afloorsServed;
        }
    }
}
