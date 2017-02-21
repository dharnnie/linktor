package handlers

import (
	"log"
	
)

func smplErr(e error, m string){
	if e != nil{
		log.Println(m, e)
	}
}