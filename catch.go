package utils

import "fmt"

func TryCatch(f func()) {
	defer func() {
		if err := recover(); err != nil {
			err = fmt.Errorf("internal error: %v", err)
			fmt.Printf("TryCatch Error：%v\n", err)
		}
	}()
	f()
}
