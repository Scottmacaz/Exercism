def is_pangram(sentence):
  cSet = set()
  for c in sentence: 
    if c.isalpha():
      cSet.add(c.lower())

  return len(cSet) == 26