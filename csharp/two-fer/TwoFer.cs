using System;

public static class TwoFer
{
    public static string Speak(string name=null)
    {
      return name == null ? respond("you") : respond(name);
    }
    private static string respond(string name)
    {
      return $"One for {name}, one for me.";
    }
}
