package main

import ("fmt"
        "io/ioutil"
		"strings"
		"strconv"
)

//File import function.
func fileimport(a string) []uint8 {

b, err := ioutil.ReadFile(a)
		
	if err != nil {
        	fmt.Print(err)
    	}
    	return b
}
//Split the input into a list on line breaks.
func listmaker(a []uint8) []string {
	b := strings.Split(string(a), "\n")
	return b
}

//Straightforward conversion of rock/paper/scissors into 1/2/3.
func numberify_plays_2a(a string) string {
    a = strings.Replace(a, "A", "1", -1)
    a = strings.Replace(a, "X", "1", -1)
    a = strings.Replace(a, "B", "2", -1)
    a = strings.Replace(a, "Y", "2", -1)
    a = strings.Replace(a, "C", "3", -1)
    a = strings.Replace(a, "Z", "3", -1)
    return a
}

//More complex conversion. We're using this procedure to figure out the correct play.
//And then we're inserting it into the old model.
//This is so that we can reuse the scoring code from 2A.
func numberify_plays_2b(a string) string {
    a = strings.Replace(a, "A", "1", -1)
    a = strings.Replace(a, "B", "2", -1)
    a = strings.Replace(a, "C", "3", -1)
    b := strings.Split(string(a)," ")
    their_play, err := strconv.Atoi(b[0])
    my_play := 0
    //Handle errors.
    if err != nil {
        fmt.Print(err)
       }
    if b[1] == "X" {
       my_play = their_play - 1
       if my_play == 0 {
           my_play = 3
       }
    } else if b[1] == "Y" {
        my_play = their_play
    } else if b[1] == "Z" {
        my_play = ( their_play  % 3 ) + 1
    }
    my_string_play := strconv.Itoa(my_play)
    a = strings.Replace(a,b[1],my_string_play,-1)
    return a
}

func get_score(a string) int {
    //fmt.Println(a)
    b := strings.Split(string(a)," ")
    my_play, err := strconv.Atoi(b[1])
    //Handle errors.
    if err != nil {
        fmt.Print(err)
       }
    their_play, err := strconv.Atoi(b[0])
    //Handle errors.
    if err != nil {
        fmt.Print(err)
       }
    my_score := my_play
    if my_play - their_play == 0 {
        my_score += 3
    } else if my_play - their_play == 1 || my_play - their_play == -2 {
        my_score += 6
    }
    return my_score
}
    
func run_through_games_2a(a []string)  int {
    running_score := 0
    for i:= 0; i < len(a) - 1; i++ {
        numberified_record := numberify_plays_2a(a[i])
        running_score += get_score(numberified_record)
    }
    return running_score
}

func run_through_games_2b(a []string)  int {
    running_score := 0
    for i:= 0; i < len(a) - 1; i++ {
        numberified_record := numberify_plays_2b(a[i])
        running_score += get_score(numberified_record)
    }
    return running_score
}

func main() {
    inputtext := fileimport("dcs_day2_input.txt")
    instructions := listmaker(inputtext)
    fmt.Println("2A:",run_through_games_2a(instructions))
    fmt.Println("2B:",run_through_games_2b(instructions))    
}
