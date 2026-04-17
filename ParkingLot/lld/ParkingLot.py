class ParkingLot:
    _instance = None 

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super().__new__(cls)
        return cls._instance

    def __init__(self,name_of_parking_lot,address):
        self.name_of_parking_lot= name_of_parking_lot
        self.address= address
        self.parking_floors=[]

    
