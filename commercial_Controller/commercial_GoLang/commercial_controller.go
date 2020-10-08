package main

func main() {

	door := new(Doors)
	door.open = true

}

type Battery struct {
	stories                 int
	floors                  int
	basements               int
	rc                      int
	numberOfColumns         int
	numberOfBasementColumns int
	elevatorsPerColumn      int

	columnList             []Column
	floorRequestButtonList []FloorRequestButton
	interfaceDisplay       InterfaceDisplay

	alarm bool
}
type Column struct {
	name              int
	floors            int
	numberOfElevators int
	base              int

	toBase     string
	nameLetter byte //byte replaces char c#

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
