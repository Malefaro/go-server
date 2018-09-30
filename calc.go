package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func calc(str string) (int, error) {
	arr := strings.Split(str, " ")
	var stack []int = make([]int, 0, 0)
	var i int = 0
	for _, val := range arr {
		switch val {
		case "", " ", "\n", "\t":
			continue
		case "+":
			if i < 2 {
				return 0, fmt.Errorf("ERROR on operation '+', not enough operands in stack")
			}
			stack[i-2] = stack[i-2] + stack[i-1]
			i--
		case "-":
			if i < 2 {
				return 0, fmt.Errorf("ERROR on operation '-', not enough operands in stack")
			}
			stack[i-2] = stack[i-2] - stack[i-1]
			i--
		case "*":
			if i < 2 {
				return 0, fmt.Errorf("ERROR on operation '*', not enough operands in stack")
			}
			stack[i-2] = stack[i-2] * stack[i-1]
			i--
		case "/":
			if i < 2 {
				return 0, fmt.Errorf("ERROR on operation '/', not enough operands in stack")
			}
			stack[i-2] = stack[i-2] / stack[i-1]
			i--
		case "=":
			if i == 0 {
				return 0, fmt.Errorf("ERROR in '=', stack is empty")
			} else if i > 1 {
				return 0, fmt.Errorf("stack have more thn 1 element in '='")
			}
			return stack[i-1], nil
		default:

			num, err := strconv.Atoi(val)
			if err != nil {
				return 0, err
			}
			if len(stack) <= i {
				stack = append(stack, num)
				i++
			} else {
				stack[i] = num
				i++
			}
		}
	}
	return 0, fmt.Errorf("ERROR : no '=' symbol")

}

func main() {
	//in := bufio.NewScanner(os.Stdin)
	//in.Scan()
	//var str string = in.Text()
	//result, err := calc(str)
	//if err == nil {
	//	fmt.Println(result)
	//} else {
	//	fmt.Println(err)
	//}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		result, _ := calc("3 3 * 2 + =")
		fmt.Fprintf(w,"2 + 3 * 3 = %d",result )
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

