package controllers

import (
	"fmt"
	"log"
)

type MyUtil struct {
}

var showErrorMsg bool = true

func (m MyUtil) Message(_title, _content, _detail interface{}) {
	var next string
	fmt.Println("=== ", _title, " ===")
	fmt.Println(_content)
	fmt.Println(_detail)
	fmt.Println("Please Enter to continue...")
	fmt.Scanln(&next)
}

func (m MyUtil) ErrorMsg(_title, _content, _detail interface{}) {
	var next string
	if showErrorMsg {
		fmt.Println("=== SHOW ERROR LOG ===")
		fmt.Println("=== ", _title, " ===")
		fmt.Println(_content)
		log.Println(_detail)
	} else {
		fmt.Println("=== ", _title, " ===")
		fmt.Println(_content)
	}
	fmt.Println("\nPlease Enter to continue...")
	fmt.Scanln(&next)
}
