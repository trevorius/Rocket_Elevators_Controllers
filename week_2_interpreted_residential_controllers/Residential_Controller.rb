
#----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

#DEFINE BuildingSpecs USING floors && collumns && elevators && _leavingTime && _arrivingTime
class Building   
  
  attr_accessor :floors, :columns , :columnList , :elevatorsPerColumn, :leavingTime, :arrivingTime, :aLARM, :numberOffloors, :numberOfcolumns, :_elevatorsPerColumn, :_leavingTime, :_arrivingTime

  def initialize( numberOffloors, numberOfcolumns, _elevatorsPerColumn, _leavingTime, _arrivingTime)
    @floors= numberOffloors
    @columns = numberOfcolumns
    @columnList = []
    @elevatorsPerColumn = _elevatorsPerColumn
    @leavingTime = _leavingTime
    @arrivingTime = _arrivingTime
    @aLARM = false

    create_columns 
  end


  def alarm 
    @aLARM = !@aLARM
    puts("WARNING! WARNING! WARNING! ALARM IS ON WARNING! WARNING! WARNING! ")
    for column in @columnList  
      column.online = !@aLARM
      for elevator in column.elevatorList  
        elevator.online = column.online
      end
    end
  end


=begin 
  """SEQUENCE timeCheck USING leavingTime && arrivingTime && time
    IF time IS GRETATER THAN (arrivingTime - 1hour) OR SMALLER THAN (arrivingTime + 1hour)   
      RETURN 1
    ENDIF
  ENDSEQUENCE
  """
=end
  
  #SEQUENCE timeCheck USING leavingTime && arrivingTime && time
  #  IF time IS GRETATER THAN (arrivingTime - 1hour) OR SMALLER THAN (arrivingTime + 1hour)   
  #    RETURN 1
  #  ENDIF
  #ENDSEQUENCE

  
  def create_columns 
    puts "creating column"   

    identity = 1
    while identity <= @columns 
      column = Column.new(identity, floors, elevatorsPerColumn)
      columnList << (column)
      identity += 1
    end
  end
  
end


  

