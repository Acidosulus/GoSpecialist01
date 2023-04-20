/*
Сформировать данные для отправки заказа из
магазина по накладной и вывести на экран:
1) Наименование товара (минимум 1, максимум 100)
2) Количество (только числа)
3) ФИО покупателя (только буквы)
4) Контактный телефон (10 цифр)
5) Адрес(индекс(ровно 6 цифр), город, улица, дом, квартира)

Эти данные не могут быть пустыми.
Проверить правильность заполнения полей.

реализовать несколько методов у типа "Накладная"

createReader == NewReader
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Bill struct {
	fio     string
	phone   string
	address Address
	items   []Item
}

type Address struct {
	zip    string
	sity   string
	street string
	home   string
	flat   string
}

type Item struct {
	name   string
	number int
}

// возвращает истину в случае если строка-первый параметр состоит только символов строки-второго параметра
func AreOnlyCharactersOfFirstStringIncludedInTheSecondString(sCheckedString string, sDictionaryString string) bool {
	for _, sCharacter := range sCheckedString {
		if strings.Contains(sDictionaryString, string(sCharacter)) == false {
			return false
		}
	}
	return true
}

// Отображает текст приглашения ввода, проверяет введённое по переданным параметрам, возвращается строку текста введённую пользователем
func GetStringFromConsoleWithCheck(sPromptText string, bOnlyLiterals, bOnlyNumbers bool, iMinLength, iMaxLength int) string {
	var sInputString, sErrors string
	var bIsAllRight bool = false
	for !bIsAllRight {
		if len(sErrors) > 0 {
			fmt.Println(sErrors)
		}
		sInputString = GetStringFromConsole(sPromptText)
		bIsAllRight = true
		sErrors = ""
		if bOnlyLiterals { // проверка на то, что строка сожержит только буквы
			if !AreOnlyCharactersOfFirstStringIncludedInTheSecondString(sInputString, " qwertyuiopasdfghjklzxcvbnmйцукенгшщзфывапролдячсмитьбжэЙЦУКЕНГШЩЗФЫВАПРОЛДЖЭЯЧСМИТЬБЮQWERTYUIOPASDFGHJKLZXCVBNM") {
				bIsAllRight = false
				sErrors += "Допустимы ввод лишь букв и пробела\n"
			}
		}
		if bOnlyNumbers { // проверка на то, что строка содержит только цифры
			if !AreOnlyCharactersOfFirstStringIncludedInTheSecondString(sInputString, "0123456789") {
				bIsAllRight = false
				sErrors += "Допустимы ввод лишь цифр\n"
			}
		}
		if iMinLength == iMaxLength && iMaxLength > 0 { // задана точная ненулевая длина строки
			if utf8.RuneCountInString(sInputString) != iMinLength {
				bIsAllRight = false
				sErrors += fmt.Sprintf("Строка должна быть длины %v символов\n", iMaxLength)
			}
		}
		if iMinLength != iMaxLength && iMinLength > 0 && iMaxLength <= 0 { // задана минимальная длина строки
			if utf8.RuneCountInString(sInputString) < iMinLength {
				bIsAllRight = false
				sErrors += fmt.Sprintf("Строка должна быть не короче %v символов\n", iMinLength)
			}
		}
		if iMinLength != iMaxLength && iMaxLength > 0 && iMinLength <= 0 { // задана максимальная длина строки
			if utf8.RuneCountInString(sInputString) > iMaxLength {
				bIsAllRight = false
				sErrors += fmt.Sprintf("Строка должна быть не длиннее %v символов\n", iMaxLength)
			}
		}
		if iMinLength != iMaxLength && iMaxLength > 0 && iMinLength > 0 { // задана максимальная и минимальная длина строки
			if utf8.RuneCountInString(sInputString) < iMinLength || utf8.RuneCountInString(sInputString) > iMaxLength {
				bIsAllRight = false
				fmt.Println(utf8.RuneCountInString(sInputString))
				sErrors += fmt.Sprintf("Длина строки должна быть в промежутке межу %v и %v символов\n", iMinLength, iMaxLength)
			}
		}
	}
	return sInputString
}

// Отображает текст приглашения ввода текста, возвращается строку текста введённую пользователем
func GetStringFromConsole(sPromptText string) string {
	if len(sPromptText) > 0 {
		fmt.Printf("%v:", sPromptText)
	}
	var sInputString string = ""
	in := bufio.NewReader(os.Stdin)
	sInputString, _ = in.ReadString('\n')
	sInputString = strings.Replace(strings.Replace(sInputString, string(12), "", -1), string(13), "", -1) //sSource[0 : len(sSource)-1]
	return strings.TrimSpace(sInputString)
}

// метод Bill, организует пользовательский интерфес ввода списка товаров
func (self *Bill) InputItems() {
	var item Item
	//	var items []Item
	var sAnswer string
	for {
		sAnswer = GetStringFromConsole("Добавить товар в накладную: 'y/Y'")
		if sAnswer == "y" || sAnswer == "Y" {
			item.name = GetStringFromConsoleWithCheck("Название товара (от 1 до 100 символов)", false, false, 1, 100)
			item.number, _ = strconv.Atoi(GetStringFromConsoleWithCheck("Количество товара", false, true, 0, 0))
			self.items = append(self.items, item)
		} else {
			break
		}
	}
	//return items
}

// метод Bill, организует пользовательский интефейс заполнения накладной
func (self *Bill) InputBill() {
	self.fio = GetStringFromConsoleWithCheck("ФИО покупателя (только буквы)", true, false, 1, 0)           // только символы, не пусто
	self.phone = GetStringFromConsoleWithCheck("Контактный телефон (от 3 до 20 цифр)", false, true, 3, 20) // только цифры от 3 до 20
	self.address.zip = GetStringFromConsoleWithCheck("индекс(ровно 6 цифр)", false, true, 6, 6)            // только цифры, ровно 6
	self.address.sity = GetStringFromConsoleWithCheck("Город", true, false, 1, 0)                          // только символы, не пусто
	self.address.street = GetStringFromConsoleWithCheck("Улица", false, false, 1, 0)                       // любые символы, не пусто
	self.address.home = GetStringFromConsoleWithCheck("Дом", false, false, 1, 0)                           // любые символы, не пусто
	self.address.flat = GetStringFromConsoleWithCheck("Кварира", false, false, 0, 0)                       // любые символы, может быть пусто
	self.InputItems()
}

// метод Bill, предзаполняет поля объекта, чтобы при тестовых запусках не делать это каждый раз в консоли
func (self *Bill) FillData() {
	self.fio, self.phone, self.address.zip, self.address.sity, self.address.street, self.address.home, self.address.flat = "Василий", "2223", "625624", "Посад", "Ленина", "16А", "245"
	self.items = []Item{{name: "Молоток", number: 1}, {name: "Гвоздодёр", number: 2}, {name: "Гвозди", number: 500}, {name: "Рулетка", number: 1}}
}

// метод Bill, возвращает инфомацию по накладной в виде текстовой строки
func (self *Bill) GetInformation() string {
	var sResult string
	sResult = "Накладная:\n" +
		fmt.Sprintf("\tФИО:%v\n", self.fio) +
		fmt.Sprintf("\tТелефон:%v\n", self.phone) +
		"\tАдрес:\n" +
		fmt.Sprintf("\t\tИндекс:%v\n", self.address.zip) +
		fmt.Sprintf("\t\tГород:%v\n", self.address.sity) +
		fmt.Sprintf("\t\tУлица:%v\n", self.address.street) +
		fmt.Sprintf("\t\tДом:%v\n", self.address.home) +
		fmt.Sprintf("\t\tКвартира:%v\n", self.address.flat)
	for index, position := range self.items {
		if index == 0 {
			sResult += "\tТовары:\n"
		}
		sResult += "\t\t-------------------------------\n" +
			fmt.Sprintf("\t\tНазвание:%v\n", position.name) +
			fmt.Sprintf("\t\tКоличество:%v\n", position.number)
	}
	return sResult
}

// метод Bill, записывает информацию о заказе в текствый файл
func (self *Bill) StoreBillIntoFile(sFilePath string) {
	file, _ := os.Create(sFilePath)
	file.WriteString(self.GetInformation())
	file.Close()
}

var OneBill Bill

func main() {
	OneBill.InputBill() // ввод пользователем
	//OneBill.FillData() // заполнить готовыми данными
	fmt.Println(OneBill.GetInformation())
	OneBill.StoreBillIntoFile("bill.txt")
}
