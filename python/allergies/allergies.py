import math

class Allergies:
    NAMES = ["eggs", "peanuts", "shellfish", "strawberries",
             "tomatoes", "chocolate", "pollen", "cats"]

    def __init__(self, score):
        self.scores = []
        for i, name in enumerate(self.NAMES):
            if (1 << i) & score:
                self.scores.append(name)
                
    def allergic_to(self, item):
       return item in self.scores
   
    @property
    def lst(self):
        return self.scores

c = Allergies(2)
print(c.allergic_to("eggs"))
