def recite(start_verse, end_verse):
    days = ['first', 'second', 'third', 'fourth', 'fifth', 'sixth',
      'seventh', 'eighth', 'ninth', 'tenth', 'eleventh', 'twelfth']

    gifts = ["twelve Drummers Drumming, ",
            "eleven Pipers Piping, ",
            "ten Lords-a-Leaping, ",
            "nine Ladies Dancing, ",
            "eight Maids-a-Milking, ",
            "seven Swans-a-Swimming, ",
            "six Geese-a-Laying, ",
            "five Gold Rings, ",
            "four Calling Birds, ",
            "three French Hens, ",
            "two Turtle Doves, ",
            "a Partridge in a Pear Tree."
            ]
    
    song = ""
    for verse in range(start_verse, end_verse+1):
        song += f"On the {days[verse-1]} day of Christmas my true love gave to me: "
        for day in range(12-int(verse), 12):
            if day == 11 and end_verse > 1:
                    song += f"and {gifts[day]}"
            else :
                song += f"{gifts[day]}"
    return [song]


print (recite(1,1))
# print("--------------------------")
# recite(1,2)
# print("--------------------------")
# recite(9,12)