using System;
using System.Collections.Generic;

public static class Pangram
{
    public static bool IsPangram(string input)
    {
      var charsFound = new HashSet<char>();
        foreach (var c in input)
        {
          if (Char.IsLetter(c))
          {
            charsFound.Add(char.ToLower(c));
          }
        }
      return charsFound.Count == 26;
    }
}
