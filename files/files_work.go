package files

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"person_menager/account"
	"strings"
	"time"
)

func ReсordFile(ref []byte) (bool, error) {
	var existingData []json.RawMessage

	file, err := os.ReadFile("buffer.json")
	if err == nil {
		json.Unmarshal(file, &existingData)
	}

	existingData = append(existingData, json.RawMessage(ref))
	newData, err := json.MarshalIndent(existingData, "", "  ")
	if err != nil {
		return false, err
	}

	err = os.WriteFile("buffer.json", newData, 0644)
	if err != nil {
		return false, err
	}

	return true, nil
}

func Addacc() error {
	var pr account.Person
	var err error

	pr.Account.Name = account.Record("Введите имя:")
	if pers := account.Search_to_name(pr.Account.Name); pers.Account.Name != "" {
		fmt.Println("Имя занято!")
	} else {
		pr.Account.Age, err = account.Record_int("Введите возраст:")
		if err == nil {
			pr.Account.Email = account.Record("Введите email:")
			chek := account.Email_right(pr.Account.Email)
			pr.Account.Time_creat = time.Now().Format("2006-01-02 15:04:05")
			pr.Quantity = account.Quantity_counter()

			if chek {
				ref, err := account.ToJson(pr)
				if err != nil {
					fmt.Println(err)
				}

				_, err = ReсordFile(ref)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Данные успешно добавлены!")
				}
				Back_to_menu()
			}
		} else {
			return err
		}
	}
	return nil
}

func Searchacc() error {
	var err error
	

	name := account.Record("Введите имя:")
	result_struct := account.Search_to_name(name)

	if result_struct.Account.Name != "" {
		fmt.Println("Ваш возраст: ", result_struct.Account.Age)
		fmt.Println("Ваш email: ", result_struct.Account.Email)
		fmt.Println("Ваше место в базе: ", result_struct.Quantity)

		Back_to_menu()
	} else {
		fmt.Println("Акаунт не найден, попробуйте ввести другое имя!")
	}

	return err
}

func Back_to_menu() {
	exit := false
	for exit != true {
		fmt.Print("Нажмите enter чтобы вернуться в меню")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		if strings.TrimSpace(input) == "" {
			exit = true
		}
	}
}

func Deleteacc() error {
	name := account.Record("Введите имя для удаления:")
	data := account.Extract_byte()
	var newData []account.Person

	found := false
	for _, item := range data {
		if item.Account.Name != name {
			newData = append(newData, item)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Println("Акаунт не найден!")
		return nil
	}

	file, err := json.MarshalIndent(newData, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("buffer.json", file, 0644)
}

