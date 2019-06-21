package main

import (
	"errors"
	"syscall/js"
)

var (
	global js.Value
)

func init() {
	global = js.Global()
}

func main() {
	global.Set("calculate", js.FuncOf(calculate))

	select {}
}

func calculate(_ js.Value, args []js.Value) interface{} {
	// Here goes the glue code
	a := args[0].Float()
	b := args[1].Float()
	operation := args[2].Int()

	var result float64
	var err error

	switch operation {
	case 1:
		result, err = add(a, b)
	case 2:
		result, err = rest(a, b)
	case 3:
		result, err = mult(a, b)
	case 4:
		result, err = div(a, b)
	default:
		err = errors.New("that operation is not known")
	}

	if err != nil {
		return map[string]interface{}{
			"result": nil,
			"error":  err.Error(),
		}
	}

	return map[string]interface{}{
		"result": result,
		"error":  nil,
	}
}

func add(a, b float64) (float64, error) {
	return a + b, nil
}

func rest(a, b float64) (float64, error) {
	return a - b, nil
}

func mult(a, b float64) (float64, error) {
	return a * b, nil
}

func div(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, errors.New("the divider can't be zero")
	}

	return a / b, nil
}
