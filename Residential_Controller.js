//import time
//import math

//----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

class Building{
    

    constructor( numberOfFloors, numberOfColumns, _elevatorsPerColumn, _leavingTime, _arrivingTime){
        this.Floors = numberOfFloors ;
        this.Columns = numberOfColumns;
        this.ColumnList = [];
        this.elevatorsPerColumn = _elevatorsPerColumn;
        this.leavingTime = _leavingTime;
        this.arrivingTime = _arrivingTime;
        this.ALARM = false;

        this.create_Columns();
    }

    alarm() { 
        for (i = 0; i < this.ColumnList.length; i++) {
            var column = this.ColumnList[i] ; 
            column.Online = !this.ALARM;
            
            for ( n = 0; n < column.ElevatorList; n++) {
                 var elevator = column.ElevatorList[n] ; 
                elevator.Online = column.Online;
            }
        }
    }
 
    /*"""SEQUENCE timeCheck USING leavingTime && arrivingTime && time
        IF time IS GRETATER THAN (arrivingTime - 1hour) || SMALLER THAN (arrivingTime + 1hour)   
            RETURN 1
        ENDIF
    ENDSEQUENCE
    """*/

    
    /*
    //SEQUENCE timeCheck USING leavingTime && arrivingTime && time
    //    IF time IS GRETATER THAN (arrivingTime - 1hour) || SMALLER THAN (arrivingTime + 1hour)   
    //        RETURN 1
    //    ENDIF
    //ENDSEQUENCE
    */

    
    create_Columns (){
        var ID = 1;
        while (ID <= this.Columns){
            var column = new Column(ID, this.Floors, this.elevatorsPerColumn);
            this.ColumnList.push(column);
            ID += 1;

        }
    }

};

var testbuild = new Building (10, 1, 2, 8, 18)
console.log(testbuild.ColumnList)


//DEFINE Collumn USING floors && elevators
class Column {

    constructor ( ID, _floors, _elevators){
        this.ID = ID ;
        this.Floors = _floors;
        this.numberOfElevators = _elevators;
        this.CallButtonList = [];
        this.ElevatorList = [];
        this.Online = True;
        this.OnlineElevatorList = [];
        this.CallList = [];


        this.create_Elevators();
        this.create_callButtons();
    }

    create_callButtons(){
        var ID = 1
        while (ID <= this.Floors){
            var callButton = new CallButton(ID) ;
            this.CallButtonList.push(callButton) ;
            ID += 1 ;
        }
    }
        
    create_Elevators(){
        ID = 1
        while (ID <= this.numberOfElevators){
            elevator = new Elevator(ID, this.Floors);
            this.ElevatorList.push(elevator);
            ID += 1;
        }
    }

    create_CallList() {
        for (i = 0 ; i < this.CallButtonList.length ; i++ ) {
            const callButton = this.CallButtonList[i];
            if (callButton.IsPressed  && !(this.CallList.includes(callButton))){
                this.CallList.push(callButton);
            }
        }                
    }
    
    create_OnlineElevatorList (){
        this.OnlineElevatorList = []
        if (this.Online) {
            for (i = 0; i < this.ElevatorList.length ; i ++){ 
                const elevator = this.ElevatorList [i];
                if (elevator.Online) {
                    this.OnlineElevatorList.push(elevator);
                }
            }
        }
    }                    

    sortElevatorsByDistance (destination){
        
        for (i = 0 ; i < this.OnlineElevatorList.length ; i++ ){
            const elevator = this.OnlineElevatorList[i] ;
            elevator.Distance = Math.abs(destination - elevator.FloorNumber );
        }
        function Distance (a,b){ 
            if(a.Distance > b.Distance){
                return 1;
            }else {
                return -1;
            }
        }
        this.OnlineElevatorList.sort(Distance(a,b))
        console.log("closest elevator is " + this.OnlineElevatorList[0].ID)
    }

