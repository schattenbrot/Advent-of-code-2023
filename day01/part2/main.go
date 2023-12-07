package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
  file, err := os.Open("values.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  sum := 0

  for scanner.Scan() {
    line := scanner.Text()

    re := regexp.MustCompile("([0-9]|one|two|three|four|five|six|seven|eight|nine)")
    re2 := regexp.MustCompile("([0-9]|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)")
    firstNum := re.FindString(string(line))


    line2 := reverseLine(line)
    lastNum := re2.FindString(string(line2))

    convertStringNum(&firstNum)
    convertStringNum(&lastNum)

//    log.Println(string(firstNum))
  //  log.Println(string(lastNum))

    num,err := strconv.Atoi(string(firstNum) + string(lastNum))
    if err != nil {
      log.Fatal(err)
    }

    //log.Println(num)

    sum = sum + num
  }

  log.Println("The solution is:", sum)
}

func reverseLine(line string) string {
  var result string
  for _, r := range line {
    result = string(r) + result
  }
  return result
}

func convertStringNum(convertableNum *string) {
    switch string(*convertableNum) {
      case "one":
        *convertableNum = "1"
        break
      case "eno":
        *convertableNum = "1"
        break
      case "two":
        *convertableNum = "2"
        break
      case "owt":
        *convertableNum = "2"
        break
      case "three":
        *convertableNum = "3"
        break
      case "eerht":
        *convertableNum = "3"
        break
      case "four":
        *convertableNum = "4"
        break
      case "ruof":
        *convertableNum = "4"
        break
      case "five":
        *convertableNum = "5"
        break
      case "evif":
        *convertableNum = "5"
        break
      case "six":
        *convertableNum = "6"
        break
      case "xis":
        *convertableNum = "6"
        break
      case "seven":
        *convertableNum = "7"
        break
      case "neves":
        *convertableNum = "7"
        break
      case "eight":
        *convertableNum = "8"
        break
      case "thgie":
        *convertableNum = "8"
        break
      case "nine":
        *convertableNum = "9"
        break
      case "enin":
        *convertableNum = "9"
        break
    }
}

