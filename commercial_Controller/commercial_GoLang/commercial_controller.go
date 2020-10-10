package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"
)

func main() {

	testBattery := new(Battery)
	testBattery.batteryConstructor(66, 6, 1, 1, 5)
	var selection string

	for selection != "leave" {
		fmt.Println("\n select scenario to run 1 - 2 - 3 - 4 - Drill   leave")

		fmt.Scan(&selection)

		if selection == "1" {
			testBattery.scenario1()
		} else if selection == "2" {
			testBattery.scenario2()
		} else if selection == "3" {
			testBattery.scenario3()

		} else if selection == "4" {
			testBattery.scenario4()
		} else if selection == "Drill" {
			testBattery.scenarioDrill()
		} else if selection == "exit" {
			return
		}
	}
	/*
		fmt.Println("To demonstrate reusability i will create a new building with the following stats : ")
		fmt.Println(" stories : 80, basements : 22, rc is 0, 8 elevators per colomn")

		finalBattery := new(Battery)
		finalBattery.batteryConstructor(80, 22, 0, 0, 8)

		fmt.Println("floors served :")
		for i := 0; i < len(finalBattery.floorRequestButtonList); i++ {
			fmt.Printf(" %s", testBattery.floorRequestButtonList[i].name)
		}
	*/

}

//Scenarios
func (battery *Battery) scenario1() {
	battery.setup1()
	battery.AssignElevator(20)
	fmt.Println("\n\n\n")
	fmt.Printf("column %v elevator %v was selected", battery.selectedColumn.nameLetter, battery.selectedElevator.name)
	fmt.Println("elevator B5 was expected to be sent...")
	fmt.Println("\n\n\n")
}
func (battery *Battery) scenario2() {
	battery.setup2()
	battery.AssignElevator(36)
	fmt.Println("\n\n\n")
	fmt.Printf("column %v elevator %v was selected", battery.selectedColumn.nameLetter, battery.selectedElevator.name)
	fmt.Println("elevator C1 was expected to be sent...")
	fmt.Println("\n\n\n")
}
func (battery *Battery) scenario3() {
	battery.setup3()
	battery.columnList[3].RequestElevator(54)
	fmt.Println("\n\n\n")
	fmt.Printf("column %v elevator %v was selected", battery.columnList[3].nameLetter, battery.columnList[3].selectedElevator.name)
	fmt.Println("elevator D1 was expected to be sent...")
	fmt.Println("\n\n\n")
}
func (battery *Battery) scenario4() {
	battery.setup4()
	battery.columnList[0].RequestElevator(-3)

	fmt.Println("\n\n\n")
	fmt.Printf("column %v elevator %v was selected", battery.columnList[0].nameLetter, battery.columnList[0].selectedElevator.name)
	fmt.Println("elevator A4 was expected to be sent...")
	fmt.Println("\n\n\n")
}
func (battery *Battery) scenarioDrill() {
	battery.pullAlarm()
	battery.setupDrill()
	battery.columnList[0].RequestElevator(-3)

	fmt.Println("\n\n\n")
	fmt.Println("Drill calls the alarm method thereby toggleing alarm status for the building, which in turn will toggle online status.")
	fmt.Println(" I tried to run Request 3rd basement.. if alarm's were on you would have got an error message.\n If not an elevator should have taken you to floor -3...")
	fmt.Println("\n\n\n")
}

