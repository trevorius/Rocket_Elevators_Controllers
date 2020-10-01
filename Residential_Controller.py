import time
import math




#----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

#DEFINE BuildingSpecs USING floors AND collumns AND elevators AND _leavingTime AND _arrivingTime
 
class Building:
    

    def __init__(self, numberOfFloors, numberOfColumns, _elevatorsPerColumn, _leavingTime, _arrivingTime):
        self.Floors= numberOfFloors
        self.Columns = numberOfColumns
        self.ColumnList = []
        self.elevatorsPerColumn = _elevatorsPerColumn
        self.leavingTime = _leavingTime
        self.arrivingTime = _arrivingTime
        self.ALARM = False

        self.create_Columns()

    def alarm(self) : 
        self.ALARM= not self.ALARM
        print("WARNING! WARNING! WARNING! ALARM IS ON WARNING! WARNING! WARNING! ")
        for column in self.ColumnList : 
            column.Online = not self.ALARM
            for elevator in column.ElevatorList : 
                elevator.Online = column.Online
 
    """SEQUENCE timeCheck USING leavingTime AND arrivingTime AND time
        IF time IS GRETATER THAN (arrivingTime - 1hour) OR SMALLER THAN (arrivingTime + 1hour)   
            RETURN 1
        ENDIF
    ENDSEQUENCE
    """

    
    #SEQUENCE timeCheck USING leavingTime AND arrivingTime AND time
    #    IF time IS GRETATER THAN (arrivingTime - 1hour) OR SMALLER THAN (arrivingTime + 1hour)   
    #        RETURN 1
    #    ENDIF
    #ENDSEQUENCE

    
    def create_Columns (self):
        ID = 1
        while ID <= self.Columns:
            column = Column(ID, self.Floors, self.elevatorsPerColumn)
            self.ColumnList.append(column)
            ID += 1

