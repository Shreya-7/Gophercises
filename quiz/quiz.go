package main

import (
	"encoding/csv"
    "fmt"
    "os"
	"flag"
	"strings"
	"time"
	"math/rand"
)

type Question struct {
	ques, ans string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func makeQuestions(csvLines [][]string, shuffle bool) []Question {

	questions := make([]Question, len(csvLines))

    for i, line := range csvLines {
        questions[i] = Question{
            ques: line[0],
            ans: strings.TrimSpace(line[1]),
		}
    }

	if shuffle {
		rand.Shuffle(len(questions), func(i, j int) {
			questions[i], questions[j] = questions[j], questions[i]
		})
	}

	return questions
}

func main() {

	// creation of command line flags

	fileName := flag.String("csv", "problems.csv", "A CSV file of the format question,answer")
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "Whether the questions should be shuffled")
	flag.Parse()

	// opening the given file

	csvFile, err := os.Open(*fileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the file %s\n", *fileName))
	}
	defer csvFile.Close()

	// parsing the given file 

	csvLines, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
        exit("Failed to parse the provided CSV file :(")
    }    

	// converting the file content into a structured form
	questions := makeQuestions(csvLines, *shuffle)

	// timer starts only after user presses `Enter`
	fmt.Printf("Press Enter to start the quiz!")
	fmt.Scanln()
	
	// create a timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// counts number of right answers
	var count int
	for i, text := range questions {

		fmt.Printf("Question %d: %s\n", i+1, text.ques)

		// channel which will handle the answer
		answerChannel := make(chan string)

		// goroutine to handle the blocking nature of `fmt.Scanln`
		go func() {
			var answer string
			fmt.Scanln(&answer)

			// send the user input into the channel
			answerChannel <- answer
		}()

		select {

		// if the timer is up - end quiz
		case <-timer.C :
			fmt.Printf("Total questions: %d, Correct answers: %d\n", len(questions), count)
			return

		// if user gives an answer
		case answer := <- answerChannel :	

			// check answer
			if strings.EqualFold(answer, text.ans) {
				count += 1
			}
		}		
	}

	// displayed if user finished quiz before time ends
	fmt.Printf("Total questions: %d, Correct answers: %d\n", len(questions), count)	
}