func (battery *Battery) setup1() {
	battery.columnList[1].elevatorList[0].floorNumber = 20
	battery.columnList[1].elevatorList[0].movement = "DOWN"
	battery.columnList[1].elevatorList[0].destinationFloor = 5
	battery.columnList[1].elevatorList[0].requestList = append(battery.columnList[1].elevatorList[0].requestList, battery.columnList[1].elevatorList[0].destinationFloor)

	battery.columnList[1].elevatorList[1].floorNumber = 3
	battery.columnList[1].elevatorList[1].movement = "UP"
	battery.columnList[1].elevatorList[1].destinationFloor = 15
	battery.columnList[1].elevatorList[1].requestList = append(battery.columnList[1].elevatorList[1].requestList, battery.columnList[1].elevatorList[1].destinationFloor)

	battery.columnList[1].elevatorList[2].floorNumber = 13
	battery.columnList[1].elevatorList[2].movement = "DOWN"
	battery.columnList[1].elevatorList[2].destinationFloor = 1
	battery.columnList[1].elevatorList[2].requestList = append(battery.columnList[1].elevatorList[2].requestList, battery.columnList[1].elevatorList[2].destinationFloor)

	battery.columnList[1].elevatorList[3].floorNumber = 15
	battery.columnList[1].elevatorList[3].movement = "DOWN"
	battery.columnList[1].elevatorList[3].destinationFloor = 2
	battery.columnList[1].elevatorList[3].requestList = append(battery.columnList[1].elevatorList[3].requestList, battery.columnList[1].elevatorList[3].destinationFloor)

	battery.columnList[1].elevatorList[4].floorNumber = 6
	battery.columnList[1].elevatorList[4].movement = "DOWN"
	battery.columnList[1].elevatorList[4].destinationFloor = 1
	battery.columnList[1].elevatorList[4].requestList = append(battery.columnList[1].elevatorList[4].requestList, battery.columnList[1].elevatorList[4].destinationFloor)

}
func (battery *Battery) setup2() {
	battery.columnList[2].elevatorList[0].floorNumber = 1
	battery.columnList[2].elevatorList[0].movement = "IDLE"
	battery.columnList[2].elevatorList[0].destinationFloor = 21
	battery.columnList[2].elevatorList[0].requestList = append(battery.columnList[2].elevatorList[0].requestList, battery.columnList[2].elevatorList[0].destinationFloor)

	battery.columnList[2].elevatorList[1].floorNumber = 23
	battery.columnList[2].elevatorList[1].movement = "UP"
	battery.columnList[2].elevatorList[1].destinationFloor = 28
	battery.columnList[2].elevatorList[1].requestList = append(battery.columnList[2].elevatorList[1].requestList, battery.columnList[2].elevatorList[1].destinationFloor)

	battery.columnList[2].elevatorList[2].floorNumber = 33
	battery.columnList[2].elevatorList[2].movement = "DOWN"
	battery.columnList[2].elevatorList[2].destinationFloor = 1
	battery.columnList[2].elevatorList[2].requestList = append(battery.columnList[2].elevatorList[2].requestList, battery.columnList[2].elevatorList[2].destinationFloor)

	battery.columnList[2].elevatorList[3].floorNumber = 40
	battery.columnList[2].elevatorList[3].movement = "DOWN"
	battery.columnList[2].elevatorList[3].destinationFloor = 24
	battery.columnList[2].elevatorList[3].requestList = append(battery.columnList[2].elevatorList[3].requestList, battery.columnList[2].elevatorList[3].destinationFloor)

	battery.columnList[2].elevatorList[4].floorNumber = 39
	battery.columnList[2].elevatorList[4].movement = "DOWN"
	battery.columnList[2].elevatorList[4].destinationFloor = 1
	battery.columnList[2].elevatorList[4].requestList = append(battery.columnList[2].elevatorList[4].requestList, battery.columnList[2].elevatorList[4].destinationFloor)

}
func (battery *Battery) setup3() {
	battery.columnList[3].elevatorList[0].floorNumber = 58
	battery.columnList[3].elevatorList[0].movement = "DOWN"
	battery.columnList[3].elevatorList[0].destinationFloor = 1
	battery.columnList[3].elevatorList[0].requestList = append(battery.columnList[3].elevatorList[0].requestList, battery.columnList[3].elevatorList[0].destinationFloor)

	battery.columnList[3].elevatorList[1].floorNumber = 50
	battery.columnList[3].elevatorList[1].movement = "UP"
	battery.columnList[3].elevatorList[1].destinationFloor = 60
	battery.columnList[3].elevatorList[1].requestList = append(battery.columnList[3].elevatorList[1].requestList, battery.columnList[3].elevatorList[1].destinationFloor)

	battery.columnList[3].elevatorList[2].floorNumber = 46
	battery.columnList[3].elevatorList[2].movement = "UP"
	battery.columnList[3].elevatorList[2].destinationFloor = 58
	battery.columnList[3].elevatorList[2].requestList = append(battery.columnList[3].elevatorList[2].requestList, battery.columnList[3].elevatorList[2].destinationFloor)

	battery.columnList[3].elevatorList[3].floorNumber = 1
	battery.columnList[3].elevatorList[3].movement = "UP"
	battery.columnList[3].elevatorList[3].destinationFloor = 54
	battery.columnList[3].elevatorList[3].requestList = append(battery.columnList[3].elevatorList[3].requestList, battery.columnList[3].elevatorList[3].destinationFloor)

	battery.columnList[3].elevatorList[4].floorNumber = 60
	battery.columnList[3].elevatorList[4].movement = "DOWN"
	battery.columnList[3].elevatorList[4].destinationFloor = 1
	battery.columnList[3].elevatorList[4].requestList = append(battery.columnList[3].elevatorList[4].requestList, battery.columnList[3].elevatorList[4].destinationFloor)

}
func (battery *Battery) setup4() {
	battery.columnList[0].elevatorList[0].floorNumber = -4
	battery.columnList[0].elevatorList[0].movement = "IDLE"
	//battery.columnList[0].elevatorList[0].destinationFloor = null;

	battery.columnList[0].elevatorList[1].floorNumber = 1
	battery.columnList[0].elevatorList[1].movement = "IDLE"
	//battery.columnList[0].elevatorList[1].destinationFloor = null;

	battery.columnList[0].elevatorList[2].floorNumber = -3
	battery.columnList[0].elevatorList[2].movement = "DOWN"
	battery.columnList[0].elevatorList[2].destinationFloor = -5
	battery.columnList[0].elevatorList[2].requestList = append(battery.columnList[0].elevatorList[2].requestList, battery.columnList[0].elevatorList[2].destinationFloor)

	battery.columnList[0].elevatorList[3].floorNumber = -6
	battery.columnList[0].elevatorList[3].movement = "UP"
	battery.columnList[0].elevatorList[3].destinationFloor = 1
	battery.columnList[0].elevatorList[3].requestList = append(battery.columnList[0].elevatorList[3].requestList, battery.columnList[0].elevatorList[3].destinationFloor)

	battery.columnList[0].elevatorList[3].load = 100000000000000000

	battery.columnList[0].elevatorList[4].floorNumber = -1
	battery.columnList[0].elevatorList[4].movement = "DOWN"
	battery.columnList[0].elevatorList[4].destinationFloor = -6
	battery.columnList[0].elevatorList[4].requestList = append(battery.columnList[0].elevatorList[4].requestList, battery.columnList[0].elevatorList[4].destinationFloor)
}
func (battery *Battery) setupDrill() {
	battery.columnList[0].elevatorList[0].floorNumber = -4
	battery.columnList[0].elevatorList[0].movement = "DOWN"
	battery.columnList[0].elevatorList[0].destinationFloor = -5
	battery.columnList[0].elevatorList[0].requestList = append(battery.columnList[0].elevatorList[0].requestList, battery.columnList[0].elevatorList[0].destinationFloor)

	battery.columnList[0].elevatorList[1].floorNumber = -5
	battery.columnList[0].elevatorList[1].movement = "DOWN"
	battery.columnList[0].elevatorList[1].destinationFloor = -6
	battery.columnList[0].elevatorList[1].requestList = append(battery.columnList[0].elevatorList[1].requestList, battery.columnList[0].elevatorList[1].destinationFloor)

	battery.columnList[0].elevatorList[2].floorNumber = -4
	battery.columnList[0].elevatorList[2].movement = "DOWN"
	battery.columnList[0].elevatorList[2].destinationFloor = -6
	battery.columnList[0].elevatorList[2].requestList = append(battery.columnList[0].elevatorList[2].requestList, battery.columnList[0].elevatorList[2].destinationFloor)

	battery.columnList[0].elevatorList[3].floorNumber = -1
	battery.columnList[0].elevatorList[3].movement = "DOWN"
	battery.columnList[0].elevatorList[3].destinationFloor = -6
	battery.columnList[0].elevatorList[3].requestList = append(battery.columnList[0].elevatorList[3].requestList, battery.columnList[0].elevatorList[3].destinationFloor)

	battery.columnList[0].elevatorList[4].floorNumber = -2
	battery.columnList[0].elevatorList[4].movement = "DOWN"
	battery.columnList[0].elevatorList[4].destinationFloor = -5
	battery.columnList[0].elevatorList[4].requestList = append(battery.columnList[0].elevatorList[4].requestList, battery.columnList[0].elevatorList[4].destinationFloor)

}