#DEFINE Collumn USING floors AND elevators
class Column:

    def __init__ (self, ID, _floors, _elevators):
        self.ID = ID
        self.Floors = _floors
        self.numberOfElevators = _elevators
        self.CallButtonList = []
        self.ElevatorList = []
        self.Online = True
        self.OnlineElevatorList = []
        self.CallList = []


        self.create_Elevators()
        self.create_callButtons()

    def create_callButtons(self):
        ID = 1
        while ID <= self.Floors:
            callButton = CallButton(ID)
            self.CallButtonList.append(callButton)
            ID += 1
        
    def create_Elevators(self):
        ID = 1
        while ID <= self.numberOfElevators:
            elevator = Elevator(ID, self.Floors)
            self.ElevatorList.append(elevator)
            ID += 1

    def create_CallList(self) :
        for callButton in self.CallButtonList :
            if callButton.IsPressed  and not (callButton in self.CallList):
                self.CallList.append(callButton)
    
    def create_OnlineElevatorList (self):
        self.OnlineElevatorList = []
        if self.Online :
            for elevator in self.ElevatorList :
                if elevator.Online :
                    self.OnlineElevatorList.append(elevator)

    def sortElevatorsByDistance (self, destination):
        
        for elevator in self.OnlineElevatorList :
            elevator.Distance = abs(destination - elevator.FloorNumber )
        def Distance (element): 
            return element.Distance
        self.OnlineElevatorList.sort(key=Distance)
        print("closest elevator is " + str(self.OnlineElevatorList[0].ID))

    def RequestElevator(self, RequestedFloor, Direction):
        self.create_OnlineElevatorList()
        if len(self.OnlineElevatorList) != 0 :     
            self.sortElevatorsByDistance ( RequestedFloor)
            selectedElevator = None
            for elevator in self.OnlineElevatorList :

                if elevator.DestinationFloor != None:
                    distanceToDestination = abs(elevator.FloorNumber - elevator.DestinationFloor)
                else : 
                    distanceToDestination = 0


                
                if  ((elevator.FloorNumber >= RequestedFloor and elevator.Movement == Direction == "DOWN") or (elevator.FloorNumber <= RequestedFloor and elevator.Movement == Direction == "UP")) :
                    if elevator.Distance <= distanceToDestination :     #'IF an elevator's travel takes it through button's floor on correct direction
                        self.move(elevator)
                    else :                                          # IF elevator travelling towards RequestedFloor on correct direction
                        elevator.DestinationList.append(RequestedFloor)
                        
                    
                    selectedElevator = elevator
                    print("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          elevator "+str(selectedElevator.ID)+ " was chosen XXXXXXXXXXXXXXXXXXXXXXXXXX")
                    self.move(selectedElevator)
                    return selectedElevator

                
                elif elevator.Movement == "IDLE" :
                    if RequestedFloor > elevator.FloorNumber :
                        elevator.Movement = "UP"
                    else : 
                        elevator.Movement = "DOWN"
                    elevator.DestinationList.append(RequestedFloor)
                    
                    selectedElevator = elevator
                    print("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          elevator "+str(selectedElevator.ID)+ " was chosen XXXXXXXXXXXXXXXXXXXXXXXXXX")
                    self.move(selectedElevator)
                    return selectedElevator

                
            if selectedElevator == None :
                selectedElevator = self.OnlineElevatorList[-1]
                selectedElevator.DestinationList.append(RequestedFloor)
                    
                print("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX          elevator "+str(selectedElevator.ID)+ " was chosen XXXXXXXXXXXXXXXXXXXXXXXXXX")
            self.move(selectedElevator)
            return selectedElevator
        else :
            print("elevators of column "+ str(self.ID)+ " are offline" )
            return


    def move (self, elevator):
        
        elevator.DestinationList.sort()
        if elevator.Movement == "UP" : 
            elevator.DestinationFloor = elevator.DestinationList[-1]
        elif elevator.Movement == "DOWN" :
            elevator.DestinationFloor = elevator.DestinationList[0]
        else :
            elevator.DestinationFloor = elevator.DestinationList[0]

        destination = elevator.DestinationFloor
        print ("elevator "+str(elevator.ID)+" is moving from " + str(elevator.FloorNumber)+" to " + str(destination))

        while not elevator.FloorNumber == destination : 

            
            

            if elevator.Movement == "UP" :
                elevator.FloorNumber += 1
                
            elif elevator.Movement == "DOWN" :
                elevator.FloorNumber -= 1
            

                
            print("elevator "+str(elevator.ID)+ " is on floor"+str(elevator.FloorNumber))
            time.sleep(0.1)
            
            if elevator.FloorNumber in elevator.DestinationList :
                elevator.openDoors()

                request = FloorRequestButton (elevator.FloorNumber, elevator.ID)
                request.IsPressed = True
                
                if request in elevator.FloorRequestButtonList :     
                    index = elevator.FloorRequestButtonList.index(request)
                    elevator.FloorRequestButtonList[index].IsPressed = False
                    elevator.DestinationList.remove(elevator.FloorNumber)
                    
                call = CallButton(elevator.FloorNumber)
                call.Direction = elevator.Movement
                call.IsPressed = True

                if call in self.CallList : 
                    index = self.CallButtonList.index(call)
                    self.CallButtonList[index].IsPressed = False
                    self.CallList.remove(call)
                    

                elevator.FloorRequestButtonList[elevator.FloorNumber-1].IsPressed = False
            if len(elevator.DestinationList) != 0 :
                destination = elevator.DestinationList[0]
                if destination > elevator.FloorNumber :
                    elevator.Movement = "UP"
                else : 
                    elevator.Movement = "DOWN"

            
            
        if elevator.FloorNumber == destination :
            elevator.Movement = "IDLE" 
            print("elevator " + str(elevator.ID) + " is IDLE")
            
            elevator.startTimer()
            elevator.DestinationList.remove(elevator.FloorNumber)
            
            # self.goToIdle)





    def RequestFloor(self, elevator, RequestedFloor):
        if elevator != None : 
            if elevator.Online == True :
                elevator.DestinationList.append(RequestedFloor)
                self.move(elevator)
        else : 
            print("request can't be dealt, there is no elevator")
            return

    def goToIdle (self) :

        self.create_OnlineElevatorList()
        counter = 0
        for elevator in self.OnlineElevatorList : 
            if elevator.Movement == "IDLE" :
                #if int(elevator.Timer) - int(time.time()) > 300 : 

                counter += 1
                

                """
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
                """      
            idleFloor = math.floor(self.Floors / (counter +1))
            counter = 1
            for elevator in self.OnlineElevatorList : 
                if elevator.Movement == "IDLE" :
                    if elevator.FloorNumber == idleFloor : 
                        return
                    else :  
                        elevator.DestinationList.append(idleFloor * counter)
                        self.move(elevator)
                        counter += 1

#DEFINE CallButton USING floor AND direction
class CallButton :
    def __init__ (self, floor):
        self.Number = floor
        self.Direction = None
        self.IsPressed = False


