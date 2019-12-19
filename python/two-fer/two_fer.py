def two_fer(name=""):
  rsp = "One for %s, one for me."
  name = name if name else "you"
  return rsp % name

two_fer("Test")
two_fer("")