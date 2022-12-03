package main

import ("fmt"
        "io/ioutil"
		"strings"
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
func listmaker(a []uint8) []string {
	b := strings.Split(string(a), "\n")
	return b
}

//This is to make 3B easier.
//Going to put the "sets" together in a pipe-delimited slice.
//Each item in the slice will be the set for comparison.
func listmaker_3b(a []string) []string {
    new_string := ""
    c := make([]string, 0)
    for i := 0; i < len(a) - 1; i++ {
        new_string = new_string + "|" + a[i]
        if (i + 1) % 3 == 0 {
            c = append(c,new_string[1:])
            new_string = ""
        }
    }
    return c
}

//We need to break the rucksack in two; this does that and makes it a slice of two objects.
func split_string_in_half(a string) []string {
    halved_string := make([]string, 2)
    halved_string[0] = a[0:len(a) / 2]
    halved_string[1] = a[len(a) / 2:]
    return halved_string
}

//This procedure finds the character that is duplicated in both sides of the rucksack.
func locate_duplicate(a []string) string {
    my_comparison_map := make(map[string]int)
    my_duplicate := ""
    for i := 0; i < len(a[0]); i++ {
        my_comparison_map[string(a[0][i])] = 1
    }
    for i := 0; i < len(a[1]); i++ {
        _, isPresent := my_comparison_map[string(a[1][i])]
        if isPresent {
            my_duplicate = string(a[1][i])
        } 
    }
    return my_duplicate
}

//This procedure converts a character into the code equivalent, via the character's ASCII code.
func convert_char_to_aoc_ascii(my_char string) int {
    aoc_char_code := 0
    my_rune := []rune(my_char)
    if int(my_rune[0]) >= 97 {
        aoc_char_code = int(my_rune[0]) - 96
    } else {
        aoc_char_code = int(my_rune[0]) - 38
    }
    return aoc_char_code
}

//Just score it with the other processes.
func score_3a(rucksacks []string) int {
    my_score := 0
    for i := 0; i < len(rucksacks) - 1; i++ {
        rucksack_split := split_string_in_half(rucksacks[i])
        dupe_char := locate_duplicate(rucksack_split)
        my_score += convert_char_to_aoc_ascii(dupe_char)
    }
    return my_score
}
//Making maps. Lots of maps.
//Logic here: maps 1-3 are basically SELECT DISTINCT on each of the strings.
//Map 4 is a map that combines each. Only one should get to 3.
func make_maps(three_set string) string {
    //fmt.Println(string(three_set))
    b := strings.Split(string(three_set), "|")
    my_comparison_map_1 := make(map[string]int)
    my_comparison_map_2 := make(map[string]int)
    my_comparison_map_3 := make(map[string]int)
    my_comparison_map_4 := make(map[string]int)
    my_duplicate := ""    
    for i := 0; i < len(b[0]); i++ {
        my_comparison_map_1[string(b[0][i])] = 1
    }
     for i := 0; i < len(b[1]); i++ {
         my_comparison_map_2[string(b[1][i])] = 1
     }
    for i := 0; i < len(b[2]); i++ {
        my_comparison_map_3[string(b[2][i])] = 1
    }
    for k := range my_comparison_map_1 {
        my_comparison_map_4[k] += 1
    }
    for k := range my_comparison_map_2 {
        my_comparison_map_4[k] += 1
    }
    for k := range my_comparison_map_3 {
        my_comparison_map_4[k] += 1
    }
    //fmt.Println(my_comparison_map_4)
    for k, v := range my_comparison_map_4 {
        if v == 3 {
            my_duplicate = k
        }
   }
    return my_duplicate
    }
    
func get_score_3b(my_instructions []string) int {
    score_3b := 0
    for _, y := range my_instructions {
        my_char := make_maps(string(y))
        score_3b += convert_char_to_aoc_ascii(my_char)
    }
    return score_3b
}

func main() {
    inputtext := fileimport("dcs_day3_input.txt")
    instructions := listmaker(inputtext)
    fmt.Println("3A:",score_3a(instructions))
    instructions_3b := listmaker_3b(instructions)
    fmt.Println("3B:",get_score_3b(instructions_3b))

}