func (b *Battery) batteryConstructor(aStories int, aBasements int, aBaseFloor int, aRc int, aElevatorsPerColumn int) {
	b.FloorsPerColumn = 20
	b.stories = aStories
	b.basements = aBasements
	b.rc = aRc
	b.baseFloor = aBaseFloor
	b.elevatorsPerColumn = aElevatorsPerColumn
	b.alarm = false

	b.floors = b.stories - b.basements
	b.numberOfColumns = int(math.Ceil(float64(b.floors) / float64(b.FloorsPerColumn)))
	b.numberOfBasementColumns = int(math.Ceil(float64(b.basements) / float64(b.FloorsPerColumn)))

	b.interfaceDisplay.goTo = "go to floor"

	fmt.Printf("\n there are %v basementColumns. \n", b.numberOfBasementColumns)
	fmt.Printf("\n there are %v columns above ground", b.numberOfColumns)
	fmt.Printf("\n floors : %v basements : %v", b.floors, b.basements)
	//createColumnList()
	b.createFloorRequestList()
	b.createColumnList()
	//will use alarm method to activate the building this may be loud SORRY
	b.pullAlarm()
	b.pullAlarm()
}
func (b *Battery) createColumnList() {
	b.createBasementColumns()
	b.createAboveGroundColumns()

}
func (b *Battery) createAboveGroundColumns() {
	floorsPerColumn := int(b.floors / b.numberOfColumns)
	remFloorsPerColumn := b.floors % b.numberOfColumns
	columnCounter := 0
	for columnCounter < b.numberOfColumns {
		floorsServed := []int{}
		if b.floors <= b.FloorsPerColumn {
			for i := 1; i >= (b.floors); i++ {
				floorsServed = append(floorsServed, i)
			}
		} else { //if there are more floors to serve after this column
			if (columnCounter+1)*floorsPerColumn < b.floors-remFloorsPerColumn {
				for i := columnCounter * floorsPerColumn * +1; i <= (columnCounter+1)*floorsPerColumn; i++ {
					floorsServed = append(floorsServed, i)
				}
			} else { // last column of floors
				for i := ((columnCounter * floorsPerColumn) + 1); i <= (b.floors); i++ {
					floorsServed = append(floorsServed, i)
				}
			}
		}
		if b.rc != 0 {

			for i, other := range floorsServed {
				if other == 0 {
					floorsServed = append(floorsServed[:i], floorsServed[i+1:]...)
				}
			}

		}
		floorsServed = append(floorsServed, b.baseFloor)
		b.columnList = append(b.columnList, Column{name: columnCounter + 1})
		b.columnList[columnCounter+b.numberOfBasementColumns].columnConstructor(columnCounter+1+b.numberOfBasementColumns, floorsServed, b.elevatorsPerColumn, b.baseFloor)
		columnCounter++

		fmt.Println("\n")
		fmt.Println(floorsServed)

	}
}
func (b *Battery) createBasementColumns() {
	floorsPerColumn := int(b.basements / b.numberOfBasementColumns)
	remFloorsPerColumn := b.basements % b.numberOfBasementColumns
	columnCounter := 0
	for columnCounter < b.numberOfBasementColumns {
		basementsServed := []int{}
		if b.basements <= b.FloorsPerColumn {
			for i := -1; i >= (b.basements * -1); i-- {
				basementsServed = append(basementsServed, i)
			}
		} else { //if there are more basements to serve after this column
			if (columnCounter+1)*floorsPerColumn < b.basements-remFloorsPerColumn {
				for i := columnCounter * floorsPerColumn * -1; i >= (columnCounter+1)*floorsPerColumn*-1; i-- {
					basementsServed = append(basementsServed, i)
				}
			} else { // last column of basements
				for i := ((columnCounter * floorsPerColumn) + 1) * -1; i >= (b.basements * -1); i-- {
					basementsServed = append(basementsServed, i)
				}
			}
		}
		if b.rc != 0 {

			for i, other := range basementsServed {
				if other == 0 {
					basementsServed = append(basementsServed[:i], basementsServed[i+1:]...)
				}
			}

		}
		basementsServed = append(basementsServed, b.baseFloor)
		b.columnList = append(b.columnList, Column{name: columnCounter + 1})
		b.columnList[columnCounter].columnConstructor(columnCounter+1, basementsServed, b.elevatorsPerColumn, b.baseFloor)
		columnCounter++

		//fmt.Println("\n createBasementColumn")

	}
}
func (b *Battery) createFloorRequestList() {
	//fmt.Println("\nfloorRequestlist")
	fmt.Println(b.floors)
	for i := b.basements * -1; i <= b.floors; i++ {
		j := strconv.Itoa(i)

		b.floorRequestButtonList = append(b.floorRequestButtonList, FloorRequestButton{name: j, nameint: i})
	}
	if b.rc != 0 {
		tmp1 := b.floorRequestButtonList[:b.basements]
		tmp2 := b.floorRequestButtonList[b.basements+1:]
		b.floorRequestButtonList = append(tmp1, tmp2...)
		fmt.Println("floors served :")
		for i := 0; i < len(b.floorRequestButtonList); i++ {
			fmt.Printf(" %s", b.floorRequestButtonList[i].name)
		}
	}

}
func (b *Battery) pullAlarm() {
	b.alarm = !b.alarm
	if b.alarm {
		fmt.Println("\n WARNING! WARNING! WARNING! WARNING! ALARMS ARE RINGING WARNING! WARNING! WARNING! WARNING! WARNING! ")
	}
	for _, column := range b.columnList {
		column.online = !b.alarm
		for _, elevator := range column.elevatorList {
			elevator.online = !b.alarm
		}
	}
}
func (b *Battery) selectColumn(floor int) Column {
	for _, column := range b.columnList {
		floorIsInFloorsServed := isIntIn(floor, column.floorsServed)

		if floorIsInFloorsServed {
			return column
		}
	}
	columnS := Column{}
	return columnS
}
func (b *Battery) AssignElevator(RequestedFloor int) {
	SelectedElevator := Elevator{}
	var indexOfRequestedFloor int
	if b.rc != 0 {
		indexOfRequestedFloor = RequestedFloor + b.basements - 1
	} else {
		indexOfRequestedFloor = RequestedFloor + b.basements
	}
	b.floorRequestButtonList[indexOfRequestedFloor].isPressed = true

	SelectedColumn := Column{}

	SelectedColumn = b.selectColumn(RequestedFloor)

	if SelectedColumn.online {
		SelectedColumn.createOnlineElevatorList()

		SelectedColumn.sortElevatorByDistance(b.baseFloor, SelectedColumn.onlineElevatorList)

		for _, elevator := range SelectedColumn.onlineElevatorList {
			if elevator.distance == 0 && elevator.movement == "IDLE" {
				SelectedElevator = *elevator
			}
		}
		if SelectedColumn.checkIfElevatorOnItsWay(b.baseFloor) {
			SelectedElevator = *SelectedColumn.selectedElevator
		}
		for _, elevator := range SelectedColumn.onlineElevatorList {
			if elevator.movement == "IDLE" {
				SelectedElevator = *elevator
			}
		}
		if SelectedElevator.movement == "" {
			//all elevators are moving away from basefloor we will selectthe one whose last request is the smallest value.
			fmt.Println("stop for test reasons")
			//SelectedColumn.onlineElevatorList =
			sort.Slice(SelectedColumn.onlineElevatorList, func(i, j int) bool {
				return SelectedColumn.onlineElevatorList[i].requestList[len(SelectedColumn.onlineElevatorList[i].requestList)-1] < SelectedColumn.onlineElevatorList[j].requestList[len(SelectedColumn.onlineElevatorList[j].requestList)-1]
			})

			SelectedElevator = *SelectedColumn.onlineElevatorList[0]

			b.interfaceDisplay.goToColumn = SelectedColumn.nameLetter
			b.interfaceDisplay.goToElevator = SelectedElevator.name
			b.interfaceDisplay.goToRequest = RequestedFloor

			b.interfaceDisplay.displaygoto()

			SelectedColumn.move(SelectedElevator)

		} else {

			SelectedElevator.requestList = append(SelectedElevator.requestList, RequestedFloor)

			if SelectedElevator.toBase == "UP" {
				sort.Sort(sort.Reverse(sort.IntSlice(SelectedElevator.requestList))) // sortdescending
			} else {
				sort.Ints(SelectedElevator.requestList) //sort ascending
			}

			b.interfaceDisplay.goToColumn = SelectedColumn.nameLetter
			b.interfaceDisplay.goToElevator = SelectedElevator.name
			b.interfaceDisplay.goToRequest = RequestedFloor

			b.interfaceDisplay.displaygoto()

		}
		SelectedElevator.destinationFloor = b.baseFloor
		//selectedElevator.requestList.Add(baseFloor)
		SelectedElevator.requestList = append(SelectedElevator.requestList, RequestedFloor)
		SelectedColumn.move(SelectedElevator)

		//selectedElevator.destinationFloor = selectedElevator.requestList[selectedElevator.requestList.Count - 1];
		//selectedColumn.move(selectedElevator);

		b.selectedElevator = SelectedElevator
		b.selectedColumn = SelectedColumn

	} else {
		b.interfaceDisplay.display("SORRY the column to your desired floor is offline.")
	}

}

