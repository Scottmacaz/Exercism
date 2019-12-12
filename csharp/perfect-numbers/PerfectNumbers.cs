using System;
using System.Collections.Generic;
using System.Linq;

public enum Classification
{
    Perfect,
    Abundant,
    Deficient
}

public static class PerfectNumbers
{
    public static Classification Classify(int number)
    {
      if (number <= 0)
      {
        throw new ArgumentOutOfRangeException("number must be greater than zero");
      }
        var factors = FindFactors(number);
        var aliquotSum = factors.Sum();
       
        if (aliquotSum == number) 
        {
          return Classification.Perfect;
        }
        if (aliquotSum > number)
        {
          return Classification.Abundant;
        }
        return Classification.Deficient;
    }
    private static IEnumerable<int> FindFactors(int number)
    {
      var factors = new List<int>();
      for (var x=1; x < number; x++)
      {
          if (number % x == 0)
          {
            factors.Add(x);
          }
      }
      return factors;
    }
    
}
