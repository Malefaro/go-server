package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestCalc(t *testing.T) {
	var Tests = struct {
		input    []string
		expected []int
	}{

		input: []string{
			"1 2 3 4 + * + =",
			"1 2 + 3 4 + * =",
			"2 3 + =",
			"2 3 * 4 5 * + =",
			"1 4 + 6 2 - * = ",
			"2 6 + 2 / =",
		},
		expected: []int{
			16, //15
			21,
			5,
			26,
			20,
			4,
		},
	}
	for i, str := range Tests.input {
		result, err := calc(str)
		if !reflect.DeepEqual(result, Tests.expected[i]) || err != nil {
			t.Error("expected :", Tests.expected[i], "but have :", result)
		}
	}
}

func TestCalcErrors(t *testing.T) {
	_, errorInStrconv := strconv.Atoi("qwe")
	var Tests = struct {
		input    []string
		expected []error
	}{

		input: []string{
			"1 2 3 4 + * + ",
			"1 2 + 3 4 + =",
			"2 3 + - =",
			"=",
			"1 / =",
			"1 + =",
			"2 * =",
			"1 qwe + =",
		},
		expected: []error{
			fmt.Errorf("ERROR : no '=' symbol"),
			fmt.Errorf("stack have more thn 1 element in '='"),
			fmt.Errorf("ERROR on operation '-', not enough operands in stack"),
			fmt.Errorf("ERROR in '=', stack is empty"),
			fmt.Errorf("ERROR on operation '/', not enough operands in stack"),
			fmt.Errorf("ERROR on operation '+', not enough operands in stack"),
			fmt.Errorf("ERROR on operation '*', not enough operands in stack"),
			errorInStrconv,
		},
	}
	for i, str := range Tests.input {
		_, err := calc(str)
		if !reflect.DeepEqual(err, Tests.expected[i]) {
			t.Error("\nexpected:", Tests.expected[i], "\nbut have:", err)
		}
	}
}

func TestCalcSpacesInput(t *testing.T) {
	var Tests = struct {
		input    []string
		expected []int
	}{

		input: []string{
			"1    2 3    4  +  * \t + =",
			"1 2 + 3 4 + * \n =",
			"2   3    +     =  ",
			"2 3 * 4 5 * +   =",
			"1   4  +   6  2 - * = ",
		},
		expected: []int{
			15,
			21,
			5,
			26,
			20,
		},
	}
	for i, str := range Tests.input {
		result, err := calc(str)
		if !reflect.DeepEqual(result, Tests.expected[i]) || err != nil {
			t.Error("expected :", Tests.expected[i], "but have :", result)
		}
	}
}
