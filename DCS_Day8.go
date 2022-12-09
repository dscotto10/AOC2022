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


 func build_my_maps(my_input []string) (map[string]tree)  {
      my_trees := make(map[string]tree)
      for i := 0; i < len(my_input); i++ {
          my_row := strings.Split(my_input[i],"")
          for j, val := range my_row {
              point_name := build_tree_name(j,i)
              my_trees[point_name] = tree{name:point_name, x_pos: j, y_pos:i, height:cstri(val)}
          }
      }
     return my_trees
}

func get_array_size(my_input []string) (int, int) {
    my_max_x := len(my_input[0]) - 1
    my_max_y := len(my_input) - 1
    return my_max_x, my_max_y
}

func build_tree_name(x int, y int) string {
    my_point_name := "x" + strconv.Itoa(x) + "y" + strconv.Itoa(y)
    return my_point_name
}

func check_if_visible(the_tree string, my_trees map[string]tree, my_max_x int, my_max_y int) map[string]tree {
    checking_tree := my_trees[the_tree]
    checking_tree.is_visible = 1
    is_visible_n := 1
    is_visible_s := 1
    is_visible_e := 1
    is_visible_w := 1
    trees_n := 0
    trees_s := 0
    trees_e := 0
    trees_w := 0
    is_edge := 0
    if checking_tree.y_pos == 0 || checking_tree.y_pos == my_max_y || checking_tree.x_pos == 0 || checking_tree.x_pos == my_max_x {
        is_edge = 1
    }
    if is_edge == 0 {
        for i := checking_tree.y_pos - 1; i >= 0; i-- {
            checker := build_tree_name(checking_tree.x_pos,i)
            trees_n += 1
            if checking_tree.height <= my_trees[checker].height {
                is_visible_n = 0
                break
            }
        }
        for i := checking_tree.y_pos + 1; i <= my_max_y; i++ {
            checker := build_tree_name(checking_tree.x_pos,i)
            trees_s += 1
            if checking_tree.height <= my_trees[checker].height {
                is_visible_s = 0
                break
            }
        }
        for i := checking_tree.x_pos - 1; i >= 0; i-- {
            checker := build_tree_name(i, checking_tree.y_pos)
            trees_w += 1
            if checking_tree.height <= my_trees[checker].height {
                is_visible_w = 0
                break
            }
        }
        for i := checking_tree.x_pos + 1; i <= my_max_x; i++ {
            checker := build_tree_name(i, checking_tree.y_pos)
            trees_e += 1
            if checking_tree.height <= my_trees[checker].height {
                is_visible_e = 0
                break
            }
        }
        if is_visible_e + is_visible_n + is_visible_s + is_visible_w == 0 {
            checking_tree.is_visible = 0
        }
        checking_tree.scenic_score = trees_n * trees_s * trees_e * trees_w
    }
    my_trees[the_tree] = checking_tree
    return my_trees
}

func count_visibles(my_trees map[string]tree) int {
    visible_count := 0
    for k := range my_trees {
        if my_trees[k].is_visible == 1 {
            visible_count += 1
        }
    }
    return visible_count
}

func get_max_scenic_score(my_trees map[string]tree) int {
    max_scenic_score := 0
    for k := range my_trees {
        if my_trees[k].scenic_score > max_scenic_score {
            max_scenic_score = my_trees[k].scenic_score
        }
    }
    return max_scenic_score
}

func main() {
    inputtext := fileimport("dcs_day8_input.txt")
    tree_slice := listmaker(inputtext)
    tree_map := build_my_maps(tree_slice)
    array_width, array_height := get_array_size(tree_slice)
    for k, _ := range tree_map {
        tree_map = check_if_visible(k,tree_map,array_width,array_height)
    }
    fmt.Println("8A:",count_visibles(tree_map))
    fmt.Println("8B:",get_max_scenic_score(tree_map))
}
