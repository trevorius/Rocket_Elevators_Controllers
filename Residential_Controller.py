import time

#----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
'''
SET BUILDING to INSTANTIATE BuildingSpecs WITH 10 AND 1 AND 2 AND AND 8 AND 18
SET COLLUMN to INSTANTIATE Collumn WITH BUILDING Floors AND BUILDING Elevators
#--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
TEST SECTION

    SET SelectedElevator to null

    Scenario 1:
        'Elevator A is Idle at floor 2'
        SET floorNumber AND Movement  of elevator WITH ID:1 TO 2 AND false
        'Elevator B is Idle at floor 6'
        SET floorNumber AND Movement of elevator WITH ID:2 TO 6 AND false
        'Someone is on floor 3 and wants to go to the 7th floor. '
        SET IsPressed  of callButtonUP WITH number: 3 to true
        CALL CallElevatorCheck WITH COLLUMN
        SET IsPressed of FloorRequestButton With Number : 7 to true
        CALL usersDestination WITH COLLUMN
        'Elevator A is expected to be sent.
        DISPLAY SelectedElevator


    Scenario 2:
        'Elevator A is Idle at floor 10 
        SET floorNumber AND Movement of elevator WITH ID:1 TO 10 AND false
        'Elevator B is idle at floor 3
        SET floorNumber AND Movement of elevator WITH ID:2 TO 3 AND false
        'Someone is on the 1st floor and requests the 6th floor. 
        SET IsPressed  of callButtonUP WITH number: 1 to true
        CALL CallElevatorCheck WITH COLLUMN
        SET IsPressed of FloorRequestButton With Number : 6 to true
        CALL usersDestination WITH COLLUMN
        'Elevator B should be sent. 
        DISPLAY SelectedElevator


        '2 minutes later, someone else is on the 3rd floor and requests the 5th floor. 
        SET IsPressed  of callButtonUP WITH number: 3 to true
        CALL CallElevatorCheck WITH COLLUMN
        SET IsPressed of FloorRequestButton With Number : 5 to true
        CALL usersDestination WITH COLLUMN
        'Elevator B should be sent.
        DISPLAY SelectedElevator


        'Finally, a third person is at floor 9 and wants to go down to the 2nd floor. 
        SET IsPressed  of callButtonUP WITH number: 9 to true
        CALL CallElevatorCheck WITH COLLUMN
        SET IsPressed of FloorRequestButton With Number : 2 to true
        CALL usersDestination WITH COLLUMN

        'Elevator A should be sent.
        DISPLAY SelectedElevator

    Scenario 3:
        'Elevator A is Idle at floor 10 
        SET floorNumber AND Movement  of elevator WITH ID:1 TO 10 AND false
        'Elevator B is Moving from floor 3 to floor 6
        SET floorNumber AND Movement AND DestinationFloor of elevator WITH ID:1 TO 3 AND up AND 6
        'Someone is on floor 3 and requests the 2nd floor. 
        SET IsPressed  of callButtonUP WITH number: 3 to true
        CALL CallElevatorCheck WITH COLLUMN
        SET IsPressed of FloorRequestButton With Number : 2 to true
        CALL usersDestination WITH COLLUMN
        'Elevator A should be sent. 
        DISPLAY SelectedElevator

        '5 minutes later, someone else is on the 10th floor and wants to go to the 3rd. 
        SET IsPressed  of callButtonUP WITH number: 10 to true
        CALL CallElevatorCheck WITH COLLUMN
        SET IsPressed of FloorRequestButton With Number : 3 to true
        CALL usersDestination WITH COLLUMN

        'Elevator B should be sent.
        DISPLAY SelectedElevator
'''
#----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

#DEFINE BuildingSpecs USING floors AND collumns AND elevators AND _leavingTime AND _arrivingTime
 
