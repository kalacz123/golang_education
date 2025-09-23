package main

import (
	"errors"
	"fmt"
)

// func sqrt(x float64) (float64, error) {
// 	if x < 0 {
// 		return 0, errors.New("negative number")
// 	}
// 	return math.Sqrt(x), nil
// }

// func process(data []byte) error {
// 	if len(data) == 0 {
// 		return errors.New("data is empty")
// 	}
// 	return nil
// }

type customError struct {
	message string
}

func (e *customError) Error() string {
	return fmt.Sprintf("custom error: %s", e.message)
}

// func eprocess() error {
// 	return &customError{message: "custom error"}
// }

func readConfig() error {
	return errors.New("config error")
}

func readData() error {
	err:= readConfig()
	if err != nil {
		return fmt.Errorf("readConfig: %w", err)
	}
	return nil
}

func main() {
	// result, err := sqrt(16)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }

	// data := []byte{}
	// if err := process(data); err != nil {
	// 	fmt.Println(err)
	// }

	// error interface of built-in package
	// if err := eprocess(); err != nil {
	// 	fmt.Println(err)
	// }

	if err:= readData(); err != nil {
		fmt.Println(err)
	}
}
