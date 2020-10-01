//import time
//import math

//----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
class FloorRequestButton{
    constructor ( _floor, ID){
        this.ID = ID;
        this.Number = _floor;
        this.IsPressed = false ;
    }
};

class Doors{
    constructor(){
        this.Open = false
        this.OpenTime = 1
        this.SafeToClose = true
        this.PassengerDetector = false
    }
    checkSafeToClose () { 
        if (!(this.PassengerDetector)) {
            this.SafeToClose = true;
        } else { 
            this.SafeToClose = false;
        }
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
        this.DestinationFloor = null ;
        this.DestinationList =[];
        this.Distance = null;
        this.Doors = null;
        this.Timer = clocktime.getTime();
        this.Online = true;
        this.RequestList = [];

        this.MAXLOAD = 10000 ;
        this.LOAD = 0;

        this.create_FloorRequestButtons();
        this.create_Doors();
    }

    create_FloorRequestButtons (){
        let number = 1
        while (number <= this.Floors){
            const floorRequestButton = new FloorRequestButton (number, this.ID);
            this.FloorRequestButtonList.push(floorRequestButton);
            number +=1 ;
        }
    }

    create_Doors (){
        const doors = new Doors () ;
        this.Doors = doors ;
    }
    isElevatorFull (){
        if (this.LOAD < this.MAXLOAD){
            return false;
        } else{
            return true;
        }
    }

    openDoors() {
        console.log("openning doors");
        this.Doors.Open = true;
        this.Doors.SafeToClose = false;
        //console.log(this.Doors.OpenTime+ " door open time")
        
        
        while (!(this.Doors.SafeToClose)){ 
            sleep(this.Doors.OpenTime * 1000);
            this.Doors.checkSafeToClose();
        }

        if (this.Doors.SafeToClose) {
            this.Doors.Open = true;
            console.log("closing Doors");
        }
    };    


    startTimer() {
        this.Timer = new Date().getTime();
    };
};

//DEFINE CallButton USING floor && direction
class CallButton {
    constructor ( floor){
        this.Number = floor
        this.Direction = null
        this.IsPressed = false
    }
};

//DEFINE Collumn USING floors && elevators
class Column {

    constructor ( ID, _floors, _elevators, building){
        this.ID = ID ;
        this.Floors = _floors;
        this.numberOfElevators = _elevators;
        this.CallButtonList = [];
        this.ElevatorList = [];
        this.Online = true;
        this.OnlineElevatorList = [];
        this.CallList = [];

        this.building = building;

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
        let ID = 1
        while (ID <= this.numberOfElevators){
            const elevator = new Elevator(ID, this.Floors);
            this.ElevatorList.push(elevator);
            ID += 1;
        }
    }

    create_CallList() {
        for (let i = 0 ; i < this.CallButtonList.length ; i++ ) {
            const callButton = this.CallButtonList[i];
            if (callButton.IsPressed  && !(this.CallList.includes(callButton))){
                this.CallList.push(callButton);
            }
        }                
    }
    
    create_OnlineElevatorList (){
        this.OnlineElevatorList = []
        if (this.Online) {
            for (let i = 0; i < this.ElevatorList.length ; i ++){ 
                const elevator = this.ElevatorList [i];
                if (elevator.Online) {
                    this.OnlineElevatorList.push(elevator);
                }
            }
        }
    }                    

    sortElevatorsByDistance (destination){
        
        for (let i = 0 ; i < this.OnlineElevatorList.length ; i++ ){
            const elevator = this.OnlineElevatorList[i] ;
            elevator.Distance = Math.abs(destination - elevator.FloorNumber );
        }
        this.OnlineElevatorList.sort((a,b) => { 
            if(a.Distance > b.Distance){
                return 1;
            }else {
                return -1;
            }
        })
        console.log("closest elevator is " + this.OnlineElevatorList[0].ID)
    }

