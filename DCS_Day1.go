package main

import ("fmt"
        "io/ioutil"
	"strings"
	"strconv"
        "sort"
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

//Find a max value in a slice. This is the 1A solution.
 func find_max(a []int) int {
     my_max := a[0]
     for i := 0; i < len(a); i++ {
         if a[i] > my_max {
             my_max = a[i]
         }
     }
     return my_max
 }

//Sort the slice descending, and sum up the first 3 items in it.
func get_top_3(a []int) int {
    sort.Slice(a, func(i, j int) bool {
        return a[j] < a[i]
    })
    return a[0] + a[1] + a[2] 
    
}
 
//Most of the work is here. It's making an array of all of the users.
func find_most_calories(a []string) []int {
    my_instructions := a
    calorie_sum := 0
    calorie_sums := make([]int, 1)
    for i := 0; i < len(my_instructions); i++    {
        if len(my_instructions[i]) > 0 {
            calorie_int, err := strconv.Atoi(my_instructions[i])
            //Handle errors.
            if err != nil {
                fmt.Print(err)
            }
            calorie_sum += calorie_int
        } else {
            calorie_sums = append(calorie_sums,calorie_sum)
            calorie_sum = 0
        }
    }
    return calorie_sums
}

func main() {
    inputtext := fileimport("dcs_day1_input.txt")
    
    instructions := listmaker(inputtext)
    elf_calories := find_most_calories(instructions)
    max_calories := find_max(elf_calories)
    top_3_elves := get_top_3(elf_calories)
    
    fmt.Println("1A:",max_calories)
    fmt.Println("1B:",top_3_elves)

    
}
