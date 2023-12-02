package main

import (
	"log"

	function "func.go"
	controller "func.go/controller03"
)

func main() {
	uuid := "8433f2c4-6446-4994-b66d-86e67c81e056"
	commonlogpath := "C://Users//HP//OneDrive//Desktop//Task_2//log"
	dir := "info"
	data, err := function.Readappdata(uuid, dir, commonlogpath)
	if err != nil {
		log.Fatal("Error in Reading app data")
	}
	controller.Insert_Mongodb(data)
}