    RequestElevator( RequestedFloor, Direction){
        this.create_OnlineElevatorList();
        if(this.OnlineElevatorList.length !== 0){
            this.sortElevatorsByDistance ( RequestedFloor);
            let selectedElevator = null;
            for (let i = 0 ; i < this.OnlineElevatorList.length ; i ++){ 
                var elevator = this.OnlineElevatorList[i] ;

                if (elevator.DestinationFloor != null){
                    const distanceToDestination = Math.abs(elevator.FloorNumber - elevator.DestinationFloor);
                } else { 
                    const distanceToDestination = 0;
                }


                
                if  ((elevator.FloorNumber >= RequestedFloor && elevator.Movement === Direction === "DOWN") || (elevator.FloorNumber <= RequestedFloor && elevator.Movement === Direction === "UP")) {
                    if (elevator.Distance <= distanceToDestination) {     //'IF an elevator's travel takes it through button's floor on correct direction
                        this.move(elevator);
                    } else {                                          // IF elevator travelling towards RequestedFloor on correct direction
                        elevator.DestinationList.push(RequestedFloor);
                    }    
                    
                    selectedElevator = elevator;
                    selectedElevator.DestinationList.push(RequestedFloor);

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
                    selectedElevator.DestinationList.push(RequestedFloor);

                    console.log("           XXXXXXXXXXXXXXXXXXXXXX          elevator "+(selectedElevator.ID)+ " was chosen           XXXXXXXXXXXXXXXXXXXXXX          ");
                    this.move(selectedElevator);
                    return selectedElevator;
                }

            } 

            if (selectedElevator === null) {
                selectedElevator = this.OnlineElevatorList[this.OnlineElevatorList.length-1];
                selectedElevator.DestinationList.push(RequestedFloor);
            }        
            console.log("           XXXXXXXXXXXXXXXXXXXXXX          elevator "+(selectedElevator.ID)+ " was chosen           XXXXXXXXXXXXXXXXXXXXXX          ")
            this.move(selectedElevator)
            return selectedElevator
        }else {console.log("elevators of column "+ this.ID+ " are offline" );
            return;
        }
    };

