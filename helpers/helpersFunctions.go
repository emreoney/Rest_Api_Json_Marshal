package helpers

import "fmt"

func CheckError(err error) {
	if err != nil {
		fmt.Println("Hata: ", err.Error())
	}
}