type Battery struct {
	FloorsPerColumn int

	stories                 int
	floors                  int
	basements               int
	rc                      int
	baseFloor               int
	numberOfColumns         int
	numberOfBasementColumns int
	elevatorsPerColumn      int

	columnList             []Column
	floorRequestButtonList []FloorRequestButton
	interfaceDisplay       InterfaceDisplay

	alarm bool

	selectedColumn   Column
	selectedElevator Elevator
}

func (c *Column) columnConstructor(aname int, afloorsServed []int, elevatorsPerColumn int, aBase int) {
	//fmt.Println("\n create columns")
	c.name = aname
	c.floorsServed = afloorsServed
	c.numberOfElevators = elevatorsPerColumn
	c.floors = len(c.floorsServed)
	c.online = true
	c.base = aBase

	c.nameLetter = stringValueOf(c.name)
	//fmt.Println(c.nameLetter)
	c.calculateToBase()
	c.createElevatorList()
	c.createCallbuttonList()

	c.selectedElevator = (&Elevator{
		toBase: "this is a test",
	})
}
func (c *Column) calculateToBase() {
	if c.floorsServed[0] < c.base {
		c.toBase = "UP"
	} else {
		c.toBase = "DOWN"
	}

}
func (c *Column) createElevatorList() {
	for elevatorName := 1; elevatorName < c.numberOfElevators+1; elevatorName++ {
		c.elevatorList = append(c.elevatorList, &Elevator{})
		c.elevatorList[elevatorName-1].ElevatorConstructor(elevatorName, c.toBase)
	}
}
func (c *Column) createCallbuttonList() {
	i := 0
	for _, floor := range c.floorsServed {

		j := strconv.Itoa(floor)
		c.callButtonList = append(c.callButtonList, CallButton{})

		c.callButtonList[i].name = j
		c.callButtonList[i].nameint = floor
		c.callButtonList[i].isPressed = false
		i++
	}
}
func (c *Column) createOnlineElevatorList() {
	for _, elevator := range c.elevatorList {
		if elevator.online {
			c.onlineElevatorList = append(c.onlineElevatorList, elevator)
		}
	}
}
func (c *Column) RequestElevator(FloorNumber int) { //c#55
	c.createOnlineElevatorList()
	c.callList = append(c.callList, FloorNumber)

	if len(c.onlineElevatorList) != 0 {
		if !c.checkIfElevatorOnItsWay(FloorNumber) {

			c.selectedElevator = (&Elevator{})

			c.sortElevatorByDistance(FloorNumber, c.onlineElevatorList)

			for _, elevator := range c.onlineElevatorList {
				if elevator.movement == "IDLE" {
					c.selectedElevator = elevator
					c.selectedElevator.selected = true
					c.selectedElevator.requestList = append(c.selectedElevator.requestList, FloorNumber)

				}
			}
			if !c.selectedElevator.selected {
				c.selectedElevator = c.onlineElevatorList[len(c.onlineElevatorList)-1]
				c.selectedElevator.requestList = append(c.selectedElevator.requestList, FloorNumber)
			}
		}
		fmt.Printf("\nxxxxxxxxxxxxxxxxxxxx         elevator %v was selected       xxxxxxxxxxxxxxxxxxxxx\n", c.selectedElevator.name)
		c.selectedElevator.destinationFloor = c.base
		c.move(*c.selectedElevator)

	} else {
		for _, elevator := range c.elevatorList {
			elevator.floorDisplay.message = "OFFLINE"
			elevator.floorDisplay.messageDisplay()
		}

	}

}
func (c *Column) checkIfElevatorOnItsWay(floor int) bool {
	for _, elevator := range c.onlineElevatorList {
		absFloorNumber := int(math.Abs(float64(elevator.floorNumber)))
		distanceToCall := int(math.Abs(math.Abs(float64(elevator.floorNumber)) - math.Abs(float64(floor))))
		distanceToDestination := int(math.Abs(math.Abs(float64(elevator.floorNumber)) - math.Abs(float64(elevator.destinationFloor))))

		//if elevators movement takes it through floor
		if absFloorNumber >= int(math.Abs(float64(floor))) && elevator.movement == elevator.toBase && distanceToCall <= distanceToDestination {
			c.selectedElevator = elevator
			return true
		}
	}
	return false
}
func (c *Column) sortElevatorByDistance(floor int, list []*Elevator) {
	fmt.Println("sortingElevators") //c# 172
	for _, elevator := range list {
		elevator.distance = int(math.Abs(math.Abs(float64(elevator.floorNumber)) - math.Abs(float64(floor))))

	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].distance < list[j].distance
	})

}
func (c *Column) move(elevator Elevator) { //c# 182
	fmt.Printf("column %v elevator %v is on floor : %v \n", c.name, elevator.name, elevator.floorNumber)
	if elevator.floorNumber == elevator.destinationFloor {
		//elevator is at it's current destination, needs to check if it has other destinations lined up.
		if len(elevator.requestList) != 0 {
			elevator.destinationFloor = elevator.requestList[0]
		} else if len(c.callList) != 0 {
			elevator.destinationFloor = c.callList[0]
		}
		//opens doors for potential passengers
	}
	for elevator.floorNumber != elevator.destinationFloor {
		if !elevator.isElevatorFull() {
			//set movement direction of the elevator
			if elevator.destinationFloor > elevator.floorNumber {
				elevator.movement = "UP"
				elevator.floorNumber++
				for _, item := range c.floorsServed {
					if item == 0 && elevator.floorNumber == 0 {
						elevator.floorNumber++
					}
				}
			} else {
				elevator.movement = "DOWN"
				elevator.floorNumber--

				for _, item := range c.floorsServed {
					if item == 0 && elevator.floorNumber == 0 {
						elevator.floorNumber--
					}
				}
			}
			for _, floor := range c.floorsServed {
				if floor == elevator.floorNumber {
					time.Sleep(500 * time.Millisecond)
				}
			}
			elevator.FloorDisplayUpdate(elevator.floorNumber, elevator.movement)

			floorIsInCallList := isIntIn(elevator.floorNumber, c.callList)
			floorIsInRequestList := isIntIn(elevator.floorNumber, elevator.requestList)

			//if elevator crosses a floor it should stop at (on a call or request list)
			if floorIsInCallList && elevator.movement == elevator.toBase { //|| (floorIsInRequestList && elevator.movement != elevator.toBase))
				elevator.doors.action()

				indexOfCall := IndexOf(elevator.floorNumber, c.floorsServed)

				c.callButtonList[indexOfCall].isPressed = false

				if floorIsInCallList {
					for i, other := range c.callList {
						if other == elevator.floorNumber {
							c.callList = append(c.callList[:i], c.callList[i+1:]...)
						}
					}
				}
				if floorIsInRequestList {
					for i, other := range elevator.requestList {
						if other == elevator.floorNumber {
							elevator.requestList = append(elevator.requestList[:i], elevator.requestList[i+1:]...)
						}
					}
				}

			}
			if elevator.floorNumber == elevator.destinationFloor {
				//elevator has reached it's destination

				elevator.doors.action()

			}
			tmpFloor := elevator.floorNumber
			for elevator.floorNumber == elevator.destinationFloor {
				//if all calls or requests have not been fullfilled this will catch them
				if len(c.callList) != 0 {
					elevator.destinationFloor = c.callList[0]
				} else if len(elevator.requestList) != 0 {
					elevator.destinationFloor = elevator.requestList[0]
				}
				if elevator.floorNumber == elevator.destinationFloor {
					if len(elevator.requestList) != 0 {
						elevator.requestList = append(elevator.requestList[:0], elevator.requestList[1:]...)
					} else {
						elevator.floorNumber = 0
					}

				}
			}
			elevator.floorNumber = tmpFloor
		} else {
			elevator.floorDisplay.message = "FULL!"
			elevator.floorDisplay.messageDisplay()

			var newLoad int
			fmt.Println("\n enter a value lower than %v", elevator.MAXLOAD)
			fmt.Scan(&newLoad)
			elevator.load = (newLoad)
			elevator.FloorDisplayUpdate(elevator.floorNumber, elevator.movement)

		}

	}

	if elevator.floorNumber == elevator.destinationFloor {
		elevator.movement = "IDLE"
		//elevator.Doors.action();
		indexOfCall := IndexOf(elevator.floorNumber, c.floorsServed)
		c.callButtonList[indexOfCall].isPressed = false
		c.goToIdle()

	}
}
func (c *Column) goToIdle() {
	fmt.Printf("\n column %v Idle elevators going to Idle levels \n", c.nameLetter)
}