    RequestElevator( RequestedFloor, Direction){
        this.create_OnlineElevatorList();
        this.sortElevatorsByDistance ( RequestedFloor);
        let selectedElevator = null;
        for (i = 0 ; i < this.OnlineElevatorList.length ; i ++){ 
            var elevator = this.OnlineElevatorList[i] ;

            if (elevator.DestinationFloor != null){
                distanceToDestination = Math.abs(elevator.FloorNumber - elevator.DestinationFloor);
            } else { 
                distanceToDestination = 0;
            }


            
            if  ((elevator.FloorNumber >= RequestedFloor && elevator.Movement === Direction === "DOWN") || (elevator.FloorNumber <= RequestedFloor && elevator.Movement === Direction === "UP")) {
                if (elevator.Distance <= distanceToDestination) {     //'IF an elevator's travel takes it through button's floor on correct direction
                    this.move(elevator);
                } else {                                          // IF elevator travelling towards RequestedFloor on correct direction
                    elevator.DestinationList.push(RequestedFloor);
                }    
                
                selectedElevator = elevator;
                console.log("           XXXXXXXXXXXXXXXXXXXXXX          elevator "+(selectedElevator.ID)+ " was chosen XXXXXXXXXXXXXXXXX");
                this.move(selectedElevator);
                return selectedElevator;
            }else if (elevator.Movement === "IDLE") {
                if (RequestedFloor > elevator.FloorNumber){
                    elevator.Movement = "UP";
                }else { 
                    elevator.Movement = "DOWN";
                    elevator.DestinationList.push(RequestedFloor);
                }
                selectedElevator = elevator;
                console.log("           XXXXXXXXXXXXXXXXXXXXXX          elevator "+(selectedElevator.ID)+ " was chosen           XXXXXXXXXXXXXXXXXXXXXX          ");
                this.move(selectedElevator);
                return selectedElevator;
            }

        }    
        if (selectedElevator === None) {
            selectedElevator = this.OnlineElevatorList[this.OnlineElevatorList.length-1];
            selectedElevator.DestinationList.push(RequestedFloor);
        }        
        console.log("           XXXXXXXXXXXXXXXXXXXXXX          elevator "+(selectedElevator.ID)+ " was chosen           XXXXXXXXXXXXXXXXXXXXXX          ")
        this.move(selectedElevator)
        return selectedElevator
    }

    move ( elevator){
        
        elevator.DestinationList.sort();
        if (elevator.Movement === "UP"){ 
            elevator.DestinationFloor = elevator.DestinationList[elevator.DestinationList.length-1];
        } else if (elevator.Movement === "DOWN") {
            elevator.DestinationFloor = elevator.DestinationList[0];
        } else{ 
            elevator.DestinationFloor = elevator.DestinationList[0];
        }

        let destination = elevator.DestinationFloor ;
        console.log ("elevator "+(elevator.ID)+" is moving from " + (elevator.FloorNumber)+" to " + (destination));

        while  (!(elevator.FloorNumber === destination)){            

            if (elevator.Movement === "UP"){
                elevator.FloorNumber += 1;
            } else if (elevator.Movement === "DOWN") {
                elevator.FloorNumber -= 1;
            }

            console.log("elevator "+(elevator.ID)+ " is on floor"+(elevator.FloorNumber));
            time.sleep(0.1);
            
            if ( elevator.DestinationList.includes(elevator.FloorNumber) ) {
                elevator.openDoors();

                const request = new FloorRequestButton (elevator.FloorNumber, elevator.ID);
                    request.IsPressed = True;
                
                if ( elevator.FloorRequestButtonList.includes(request) ) {     
                    const index = elevator.FloorRequestButtonList.index(request);
                    elevator.FloorRequestButtonList[index].IsPressed = False;
                    elevator.DestinationList.remove(elevator.FloorNumber);
                }    
                const call = new CallButton(elevator.FloorNumber);
                    call.Direction = elevator.Movement;
                    call.IsPressed = True;

                if (this.CallList.includes(call)) { 
                    const index = this.CallButtonList.index(call);
                    this.CallButtonList[index].IsPressed = False;
                    this.CallList.remove(call);
                }    

                elevator.FloorRequestButtonList[elevator.FloorNumber-1].IsPressed = False;
            }
            if (elevator.DestinationList.length != 0) {
                destination = elevator.DestinationList[0];
                if (destination > elevator.FloorNumber){
                    elevator.Movement = "UP";
                }else { 
                    elevator.Movement = "DOWN";
                }
            }

        }    
            
        if (elevator.FloorNumber === destination) {
            elevator.Movement = "IDLE" ;
            console.log("elevator " + (elevator.ID) + " is IDLE");
            
            elevator.startTimer();
            elevator.DestinationList.remove(elevator.FloorNumber);
            //this.goToIdle()
        }
    }



