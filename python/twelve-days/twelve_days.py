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


def recite(start_verse, end_verse):
    return [make_verse(i) for i in range(start_verse, end_verse+1)]
   
def make_verse(verse_num):
    verse = ""
    verse += f"On the {days[verse_num-1]} day of Christmas my true love gave to me: "
    for day in range(12-int(verse_num), 12):
        if day == 11 and verse_num > 1:
            verse += f"and {gifts[day]}"
        else:
            verse += f"{gifts[day]}"
    return verse