#DEFINE cOLLUMN USING floors && elevators
class Column

  attr_accessor :identity, :floors, :numberOfElevators, :callButtonList, :elevatorList, :online, :onlineelevatorList, :callList, :floors, :elevatorsPerColumn


  def initialize(identity, floors, elevatorsPerColumn)
    @identity = identity
    @floors = floors
    @numberOfElevators = elevatorsPerColumn
    @callButtonList = []
    @elevatorList = []
    @online = true
    @onlineelevatorList = []
    @callList = []


    create_Elevators
    create_callButtons
  end

  def create_callButtons
    identity = 1
    while identity <= @floors
      callButton = CallButton.new(identity)
      @callButtonList << (callButton)
      identity += 1
    end
  end

  def create_Elevators
    identity = 1
    while identity <= numberOfElevators
      elevator = Elevator.new(identity, floors)
      @elevatorList << (elevator)
      identity += 1
    end
  end

  def create_callList 
    for callButton in @callButtonList 
      if callButton.isPressed && !@callList.include?(callButton)
        @callList << (callButton)
      end
    end
  end
  
  def create_onlineelevatorList 
    @onlineelevatorList = []
    if @online 
      for elevator in @elevatorList 
        if elevator.online 
          @onlineelevatorList << (elevator)
        end
      end
    end
  end


  def sortElevatorsBydistance ( destination)
    
    for elevator in @onlineelevatorList 
      elevator.distance = (destination - elevator.floorNumber ).abs
      #print "elevator distance is :"
      #puts elevator.distance
    end

    

    @onlineelevatorList.sort!{|elevatorA, elevatorB| elevatorA.distance <=> elevatorB.distance }  #(key=distance)
    puts("closest elevator is #{@onlineelevatorList[0].identity}")
  end

  def RequestElevator( requestedFloor, direction)
    create_onlineelevatorList()
    if @onlineelevatorList.length != 0    
      sortElevatorsBydistance(requestedFloor)
      selectedElevator = nil
      for elevator in @onlineelevatorList 
        #print " checking "
        #print elevator.identity
        #print "#{elevator.movement}\n" 

        if elevator.destinationFloor != nil
          distanceToDestination = (elevator.floorNumber - elevator.destinationFloor).abs
        else  
          distanceToDestination = 0
        end

        if  ((elevator.floorNumber >= requestedFloor && elevator.movement == direction && direction == "DOWN") || (elevator.floorNumber <= requestedFloor && elevator.movement == direction && direction == "UP"))
          puts "elevator is moving in the same direction as call" 
          if elevator.distance <= distanceToDestination    #'IF an elevator's travel takes it through button's floor on correct direction
            #puts "elevator is moving through destination"
            move(elevator)
          else                       # IF elevator travelling towards requestedFloor on correct direction
            elevator.destinationList << (requestedFloor)
            #puts " elevator is going towards floor "
          end
          
          selectedElevator = elevator
          puts("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX      elevator #{selectedElevator.identity} was chosen XXXXXXXXXXXXXXXXXXXXXXXXXX")
          move(selectedElevator)
          return selectedElevator
        
        

        
        elsif elevator.movement == "IDLE" 
          #puts "elevator was idle"
          if requestedFloor > elevator.floorNumber 
            elevator.movement = "UP"
          else  
            elevator.movement = "DOWN"
          end
          elevator.destinationList << (requestedFloor)
          
          
          selectedElevator = elevator
          puts("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX      elevator #{selectedElevator.identity} was chosen XXXXXXXXXXXXXXXXXXXXXXXXXX")
          move(selectedElevator)
          return selectedElevator
        end
      end

        
      if selectedElevator == nil
        #puts " all elevators are moving away from call"
         
        selectedElevator = @onlineelevatorList[-1]
        selectedElevator.destinationList << (requestedFloor)
          
        puts("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX      elevator #{selectedElevator.identity} was chosen XXXXXXXXXXXXXXXXXXXXXXXXXX")
        move(selectedElevator)
        return selectedElevator
      end
      
    else 
      puts("elevators of column #{@identity} are offline" )
      return
    end
  end



  def move ( elevator)

    while elevator.isElevatorFull() 

      puts("lOAD IS  #{elevator.lOAD}")
      puts ("ELEVATOR FULL!!!!")
      puts("Reduce load to less than 10000 ")
      elevator.lOAD = gets.to_i
    end

      
    elevator.destinationList.sort!()
    if elevator.movement == "UP"  
      elevator.destinationFloor = elevator.destinationList[-1]

    elsif elevator.movement == "DOWN" 
      elevator.destinationFloor = elevator.destinationList[0]
    else 
      elevator.destinationFloor = elevator.destinationList[0]
    end


    destination = elevator.destinationFloor
    puts ("elevator #{elevator.identity} is moving from #{elevator.floorNumber} to #{destination}")

    while elevator.floorNumber != destination   
      

      if elevator.movement == "UP" 
        elevator.floorNumber += 1
        
      elsif elevator.movement == "DOWN" 
        elevator.floorNumber -= 1
      end
      

        
      puts("elevator #{elevator.identity} is on floor#{elevator.floorNumber}")
      #Time.sleep(0.1)
      sleep(0.5)
      
      if elevator.destinationList.include?(elevator.floorNumber) 
        elevator.opendoors()

        request = FloorRequestButton.new(elevator.floorNumber, elevator.identity)
        request.isPressed = true
        
        if elevator.floorRequestButtonList.include?(request)    
          index = elevator.floorRequestButtonList.index(request)
          elevator.floorRequestButtonList[index].isPressed = false
          elevator.destinationList.remove(elevator.floorNumber)
        end
         
        # create call to check if its in the list
        call = CallButton.new(elevator.floorNumber)
        call.direction = elevator.movement
        call.isPressed = true

        if @callList.include?(call)  
          index = @callButtonList.index(call)
          @callButtonList[index].isPressed = false
          @callList.remove(call)
        end
      end
          

        elevator.floorRequestButtonList[elevator.floorNumber-1].isPressed = false
      if elevator.destinationList.length != 0 
        destination = elevator.destinationList[0]
        if destination > elevator.floorNumber 
          elevator.movement = "UP"
        else  
          elevator.movement = "DOWN"
        end
      end
    end

      
      
    if elevator.floorNumber == destination 
      elevator.movement = "IDLE" 
      puts("elevator #{elevator.identity} is IDLE")
      
      elevator.starttimer

      elevator.destinationList.delete(elevator.floorNumber)
      
      # @goToidle
    end
  end
     
  def RequestFloor( elevator, requestedFloor)
    if elevator != nil  
      if elevator.online == true 
        elevator.destinationList << (requestedFloor)
        move(elevator)
      end
    else  
      puts("request can't be dealt, there is no elevator")
      return
    end
  end


  def goToidle  

    create_onlineelevatorList()
    counter = 0
    for elevator in @onlineelevatorList  
      if elevator.movement == "IDLE" 
        #if int(elevator.timer) - int(time.now.to_i) > 300  

        counter += 1
      end
    
      #use timecheck day sequence here find at the end of file 


      idleFloor = math.floor(@floors / (counter +1))
      counter = 1
      for elevator in @onlineelevatorList  
        if elevator.movement == "IDLE" 
          if elevator.floorNumber == idleFloor  
            return
          else   
            elevator.destinationList << (idleFloor * counter)
            move(elevator)
            counter += 1
          end
        end
      end
    end
  end
