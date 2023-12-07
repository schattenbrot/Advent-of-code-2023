package main

type Round struct {
  Red int
  Green int
  Blue int
}

type Game struct {
  Id int
  Rounds []Round
}

func (g *Game) addRound(round Round) {
  g.Rounds = append(g.Rounds, round)
}

func (g *Game) checkValidity() int {
  maxRed := 12
  maxGreen := 13
  maxBlue := 14

  for _, round := range g.Rounds {
    if round.Red > maxRed || round.Green > maxGreen || round.Blue > maxBlue {
      return 0
    }
  }
  return g.Id
}


