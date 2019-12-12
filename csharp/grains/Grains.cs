using System;
using System.Collections.Generic;

public static class Grains
{
  private static Dictionary<int, ulong> _boardPosition;

  static Grains()
  {
    _boardPosition = new Dictionary<int, ulong>();
    ulong grainCount = 1;
    _boardPosition[1] = grainCount;
    for (var i = 2; i <= 64; i++)
    {
      grainCount *= 2;
      _boardPosition[i] = grainCount;
    }
  }
  public static ulong Square(int n)
  {
    if (n <= 0 || n > 64)
    {
      throw new ArgumentOutOfRangeException("n must be between 1 and 64");
    }
    return _boardPosition[n];
  }

  public static ulong Total()
  {
    ulong total = 0;
    foreach (var value in _boardPosition.Values)
    {
      total += value;
    }
    return total;
  }

}