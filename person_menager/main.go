package main

import (
	"fmt"
	"person_menager/files"
)

func menu() int {
	fmt.Println("1. Добавить акаунт")
	fmt.Println("2. Найти акаунт")
	fmt.Println("3. Удолить акаунт")
	fmt.Println("4. Выход")
	var ret int
	if _, err := fmt.Scan(&ret); err != nil {
    fmt.Println("Ошибка ввода, попробуйте снова")
    return 0
}
	if ret >= 1 && ret <= 4 {
		return ret
	}
	return 0
}

func main() {
	exit := true
	var err error

	for exit == true {
		keys := menu()
		switch keys {

		case 1:
			err = files.Addacc()
		case 2:
			err = files.Searchacc()
		case 3:
			err = files.Deleteacc()
		case 4:
			exit = false
		default:
			fmt.Println("Неверное значение!")
		}
		if err != nil {
			fmt.Println("Oшибка:", err)
		}
	}
}