end

#DEFINE CallButton USING floor && direction
class CallButton
  attr_accessor  :number, :direction, :isPressed

  def initialize ( floor)
    @number = floor
    @direction = nil
    @isPressed = false
  end
end


#DEFINE Elevator USING identity && Location && floors
class Elevator

  attr_accessor :identity, :floors, :floorNumber, :movement, :floorRequestButtonList, :destinationFloor, :destinationList, :distance, :doors, :timer, :online, :requestList, :mAXlOAD, :lOAD, :floorRequestButton

  def initialize(identity, floors)
    @identity = identity
    @floors = floors
    @floorNumber = 1
    @movement = "IDLE" #can be  up, down, or IDLE
    @floorRequestButtonList = []
    @destinationFloor = nil 
    @destinationList =[]
    @distance = nil
    @doors = nil
    @timer = Time.now.to_i
    @online = true
    @requestList = []

    @mAXlOAD = 10000 
    @lOAD = 0

    create_FloorRequestButtons()
    create_doors()
  end

  
 
  def create_FloorRequestButtons 
    number = 1
    while number <= floors
      floorRequestButton = FloorRequestButton.new(number, identity)
      #puts(floorRequestButton)
      #puts(@floorRequestButtonList)
      @floorRequestButtonList << (floorRequestButton)
      number +=1
    end
  end

  def create_doors 
    doors = Doors.new()
    @doors = doors
  end

  def isElevatorFull 
    if @lOAD < @mAXlOAD
      return false
    else 
      return true
    end
  end

  def opendoors 
    puts("openning doors")
    @doors.open = true
    @doors.safeToClose = false
    #time.sleep(@doors.openTime)
    sleep(@doors.openTime)
    
    while @doors.safeToClose == false 
      @doors.checksafeToClose 
    end
    
    if @doors.safeToClose == true 
      @doors.open = true
      puts("closing doors")
    end
  end
    


  def starttimer  
    @timer = Time.now.to_i
  end
end

class FloorRequestButton

  attr_accessor :identity, :number, :isPressed, :_floor 
  
  def initialize ( _floor, identity)
    @identity = identity
    @number = _floor
    @isPressed = false 
  end
end

class Doors
  attr_accessor :open, :openTime, :safeToClose, :passengerDetector
  def initialize  
    @open = false
    @openTime = 1
    @safeToClose = true
    @passengerDetector = false
  end

  def checksafeToClose   
    if @passengerDetector == false 
      @safeToClose = true 
    else  
      @safeToClose = false
    end
  end
end

#SET time to clockTIME
clocktime = Time.now.to_i



#------------------------------------------------------------------------------------------------------------------------
#    TEST ZONE
def run_test column, calls, requestList
  #puts"run_test"
  #for call in calls
    #print call.number
    #print "   "
  #end

  calls.reverse!()
  #puts "reversed : "

  #for call in calls
    #print call.number
    #print "   "
  #end

  
  counter = 0
  while calls.length() != 0 
    call = calls[-1]
    puts("checking call # #{counter +1}")
    tmp = column.RequestElevator(call.number,call.direction)

    calls.delete_at(-1)
    column.RequestFloor(tmp, requestList[counter])
    for elevator in column.onlineelevatorList  
      if elevator.destinationList != []       
        column.move(elevator)
      end
    end
    counter +=1
  end
end



build = Building.new(10, 1, 2, 8, 18)
cOLLUMN = build.columnList[0]
cALLLIST = cOLLUMN.callList

build.alarm()


build.columnList[0].callButtonList[5].isPressed = true
build.columnList[0].callButtonList[5].direction = "DOWN"  # to 3
build.columnList[0].create_callList()
build.columnList[0].callButtonList[7].isPressed = true
build.columnList[0].callButtonList[7].direction = "DOWN"  # to 5
build.columnList[0].create_callList()
build.columnList[0].callButtonList[4].isPressed = true
build.columnList[0].callButtonList[4].direction = "UP"    # to 10
build.columnList[0].create_callList()
build.columnList[0].callButtonList[6].isPressed = true
build.columnList[0].callButtonList[6].direction = "DOWN"  # to 2

cOLLUMN.elevatorList[0].floorNumber = 2
cOLLUMN.elevatorList[1].floorNumber = 3
requestList = [3,5,10,2]

build.columnList[0].create_callList()


run_test cOLLUMN, cALLLIST, requestList

#--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
#TEST SECTION


puts ("_______________________________________   SCENARIO I ________________________________________________")
scenario1 = Building.new(10, 1, 2, 8, 18)
scen1Column = scenario1.columnList[0]
scen1CALLS = scen1Column.callList
  
