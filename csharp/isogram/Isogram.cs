using System;
using System.Collections.Generic;
using System.Linq;

public static class Isogram
{
  public static bool IsIsogram(string word)
  {
    var chars = word.ToCharArray().Where(c=>char.IsLetter(c)).Select(c => char.ToLower(c));
    return chars.Distinct().Count() == chars.Count();
   }
}