#DEFINE Elevator USING id AND Location AND floors
class Elevator:



    def __init__ (self, ID, _floors):
        self.ID = ID
        self.Floors = _floors
        self.FloorNumber = 1
        self.Movement = "IDLE" #can be : up, down, or IDLE
        self.FloorRequestButtonList = []
        self.DestinationFloor = None 
        self.DestinationList =[]
        self.Distance = None
        self.Doors = None
        self.Timer = time.time()
        self.Online = True
        self.RequestList = []

        self.MAXLOAD = 10000 
        self.LOAD = 0

        self.create_FloorRequestButtons()
        self.create_Doors()

    
 
    def create_FloorRequestButtons (self):
        number = 1
        while number <= self.Floors:
            floorRequestButton = FloorRequestButton (number, self.ID)
            self.FloorRequestButtonList.append(floorRequestButton)
            number +=1

    def create_Doors (self):
        doors = Doors ()
        self.Doors = doors

    def isElevatorFull (self):
        if self.LOAD < self.MAXLOAD:
            return False
        else: return True

    def openDoors(self) :
        print("openning doors")
        self.Doors.Open = True
        self.Doors.SafeToClose = False
        time.sleep(self.Doors.OpenTime)
        
        while self.Doors.SafeToClose == False :
            self.Doors.checkSafeToClose ()
        
        if self.Doors.SafeToClose == True :
            self.Doors.Open = True
            print("closing Doors")
        


    def startTimer(self) : 
        self.Timer = time.time()

class FloorRequestButton:
    
    def __init__ (self, _floor, ID):
        self.ID = ID
        self.Number = _floor
        IsPressed = False 

class Doors:
    def __init__(self) : 
        self.Open = False
        self.OpenTime = 1
        self.SafeToClose = True
        self.PassengerDetector = False

    def checkSafeToClose (self) : 
        if self.PassengerDetector == False :
            self.SafeToClose = True 
        else : 
            self.SafeToClose = False

#SET time to clockTIME
clocktime = time.time()



#------------------------------------------------------------------------------------------------------------------------
#      TEST ZONE


build = Building (10, 1, 2, 8, 18)
COLLUMN = build.ColumnList[0]
CALLLIST = COLLUMN.CallList

build.alarm()


build.ColumnList[0].CallButtonList[5].IsPressed = True
build.ColumnList[0].CallButtonList[5].Direction = "DOWN"    # to 3
build.ColumnList[0].create_CallList()
build.ColumnList[0].CallButtonList[7].IsPressed = True
build.ColumnList[0].CallButtonList[7].Direction = "DOWN"    # to 5
build.ColumnList[0].create_CallList()
build.ColumnList[0].CallButtonList[4].IsPressed = True
build.ColumnList[0].CallButtonList[4].Direction = "UP"      # to 10
build.ColumnList[0].create_CallList()
build.ColumnList[0].CallButtonList[6].IsPressed = True
build.ColumnList[0].CallButtonList[6].Direction = "DOWN"    # to 2

COLLUMN.ElevatorList[0].FloorNumber = 2
COLLUMN.ElevatorList[1].FloorNumber = 3
requestList = [3,5,10,2]

build.ColumnList[0].create_CallList()

CALLLIST.reverse()

counter = 0
while len(CALLLIST) != 0 :
    call = CALLLIST[-1]
    tmp = COLLUMN.RequestElevator(call.Number,call.Direction)
    CALLLIST.pop(-1)
    COLLUMN.RequestFloor(tmp,requestList[counter])

    counter +=1



#--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
#TEST SECTION


print ("_______________________________________     SCENARIO I ________________________________________________")
Scenario1 = Building (10, 1, 2, 8, 18)
scen1Column = Scenario1.ColumnList[0]
scen1CALLS = scen1Column.CallList
    
#'Elevator A is Idle at floor 2'
scen1Column.ElevatorList[0].FloorNumber = 2
scen1Column.ElevatorList[0].Movement = "IDLE"
#'Elevator B is Idle at floor 6'
scen1Column.ElevatorList[1].FloorNumber = 6
scen1Column.ElevatorList[1].Movement = "IDLE"

#'Someone is on floor 3 and wants to go to the 7th floor. '
scen1Column.CallButtonList[2].IsPressed = True
scen1Column.CallButtonList[2].Direction = "UP"
scen1Column.create_CallList()
requestList = [7]
    
scen1CALLS.reverse()