#'Elevator A is IDLE at floor 2'
scen1Column.elevatorList[0].floorNumber = 2
scen1Column.elevatorList[0].movement = "IDLE"
#'Elevator B is IDLE at floor 6'
scen1Column.elevatorList[1].floorNumber = 6
scen1Column.elevatorList[1].movement = "IDLE"

#'Someone is on floor 3 && wants to go to the 7th floor. '
scen1Column.callButtonList[2].isPressed = true
scen1Column.callButtonList[2].direction = "UP"
scen1Column.create_callList()
requestList = [7]
  
run_test scen1Column, scen1CALLS, requestList

  puts('Elevator A was expected to be sent from 3 tom 7.')
     

puts ("_______________________________________   SCENARIO II ________________________________________________")
scenario2 = Building.new(10, 1, 2, 8, 18)
scen2Column = scenario2.columnList[0]
scen2CALLS = scen2Column.callList

#'Elevator A is IDLE at floor 10 
scen2Column.elevatorList[0].floorNumber = 10
scen2Column.elevatorList[0].movement = "IDLE"
#'Elevator B is IDLE at floor 3
scen2Column.elevatorList[1].floorNumber = 3
scen2Column.elevatorList[1].movement = "IDLE"

#'Someone is on the 1st floor && requests the 6th floor. 
scen2Column.callButtonList[0].isPressed = true
scen2Column.callButtonList[0].direction = "UP"
scen2Column.create_callList()
requestList = []
requestList << (6)
#'2 minutes later, someone else is on the 3rd floor && requests the 5th floor. 
scen2Column.callButtonList[2].isPressed = true
scen2Column.callButtonList[2].direction = "UP"
scen2Column.create_callList()

requestList << (5)

#'Finally, a third person is at floor 9 && wants to go down to the 2nd floor. 
scen2Column.callButtonList[8].isPressed = true
scen2Column.callButtonList[8].direction = "DOWN"
scen2Column.create_callList()

requestList << (2)

run_test scen2Column, scen2CALLS, requestList 

puts('Elevator B (1to6) then B (3to5) && A (9to2) were expected to be sent.')



puts ("_______________________________________   SCENARIO III ________________________________________________")

scenario3 = Building.new(10, 1, 2, 8, 18)
scen3Column = scenario3.columnList[0]
scen3CALLS = scen3Column.callList

#'Elevator A is IDLE at floor 10 
scen3Column.elevatorList[0].floorNumber = 10
scen3Column.elevatorList[0].movement = "IDLE"
scen3Column.elevatorList[0].lOAD = 200000000000000
#'Elevator B is Moving from floor 3 to floor 6
scen3Column.elevatorList[1].floorNumber = 3
scen3Column.elevatorList[1].movement = "UP"
scen3Column.elevatorList[1].destinationFloor = 6
scen3Column.elevatorList[1].destinationList = [6]



#'Someone is on floor 3 && requests the 2nd floor. 

scen3Column.callButtonList[2].isPressed = true
scen3Column.callButtonList[2].direction = "DOWN"
scen3Column.create_callList()
requestList = []
requestList << (2)



#'5 minutes later, someone else is on the 10th floor && wants to go to the 3rd. 
scen3Column.callButtonList[9].isPressed = true
scen3Column.callButtonList[9].direction = "DOWN"
scen3Column.create_callList()

requestList << (3)




run_test scen3Column, scen3CALLS, requestList

puts("Elevator A (3to2) Then Elevator B (10to3) was expected to be sent.")
 
 #----------


puts " .\n \n \n NEW TEST TO CREATE A HUGE BUILDING"


scenB = Building.new(66, 4, 5, 8, 18)

scenB.alarm

=begin
#----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
#EVENT LISTENERS
'''
 
    WHEN a Callbutton IS pressed
      CALL CallElevatorCheck WITH cOLLUMN

    WHEN a FloorRequestButton IS pressed
      CALL usersDestination WITH cOLLUMN

    WHEN 1 second passes
      CALL timerAdd WITH BUILDING

    WHEN 'elevator is IDLE for 5 mins'
      a  movement of Elevator of cOLLUMN IS false && timer of Elevator of cOLLUMN  is 300 
      CALL goToidle WITH cOLLUMN

    WHEN BUILDING aLARM IS CHANGED
      CALL alarm WITH cOLLUMN && BUILDING

'''
#----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
=end

=begin        """
        Call timeCheck of BuildingSpecs RETURNING idleFloor
        IF idleFloor ! null THEN
          FOR every elevator of cOLLUMN 
            IF movement IS false
              THEN
                SET  destinationFloor of elevator of cOLLUMN to idleFloor 
                CALL moveElevator WITH elevator of cOLLUMN 
            ENDIF
          ENDFOR
        ELSE 
        """
=end
 