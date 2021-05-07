package main

import (
	"encoding/csv"
    "fmt"
    "os"
	"strings"
)

type Question struct {
	ques, ans string
}

func main() {

	var questions []Question
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
        fmt.Println(err)
    }    
    for _, line := range csvLines {
        questions = append(questions, Question{
            ques: line[0],
            ans: line[1],
		})
    }
	var count int
	for i, text := range questions {
		var answer string
		fmt.Printf("Question %d: %s\n", i+1, text.ques)
		fmt.Scanln(&answer)

		if strings.Compare(answer, text.ans) == 0 {
			count += 1
		}
	}
	fmt.Printf("Total questions: %d, Correct answers: %d\n", len(questions), count)
}
