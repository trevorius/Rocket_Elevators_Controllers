this is the first week of Odysey program 

the exercise consists of writing the algorythms to controle 2 elevator situations: one residential and one corporate (which will be derived more than likeley from the first example).

these algorythms are to be written in Pseudocode according to codeboxx standards in 2 files , one for each situation.
a video presenting the project is also to be made 2-5 minutes

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
the algorythm controlls :   - 1 battery
                            - 4 collumns
                            - 3 elevators per collumn (12 total)
                            - call buttons
                            - Doors
                            - Floor request buttons
                            - Floor display

ADDITIONAL REQUIREMENTS : 

    //Logic of priritization of elevators
            
            - elevators will be called  first if moving towards the call
                                        second if stationary
                                        third if closer to their destination    
    //Logic back to origin
            - positions for idle lifts
    
    Security Logic status online/offline to be added and circomstances to be used

    Temporal logic
            - use building schedual to strategically position elevators
            - measure waiting times for calls
    Loading Logic
            - detect if full (link to security)
            -use with waiting time to get waiting time closer to average waiting time. 