    RequestFloor( elevator, RequestedFloor){
        if (elevator.Online ) {
            elevator.DestinationList.push(RequestedFloor);
            this.move(elevator);
        }
    }

    goToIdle () {

        this.create_OnlineElevatorList();
        var counter = 0;
        for (i = 0 ; i < this.OnlineElevatorList.lenght ; i++){ 
            let elevator = this.OnlineElevatorList[i] ;
            if (elevator.Movement === "IDLE") {
                if ((elevator.Timer - clocktime.getTime()) > 300000 ){

                    counter += 1
                }

                /*
                Call timeCheck of BuildingSpecs RETURNING idleFloor
                IF idleFloor not null THEN
                    FOR every elevator of collumn 
                        IF Movement IS false
                            THEN
                                SET  destinationFloor of elevator of collumn to idleFloor 
                                CALL moveElevator WITH elevator of collumn 
                        ENDIF
                    ENDFOR
                ELSE 
                */      
                var idleFloor = math.floor(this.Floors / (counter +1)) ;
                counter = 1 ;
                for ( i = 0 ; i < this.OnlineElevatorList.length; i++){ 
                    const elevator = this.OnlineElevatorList[i] ; 
                    if (elevator.Movement === "IDLE") {
                        if (elevator.FloorNumber !== idleFloor){ 
                            elevator.DestinationList.push(idleFloor * counter);
                            this.move(elevator)
                            counter += 1;
                        } else { return;
                        } 
                    }
                }
            }
        }

    }
}
//DEFINE CallButton USING floor && direction
class CallButton {
    constructor ( floor){
        this.Number = floor
        this.Direction = None
        this.IsPressed = False
    }
};

//DEFINE Elevator USING id && Location && floors
class Elevator{
    constructor ( ID, _floors){
        this.ID = ID;
        this.Floors = _floors;                              
        this.FloorNumber = 1;
        this.Movement = "IDLE"; //can be : up, down, || IDLE
        this.FloorRequestButtonList = [];
        this.DestinationFloor = None ;
        this.DestinationList =[];
        this.Distance = None;
        this.Doors = None;
        this.Timer = clocktime.getTime();
        this.Online = True;
        this.RequestList = [];

        this.MAXLOAD = 10000 ;
        this.LOAD = 0;

        this.create_FloorRequestButtons();
        this.create_Doors();
    }

    create_FloorRequestButtons (){
        number = 1
        while (number <= this.Floors){
            const floorRequestButton = new FloorRequestButton (number, this.ID);
            this.FloorRequestButtonList.push(floorRequestButton);
            number +=1 ;
        }
    }

    create_Doors (){
        doors = Doors () ;
        this.Doors = doors ;
    }
    isElevatorFull (){
        if (this.LOAD < this.MAXLOAD){
            return False;
        } else{
            return True;
        }
    }

    openDoors() {
        console.log("openning doors");
        this.Doors.Open = True;
        this.Doors.SafeToClose = False;
        (this.Doors.OpenTime)
        
        while (!(this.Doors.SafeToClose)){ 
            setTimeout(this.Doors.checkSafeToClose(),(this.Doors.OpenTime * 1000)) ;
            }
        
        if (this.Doors.SafeToClose) {
            this.Doors.Open = True;
            console.log("closing Doors");
        }
    }    


    startTimer() {
        this.Timer = clocktime.getTime();
    }
}

class FloorRequestButton{
    constructor ( _floor, ID){
        this.ID = ID;
        this.Number = _floor;
        IsPressed = False ;
    }
}

class Doors{
    constructor(){
        this.Open = False
        this.OpenTime = 1
        this.SafeToClose = True
        this.PassengerDetector = False
    }
    checkSafeToClose () { 
        if (!(this.PassengerDetector)) {
            this.SafeToClose = True;
        } else { 
            this.SafeToClose = False;
        }
    }
}
//SET time to clockTIME
var clocktime = new Date()



