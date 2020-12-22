package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"regexp"
	"strings"
)

const simpleParenthesysExp = `\([^\(\)]+\)`

func evaluate(expression string) int {
	re := regexp.MustCompile(simpleParenthesysExp)
	for re.MatchString(expression) {
		expression = re.ReplaceAllStringFunc(expression, func(s string) string {
			return strconv.Itoa(evaluate(s[1 : len(s)-1]))
		})
	}
	return left2right(expression)
}

func evaluateP2(expression string) int {
	re := regexp.MustCompile(simpleParenthesysExp)
	for re.MatchString(expression) {
		expression = re.ReplaceAllStringFunc(expression, func(s string) string {
			return strconv.Itoa(evaluateP2(s[1 : len(s)-1]))
		})
	}
	re = regexp.MustCompile(`\d+ \+ \d+`)
	for re.MatchString(expression) {
		expression = re.ReplaceAllStringFunc(expression, func(s string) string {
			return strconv.Itoa(left2right(s))
		})
	}
	return left2right(expression)
}

func left2right(expression string) int {
	ans := 0

	operator := ""
	for _, s := range strings.Split(expression, " ") {
		if s == "+" || s == "*" {
			operator = s
		} else {
			i, _ := strconv.Atoi(s)
			if operator == "+" {
				ans += i
			} else if operator == "*" {
				ans *= i
			} else {
				ans = i
			}
			operator = ""
		}
	}

	return ans
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	ansP1 := 0
	ansP2 := 0
	for _, line := range lines {
		ansP1 += evaluate(line)
		ansP2 += evaluateP2(line)
	}
	fmt.Println("P1:", ansP1)
	fmt.Println("P2:", ansP2)
}