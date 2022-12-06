package main

import ("fmt"
        "io/ioutil"
		"strings"
        //"math"
		//"strconv"
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
func listmaker(a []uint8) string {
	b := strings.Split(string(a), "\n")
	return b[0]
}

//This one sets up a map, and if it finds the same item twice in the map, it breaks and moves on.
func check_for_no_matches(a string) int {
    my_map_check := make(map[string] int)
    my_match := 0
    for i:=0; i < len(a); i++ {
        _, isPresent := my_map_check[string(a[i])]
        if isPresent {
            my_match = 1
            break
        } else {
            my_map_check[string(a[i])] = 1
        }
    }
    return my_match
    }

//Function with two parameters: the input string, and the target set length of distincts.
//Basically: run a loop through every possible character string of length "set_length."
//Then break it when you find a unique set.
//That relies on check_for_no_matches.
func run_string_check(my_string string, set_length int) int {
    my_answer := 0
    for x:=0; x < len(my_string) - set_length; x++ {
        partial_string := my_string[x:x + set_length]
        if check_for_no_matches(partial_string) == 0 {
            my_answer = x + set_length
            break
        }
    }
    return my_answer
}
    

func main() {
    inputtext := fileimport("dcs_day6_input.txt")
    string_to_test := listmaker(inputtext)
    fmt.Println("6A:",run_string_check(string_to_test,4))
    fmt.Println("6B:",run_string_check(string_to_test,14))    
}
