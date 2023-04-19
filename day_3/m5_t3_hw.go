/*
Задача №2
Вход:
Пользователь должен ввести правильный пароль, состоящий из:
цифр,
букв латинского алфавита(строчные и прописные) и
специальных символов  special = "_!@#$%^&"

Всего 4 набора различных символов.
В пароле обязательно должен быть хотя бы один символ из каждого набора.
Длина пароля от 8(мин) до 15(макс) символов.
Максимальное количество попыток ввода неправильного пароля - 5.
Каждый раз выводим номер попытки.
*Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

digits = "0123456789"
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
special = "_!@#$%^&"

Выход:
Написать, что ввели правильный пароль.

Пример:
хороший пароль -> o58anuahaunH!
хороший пароль -> aaaAAA111!!!
плохой пароль -> saucacAusacu8
*/

package main

import (
	"fmt"
)

// проверяет пароль на соответствие максимальной и минимальной длины
func CheckForLength(sCheckedPassword string, iMinLength, iMaxLength int) (bool, string) {
	var sResult string
	var bResult bool = true
	if len(sCheckedPassword) < iMinLength {
		sResult += fmt.Sprintf("Минимальная длина пароля должна быть %v символов", iMinLength)
		bResult = false
	}
	if len(sCheckedPassword) > iMaxLength {
		sResult += fmt.Sprintf("Максимальная длина пароля должна быть %v символов", iMaxLength)
		bResult = false
	}
	return bResult, sResult
}

// проверяет пароль на то входит ли в него хоть один символ из переданной функции строки
func CheckForDictionary(sCheckedPassword, sDictionary string) (bool, string) {
	var sResult string
	var bResult bool = false
	for _, rPasswordCharacter := range sCheckedPassword {
		for _, rDictionatyCharacter := range sDictionary {
			if rDictionatyCharacter == rPasswordCharacter {
				bResult = true
			}
		}
	}
	if !bResult {
		sResult = fmt.Sprintf("В пароль дожен входить хоть один символ из множества: %v", sDictionary)
	}
	return bResult, sResult
}

// спрашивает у пользователя пароль
func GetPasswordFromConsole() string {
	var sPass string
	fmt.Print("Введите пароль:")
	fmt.Scan(&sPass)
	return sPass
}

var sPassword string
var bResult, bResultTmp bool = true, true
var sResult, sResultTmp string
var iAttemptCount int = 0

func main() {
	for true {
		iAttemptCount++
		fmt.Println("Попытка № ", iAttemptCount)
		bResult = true
		sResult = ""
		sPassword = GetPasswordFromConsole()

		bResultTmp, sResultTmp = CheckForLength(sPassword, 8, 15)
		if bResultTmp == false {
			bResult = false
			sResult += sResultTmp + "\n"
		}

		bResultTmp, sResultTmp = CheckForDictionary(sPassword, "0123456789")
		if bResultTmp == false {
			bResult = false
			sResult += sResultTmp + "\n"
		}

		bResultTmp, sResultTmp = CheckForDictionary(sPassword, "abcdefghiklmnopqrstvxyz")
		if bResultTmp == false {
			bResult = false
			sResult += sResultTmp + "\n"
		}

		bResultTmp, sResultTmp = CheckForDictionary(sPassword, "ABCDEFGHIKLMNOPQRSTVXYZ")
		if bResultTmp == false {
			bResult = false
			sResult += sResultTmp + "\n"
		}

		bResultTmp, sResultTmp = CheckForDictionary(sPassword, "_!@#$%^&")
		if bResultTmp == false {
			bResult = false
			sResult += sResultTmp + "\n"
		}

		if bResult == false {
			fmt.Print(sResult)
			if iAttemptCount < 5 {
				continue
			}
		} else {
			fmt.Println("Пароль принят")
			break
		}

		if iAttemptCount >= 5 {
			fmt.Println("Превышено число неудачных попыток ввода")
			break
		}

	}
}
