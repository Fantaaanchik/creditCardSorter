package main

import (
	"fmt"
	"log"
)

func main() {
	cardList1 := []string{"0123-4867", "7123-4567", "5123-4567", "0143-4561", "0113-4527"}
	cardList2 := []string{"0123-1867", "7153-4567", "6123-4569", "9143-4561", "0143-4527"}

	humoMap := make(map[string]string)
	othersBanksMap := make(map[string]string)

	for _, v := range cardList1 {
		if v[0] >= 53 {
			humoMap[v] = "HUMO"
		} else {
			othersBanksMap[v] = "Other Bank"
		}
	}
	for _, v := range cardList2 {
		if v[0] >= 53 {
			humoMap[v] = "HUMO\n"
		} else {
			othersBanksMap[v] = "Other Bank\n"
		}
	}
	fmt.Println("Список карточек Хумо:")
	fmt.Println(humoMap)
	input := ""
	fmt.Printf("Введите счет карты: \n")
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal("НЕ верный тип данных!", err.Error())
	}

	fmt.Println("Введено значение: ", input)

	if len(input) != 9 {
		log.Fatal("Длинна не соответствует формату счета карты! (Количество символов = 9)")
	}
	if input[4] != 45 {
		log.Fatal(`Знак "-" не на своём месте или отсуствует`)
	}

	message := false
	for _, value := range cardList1 {
		if value == input {
			message = true
			break
		} else if value != input {
			for _, val := range cardList2 {
				if val == input {
					message = true
					break
				}
			}
		}
	}

	if !message && input[0] >= 53 && input[0] <= 57 {
		humoMap[""] = input + ":HUMO \n"
	} else if !message && input[0] >= 49 && input[0] <= 52 {
		othersBanksMap[""] = input + ":Other Bank \n"
	} else {
		fmt.Println("Ошибка при добавлении, что-то пошло не так")
	}
	fmt.Println("Это список карт Хумо")
	fmt.Println(humoMap)
	fmt.Println("Это список других карт")
	fmt.Println(othersBanksMap)
}
