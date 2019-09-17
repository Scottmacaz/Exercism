using System;

public static class Leap
{
    public static bool IsLeapYear(int year)
    {
        var yearDivisibleBy4 = year % 4 == 0;
        var yearDivisibleBy100 = year % 100 == 0;
        var yearDivisibleBy400 = year % 400 == 0;

        if (yearDivisibleBy4)
        {
            if (yearDivisibleBy100)
            {
                if (yearDivisibleBy400)
                {
                    return true;
                }
                return false;
            }
            return true;
        }
        return false;
    }
}