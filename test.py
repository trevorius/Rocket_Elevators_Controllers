#DEFINE BuildingSpecs USING floors AND columns AND elevators AND _leavingTime AND _arrivingTime
 
class Building:
    ALARM = False
    ColumnList = []
    def __init__(self, numberOfFloors, numberOfColumns, _elevatorsPerColumn, _leavingTime, _arrivingTime):
        self.Floors = numberOfFloors
        self.Columns = numberOfColumns
        
        self.elevatorsPerColumn = _elevatorsPerColumn
        self.leavingTime = _leavingTime
        self.arrivingTime = _arrivingTime

        self.create_Columns()
    
    #SEQUENCE timeCheck USING leavingTime AND arrivingTime AND time
    #   IF time IS GRETATER THAN (arrivingTime - 1hour) OR SMALLER THAN (arrivingTime + 1hour)   
    #        RETURN 1
    #    ENDIF
    #ENDSEQUENCE

    
    def create_Columns(self):
        ID = 1
        while ID <= self.Columns:
            column = Column(ID, self.Floors, self.elevatorsPerColumn)
            self.ColumnList.append(column)
            ID += 1
    

        

class Column:
    def __init__(self, ID, _floors, _elevators):
        self.ID = ID
        self.Floors = _floors
        self.elevators = _elevators
        self.elevatorList = []

        self.create_elevators()

    def create_elevators(self):
        
        for ID in range(self.elevators):
            elevator = Elevator(ID+1,self.Floors)
            self.elevatorList.append(elevator)
            

class Elevator:
    location = 1
    def __init__(self, ID, _floors):
        self.ID = ID
        self.Floors = _floors
    



build = Building (10, 1, 2, 8, 18)

'''
print (build.Columns)
print (build.ColumnList[0].ID)
print (build.ColumnList[0].elevatorList[1].ID)

rang = range(build.Floors)
print (rang)

for x in rang :
    print (x)
'''

import time
timeIs = time.time()
print (time.ctime(timeIs))