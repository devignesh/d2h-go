package main

import (
	"fmt"
	"os"
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

// getting Current Balance
func (d *Dth) getBalance() {

	fmt.Println("View current balance in the account")
	fmt.Println("Current balance is", d.balance, "Rs.\n")

	return

}

//Adding New Balance
func (d *Dth) recharge() {

}

//Showing list of services, packs and chanels
func (d *Dth) listOffers() {

}

//Subscribe new packages
func (d *Dth) subscribrChannelPacks() {

}

//Adding channels
func (d *Dth) addChannelToexsistSubscription() {

}

//Subscribe New services
func (d *Dth) subscribeNewService() {

}

//Listing account Details
func (d *Dth) getSubscriptionDetails() {

}

//Updating email and phone nuber
func (d *Dth) updateEmailandPhone() {

}

//Function for option in D2h
func help() {

	fmt.Println("1: View current balance in the account")
	fmt.Println("2: Recharge Account")
	fmt.Println("3: View available packs, channels and services")
	fmt.Println("4: Subscribe to base packs")
	fmt.Println("5: Add channels to an existing subscription")
	fmt.Println("6: Subscribe to special services")
	fmt.Println("7: View current subscription details")
	fmt.Println("8: Update email and phone number for notifications")
	fmt.Println("9: Exit\n")
	fmt.Println("Enter the option")

}

//Main function
func main() {

	fmt.Println("Welcome to DishTV\n")

	dt := &Dth{name: "Vignesh", email: "vinesh1865@gmail.com", phone: 9047660920, balance: 100, packs: "Gold", chanels: "Zee, Sony", services: "LearnEnglish"}
	// fmt.Println(dt)
	help()

	for {
		var option int
		_, err := fmt.Scanf("%d", &option) //Getting user input for options

		if err != nil {
			fmt.Println(err)
		}

		if option > 9 {

			fmt.Println("Please Enter valid option\n")
			help()

		} else {

			switch option {

			case 1:

				dt.getBalance()

			case 2:

				dt.recharge()

			case 3:

				dt.listOffers()

			case 4:

				dt.subscribrChannelPacks()

			case 5:

				dt.addChannelToexsistSubscription()

			case 6:

				dt.subscribeNewService()

			case 7:

				dt.getSubscriptionDetails()

			case 8:

				dt.updateEmailandPhone()

			case 9:

				os.Exit(1)
			}
			help()
			continue
		}
	}
}