class Building:
    ALARM : False

    def __init__(self, numberOfFloors, numberOfColumns, _elevatorsPerColumn, _leavingTime, _arrivingTime):
        self.Floors= numberOfFloors
        self.Columns = numberOfColumns
        self.ColumnList = []
        self.elevatorsPerColumn = _elevatorsPerColumn
        self.leavingTime = _leavingTime
        self.arrivingTime = _arrivingTime

        self.create_Columns()
 
    
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
            #callButtonDown = CallButton(ID, "DOWN")
            #self.CallButtonList.append(callButtonDown)
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
        if self.Online :
            for elevator in self.ElevatorList :
                if elevator.Online :
                    self.OnlineElevatorList.append(elevator)

    def sortElevatorsByDistance (self, destination):
        
        for elevator in self.OnlineElevatorList :
            elevator.Distance = abs(elevator.FloorNumber - destination)
        def Distance (element): 
            return element.Distance
        self.OnlineElevatorList.sort(key=Distance)

        
       
                    

    def RequestElevator(self, RequestedFloor, Direction):
        self.create_OnlineElevatorList()
        self.sortElevatorsByDistance ( RequestedFloor)
        selectedElevator = None
        #print (RequestedFloor)
        #print(self.OnlineElevatorList[0].ID)
        for elevator in self.OnlineElevatorList :

            #CALL (magnitude WITH (FloorNumber of Elevator MINUS callButton Number)) RETURNING distanceToCall
            distanceToCall = abs(elevator.FloorNumber - RequestedFloor)
            #CALL (magnitude WITH (FloorNumber of Elevator MINUS Elevator DestinationFloor)) RETURNING distanceToDestination
            distanceToDestination = abs(elevator.FloorNumber - elevator.DestinationFloor)


            
            if  ((elevator.FloorNumber >= RequestedFloor and elevator.Movement == Direction == "DOWN")
                or (elevator.FloorNumber <= RequestedFloor and elevator.Movement == Direction == "UP")) 
                if distanceToCall <= distanceToDestination :     #'IF an elevator's travel takes it through button's floor on correct direction
                    elevator.move(elevator.DestinationFloor)
                else :                                          # IF elevator travelling towards RequestedFloor on correct direction
                    elevator.DestinationFloor = RequestedFloor
                    elevator.move ()
                
                selectedElevator = elevator
            
            elif elevator.Movement is "IDLE" :
                elevator.DestinationFloor = RequestedFloor
                elevator.move ()
                selectedElevator = elevator
            
        if selectedElevator is None
            selectedElevator = self.OnlineElevatorList[-1]

                
                    
                    CASE Elevator Movement is false
                            SET DestinationFloor of Elevator to Number of callElement 
                            call moveElevator WITH Elevator
                            SET SelectedElevator to elevator of Online_elevators of collumn
                                
                    CASE FloorNumber OF Elevator GREATER THAN Number of callElement  AND Movement of Elevator  IS up 
                        OR FloorNumber of Elevator SMALLER THAN Number ofcallElement  AND Movement of Elevator  IS down
                        go to next Elevator
                            
                    ENDCASE
                ENDFOR
            ENDIF
        ENDFOR  
        """      





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
        self.Distance = None
        self.Doors = None
        self.Timer = 0
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
            #print(str(self.ID)+" "+str(number))
            number +=1

    def create_Doors (self):
        doors = Doors ()
        self.Doors = doors

    def isElevatorFull (self):
        if self.LOAD < self.MAXLOAD:
            return False
        else: return True

    def move (self, destination =self.DestinationFloor )

class FloorRequestButton:
    IsPressed = False
    def __init__ (self, _floor, ID):
        self.ID = ID
        self.Number = _floor 

class Doors:
    Open = False
    OpenTime = 5
    SafeToClose = True
    PassengerDetector : False

#SET time to clockTIME
time = time.time()

build = Building (10, 1, 2, 8, 18)
COLLUMN = build.ColumnList[0]
CALLLIST = COLLUMN.CallList


build.ColumnList[0].CallButtonList[5].IsPressed = True
build.ColumnList[0].CallButtonList[5].Direction = "DOWN"
build.ColumnList[0].create_CallList()
build.ColumnList[0].CallButtonList[7].IsPressed = True
build.ColumnList[0].CallButtonList[7].Direction = "DOWN"
build.ColumnList[0].create_CallList()
build.ColumnList[0].CallButtonList[4].IsPressed = True
build.ColumnList[0].CallButtonList[4].Direction = "UP"
build.ColumnList[0].create_CallList()
build.ColumnList[0].CallButtonList[6].IsPressed = True
build.ColumnList[0].CallButtonList[6].Direction = "DOWN"

COLLUMN.ElevatorList[0].FloorNumber = 2
COLLUMN.ElevatorList[1].FloorNumber = 3

COLLUMN.create_OnlineElevatorList()

COLLUMN.sortElevatorsByDistance(4)
print (COLLUMN.ElevatorList[0].FloorNumber)
print (COLLUMN.OnlineElevatorList[0].FloorNumber)




#print(dir(build.ColumnList[0].CallButtonList[3]))

build.ColumnList[0].create_CallList()
#print(len(build.ColumnList[0].CallList))
#print(build.ColumnList[0].CallList[0].Number)


#build.ColumnList[0].RequestElevator(build.ColumnList[0].CallList[2].Number, build.ColumnList[0].CallList[2].Direction)

#for call in CALLLIST:
#    COLLUMN.RequestElevator(call.Number,call.Direction)

""" object construction printouts
    #print("column ID is :")
    #print (build.ColumnList[0].ID)

    #print("column contains Elevators:")
    #print(len(build.ColumnList[0].ElevatorList))



    #print("elevatorList[1].Floors")
    #print(build.ColumnList[0].ElevatorList[1].Floors)
    #print("there are request buttons :")
    #print(len(build.ColumnList[0].ElevatorList[0].FloorRequestButtonList))
    #print(build.ColumnList[0].ElevatorList[1].FloorRequestButtonList[5].ID)
    #print(build.ColumnList[0].ElevatorList[0].ID)



    #print (range(build.ColumnList[0].ElevatorList[0].Floors))
    #print("testElevator")

    #testElevator = Elevator(1,10)
    #print("testElevator.Floors :")
    #print (testElevator.Floors)
    #print("there are requestbuttons in testElevator")
    #print (len(testElevator.FloorRequestButtonList))

    #print("testColumn")
    #testColumn = Column(1, 10, 2)
    #print ("testColumn.ElevatorList[1].Floors")
    #print (testColumn.ElevatorList[1].Floors)
    #print ("len(testColumn.ElevatorList[1].FloorRequestButtonList)")
    #print (len(testColumn.ElevatorList[1].FloorRequestButtonList))
