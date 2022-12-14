package main

import ("fmt"
        "io/ioutil"
		"strings"
        "sort"
        //"math"
		"strconv"
)

type monkey struct {
    name int
    items_list []int
    operation string
    test_divisible int
    test_true_target int
    test_false_target int
    inspect_count int
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
    c := make([]string,0)
    for _, item := range b {
        item = strings.Replace(item,":","",-1)
        c = append(c,strings.TrimSpace(item))
        //fmt.Println(item)
    }
	return c
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

func parse_input(a []string) map[int]monkey {
    my_monkey_map := make(map[int]monkey)
    active_monkey := -1
    for _, item := range a {
        b := strings.Split(string(item)," ")
        if len(b) > 1 {
            if b[0] == "Monkey" {
                my_monkey_map[cstri(b[1])] = monkey{name:cstri(b[1])}
                active_monkey = cstri(b[1])
            } else if b[0] == "Operation" {
                item = strings.Replace(item,"Operation new = old ","",-1)
                current_monkey := my_monkey_map[active_monkey]
                current_monkey.operation = item
                my_monkey_map[active_monkey] = current_monkey
            } else if b[0] == "Starting" {
                item = strings.Replace(item,"Starting items ","",-1)
                b := strings.Split(string(item),", ")
                current_monkey := my_monkey_map[active_monkey]
                current_monkey.items_list = make_slice_of_strings_into_ints(b)
                my_monkey_map[active_monkey] = current_monkey
            } else if b[0] == "Test" {
                current_monkey := my_monkey_map[active_monkey]
                current_monkey.test_divisible = cstri(b[3])
                my_monkey_map[active_monkey] = current_monkey
            } else if b[1] == "true" {
                current_monkey := my_monkey_map[active_monkey]
                current_monkey.test_true_target = cstri(b[5])
                my_monkey_map[active_monkey] = current_monkey
            } else if b[1] == "false" {
                current_monkey := my_monkey_map[active_monkey]
                current_monkey.test_false_target = cstri(b[5])
                my_monkey_map[active_monkey] = current_monkey
                //my_monkey_map = update_item_in_monkey(my_monkey_map, active_monkey, test_false_target, cstri(b[5]))
            } 
        }
    }
    return my_monkey_map
}

func make_slice_of_strings_into_ints(a []string) []int {
    z := make([]int,0)
    for _, y := range a {
        z = append(z,cstri(y))
    }
    return z
}

func figure_out_operation(instruction string, old_worry_level int) int {
    var new_worry_level int
    b := strings.Split(string(instruction)," ")
    if b[1] == "old" {
        b[1] = strconv.Itoa(old_worry_level)
    }
    if b[0] == "*" {
        new_worry_level = cstri(b[1]) * old_worry_level
    } else if b[0] == "+" {
        new_worry_level = cstri(b[1]) + old_worry_level
    }
    new_worry_level = new_worry_level / 3
    //fmt.Println(string_slice_print(b))
    //fmt.Println(new_worry_level)
    return new_worry_level
}

func delete_item_from_slice(a []int, b int) []int {
    c := make([]int,0)
    for _, item := range a {
        if item != b {
            c = append(c,item)
        }
    }
    return c
}

func run_round(my_map map[int]monkey) map[int]monkey {
    for x:=0; x < len(my_map); x++ {
        current_monkey := my_map[x]
        for _, item := range current_monkey.items_list {
            current_monkey.inspect_count += 1
            //fmt.Println(item)
            new_item := figure_out_operation(current_monkey.operation,item)
            //fmt.Println(new_item)
            if new_item % current_monkey.test_divisible == 0 {
                receiving_monkey := my_map[current_monkey.test_true_target]
                receiving_monkey.items_list = append(receiving_monkey.items_list,new_item)
                current_monkey.items_list = delete_item_from_slice(current_monkey.items_list,item)
                my_map[my_map[x].test_true_target] = receiving_monkey
            } else {
                receiving_monkey := my_map[current_monkey.test_false_target]
                receiving_monkey.items_list = append(receiving_monkey.items_list,new_item)
                current_monkey.items_list = delete_item_from_slice(current_monkey.items_list,item)
                my_map[my_map[x].test_false_target] = receiving_monkey
            }
            my_map[x] = current_monkey
            //fmt.Println(my_map)
        }
    }
    return my_map
}

func get_top_two_inspect_counts(my_map map[int]monkey) int {
    inspect_counts_slice := make([]int,0)
    for _, v := range my_map {
        inspect_counts_slice = append(inspect_counts_slice, v.inspect_count)
    }  
    sort.Slice(inspect_counts_slice, func(i, j int) bool {
    return inspect_counts_slice[i] > inspect_counts_slice[j]
    })
    return inspect_counts_slice[0] * inspect_counts_slice[1]
}

// func update_item_in_monkey(my_map monkey, active_monkey, attribute, update_val) monkey {
//     current_monkey := my_map[active_monkey]
//     current_monkey.attribute = update_val
//     my_map[active_monkey] = current_monkey
//     return my_map
// }

func main() {
    inputtext := fileimport("dcs_day11_input.txt")
    instructions := listmaker(inputtext)
    //fmt.Println(string_slice_print(instructions))
    monkeys := parse_input(instructions)
    //fmt.Println(monkeys)
    for i := 0; i < 20; i++ {
        monkeys = run_round(monkeys)
    }
    my_number := get_top_two_inspect_counts(monkeys)
    fmt.Println(my_number)
}