type Column struct {
	name              int
	floors            int
	numberOfElevators int
	base              int

	toBase     string
	nameLetter string

	online bool

	callButtonList     []CallButton
	elevatorList       []*Elevator
	onlineElevatorList []*Elevator

	floorsServed []int
	callList     []int

	selectedElevator *Elevator
}

func (e *Elevator) ElevatorConstructor(elevatorname int, atoBase string) {
	e.name = elevatorname
	e.toBase = atoBase
	e.floorNumber = 0
	e.timer = 0
	e.load = 100
	e.online = true

	e.MAXLOAD = 10000

	e.doors = (Doors{
		open:              false,
		safeToClose:       true,
		passengerDetector: false,
		openTime:          5,
	})
	e.floorDisplay = (FloorDisplay{
		number:    0,
		direction: "",
		message:   "",
	})
}
func (e *Elevator) FloorDisplayUpdate(Number int, direction string) {
	e.floorDisplay.number = Number
	e.floorDisplay.direction = direction
	e.floorDisplay.noProblem()
}
func (e *Elevator) isElevatorFull() bool {
	if e.load <= e.MAXLOAD {
		return false
	} else {
		return true
	}
}

type Elevator struct {
	MAXLOAD int

	name             int
	floorNumber      int
	destinationFloor int
	distance         int
	timer            int
	load             int

	online   bool
	selected bool

	doors        Doors
	floorDisplay FloorDisplay

	movement string
	toBase   string

	requestList []int
}