"""

#----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
#sequences used as tools.
"""
    'sort SEQUENCE EDX week 3'
    'compare values array'
    SEQUENCE IS_IN USING value AND list     
        FOR every element of list
            IF value is SAME AS element
                THEN RETURN true
            ELSE RETURN false
            ENDIF
        ENDFOR
    ENDSEQUENCE
        

    'magnitude is the unsigned value of the variable'
    SEQUENCE magnitude USING number     
        IF number SMALLER THAN 0
            THEN RETURN number MULTIPLIES -1
        ELSE RETURN number
        ENDIF
    ENDSEQUENCE
"""
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

'''
SEQUENCE CallElevatorCheck USING collumn#to become RequestElevator(RequestedFloor, Direction)

    //CALL checkIfElevatorOnline WITH collumn
    
    //CALL pressButton WITH CallButtons of collumn  
    FOR every CallButton of CallButtons of collumn
        IF CallButton IsPressed SAME AS true 
            THEN ADD callButton to calls 
        ENDIF
    ENDFOR

    //FOR every callElement of calls 

    ENDFOR
    CALL usersDestination WITH collumn

               

ENDSEQUENCE

SEQUENCE usersDestination USING collumn# to become  RequestFloor (Elevator, RequestedFloor) 

    CALL checkIfElevatorOnline WITH collumn


    CALL pressButton WITH FloorRequestButtons
    FOR every  Elevator of Collumn    
        FOR every  FloorRequestButton of Elevator of Collumn 
            IF FloorRequestButton of Elevator of Collumn IsPressed IS true                
                ADD Number of FloorRequestButton of Elevator of Collumn to requests 
                SORT requests ascending
                CASE Movement of Elevator of Collumn IS up
                    SET DestinationFloor of Elevator of Collumn to  last of requests
                CASE Movement of Elevator of Collumn IS down
                    SET DestinationFloor of Elevator of Collumn to first of requests                   
                CASE Movement of Elevator of Collumn IS false

                    CALL (magnitude WITH (FloorNumber of Elevator of Collumn MINUS first of requests)) RETURNING lowest
                    CALL (magnitude WITH (last of requests MINUS FloorNumber of Elevator of Collumn)) RETURNING highest

                    CASE lowest GREATER THAN highest
                        SET DestinationFloor of Elevator of Collumn to first of requests
                    CASE lowest SMALLER THAN highest
                        SET DestinationFloor of Elevator of Collumn to last of requests
                    CASE lowest SAME AS highest
                        SET DestinationFloor of Elevator of Collumn to first of requests
                    ENDCASE
                ENDCASE
            ENDIF
        ENDFOR
        CALL moveElevator USING  Elevator of Collumn 
    ENDFOR
ENDSEQUENCE



SEQUENCE pressButton USING buttons 
    FOR every button  of buttons
        IF button is pressed
            SET  IsPressed  of button to true
        ENDIF 
    ENDFOR
ENDSEQUENCE

SEQUENCE moveElevator USING Elevator of Collumn 
    
    WHILE FloorNumber of Elevator  IS NOT elevator DestinationFloor AND (CALL isElevatorFull OF elevator) returning false

        IF elevator destinationFloor  GREATER THAN FloorNumber of Elevator 
            THEN SET elevator Movement to up
            INCREMENT floorNumber of elevator by +1
            
        ELSE 
            SET elevator Movement to down
            INCREMENT floorNumber of elevator by -1
        ENDIF

        CALL (IS_IN WITH FloorNumber of Elevator of Collumn   AND calls callElement Numbers) RETURNING FloorNumber_IS_IN_calls
        CALL (IS_IN WITH FloorNumber of Elevator of Collumn AND requests) RETURNING FloorNumber_IS_IN_requests

        IF  FloorNumber_IS_IN_calls AND Movement of Elevator of Collumn SAME AS callElement Direction
            OR FloorNumber_IS_IN_requests 
                THEN 
                    CALL DoorsAction USING Elevator of Collumn 
                    SET callButtons (WITH Number SAME AS FloorNumber of Elevator of Collumn) IsPressed  to false
                    SET FloorRequestButton (WITH Number SAME AS FloorNumber of Elevator of Collumn) IsPressed  to false
                    REMOVE element WITH number SAME AS FloorNumber of Elevator of Collumn FROM requests AND calls
        ENDIF
    'catch missed calls'
    CASE calls IS NOT empty
        SET DestinationFloor of Elevator of Collumn to first callElement Number
    CASE requests IS NOT empty
        SET DestinationFloor of Elevator of Collumn to first request
    ENDCASE
    ENDWHILE
        
    IF FloorNumber of Elevator SAME AS destinationFloor
        THEN SET elevator Movement to false
            CALL StartTimer WITH elevator
            CALL DoorsAction WITH Elevator of Collumn 
            SET callButtons WITH Number SAME AS destinationFloor IsPressed  to false
    
    ENDIF
    SET destinationFloor of elevator  to null
ENDSEQUENCE

SEQUENCE doorsAction USING elevator
    SET  open of  Doors of elevator to true
    wait openTime
   
    SET  safeToClose of  Doors of elevator to false
    WHILE safeToClose of  Doors of elevator is false
        CALL checkSafeToClose USING elevator Doors
            IF safeToClose of  Doors of elevator is true
                THEN close elevator Doors
                SET elevator Doors open to false
            ENDIF
    ENDWHILE
    
ENDSEQUENCE

SEQUENCE checkSafeToClose USING doors
    IF PassengerDetector of doors  is false 'not activated'
        SET SafeToClose of doors  to true
ENDSEQUENCE

SEQUENCE timerAdd USING Battery
    FOR every  collumn of Battery
        FOR every Elevator of collumn of Battery
            INCREMENT Timer of elevator of collumn of Battery by +1
        ENDFOR
    ENDFOR

ENDSEQUENCE

SEQUENCE StartTimer USING elevator
    SET  Timer of elevator to 0
ENDSEQUENCE

SEQUENCE goToIdle USING collumn and BuildingSpecs
    SET counter to 0
    FOR every elevator of collumn 
        IF Movement IS false
            THEN   
                INCREMENT counter by +1
        ENDIF
    ENDFOR

    
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
 
        SET idleFloor to (floors DIVIDED BY (counter PLUS 1) ROUNDED DOWN)
        SET counter to 1
        FOR every elevator of collumn 
            IF Movement IS false
                THEN
                    SET  destinationFloor of elevator of collumn to (idleFloor MULTIPLIES counter)  
                    INCREMENT counter by +1
                    CALL moveElevator WITH elevator of collumn 
            ENDIF
        ENDFOR
    ENDIF

ENDSEQUENCE

SEQUENCE alarm USING collumn AND BUILDING
    SET collumn Online to NOT BUILDING ALARM 'true to be false and vice versa'
    FOR every elevator of collumn 
        SET Online of elevator of collumn  to  Online of collumn
    ENDFOR
ENDSEQUENCE

SEQUENCE checkIfElevatorOnline USING collumn
    FOR every elevator OF collumn
        IF Online OF elevator OF collumn
            ADD elevator OF collumn to Online_elevators of collumn 
        ENDIF
    ENDFOR
ENDSEQUENCE
'''