package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type entry struct {
	question string
	answer   string
}

func main() {
	var entries []entry
	file, err := os.Open("data.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	z := html.NewTokenizer(file)
	q := true
	var entry entry
	for z.Err() == nil {
		byteTagName, _ := z.TagName()
		if string(byteTagName) == "td" {
			if z.Token().Type.String() == "StartTag" {
				z.Next()
				data := z.Token().String()
				if q {
					entry.question = data
				} else {
					entry.answer = data
					entries = append(entries, entry)
				}
				q = !q
			}
		}
		z.Next()
	}
	for true {
		var input string
		fmt.Print("input: ")
		_, err = fmt.Scan(&input)
		if err != nil {
			panic(err)
		}
		if err := clear(); err != nil {
			panic(err)
		}

		for _, entry := range entries {
			if strings.Contains(strings.ToLower(entry.question), input) {
				fmt.Println(entry.question, "||", entry.answer)
			}
		}
	}
}

func clear() error {
	cmd, err := exec.Command("clear").Output()
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(cmd)
	if err != nil {
		return err
	}
	return nil
}

func showDeadline() {
	myLayout := "Jan 2 03PM '06"
	deadLineTime := "Oct 15 08PM '24"
	deadline, err := time.Parse(myLayout, deadLineTime)
	if err != nil {
		panic(err)
	}
	timeUntil := deadline.Local().Truncate(time.Second).Format(myLayout)
	for time.Now().Truncate(time.Hour) != deadline.Truncate(time.Hour) {
		sinceDeadline := time.Since(deadline).Truncate(time.Second)
		fmt.Printf("sinceDeadline: %v\n", sinceDeadline)

		fmt.Printf("timeUntil: %v\n", timeUntil)

		time.Sleep(time.Second)
	}
}

func hpCalc() {
	maxHP := 50000
	potion := 14000 + 6000*5
	fmt.Println(maxHP - potion)
}
