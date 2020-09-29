package main

import (
	"fmt"
)

//Declaring user defined fields
type Dth struct {
	name     string
	email    string
	phone    int
	balance  int
	packs    string
	chanels  string
	services string
}

var pack = map[string]int{
	"Silver pack:  Zee, Sony, Star Plus:":                 50,
	"Gold Pack: Zee, Sony, Star Plus, Discovery, NatGeo:": 100,
}

var chanalMap = map[string]int{
	"Zee":       10,
	"Sony":      15,
	"StarPlus":  20,
	"Discovery": 10,
	"NatGeo":    20,
}

var service = map[string]int{
	"LearnEnglish": 200,
	"LearnCooking": 100,
}

//Main function
func main() {

	fmt.Println("Welcome to DishTV\n")

	dt := &Dth{name: "Vignesh", email: "vinesh1865@gmail.com", phone: 9047660920, balance: 100, packs: "Gold", chanels: "Zee, Sony", services: "LearnEnglish"}
	fmt.Println(dt)

}