counter = 0
while len(scen1CALLS) != 0 :
    call = scen1CALLS[-1]
    print("checking call # "+ str(counter +1))
    tmp = scen1Column.RequestElevator(call.Number,call.Direction)

    scen1CALLS.pop(-1)
    scen1Column.RequestFloor(tmp,requestList[counter])

    counter +=1
        
    print('Elevator A was expected to be sent from 3 tom 7.')
       

print ("_______________________________________     SCENARIO II ________________________________________________")
Scenario2 = Building (10, 1, 2, 8, 18)
scen2Column = Scenario2.ColumnList[0]
scen2CALLS = scen2Column.CallList

#'Elevator A is Idle at floor 10 
scen2Column.ElevatorList[0].FloorNumber = 10
scen2Column.ElevatorList[0].Movement = "IDLE"
#'Elevator B is idle at floor 3
scen2Column.ElevatorList[1].FloorNumber = 3
scen2Column.ElevatorList[1].Movement = "IDLE"

#'Someone is on the 1st floor and requests the 6th floor. 
scen2Column.CallButtonList[0].IsPressed = True
scen2Column.CallButtonList[0].Direction = "UP"
scen2Column.create_CallList()
requestList = []
requestList.append(6)
#'2 minutes later, someone else is on the 3rd floor and requests the 5th floor. 
scen2Column.CallButtonList[2].IsPressed = True
scen2Column.CallButtonList[2].Direction = "UP"
scen2Column.create_CallList()

requestList.append(5)

#'Finally, a third person is at floor 9 and wants to go down to the 2nd floor. 
scen2Column.CallButtonList[8].IsPressed = True
scen2Column.CallButtonList[8].Direction = "DOWN"
scen2Column.create_CallList()

requestList.append(2)

scen2CALLS.reverse()

counter = 0
while len(scen2CALLS) != 0 :
    call = scen2CALLS[-1]
    print("checking call # "+ str(counter +1))
    tmp = scen2Column.RequestElevator(call.Number,call.Direction)

    scen2CALLS.pop(-1)
    scen2Column.RequestFloor(tmp,requestList[counter])

    counter +=1
        
print('Elevator B (1to6) then B (3to5) and A (9to2) were expected to be sent.')



print ("_______________________________________     SCENARIO III ________________________________________________")

Scenario3 = Building (10, 1, 2, 8, 18)
scen3Column = Scenario3.ColumnList[0]
scen3CALLS = scen3Column.CallList

#'Elevator A is Idle at floor 10 
scen3Column.ElevatorList[0].FloorNumber = 10
scen3Column.ElevatorList[0].Movement = "IDLE"
#'Elevator B is Moving from floor 3 to floor 6
scen3Column.ElevatorList[1].FloorNumber = 3
scen3Column.ElevatorList[1].Movement = "UP"
scen3Column.ElevatorList[1].DestinationFloor = 6
scen3Column.ElevatorList[1].DestinationList = [6]



#'Someone is on floor 3 and requests the 2nd floor. 

scen3Column.CallButtonList[2].IsPressed = True
scen3Column.CallButtonList[2].Direction = "DOWN"
scen3Column.create_CallList()
requestList = []
requestList.append(2)



#'5 minutes later, someone else is on the 10th floor and wants to go to the 3rd. 
scen3Column.CallButtonList[9].IsPressed = True
scen3Column.CallButtonList[9].Direction = "DOWN"
scen3Column.create_CallList()

requestList.append(3)




scen3CALLS.reverse()

counter = 0
while len(scen3CALLS) != 0 :
    call = scen3CALLS[-1]
    print("checking call # "+ str(counter +1))
    tmp = scen3Column.RequestElevator(call.Number,call.Direction)

    scen3CALLS.pop(-1)
    scen3Column.RequestFloor(tmp,requestList[counter])
    for elevator in scen3Column.OnlineElevatorList : 
        if elevator.DestinationList != [] :            
            scen3Column.move(elevator)
    counter +=1

print("Elevator A (3to2) Then Elevator B (10to3) was expected to be sent.")
 
 
#----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
#EVENT LISTENERS
'''
 
        WHEN a Callbutton IS pressed
            CALL CallElevatorCheck WITH COLLUMN

        WHEN a FloorRequestButton IS pressed
            CALL usersDestination WITH COLLUMN

        WHEN 1 second passes
            CALL timerAdd WITH BUILDING

        WHEN 'elevator is IDLE for 5 mins'
            a  Movement of Elevator of collumn IS false AND timer of Elevator of collumn  is 300 
            CALL goToIdle WITH collumn

        WHEN BUILDING ALARM IS CHANGED
            CALL alarm WITH COLLUMN AND BUILDING

'''
#----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

