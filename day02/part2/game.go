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

func (g *Game) checkPowers() int {
  red := 0
  green := 0
  blue := 0
  
  for _, round := range g.Rounds {
    if round.Red > red {
      red = round.Red
    }
    if round.Green > green {
      green = round.Green
    }
    if round.Blue > blue {
      blue = round.Blue
    }
  }

  return red * green * blue
}

