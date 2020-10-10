# **-----------------------**
# **WEEK 1 Algorithms and Problem Solving**
# **-----------------------**
  

*this is the first week of Odyssey program*  

the exercise consists of writing the algorithms to control 2 elevator situations: one residential and one commercial (which will be derived more than likely from the first example).
These algorithms are to be written in Pseudocode according to codeboxx standards in 2 files , one for each situation.
a video presenting the project is also to be made 2-5 minutes

https://www.youtube.com/watch?v=uw5BK0z-czg  

**FIRST SCENARIO**
Residential building:
the algorythm must be implemented for : 
- 10 floors
- 1 column
- 2 elevator cages
- call buttons
- elevator doors
- floor request buttons

**SECOND SCENARIO**
Corporate building:
the algorithm controlls :
- 66 Floors ( 6 basements included)
- 1 battery
- 4 collumns
- 3 elevators per collumn (12 total)
- call buttons
- Doors
- Floor request buttons
- Floor display

ADDITIONAL REQUIREMENTS :  

*Logic of prioritization of elevators*
- elevators will be called first if moving towards the call
- second if stationary
- third if closer to their destination

*Logic back to origin*
- positions for idle lifts evenly spread out

*Security Logic status online/offline to be added and circomstances to be used*  

*Temporal logic*
- use building schedual to strategically position elevators  

*Loading Logic*
- detect if full (link to security)

  

# **-----------------------**
# **Week 2 Mechanics of Interpreted Languages**
# **-----------------------**
  

The second week of the Odyssey program is to convert the residential controller algorithm into interpretted languages.

- **JavaScript**
- **Python**  

The user first calls an elevator and when the elevator reaches the floor the user enters the elevator and requires a floor to which the elevator then moves.

  

They must contain the following methods :

- *Method 1: RequestElevator (RequestedFloor, Direction)*
- *Method 1 must return the chosen elevator and move the elevators in its treatment.*
  - *Method 2: RequestFloor (Elevator, RequestedFloor)*
  - *Method 2 must move the elevators in its treatment.*

  

The python version runs in terminal using : python3 Residential_Controller.py.

The javascript version runs in the terminal also using nodeJS. once installed use: node Residential_Controller.js needs prompt-sync to be installed (already installed and included in the repository).  

All programs run automatically creating seperate buildings for each test scenario and displaying the elevator status and floors in the terminal logs. 

**EXTRAS**

- both JS and PY check the elevators online status and if they are full before moving. scenario III will not work unless you reduce the load in the full elevator to below 10 000 on prompt 

- if you call Building.alarm() for any Building class (Scenario1, Scenario2, Scenario3) all online status will be switched to the opposite of the buildings alarm status that will log a warning message if it's alarm is switchend on. As displayed on scenB building (an example of a possibility to make buildings bigger and more complex in this case a 66 floor building with 4 columns of 5 elevators) the alarm is also on the first test and displays the consequences of pushing buttons and trying to run the sequence on an offline column. 

- all elevators check if they are overloaded before moving in scenario III you will have to enter a valid weight < 10 000 for it to run.

- In Javascript a delayed callback function will move the idle elevators to a calculated ideal position dependant on how many elevators are Idle . in python version the go to idle function will run at the end of each similation. 

- started work on a ruby file. It can run the simulations the same way as python and javascript but doesn't use the goToIdle method. 


#  **-----------------------**

#  **Week 3 Mechanics of Compiled Languages**

#  **-----------------------**

  

This week we transcribed the commercial algorythm to 2 co;piled languages: 

 **- c#**
 **- Golang**

the program is a console app in both languages that runs a simulation with a basic interface to select the scenario to run.

to run the c# you can use visual studio and run it from there or you need to go to the commercial_Controller_cs folder and in the teminal use dotnet run comand to run the program.

to run the Golang program go to the commercial_GoLang folder and run in the terminal commercial_controller.go file using :  go run  commercial_controller.go command.

both versions contain the required methods for a modern typed building :

 - **RequestElevator(FloorNumber)**
 - **AssignElevator(RequestFloor)**

to navigate the building with ease. to or from the ground floor.

Both programs display the information the user would have via display interfaces on the ground floor or inside the elevators and on the floors at each elevator door.

**EXTRAS**  

- both languages include in scenario4 a full elevator situation where the user will be asked to enter a newload value below the maximum of 10000

- in both languages the online status of the elevators are checked before completing the request and a list of online elevator is made allowing for on or some to deactivated for maintenance or cleaning... 
- a Drill scenario has been included in both languages where the building alarms are toggled thereby toggling all the collumns and elevators online status. It works verry well in C#. however in Golang it seams to work only if it is the first scenario to be ran if you run it a second time(putting all elevators back online) in both languages it will show a 5th scenario if all elevators are moving away from the user.
- In C# if you use the interface to exit the run cycle it will display last of all an example of a larger building with many more basements and calculations are made to seperate the collumns according to Rocket Elevators quotation criteria and the different information taken during the weeks.


I lacked the time to implement the goToIdle calculations in this weeks project so where the method would be called it simply logs to the console that the elevators return to Idle position.