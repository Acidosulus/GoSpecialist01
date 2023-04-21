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
	"fmt"
	"net/http"
	"strings"
	"text/template"
	"unicode/utf8"
)

type Registry struct {
	Bills []Bill
}

type Bill struct {
	Fio     string
	Phone   string
	Address Address
	Items   []Item
}

type Address struct {
	Zip    string
	Sity   string
	Street string
	Home   string
	Flat   string
}

type Item struct {
	Name   string
	Number int
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

// проверяет строку на правильность заполнения, возращает статус проверке и текст ошибки
func CheckString(Text string, bOnlyLiterals, bOnlyNumbers bool, iMinLength, iMaxLength int) (string, bool) {
	var bIsAllRight bool = true
	sErrors := ""
	if bOnlyLiterals { // проверка на то, что строка сожержит только буквы
		if !AreOnlyCharactersOfFirstStringIncludedInTheSecondString(Text, " qwertyuiopasdfghjklzxcvbnmйцукенгшщзфывапролдячсмитьбжэЙЦУКЕНГШЩЗФЫВАПРОЛДЖЭЯЧСМИТЬБЮQWERTYUIOPASDFGHJKLZXCVBNM") {
			bIsAllRight = false
			sErrors += "Допустимы ввод лишь букв и пробела\n"
		}
	}
	if bOnlyNumbers { // проверка на то, что строка содержит только цифры
		if !AreOnlyCharactersOfFirstStringIncludedInTheSecondString(Text, "0123456789") {
			bIsAllRight = false
			sErrors += "Допустимы ввод лишь цифр\n"
		}
	}
	if iMinLength == iMaxLength && iMaxLength > 0 { // задана точная ненулевая длина строки
		if utf8.RuneCountInString(Text) != iMinLength {
			bIsAllRight = false
			sErrors += fmt.Sprintf("Строка должна быть длины %v символов\n", iMaxLength)
		}
	}
	if iMinLength != iMaxLength && iMinLength > 0 && iMaxLength <= 0 { // задана минимальная длина строки
		if utf8.RuneCountInString(Text) < iMinLength {
			bIsAllRight = false
			sErrors += fmt.Sprintf("Строка должна быть не короче %v символов\n", iMinLength)
		}
	}
	if iMinLength != iMaxLength && iMaxLength > 0 && iMinLength <= 0 { // задана максимальная длина строки
		if utf8.RuneCountInString(Text) > iMaxLength {
			bIsAllRight = false
			sErrors += fmt.Sprintf("Строка должна быть не длиннее %v символов\n", iMaxLength)
		}
	}
	if iMinLength != iMaxLength && iMaxLength > 0 && iMinLength > 0 { // задана максимальная и минимальная длина строки
		if utf8.RuneCountInString(Text) < iMinLength || utf8.RuneCountInString(Text) > iMaxLength {
			bIsAllRight = false
			sErrors += fmt.Sprintf("Длина строки должна быть в промежутке межу %v и %v символов\n", iMinLength, iMaxLength)
		}
	}

	return sErrors, bIsAllRight
}

// метод, предзаполняет поля объекта, чтобы при тестовых запусках не делать это каждый раз в консоли
func (self *Registry) FillData() {
	var bill Bill

	bill.Fio, bill.Phone, bill.Address.Zip, bill.Address.Sity, bill.Address.Street, bill.Address.Home, bill.Address.Flat = "Василий", "2223", "625624", "Посад", "Ленина", "16А", "245"
	bill.Items = []Item{{Name: "Молоток", Number: 1}, {Name: "Гвоздодёр", Number: 2}, {Name: "Гвозди", Number: 500}, {Name: "Рулетка", Number: 1}}
	self.Bills = append(self.Bills, bill)

	bill.Fio, bill.Phone, bill.Address.Zip, bill.Address.Sity, bill.Address.Street, bill.Address.Home, bill.Address.Flat = "Сергей", "2355655", "625884", "Посад", "Пушкина", "22", "75"
	bill.Items = []Item{{Name: "Суперклей", Number: 10}, {Name: "Растворитель суперклея", Number: 10}}
	self.Bills = append(self.Bills, bill)

	bill.Fio, bill.Phone, bill.Address.Zip, bill.Address.Sity, bill.Address.Street, bill.Address.Home, bill.Address.Flat = "Терентий", "3445321", "625123", "Посад", "Тихая", "2", "1"
	bill.Items = []Item{{Name: "Диэлектрические перчатки", Number: 1}, {Name: "Съемник напряжения", Number: 1}, {Name: "Нашатырь", Number: 5}}
	self.Bills = append(self.Bills, bill)

}

// страница создания нового покупателя, без списка закупок :(
func createHandler(writer http.ResponseWriter, request *http.Request) {
	var bill Bill
	var Message_text string
	bill.Fio = request.FormValue("fio")
	bill.Phone = request.FormValue("phone")
	bill.Address.Zip = request.FormValue("addresszip")
	bill.Address.Sity = request.FormValue("addresssity")
	bill.Address.Street = request.FormValue("addressstreet")
	bill.Address.Home = request.FormValue("addresshome")
	bill.Address.Flat = request.FormValue("addressflat")
	errs, err := CheckString(bill.Fio, true, false, 3, 100)
	if err == false {
		Message_text += "ФИО: " + errs + "\n"
	}

	errs, err = CheckString(bill.Phone, false, true, 2, 50)
	if err == false {
		Message_text += "Телефон: " + errs + "\n"
	}

	errs, err = CheckString(bill.Address.Zip, false, true, 6, 6)
	if err == false {
		Message_text += "Индекс: " + errs + "\n"
	}

	errs, err = CheckString(bill.Address.Sity, true, false, 2, 50)
	if err == false {
		Message_text += "Город: " + errs + "\n"
	}

	errs, err = CheckString(bill.Address.Street, false, false, 2, 50)
	if err == false {
		Message_text += "Улица: " + errs + "\n"
	}

	errs, err = CheckString(bill.Address.Home, false, false, 1, 50)
	if err == false {
		Message_text += "Дом: " + errs + "\n"
	}
	if len(Message_text) > 0 {
		Message_text = "Ошибки ввода:" + Message_text + "\n"
		html, err := template.ParseFiles("error.html")
		if err != nil {
			fmt.Println(err)
		}
		err = html.Execute(writer, Message_text)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		OneRegistry.Bills = append(OneRegistry.Bills, bill)
		http.Redirect(writer, request, "", http.StatusSeeOther)
	}

}

// страница просмотра ошибок ввода данных
func errorHandler(writer http.ResponseWriter, request *http.Request) {

	html, err := template.ParseFiles("error.html")
	if err != nil {
		fmt.Println(err)
	}
	err = html.Execute(writer, OneRegistry)
	if err != nil {
		fmt.Println(err)
	}

}

// рендер страницы добавления нового покупателя
func newHandler(writer http.ResponseWriter, request *http.Request) {

	html, err := template.ParseFiles("new_customer.html")
	if err != nil {
		fmt.Println(err)
	}
	err = html.Execute(writer, OneRegistry)
	if err != nil {
		fmt.Println(err)
	}

}

// процедура добавления покупателя
func deleteHandler(writer http.ResponseWriter, request *http.Request) {
	var bills []Bill
	for _, element := range OneRegistry.Bills {
		if element.Fio != request.FormValue("delete") {
			bills = append(bills, element)
		}
	}
	OneRegistry.Bills = bills
	http.Redirect(writer, request, "", http.StatusSeeOther)
}

// стартовая страница со списком покупателей
func viewHandler(writer http.ResponseWriter, request *http.Request) {

	html, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}
	err = html.Execute(writer, OneRegistry)
	if err != nil {
		fmt.Println(err)
	}

}

var OneRegistry Registry

func main() {
	OneRegistry.FillData()
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.ListenAndServe("localhost:8080", nil)

}
