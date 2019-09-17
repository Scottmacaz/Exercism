using System;
using System.Collections.Generic;

public static class Isogram
{
    public static bool IsIsogram(string word)
    {
      var charsFound = new List<char>();
      foreach(char c in word) 
      {
        //Only test for letters.
        if (!char.IsLetter(c))
        {
          continue;
        }
        if (charsFound.Contains(char.ToLower(c)))
        {
          return false;
        }
        charsFound.Add(char.ToLower(c));
      }
      return true;
    }
}