    goToIdle () {

        //console.log(this.building + " elevators are going to Idle positions +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++");
        sleep(1000);
        this.create_OnlineElevatorList();
        
        var counter = 0;
        for (let i = 0 ; i < this.OnlineElevatorList.length ; i++){ 
            let elevator = this.OnlineElevatorList[i] ;
            if (elevator.Movement === "IDLE") {
                //if ((elevator.Timer - clocktime.getTime()) > 300000 ){

                counter += 1
            }
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
        var idleFloor = Math.floor(this.Floors / (counter +1)) ;
        counter = 1 ;
        for ( let i = 0 ; i < this.OnlineElevatorList.length; i++){ 
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
            

    };

    move ( elevator){
        if ( !(elevator.isElevatorFull()) ){
            elevator.DestinationList.sort();
            if (elevator.Movement === "UP"){ 
                elevator.DestinationFloor = elevator.DestinationList[elevator.DestinationList.length-1];
            } else if (elevator.Movement === "DOWN") {
                elevator.DestinationFloor = elevator.DestinationList[0];
            } else{ 
                elevator.DestinationFloor = elevator.DestinationList[0];
            }

            let destination = elevator.DestinationFloor ;
            console.log (this.building+" elevator "+(elevator.ID)+" is moving from " + (elevator.FloorNumber)+" to " + (destination));

            while  (!(elevator.FloorNumber === destination)){            

                if (elevator.Movement === "UP"){
                    elevator.FloorNumber += 1;
                } else if (elevator.Movement === "DOWN") {
                    elevator.FloorNumber -= 1;
                }
                function wait(){
                    var wait = "TIME"
                }
                console.log("elevator "+(elevator.ID)+ " is on floor"+(elevator.FloorNumber));
                sleep(300);
               
                
                if ( elevator.DestinationList.includes(elevator.FloorNumber) ) {
                    elevator.openDoors();

                    const request = new FloorRequestButton (elevator.FloorNumber, elevator.ID);
                        request.IsPressed = true;
                    
                    if ( elevator.FloorRequestButtonList.includes(request) ) {     
                        const index = elevator.FloorRequestButtonList.index(request);
                        elevator.FloorRequestButtonList[index].IsPressed = false;
                        elevator.DestinationList.remove(elevator.FloorNumber);
                    }    
                    const call = new CallButton(elevator.FloorNumber);
                        call.Direction = elevator.Movement;
                        call.IsPressed = true;

                    if (this.CallList.includes(call)) { 
                        const index = this.CallButtonList.index(call);
                        this.CallButtonList[index].IsPressed = false;
                        this.CallList.remove(call);
                    }    

                    elevator.FloorRequestButtonList[elevator.FloorNumber-1].IsPressed = false;
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
                elevator.DestinationList.splice(elevator.DestinationList.indexOf(elevator.FloorNumber),1);
                
                var _this = this;
                setTimeout(function(){_this.goToIdle()},10000);
            }
        }else {
            console.log("elevator is FULL")
        }
    }



    RequestFloor( elevator, RequestedFloor){
        if (elevator !== undefined) {
            if (elevator.Online ) {
            elevator.DestinationList.push(RequestedFloor);
            this.move(elevator);
            }else {console.log("elevator off line")}
        }else {console.log("request can't be treated !!!")}
    }
};

class Building{
    

    constructor( numberOfFloors, numberOfColumns, _elevatorsPerColumn, _leavingTime, _arrivingTime, name){
        this.Floors = numberOfFloors ;
        this.Columns = numberOfColumns;
        this.ColumnList = [];
        this.elevatorsPerColumn = _elevatorsPerColumn;
        this.leavingTime = _leavingTime;
        this.arrivingTime = _arrivingTime;
        this.ALARM = false;

        this.name = name;

        this.create_Columns();
    }

    alarm() {
        this.ALARM = !this.ALARM 
        if (this.ALARM){
            console.log("WARNING! WARNING! WARNING! ALARMS ARE ON WARNING! WARNING! WARNING!")
            console.log("                                / \\   ");
            console.log("                               /   \\   ");
            console.log("                              /  |  \\   ");
            console.log("                             /   |   \\   ");
            console.log("                            /    |    \\   ");
            console.log("                           /     o     \\   ");
            console.log("                          /_____________\\   ");


        }
       for (let i = 0; i < this.ColumnList.length; i++) {
            var column = this.ColumnList[i] ; 
            column.Online = !this.ALARM;
            
            for (let n = 0; n < column.ElevatorList; n++) {
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
            var column = new Column(ID, this.Floors, this.elevatorsPerColumn, this.name);
            this.ColumnList.push(column);
            ID += 1;

        }
    }

};

//SET time to clockTIME
var clocktime = new Date() ;

//sleep function found on discord given by William Sinclair
function sleep(miliseconds) {
    const start = new Date().getTime();
    while ((new Date().getTime()-start) < miliseconds){}
}

//------------------------------------------------------------------------------------------------------------------------
//      TEST ZONE
function runTest(CALLS,COLUMN, REQUESTLIST ){
    
    CALLS.reverse();

    let counter = 0;
    while (CALLS.length !== 0) {
        const call = CALLS[CALLS.length-1];
        console.log("checking call # "+ (counter +1));
        console.log (call)
        for (let i = 0 ; i < COLUMN.OnlineElevatorList.length ; i ++){ 
            let elevator = COLUMN.OnlineElevatorList[i] ; 
            console.log("elevator " +elevator.ID + " movement is " + elevator.Movement + " on floor # " + elevator.FloorNumber )
        }
        const tmp = COLUMN.RequestElevator(call.Number,call.Direction);

        CALLS.pop();
        console.log(requestList[counter]+ " is requested")
        COLUMN.RequestFloor(tmp,REQUESTLIST[counter])
        for (let i = 0 ; i < COLUMN.OnlineElevatorList.length ; i ++){ 
            let elevator = COLUMN.OnlineElevatorList[i] ; 
            if (elevator.DestinationList.length !== 0) {            
                COLUMN.move(elevator)
            }
        }
        counter +=1
    }

}

/* THIS WAS MY TEST SEQUENCE
var build = new Building (10, 1, 2, 8, 18, "build")
var COLLUMN = build.ColumnList[0]
var CALLLIST = COLLUMN.CallList


build.ColumnList[0].CallButtonList[5].IsPressed = true;
build.ColumnList[0].CallButtonList[5].Direction = "DOWN";    // to 3
build.ColumnList[0].create_CallList();
build.ColumnList[0].CallButtonList[7].IsPressed = true;
build.ColumnList[0].CallButtonList[7].Direction = "DOWN";    // to 5
build.ColumnList[0].create_CallList();
build.ColumnList[0].CallButtonList[4].IsPressed = true;
build.ColumnList[0].CallButtonList[4].Direction = "UP" ;     // to 10
build.ColumnList[0].create_CallList();
build.ColumnList[0].CallButtonList[6].IsPressed = true;
build.ColumnList[0].CallButtonList[6].Direction = "DOWN";    // to 2

COLLUMN.ElevatorList[0].FloorNumber = 2 ;

COLLUMN.ElevatorList[1].FloorNumber = 3 ;

console.log(COLLUMN.ElevatorList[0].FloorNumber) ;

console.log(COLLUMN.ElevatorList[1].FloorNumber);

var requestList = [3,5,10,2] ;

build.ColumnList[0].create_CallList() ;


runTest(CALLLIST, COLLUMN, requestList);

 */   


//--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
    //TEST SECTION


    console.log ("_______________________________________     SCENARIO I ________________________________________________");
    var Scenario1 = new Building (10, 1, 2, 8, 18, "Scenario1");
    var scen1Column = Scenario1.ColumnList[0];
    var scen1CALLS = scen1Column.CallList;
    //Scenario1.alarm()

    console.log(Scenario1.name)
    console.log(scen1Column.building)    
    //'Elevator A is Idle at floor 2'
    scen1Column.ElevatorList[0].FloorNumber = 2;
    scen1Column.ElevatorList[0].Movement = "IDLE";
    //'Elevator B is Idle at floor 6'
    scen1Column.ElevatorList[1].FloorNumber = 6;
    scen1Column.ElevatorList[1].Movement = "IDLE";

    //'Someone is on floor 3 && wants to go to the 7th floor. '
    scen1Column.CallButtonList[2].IsPressed = true;
    scen1Column.CallButtonList[2].Direction = "UP";
    scen1Column.create_CallList();
    requestList = [7];
   
    runTest(scen1CALLS, scen1Column, requestList);
            
    console.log('                                   Elevator A was expected to be sent from 3 tom 7.');
        

    console.log ("_______________________________________     SCENARIO II ________________________________________________");
    var Scenario2 = new Building (10, 1, 2, 8, 18, "scenario2");
    var scen2Column = Scenario2.ColumnList[0];
    var scen2CALLS = scen2Column.CallList;

    //'Elevator A is Idle at floor 10 
    scen2Column.ElevatorList[0].FloorNumber = 10;
    scen2Column.ElevatorList[0].Movement = "IDLE";
    //'Elevator B is idle at floor 3
    scen2Column.ElevatorList[1].FloorNumber = 3;
    scen2Column.ElevatorList[1].Movement = "IDLE";

    //'Someone is on the 1st floor && requests the 6th floor. 
    scen2Column.CallButtonList[0].IsPressed = true;
    scen2Column.CallButtonList[0].Direction = "UP";
    scen2Column.create_CallList();
    requestList = [];
    requestList.push(6);
    //'2 minutes later, someone else is on the 3rd floor && requests the 5th floor. 
    scen2Column.CallButtonList[2].IsPressed = true;
    scen2Column.CallButtonList[2].Direction = "UP";
    scen2Column.create_CallList();

    requestList.push(5);

    //'Finally, a third person is at floor 9 && wants to go down to the 2nd floor. 
    scen2Column.CallButtonList[8].IsPressed = true;
    scen2Column.CallButtonList[8].Direction = "DOWN";
    scen2Column.create_CallList();

    requestList.push(2);

    runTest(scen2CALLS, scen2Column, requestList);
            
    console.log('                           Elevator B (1to6) then B (3to5) && A (9to2) were expected to be sent.');



    console.log ("_______________________________________     SCENARIO III ________________________________________________");

    Scenario3 = new Building (10, 1, 2, 8, 18, "scenario3");
    scen3Column = Scenario3.ColumnList[0];
    scen3CALLS = scen3Column.CallList;

    //'Elevator A is Idle at floor 10 
    scen3Column.ElevatorList[0].FloorNumber = 10;
    scen3Column.ElevatorList[0].Movement = "IDLE";
    //'Elevator B is Moving from floor 3 to floor 6
    scen3Column.ElevatorList[1].FloorNumber = 3;
    scen3Column.ElevatorList[1].Movement = "UP";
    scen3Column.ElevatorList[1].DestinationFloor = 6;
    scen3Column.ElevatorList[1].DestinationList = [6];



    //'Someone is on floor 3 && requests the 2nd floor. 

    scen3Column.CallButtonList[2].IsPressed = true;
    scen3Column.CallButtonList[2].Direction = "DOWN";
    scen3Column.create_CallList();
    requestList = [];
    requestList.push(2);



    //'5 minutes later, someone else is on the 10th floor && wants to go to the 3rd. 
    scen3Column.CallButtonList[9].IsPressed = true;
    scen3Column.CallButtonList[9].Direction = "DOWN";
    scen3Column.create_CallList();

    requestList.push(3);

    runTest(scen3CALLS, scen3Column, requestList)
    console.log("                                   Elevator A (3to2) Then Elevator B (10to3) was expected to be sent.")

    //bonus scenario
    var sceneB = new Building(66, 4, 5, 8, 18, sceneB);

    sceneB.alarm();

    
 
/*//----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
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