def convert(number):
  sound = ""
  if number % 3 == 0:
    sound += "Pling"
  if number % 5 == 0:
    sound += "Plang"
  if number % 7 == 0:
    sound += "Plong"
  
  return sound if sound else str(number)