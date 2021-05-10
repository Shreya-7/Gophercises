# Timed Quiz using Go

Program to read in a quiz provided via a CSV file and gives the quiz to the user to play.   

## Features:
1. Keeps track of how many questions were right.
2. Regardless of whether the answer is right or not, the next question is asked.
3. Quizzes have single word/number answers.
4. Displays the total number of questions and the number of correct questions at the end.
5. Input CSV file defaults to "problems.csv" and is customizable using the `-csv` flag.
6. Quiz is timed - the user will not be able to see more questions or enter answers after it is up.
7. Timer defaults to 30 and is customizable using the `-limit` flag in seconds.
8. Questions are not shuffled by default. Customizable using the `-shuffle` flag.
9. Answers are do not differentiate because of whitespaces and case sensitivity.

## Concepts Learned:
1. Writing a normal Go program.
2. Reading and parsing CSV files using the `encoding/csv` package.
3. Using the `flag` package to control command line arguments - setting names, defaults & helptext.
4. Working with structures.
5. Timers and Sleep using the `time` package.
6. Creation of and message handling using channels.
7. Goroutine basics.