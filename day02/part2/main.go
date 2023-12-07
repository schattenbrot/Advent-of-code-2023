package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)

  var games []Game

  for scanner.Scan() {
    line := scanner.Text()

    games = append(games, createGame(line))
  }

  resultPart1 := 0
  resultPart2 := 0
  for _, game := range games {
    resultPart1 = resultPart1 + game.checkValidity()
    resultPart2 = resultPart2 + game.checkPowers()
  }

  log.Println("Result Part 1:", resultPart1)
  log.Println("Result Part 2:", resultPart2)


}

func createGame(line string) Game {
  gameSplit := strings.Split(line, ":")
  gameId, _ := strconv.Atoi(strings.Split(gameSplit[0], " ")[1])

  game := Game{
    Id: gameId,
  }

  rounds := strings.Split(gameSplit[1], ";")
  for _, round := range rounds {
    colors := strings.Split(round, ",")
    red := 0
    blue := 0
    green := 0

    for _, color := range colors {
      colorSplit := strings.Split(color, " ")
      amount, _ := strconv.Atoi(colorSplit[1])
      color := colorSplit[2]

      switch color {
        case "red":
          red = amount
          break
        case "green":
          green = amount
          break
        case "blue":
          blue = amount
          break
      }
    }

    gameRound := Round{
      Red: red,
      Green: green,
      Blue: blue,
    }

    game.addRound(gameRound)
  }
  return game
}
