package main

import (
	"fmt"
	"math"
	"strconv"
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
	//c.createElevatorList()
	//c.createCallbuttonList()
}
func (c *Column) calculateToBase() {
	if c.floorsServed[0] < c.base {
		c.toBase = "UP"
	} else {
		c.toBase = "DOWN"
	}

}
func stringValueOf(i int) string {
	var foo = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return string(foo[i-1])
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
}
type Elevator struct {
	MAXLOAD int

	name             int
	floorNumber      int
	destinationFloor int
	distance         int
	timer            int
	load             int

	online bool

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
type FloorDisplay struct {
	number    int
	direction string
	message   string
}
type Doors struct {
	open              bool
	safeToClose       bool
	passengerDetector bool
	openTime          int
}