//------------------------------------------------------------------------------------------------------------------------
//      TEST ZONE


var build = new Building (10, 1, 2, 8, 18)
var COLLUMN = build.ColumnList[0]
var CALLLIST = COLLUMN.CallList


build.ColumnList[0].CallButtonList[5].IsPressed = True;
build.ColumnList[0].CallButtonList[5].Direction = "DOWN";    // to 3
build.ColumnList[0].create_CallList();
build.ColumnList[0].CallButtonList[7].IsPressed = True;
build.ColumnList[0].CallButtonList[7].Direction = "DOWN";    // to 5
build.ColumnList[0].create_CallList();
build.ColumnList[0].CallButtonList[4].IsPressed = True;
build.ColumnList[0].CallButtonList[4].Direction = "UP" ;     // to 10
build.ColumnList[0].create_CallList();
build.ColumnList[0].CallButtonList[6].IsPressed = True;
build.ColumnList[0].CallButtonList[6].Direction = "DOWN";    // to 2

COLLUMN.ElevatorList[0].FloorNumber = 2 ;
COLLUMN.ElevatorList[1].FloorNumber = 3 ;
var requestList = [3,5,10,2] ;

build.ColumnList[0].create_CallList() ;

/*CALLLIST.reverse();

counter = 0 ;
while len(CALLLIST) != 0 :
    call = CALLLIST[CALLLIST.length-1]
    tmp = COLLUMN.RequestElevator(call.Number,call.Direction)
    CALLLIST.pop(-1)
    COLLUMN.RequestFloor(tmp,requestList[counter])

    counter +=1
*/
function runTest(CALLS,COLUMN, REQUESTLIST ){
    
    CALLS.reverse();

    let counter = 0;
    while (CALLS.length !== 0) {
        const call = CALLS[CALLS.length-1];
        console.log("checking call # "+ (counter +1));
        const tmp = COLUMN.RequestElevator(call.Number,call.Direction);
    
        CALLS.pop();
        COLUMN.RequestFloor(tmp,REQUESTLIST[counter])
        for (i = 0 ; i < COLUMN.OnlineElevatorList.length ; i ++){ 
            let elevator = COLUMN.OnlineElevatorList[i] ; 
            if (elevator.DestinationList !== []) {            
                scen3Column.move(elevator)
            }
        }
        counter +=1
    }

}
runTest(CALLLIST, COLLUMN, requestList);

    


