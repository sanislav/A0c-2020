package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strings"
)

func solve() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n",err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	completeDataPassports := 0
	validPassports := 0
	info := ""
	for scanner.Scan() {
		lineText := scanner.Text()
		if len(lineText) > 0 {
			info += lineText + " "
		} else {
			spaces := regexp.MustCompile(" ")

			matches := spaces.FindAllString(info, -1)
			if (len(matches) == 8) {
				completeDataPassports++
				if (isValid(info)) {
					validPassports++
				}
			} else if (len(matches) == 7) {
				pid := strings.Contains(info, "cid")
				if ! pid {
					completeDataPassports++
					if (isValid(info)) {
						validPassports++
					}
				}
			}

			info = ""
		}
	}

	return []int{completeDataPassports, validPassports}
}

func isValid(info string) bool {
	validPid := regexp.MustCompile("pid:[0-9]{9} ")
	matchedPid := validPid.FindString(info)
	if len(matchedPid) == 0 {
		return false
	}

	validEcl := regexp.MustCompile("ecl:(amb|blu|brn|gry|grn|hzl|oth) ")
	matchedEcl := validEcl.FindString(info)
	if len(matchedEcl) == 0 {
		return false
	}

	validHcl := regexp.MustCompile("hcl:#[a-f0-9]{6} ")
	matchedHcl := validHcl.FindString(info)
	if len(matchedHcl) == 0 {
		return false
	}

	validHgt := regexp.MustCompile("hgt:1[5-8]{1}[0-9]{1}cm |19[0-3]{1}cm |hgt:59in |hgt:[6-8]{1}[0-9]{1}in |hgt:9[0-3]{1}in ")
	matchedHgt := validHgt.FindString(info)
	if len(matchedHgt) == 0 {
		return false
	}

	validEyr := regexp.MustCompile("eyr:202[0-9]{1} |eyr:2030 ")
	matchedEyr := validEyr.FindString(info)
	if len(matchedEyr) == 0 {
		return false
	}

	validIyr := regexp.MustCompile("iyr:201[0-9]{1} |iyr:2020 ")
	matchedIyr := validIyr.FindString(info)
	if len(matchedIyr) == 0 {
		return false
	}

	validByr := regexp.MustCompile("byr:19[2-9]{1}[0-9]{1} |byr:200[0-2]{1} ")
	matchedByr := validByr.FindString(info)
	if len(matchedByr) == 0 {
		return false
	}

	return true
}

func main() {
	completeData := solve()
	fmt.Println("Number of complete data passports:", completeData[0])
	fmt.Println("Number of valid data passports:", completeData[1])
}