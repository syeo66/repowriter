package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var pixels = map[rune][]string{
	'a': {
		" O ",
		"O O",
		"OOO",
		"O O",
		"O O",
	},
	'b': {
		"OO ",
		"O O",
		"OO ",
		"O O",
		"OO ",
	},
	'c': {
		" O ",
		"O O",
		"O  ",
		"O O",
		" O ",
	},
	'd': {
		"OO ",
		"O O",
		"O O",
		"O O",
		"OO ",
	},
	'e': {
		"OOO",
		"O  ",
		"OO ",
		"O  ",
		"OOO",
	},
	'f': {
		"OOO",
		"O  ",
		"OO ",
		"O  ",
		"O  ",
	},
	'g': {
		" OO",
		"O  ",
		"O O",
		"O O",
		" OO",
	},
	'h': {
		"O O",
		"O O",
		"OOO",
		"O O",
		"O O",
	},
	'i': {
		"O",
		"O",
		"O",
		"O",
		"O",
	},
	'j': {
		"  O",
		"  O",
		"  O",
		"  O",
		"OO ",
	},
	'k': {
		"O O",
		"O O",
		"OO ",
		"O O",
		"O O",
	},
	'l': {
		"O  ",
		"O  ",
		"O  ",
		"O  ",
		"OOO",
	},
	'm': {
		"O   O",
		"OO OO",
		"O O O",
		"O   O",
		"O   O",
	},
	'n': {
		"O   O",
		"OO  O",
		"O O O",
		"O  OO",
		"O   O",
	},
	'o': {
		" O ",
		"O O",
		"O O",
		"O O",
		" O ",
	},
	'p': {
		"OO ",
		"O O",
		"OO ",
		"O  ",
		"O  ",
	},
	'q': {
		" OOO ",
		"O   O",
		"O O O",
		"O  OO",
		" OOO ",
	},
	'r': {
		"OO ",
		"O O",
		"OO ",
		"O O",
		"O O",
	},
	's': {
		" OO",
		"O  ",
		" O ",
		"  O",
		"OO ",
	},
	't': {
		"OOO",
		" O ",
		" O ",
		" O ",
		" O ",
	},
	'u': {
		"O O",
		"O O",
		"O O",
		"O O",
		" O ",
	},
	'v': {
		"O   O",
		"O   O",
		"O   O",
		" O O ",
		"  O  ",
	},
	'w': {
		"O   O",
		"O   O",
		"O O O",
		"O O O",
		" O O ",
	},
	'x': {
		"O O",
		"O O",
		" O ",
		"O O",
		"O O",
	},
	'y': {
		"O O",
		"O O",
		" O ",
		" O ",
		" O ",
	},
	'z': {
		"OOO",
		"  O",
		" O ",
		"O  ",
		"OOO",
	},
	' ': {
		"  ",
		"  ",
		"  ",
		"  ",
		"  ",
	},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	output := make([]string, 5)

	for _, a := range strings.ToLower(text) {
		matrix := pixels[a]
		for i, line := range matrix {
			if len(output[i]) > 0 {
				output[i] += " "
			}
			output[i] += line
		}
	}

	for _, line := range output {
		fmt.Println(line)
	}

	datelist := createDateList(output)
	for _, date := range datelist {
		for x := 0; x < 30; x++ {
			createCommit(date)
			date = date.Add(1 * time.Minute)
		}
	}
}

func createCommit(date time.Time) {
	f, err := os.Create("date.txt")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)

	_, err = fmt.Fprintf(w, "%v\n", date)
	check(err)
	w.Flush()

	err = exec.Command("git", "add", "date.txt").Run()
	check(err)
	err = exec.Command("git", "commit", "--date", date.Format("2006-01-02 15:04:05"), "-m", "Commit for "+date.Format("2006-01-02")).Run()
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func createDateList(output []string) []time.Time {
	date := lastFriday()
	dateList := []time.Time{}

	for x := len(output[0]) - 1; x >= 0; x-- {

		for y := 4; y >= 0; y-- {

			if output[y][x] == 'O' {
				dateList = append(dateList, time.Time(date))
			}
			date = date.AddDate(0, 0, -1)
		}

		date = date.AddDate(0, 0, -2)
	}

	return dateList
}

func lastFriday() time.Time {
	today := time.Now()
	lastFriday := time.Date(today.Year(), today.Month(), today.Day(), 1, 0, 0, 0, time.Local)
	for lastFriday.Weekday() != time.Friday {
		lastFriday = lastFriday.AddDate(0, 0, -1)
	}
	return lastFriday
}
