package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func part2() {
  log.Println("=== ====== ===")
  log.Println("=== PART 2 ===")
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
    log.Println(line)

    re := regexp.MustCompile("([0-9]|one|two|three|four|five|six|seven|eight|nine)")
    test := re.FindAllString(string(line), -1)
    if test == nil {
      log.Fatal("nope")
    }
    for _, x := range test {
      log.Println(string(x))
    }

    firstNum := test[0]
    lastNum := test[len(test)-1]

    convertStringNum(&firstNum)
    convertStringNum(&lastNum)

    /*
    log.Println(string(firstNum))
    log.Println(string(lastNum))
*/
    num,err := strconv.Atoi(string(firstNum) + string(lastNum))
    if err != nil {
      log.Fatal(err)
    }

    log.Println(num)

    sum = sum + num
  }

  log.Println("The solution is:", sum)
}

func convertStringNum(convertableNum *string) {
    switch string(*convertableNum) {
      case "one":
        *convertableNum = "1"//[]byte("1")
        break
      case "two":
        *convertableNum = "2"//[]byte("2")
        break
      case "three":
        *convertableNum = "3"//[]byte("3")
        break
      case "four":
        *convertableNum = "4"//[]byte("4")
        break
      case "five":
        *convertableNum = "5"//[]byte("5")
        break
      case "six":
        *convertableNum = "6"//[]byte("6")
        break
      case "seven":
        *convertableNum = "7"//[]byte("7")
        break
      case "eight":
        *convertableNum = "8"//[]byte("8")
        break
      case "nine":
        *convertableNum = "9"//[]byte("9")
        break
    }
}