/*//--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
    //TEST SECTION


    console.log ("_______________________________________     SCENARIO I ________________________________________________")
    Scenario1 = Building (10, 1, 2, 8, 18)
    scen1Column = Scenario1.ColumnList[0]
    scen1CALLS = scen1Column.CallList
        
    //'Elevator A is Idle at floor 2'
    scen1Column.ElevatorList[0].FloorNumber = 2
    scen1Column.ElevatorList[0].Movement = "IDLE"
    //'Elevator B is Idle at floor 6'
    scen1Column.ElevatorList[1].FloorNumber = 6
    scen1Column.ElevatorList[1].Movement = "IDLE"

    //'Someone is on floor 3 && wants to go to the 7th floor. '
    scen1Column.CallButtonList[2].IsPressed = True
    scen1Column.CallButtonList[2].Direction = "UP"
    scen1Column.create_CallList()
    requestList = [7]
        
    scen1CALLS.reverse()

    counter = 0
    while len(scen1CALLS) != 0 :
        call = scen1CALLS[scen1CALLS.length-1]
        console.log("checking call // "+ (counter +1))
        tmp = scen1Column.RequestElevator(call.Number,call.Direction)

        scen1CALLS.pop(-1)
        scen1Column.RequestFloor(tmp,requestList[counter])

        counter +=1
            
        console.log('Elevator A was expected to be sent from 3 tom 7.')
        

    console.log ("_______________________________________     SCENARIO II ________________________________________________")
    Scenario2 = Building (10, 1, 2, 8, 18)
    scen2Column = Scenario2.ColumnList[0]
    scen2CALLS = scen2Column.CallList

    //'Elevator A is Idle at floor 10 
    scen2Column.ElevatorList[0].FloorNumber = 10
    scen2Column.ElevatorList[0].Movement = "IDLE"
    //'Elevator B is idle at floor 3
    scen2Column.ElevatorList[1].FloorNumber = 3
    scen2Column.ElevatorList[1].Movement = "IDLE"

    //'Someone is on the 1st floor && requests the 6th floor. 
    scen2Column.CallButtonList[0].IsPressed = True
    scen2Column.CallButtonList[0].Direction = "UP"
    scen2Column.create_CallList()
    requestList = []
    requestList.push(6)
    //'2 minutes later, someone else is on the 3rd floor && requests the 5th floor. 
    scen2Column.CallButtonList[2].IsPressed = True
    scen2Column.CallButtonList[2].Direction = "UP"
    scen2Column.create_CallList()

    requestList.push(5)

    //'Finally, a third person is at floor 9 && wants to go down to the 2nd floor. 
    scen2Column.CallButtonList[8].IsPressed = True
    scen2Column.CallButtonList[8].Direction = "DOWN"
    scen2Column.create_CallList()

    requestList.push(2)

    scen2CALLS.reverse()

    counter = 0
    while len(scen2CALLS) != 0 :
        call = scen2CALLS[scen2CALLS.length-1]
        console.log("checking call // "+ (counter +1))
        tmp = scen2Column.RequestElevator(call.Number,call.Direction)

        scen2CALLS.pop(-1)
        scen2Column.RequestFloor(tmp,requestList[counter])

        counter +=1
            
    console.log('Elevator B (1to6) then B (3to5) && A (9to2) were expected to be sent.')



    console.log ("_______________________________________     SCENARIO III ________________________________________________")

    Scenario3 = Building (10, 1, 2, 8, 18)
    scen3Column = Scenario3.ColumnList[0]
    scen3CALLS = scen3Column.CallList

    //'Elevator A is Idle at floor 10 
    scen3Column.ElevatorList[0].FloorNumber = 10
    scen3Column.ElevatorList[0].Movement = "IDLE"
    //'Elevator B is Moving from floor 3 to floor 6
    scen3Column.ElevatorList[1].FloorNumber = 3
    scen3Column.ElevatorList[1].Movement = "UP"
    scen3Column.ElevatorList[1].DestinationFloor = 6
    scen3Column.ElevatorList[1].DestinationList = [6]



    //'Someone is on floor 3 && requests the 2nd floor. 

    scen3Column.CallButtonList[2].IsPressed = True
    scen3Column.CallButtonList[2].Direction = "DOWN"
    scen3Column.create_CallList()
    requestList = []
    requestList.push(2)



    //'5 minutes later, someone else is on the 10th floor && wants to go to the 3rd. 
    scen3Column.CallButtonList[9].IsPressed = True
    scen3Column.CallButtonList[9].Direction = "DOWN"
    scen3Column.create_CallList()

    requestList.push(3)




    scen3CALLS.reverse()

    counter = 0
    while len(scen3CALLS) != 0 :
        call = scen3CALLS[scen3CALLS.length-1]
        console.log("checking call // "+ (counter +1))
        tmp = scen3Column.RequestElevator(call.Number,call.Direction)

        scen3CALLS.pop(-1)
        scen3Column.RequestFloor(tmp,requestList[counter])
        for elevator in scen3Column.OnlineElevatorList : 
            if elevator.DestinationList != [] :            
                scen3Column.move(elevator)
        counter +=1

    console.log("Elevator A (3to2) Then Elevator B (10to3) was expected to be sent.")
    
 
//----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
//EVENT LISTENERS
'''
 
        WHEN a Callbutton IS pressed
            CALL CallElevatorCheck WITH COLLUMN

        WHEN a FloorRequestButton IS pressed
            CALL usersDestination WITH COLLUMN

        WHEN 1 second passes
            CALL timerAdd WITH BUILDING

        WHEN 'elevator is IDLE for 5 mins'
            a  Movement of Elevator of collumn IS false && timer of Elevator of collumn  is 300 
            CALL goToIdle WITH collumn

        WHEN BUILDING ALARM IS CHANGED
            CALL alarm WITH COLLUMN && BUILDING

'''
//----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

*/