package main

import ("fmt"
        "io/ioutil"
		"strings"
        //"math"
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

func find_embeds(a string) bool {
    b := strings.Replace(a,"-",",",-1)
    c := strings.Split(string(b),",")
    d := conv_slice_of_strings_to_int(c)
    var is_encompassed bool
    if d[2] >= d[0] && d[3] <= d[1] {
        is_encompassed = true
    }
    if d[2] <= d[0] && d[3] >= d[1] {
        is_encompassed = true
    }
    return is_encompassed
}

func find_any_overlaps(a string) bool {
    b := strings.Replace(a,"-",",",-1)
    c := strings.Split(string(b),",")
    d := conv_slice_of_strings_to_int(c)
    var is_encompassed bool
    if d[2] <= d[1] && d[3] >= d[0] {
        is_encompassed = true
    }
    return is_encompassed
}

//TFG. Gotta convert the slice of strings to a slice of ints. Damn it.
func conv_slice_of_strings_to_int(a []string) []int {
    my_int_slice := make([]int, 0)
    for _, y := range a {
        x, err := strconv.Atoi(y)
        if err != nil {
            panic(err)
        }
        my_int_slice = append(my_int_slice, x)
    }
    return my_int_slice
}

func solve_4a(a []string) int {
    my_embed_counter := 0
    for _, y := range a {
        if len(y) > 0 {
        if find_embeds(y) {
            my_embed_counter += 1
            }
        }
    }
    return my_embed_counter
}

func solve_4b(a []string) int {
    my_overlap_counter := 0
    for _, y := range a {
        if len(y) > 0 {
        if find_any_overlaps(y) {
            my_overlap_counter += 1
            }
        }
    }
    return my_overlap_counter
}

func main() {
    inputtext := fileimport("dcs_day4_input.txt")
    instructions := listmaker(inputtext)
    fmt.Println("4A:",solve_4a(instructions))
    fmt.Println("4B:",solve_4b(instructions))
}
