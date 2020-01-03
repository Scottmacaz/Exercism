import re


class PhoneNumber:
    def __init__(self, number):
        self.number = re.sub('[^0-9]', '', number)
        if len(self.number) > 11 or len(self.number) < 10:
            raise ValueError("Phone number must be either 10 or 11 digits")
        if len(self.number) == 11:
            if self.number[0] == '1':
                self.number = self.number[1:]
            else:
                raise ValueError("Area code prefix must start with a 1.")
        if self.number[0] == '0' or self.number[0] == '1':
            raise ValueError("Area code cannot start with zero or one.")
        if self.number[3] == '0' or self.number[3] == '1':
            raise ValueError("Exchange code cannot start with zero or one.")
        self.area_code = self.number[0:3]

    def pretty(self):
        return f"({self.number[0:3]}) {self.number[3:6]}-{self.number[6:11]}"