type FloorRequestButton struct {
	name      string
	nameint   int
	isPressed bool
}
type CallButton struct {
	name      string
	nameint   int
	isPressed bool
}

func (i *InterfaceDisplay) displaygoto() {
	fmt.Println("\n interface displays : ")
	fmt.Printf("xxxxxxxxxxxxxxxxxxxxxx Go to column %v, elevator %v to go to floor : %v xxxxxxxxxxxxxxxxxxxxxxxxxx \n", i.goToColumn, i.goToElevator, i.goToRequest)
}
func (i *InterfaceDisplay) display(msg string) {
	fmt.Println("\n Interface Displays : ", msg)
}

type InterfaceDisplay struct {
	goTo          string
	floorServedBy string
	message       string
	goToColumn    string
	goToElevator  int
	goToRequest   int
}

func (f *FloorDisplay) messageDisplay() {
	fmt.Println("floor display shows : ")
	fmt.Println(f.message)
}
func (f *FloorDisplay) noProblem() {
	fmt.Printf("\n shaft floorDisplays show : %v going %v \n", f.number, f.direction)
}

type FloorDisplay struct {
	number    int
	direction string
	message   string
}

func (d *Doors) action() {
	d.open = true
	fmt.Println("open Doors")
	time.Sleep(time.Duration(d.openTime) * time.Second) // d.openTime = 5
	d.checkIfSafeToClose()
	fmt.Println("close Doors")
	d.open = false

}
func (d *Doors) checkIfSafeToClose() {

	for d.passengerDetector {
		d.safeToClose = false
		time.Sleep(5 * time.Second)
	}

	d.safeToClose = true

}

type Doors struct {
	open              bool
	safeToClose       bool
	passengerDetector bool
	openTime          int
}

func stringValueOf(i int) string {
	var foo = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return string(foo[i-1])
}
func isIntIn(Int int, Slice []int) bool {
	for _, item := range Slice {
		if item == Int {
			return true
		}
	}
	return false
}
func IndexOf(Int int, Slice []int) int {
	for i, item := range Slice {
		if item == Int {
			return i
		}
	}
	return -1
}
