using System;

public static class Acronym
{
    public static string Abbreviate(string phrase)
    {
        var words = phrase.Split(" ");
        var acronym = string.Empty;
        foreach (var word in words)
        {
          acronym += GetAcronymLetters(word);
        }
        return acronym;
    }

    private static string GetAcronymLetters(string word)
    {
      string acronymLetters = string.Empty;
      //If it's a single character and not a letter then it's not part of the acronym.
      if (word.Length == 1 && char.IsLetter(word[0]) == false)
      {
        return string.Empty;
      }
      
      var hyphenWords = word.Split("-");
      foreach(var hyphenWord in hyphenWords)
      {
        if( char.IsLetter(hyphenWord[0]))
        {
          acronymLetters +=  char.ToUpper(hyphenWord[0]);
        }
        else 
        {
          acronymLetters += char.ToUpper(hyphenWord[1]);
        }
      }
      return acronymLetters;
          
      
    }
    
}