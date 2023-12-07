package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
  log.Println("=== ====== ===")
  log.Println("=== PART 1 ===")
  log.Println("=== ====== ===")
  file, err := os.Open("values.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)

  sum := 0

  for scanner.Scan() {
    line := scanner.Text()
    lineInversed := reverseLine(line)

    re := regexp.MustCompile("[0-9]")
    firstNum := re.Find([]byte(line))
    secondNum := re.Find([]byte(lineInversed))

    num,err := strconv.Atoi(string(firstNum) + string(secondNum))
    if err != nil {
      log.Fatal(err)
    }

    sum = sum + num
  }

  log.Println("The solution is:", sum)
}

func reverseLine(line string) (string) {
  var result string
  for _, v := range line {
    result = string(v)+result
  }
  return result
}
