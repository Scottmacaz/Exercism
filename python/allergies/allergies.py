import math

class Allergies:
    names = ["eggs", "peanuts", "shellfish", "strawberries",
             "tomatoes", "chocolate", "pollen", "cats"]

    def __init__(self, score):
        self.score = score

    def allergic_to(self, item):
        
        # for i in range(len(self.names)):
        #     print(f"name: {self.names[i]}  value:{2 ** i}")

        # if self.score == 1:
        #     if item == "eggs":
        #         return True
        #     return False
        
        name_index = self.names.index(item)
        print (f"name_index: {name_index}")
        name_value = 2 ** name_index
        print(f"name_value: {name_value}")
        
         
       
        
        if s <= self.score:
            return True
        return False

    def _is_perfect_square(self,num) :
        root = math.sqrt(num)
        return int(root + .5) ** 2 == num
        

    @property
    def lst(self):
        pass


# c = Allergies(2)
# print(c.allergic_to("peanuts"))

c = Allergies(2)
print(c.allergic_to("eggs"))

# c = Allergies(2)
# print(c.allergic_to("shellfish"))