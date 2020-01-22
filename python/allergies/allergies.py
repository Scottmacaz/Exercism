class Allergies:
    names = ["eggs", "peanuts", "shellfish", "strawberries",
             "tomatoes", "chocolate", "pollen", "cats"]

    def __init__(self, score):
        self.score = score

    def allergic_to(self, item):
        if self.score == 1:
            if item == "eggs":
                return True
            return False
        
        i = self.names.index(item)
        s = (i+1) << (1)
        if s < self.score:
            return True
        print(f"({i} + 1) << (1) = {(i + 1) << (1)}")
        return False

    @property
    def lst(self):
        pass


c = Allergies(2)
print(c.allergic_to("shellfish"))
