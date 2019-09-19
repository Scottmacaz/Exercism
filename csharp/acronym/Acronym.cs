using System;
using System.Linq;

public static class Acronym
{
  public static string Abbreviate(string phrase)
  {
    var words = phrase.Split(new char[] { ' ', '_', '-' }, StringSplitOptions.RemoveEmptyEntries);
    return string.Join("", words.Select(c => char.ToUpper(c[0])));
  }
}