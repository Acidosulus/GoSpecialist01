/*
Задача №1
Написать функцию, которая расшифрует строку.
code = "220411112603141304"
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через вторую строку.

codeToString(code) -> "???????'
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Sipher(sSource string, alphabet map[string]string) string {
	var sResult, value string
	var ok bool
	for i := 0; i < len(sSource); i++ {
		value, ok = alphabet[string(sSource[i])]
		if ok {
			sResult += value
		} else {
			fmt.Printf("В строке найден символ %v ASCII:%v в позиции:%v выходящий за пределы области допустимых значений\n", value, sSource[i], i)
		}
	}
	return sResult
}

func DeSipher(sSource string, alphabet map[string]string) string {
	var i int = 0
	var sCode, value, sResult string
	var ok bool
	for i < (len(sSource))/2 {
		sCode = string(sSource[i*2]) + string(sSource[i*2+1])
		value, ok = alphabet[string(sCode)]
		if ok {
			sResult += value
		} else {
			fmt.Printf("В строке найден код %v выходящий за пределы области допустимых значений\n", sCode)
		}
		i++
	}
	return sResult
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

var mSipher, mDeSipher map[string]string
var cNumber, sSource, sAnswer string

func main() {

	//карты для шифровки дешифровки
	mSipher = make(map[string]string)
	mDeSipher = make(map[string]string)
	for i := 97; i <= 122; i++ {
		if i-97 < 10 {
			cNumber = "0" + strconv.Itoa(i-97)
		} else {
			cNumber = strconv.Itoa(i - 97)
		}
		mSipher[string(i)] = cNumber
		mDeSipher[cNumber] = string(i)
	}
	mSipher[" "] = "26"
	mDeSipher["26"] = " "

	var sSipherResult, sDeSipherResult string

	sAnswer = GetStringFromConsole("Защифровать: 'y/Y' Расшифровать - любой другой символ")
	if sAnswer == "y" || sAnswer == "Y" {
		sSource = GetStringFromConsole("Введите слово (маленькие латинские буквы или пробел)")
		sSipherResult = Sipher(sSource, mSipher)
		sDeSipherResult = DeSipher(sSipherResult, mDeSipher)
		fmt.Println("Исходная строка: ", sSource, " Закодированная строка: ", sSipherResult, " Обратно раскодированная строка: ", sDeSipherResult)
	} else {
		sSource = GetStringFromConsole("Введите код от 00 до 26")
		sDeSipherResult = DeSipher(sSource, mDeSipher)
		sSipherResult = Sipher(sDeSipherResult, mSipher)
		fmt.Println("Исходная код: ", sSource, " Раскодированная строка: ", sDeSipherResult, " Обратно закодированная строка: ", sSipherResult)
	}
}
