package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	//define a struct
	door := new(Doors)
	door.open = false

	fmt.Printf("door : %v", door)

	testBattery := new(Battery)
	testBattery.batteryConstructor(66, 6, 1, 1, 5)

	fmt.Println("floors served :")
	for i := 0; i < len(testBattery.floorRequestButtonList); i++ {
		fmt.Printf(" %s", testBattery.floorRequestButtonList[i].name)
	}

	fmt.Printf("\n stories = %v", testBattery.interfaceDisplay)

	var newLoad int
	fmt.Printf("\n enter a value lower than %v", testBattery.columnList[0].elevatorList[0].MAXLOAD)
	fmt.Scan(&newLoad)

	fmt.Println("you entered : ", newLoad)

	testBattery.columnList[1].RequestElevator(10)
	fmt.Println("\n endProgram")
	//printed the struct for testing

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
func (b*Battery) selectColumn() Column{
	for _ ,column := range   each(Column column in columnList)
	{
		if (column.floorsServed.Contains(floor))
		{
			return column;
		}
	}
	return null;
}}

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

	c.selectedElevator = (Elevator{
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
		c.elevatorList = append(c.elevatorList, Elevator{})
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
	if len(c.onlineElevatorList) != 0 {
		if !c.checkIfElevatorOnItsWay(FloorNumber) {

			c.selectedElevator = (Elevator{})

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
			fmt.Println("xxxxxxxxxxxxxxxxxxxx         elevator %v was selected       xxxxxxxxxxxxxxxxxxxxx", c.selectedElevator.name)
			c.selectedElevator.destinationFloor = c.base
			c.move(c.selectedElevator)

		} else {
			for _, elevator := range c.elevatorList {
				elevator.floorDisplay.message = "OFFLINE"
				elevator.floorDisplay.messageDisplay()
			}
		}

	}

}
func (c *Column) checkIfElevatorOnItsWay(floor int) bool {
	return false
}
func (c *Column) sortElevatorByDistance(floor int, list []Elevator) {
	fmt.Println("sortingElevators") //c# 172
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
	elevatorList       []Elevator
	onlineElevatorList []Elevator

	floorsServed []int
	callList     []int

	selectedElevator Elevator
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
type InterfaceDisplay struct {
	goTo          string
	floorServedBy string
	message       string
	goToColumn    byte
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
