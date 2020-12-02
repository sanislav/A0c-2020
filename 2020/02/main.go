package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strconv"
)


func Readln(r *bufio.Reader) (string, error) {
	var (isPrefix bool = true
		 err error = nil
		 line, ln []byte
		)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}

	return string(ln),err
}

func isCorrect(s string) []bool {
	boundaries := getCharBoundaries(s)
	char := getChar(s)
	pass := getPassword(s)

	charInPass := regexp.MustCompile(char)

	matches := charInPass.FindAllStringIndex(pass, -1)

	correctV1 := false
	correctV2 := false

	if (len(matches) >= boundaries[0] && len(matches) <= boundaries[1]) {
		correctV1 = true
	}

	if ((string(pass[boundaries[0] -1]) == char && string(pass[boundaries[1] - 1]) != char) || (string(pass[boundaries[0] -1]) != char && string(pass[boundaries[1] - 1]) == char)) {
		correctV2 = true
	}

	return []bool{correctV1, correctV2}
}

func getCharBoundaries(s string) []int {
	var re = regexp.MustCompile("([0-9]*)-([0-9]*) ")
	res := re.FindStringSubmatch(s)
	lower, _ := strconv.Atoi(res[1])
	upper, _ := strconv.Atoi(res[2])

	return []int{lower, upper}
}

func getChar(s string) string {
	var re = regexp.MustCompile("([a-z]): ")
	res := re.FindStringSubmatch(s)

	return res[1]
}

func getPassword(s string) string {
	var re = regexp.MustCompile(": ([a-z]*)")
	res := re.FindStringSubmatch(s)

	return res[1]
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v\n",err)
		os.Exit(1)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	s, e := Readln(r)
	correctPasswordsV1 := 0;
	correctPasswordsV2 := 0;
	for e == nil {
		s, e = Readln(r)

		if len(s) > 0 {
			correctness := isCorrect(s)
			if correctness[0] {
				correctPasswordsV1++
			}
			if correctness[1] {
				correctPasswordsV2++
			}
		}
	}

	fmt.Println("Number of correct V1 passwords is:", correctPasswordsV1)
	fmt.Println("Number of correct V2 passwords is:", correctPasswordsV2)
}