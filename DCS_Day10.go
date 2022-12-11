package main

import ("fmt"
        "io/ioutil"
		"strings"
        //"unicode"
        //"math"
		"strconv"
)

	
type tree struct {
    name string
    x_pos int 
    y_pos int
    height int
    is_visible int
    scenic_score int
}

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
    b = b[:len(b)-1]
	return b
}

func string_slice_print(my_slice []string) string {
    print_string := "["
    for _, y := range my_slice {
        print_string = print_string + "'" + y + "'" + ","
    }
    print_string = print_string + "]" 
    print_string = strings.Replace(print_string,",]","]",-1)
    return print_string
}

func cstri(a string) int {
    new_string, err := strconv.Atoi(strings.TrimSpace(a))
    if err != nil {
        panic(err)
    }
    return new_string
}

func run_through_list_10a(my_instructions []string) map[int]int {
    regx := 1
    cycle_counter := 1
    var line_string string
    signal_strengths := make(map[int]int)
    for _, v := range my_instructions {
        current_instruction := strings.Split(v," ")
        if cycle_counter % 40 == 0 {
            fmt.Println(line_string)
            line_string = ""
        }
        if current_instruction[0] == "noop" {
            cycle_counter += 1
            signal_strengths[cycle_counter] = cycle_counter * regx
            line_string = line_string + what_to_draw(cycle_counter,regx)
        } else {
            cycle_counter += 1
            signal_strengths[cycle_counter] = cycle_counter * regx
            line_string = line_string + what_to_draw(cycle_counter,regx)
            if cycle_counter % 40 == 0 {
                fmt.Println(line_string)
                line_string = ""
            }
            regx += cstri(current_instruction[1])
            cycle_counter += 1
            signal_strengths[cycle_counter] = cycle_counter * regx
            line_string = line_string + what_to_draw(cycle_counter,regx)
        }
    }
    fmt.Println(line_string)
    return signal_strengths
}

func what_to_draw(a int, b int) string {
    return_val := ""
    if (a - b) % 40 <= 2 && (a - b) % 40 >= -2 {
        return_val = "#"
    } else {
        return_val = "."
    }
    return return_val
}

func main() {
    inputtext := fileimport("dcs_day10_input.txt")
    instructions := listmaker(inputtext)
    strengths := run_through_list_10a(instructions)
    fmt.Println("10A:",strengths[20] + strengths[60] + strengths[100] + strengths[140] + strengths[180] + strengths[220])
}
