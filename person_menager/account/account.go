package account

import (
	"encoding/json"
	"fmt"
	"os"
	"person_menager/email"
)

type Person struct {
	Account struct {
		Name       string `json:"name"`
		Age        uint   `json:"age"`
		Email      string `json:"email"`
		Time_creat string `json:"time_creat"`
	} `json:"Account"`
	Quantity int `json:"quantity"`
}

func Record(str string) string {
	fmt.Printf("%s ", str)
	fmt.Scan(&str)
	return str
}

func Record_int(str string) (uint, error) {
	fmt.Printf("%s ", str)
	var res uint
	if _, err := fmt.Scan(&res); err != nil {
		fmt.Println("Неверный формат!")
		return 0, err
	}
	return res, nil
}

func Email_right(ema string) bool {
	api_key := "d66871ff7ca0a597cdd3a4474ab10a8a2818bbe5"
	chek_email, err := email.CheckEmailWithHunter(ema, api_key)
	if err != nil {
		fmt.Println("Ошибка проверки:", err)
	} else if !chek_email {
		fmt.Println("Email не найден или недействителен.")
		return false
	}
	return true
}

func Quantity_counter() int {
	date := Extract_byte()
	if len(date) == 0 {
    return 1
	}
	return date[len(date)-1].Quantity + 1
}

func Extract_byte() []Person {
	files, _ := os.ReadFile("buffer.json")
	var date []Person
	json.Unmarshal(files, &date)
	return date
}

func ToJson(date Person) ([]byte, error) {
	return json.Marshal(date)
}

func Search_to_name(name string) Person {
	date := Extract_byte()
	for _, item := range date {
		if item.Account.Name == name {
			return item
		}
	}
	return Person{}
}

