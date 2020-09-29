-----------------------WEEK 1 Algorithms and Problem Solving---------------------------------------

        this is the first week of Odyssey program 

        the exercise consists of writing the algorithms to control 2 elevator situations: one residential and one commercial (which will be derived more than likely from the first example).

        these algorithms are to be written in Pseudocode according to codeboxx standards in 2 files , one for each situation.
        a video presenting the project is also to be made 2-5 minutes
                https://www.youtube.com/watch?v=uw5BK0z-czg

        FIRST SCENARIO

        Residential building:
        -10 floors
        the algorythm must be implemented for : - 1 column
                                                - 2 elevator cages
                                                - call buttons
                                                - elevator doors
                                                - floor request buttons
                                        
        SECOND SCENARIO

        Corporate building:
        -66 Floors ( 6 basements included)
        the algorithm controlls :   - 1 battery
                                - 4 collumns
                                - 3 elevators per collumn (12 total)
                                - call buttons
                                - Doors
                                - Floor request buttons
                                - Floor display

        ADDITIONAL REQUIREMENTS : 

        Logic of prioritization of elevators
                
                - elevators will be called  first if moving towards the call
                                                second if stationary
                                                third if closer to their destination    
        Logic back to origin
                - positions for idle lifts evenly spread out 
        
        Security Logic status online/offline to be added and circomstances to be used

        Temporal logic
                - use building schedual to strategically position elevators

        Loading Logic
                - detect if full (link to security)

--------------------------------------- Week 2  Mechanics of Interpreted Languages ---------------------------------------------

The second week of the Odyssey program is to  convert the residential controller algorithm into interpretted languages.
        JavaScript
        Python

the user first  calls an elevator and when the elevator reaches the floor the user enters the elevator and requires a floor to which the elevator then moves.

they must contain  the following methods :

        Method 1: RequestElevator (RequestedFloor, Direction)
        Method 1 must return the chosen elevator and move the elevators in its treatment.

        Method 2: RequestFloor (Elevator, RequestedFloor)
        Method 2 must move the elevators in its treatment.

