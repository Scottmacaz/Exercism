using System;
using System.Collections.Generic;
using System.Linq;

public static class Isogram
{
  public static bool IsIsogram(string word)
  {
    var charsFound = new List<char>();
    foreach (char c in word)
    {
      if (char.IsLetter(c))
      {
        if (charsFound.Contains(char.ToLower(c)))
        {
          return false;
        }
        charsFound.Add(char.ToLower(c));
      }
    }
    return true;
  }
}
