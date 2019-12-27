import random
import string

names = set()

class Robot:
    name = ''

    def __init__(self):
        print(random.choice(string.ascii_uppercase) for i in range(2))
        self.name = self.__makeName()
        
    def reset(self):
        self.__init__()

    def __makeName(self):
        random.seed()
        newName = ''.join(random.choice(string.ascii_uppercase) for i in range(
            2)) + ''.join(random.choice(string.digits) for i in range(3))
        if newName in names:
            self.__makeName()
        names.add(newName)
        return newName
