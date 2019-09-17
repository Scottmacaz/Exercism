using System;

public static class Leap
{
    public static bool IsLeapYear(int year)
    {
        var yearDivisibleBy4 = year % 4 == 0;
        var yearDivisibleBy100 = year % 100 == 0;
        var yearDivisibleBy400 = year % 400 == 0;

        return yearDivisibleBy4 && !yearDivisibleBy100 || yearDivisibleBy400;
    }
}