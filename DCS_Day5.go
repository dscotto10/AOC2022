//This one was VERY educational. Freaking Golang.

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
    

func conv_list_into_map(a []string) (map[int][]string, int) {
    my_stack_map := make(map[int] []string)
    last_row_of_crates := 0
    for k, i := range a {
        if len(i) > 0 {
        if string(i[1]) == "1" {
            last_row_of_crates = k 
            map_key_slice := strings.Split(string(i[1:]),"   ")
            for _, j := range map_key_slice {
                x := cstri(j)
                //x, err := strconv.Atoi(strings.TrimSpace(j))
            //if err != nil {
                //panic(err)
            //}
                my_stack_map[x] = make([]string, 0)
            }
            //fmt.Println(my_stack_map)
        }
        }
    }
    for i := 0; i < last_row_of_crates; i++ {
        for j := 0; j < len(a[i]); j++ {
            if (j - 1) % 4 == 0 && string(a[i][j]) != " " {
                my_stack_map[(j + 3) / 4] = append(my_stack_map[(j+3) / 4],string(a[i][j]))
                }
            }
        }
    return my_stack_map, last_row_of_crates + 2
}

func work_instruction_5a(my_map map[int][]string, instruction string) map[int][]string {
    instruction = strings.Replace(instruction,"move ","",-1)
    instruction = strings.Replace(instruction," from ","|",-1)
    instruction = strings.Replace(instruction," to ","|",-1)
    instruction_slice := strings.Split(instruction,"|")
    for i := 0; i < cstri(instruction_slice[0]); i++ {
        dest_slice := my_map[cstri(instruction_slice[2])]
        source_slice := my_map[cstri(instruction_slice[1])]
        dest_slice = append([]string{source_slice[0]},dest_slice...)
        my_map[cstri(instruction_slice[1])] = source_slice[1:]
        my_map[cstri(instruction_slice[2])] = dest_slice 
        //slice_0_position_move(my_map[cstri(instruction_slice[2])],my_map[cstri(instruction_slice[1])])
    }
    return my_map
}

//Oh my goodness so frustrating. This is because "naive" merging of slices bleeds across other slices.
func merge_two_slices(s1 []string, s2[]string) []string {
    var slice []string
    for i := range s1 {
        slice = append(slice, string(s1[i]))
    }
    for i := range s2 {
        slice = append(slice, string(s2[i]))
    }
    return slice
}

//Needs work
 func work_instruction_5b(my_map map[int][]string, instruction string) map[int][]string {
     instruction = strings.Replace(instruction,"move ","",-1)
     instruction = strings.Replace(instruction," from ","|",-1)
     instruction = strings.Replace(instruction," to ","|",-1)
     instruction_slice := strings.Split(instruction,"|")
     temp_slice := make([]string, 20)
     source_slice := make([]string, 20)
     dest_slice := make([]string, 20)
     dest_slice = my_map[cstri(instruction_slice[2])]
         source_slice = my_map[cstri(instruction_slice[1])]
         temp_slice = source_slice[0:cstri(instruction_slice[0])]
         //Note that append(temp_slice,dest_slice...) fails REALLY badly here.
         dest_slice = merge_two_slices(temp_slice,dest_slice)
         my_map[cstri(instruction_slice[1])] = source_slice[cstri(instruction_slice[0]):]
                  my_map[cstri(instruction_slice[2])] = dest_slice
    
     return my_map
 }

func string_slice_print(my_slice []string) string {
    print_string := "["
    for _, y := range my_slice {
        print_string = print_string + "'" + y + "'" + ","
    }
    print_string = print_string + "]" 
    return print_string
}

func cstri(a string) int {
    new_string, err := strconv.Atoi(strings.TrimSpace(a))
    if err != nil {
        panic(err)
    }
    return new_string
}

func get_top_crate_on_each_stack(my_map map[int][]string) string {
    my_string := ""
    for k := 1; k <= len(my_map); k++ {
        my_string = my_string + my_map[k][0]
    }
    return my_string
}

func main() {
    inputtext := fileimport("dcs_day5_input.txt")
    //a_data := []string{"A", "B", "C", "D"}
    //b_data := []string{"E", "F", "G", "H"}
    //a_data, b_data = slice_0_position_move(a_data,b_data)
    instructions := listmaker(inputtext)
    //fmt.Println(string_slice_print(instructions))
    my_stacks, my_last_row := conv_list_into_map(instructions)
    //fmt.Println(my_stacks)
    //fmt.Println(my_last_row)
    for y:= my_last_row; y < len(instructions) - 1; y++ {
        my_stacks = work_instruction_5a(my_stacks,instructions[y])
        //fmt.Println(instructions[y])
        //fmt.Println(my_stacks)
    }
    fmt.Println("5A:",get_top_crate_on_each_stack(my_stacks))
    
    
    my_stacks, my_last_row = conv_list_into_map(instructions)
    //fmt.Println(my_stacks)
     for y:= my_last_row; y < len(instructions) - 1; y++ {
         my_stacks = work_instruction_5b(my_stacks,instructions[y])
         //fmt.Println(instructions[y])
         //fmt.Println(my_stacks)
     }
     fmt.Println("5B:",get_top_crate_on_each_stack(my_stacks))
   
    
    //     fmt.Println("4A:",solve_4a(instructions))
//     fmt.Println("4B:",solve_4b(instructions))
}

    
//  func slice_0_position_move(dest_slice []string,source_slice []string) ([]string, []string) {
//      dest_slice = append([]string{source_slice[0]},dest_slice...)
//      replace_source_slice := source_slice[1:]
//      return dest_slice, replace_source_slice
//  }
