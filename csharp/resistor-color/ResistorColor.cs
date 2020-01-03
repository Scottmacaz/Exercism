using System;

public static class ResistorColor
{
  private static string[] _resistorColors;
  static ResistorColor() => _resistorColors = new string[]
      {
            "black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"
      };

  public static int ColorCode(string color) => Array.IndexOf(_resistorColors, color);

  public static string[] Colors() => _resistorColors;
}