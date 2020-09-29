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

	fmt.Println("Recharge Account")
	var updatebalance int
	fmt.Println("Enter the amount to recharge:")
	fmt.Scanf("%d", &updatebalance)

	d.balance = d.balance + updatebalance // Balance update
	fmt.Println("Recharge completed successfully.\nCurrent balance is", d.balance, "Rs", "\n")

	return

}

//Showing list of services, packs and chanels
func (d *Dth) listOffers() {

	fmt.Println("Available packs for subscription:")

	for dth, price := range pack {
		fmt.Println(dth, price)
	}
	fmt.Println("Available channels for subscription")

	for channame, price := range chanalMap {
		fmt.Println(channame, price)
	}
	fmt.Println("Available services for subscription")

	for servicename, price := range service {
		fmt.Println(servicename, "service:", price, "\n")
	}

	return

}

//Subscribe new packages
func (d *Dth) subscribrChannelPacks() {

	fmt.Println("Subscribe to channel packs")

	var packs string
	var month int
	var montlyprice int
	silvar := "S"
	gold := "G"
	var totalAmount int

	fmt.Println("Enter the Pack you wish to subscribe: (Silver: 'S', Gold: 'G'):")
	fmt.Scanln(&packs)

	if packs == silvar {

		d.packs = "Silvar"

		montlyprice = pack["Silver pack:  Zee, Sony, Star Plus:"]
		fmt.Println("No of months:")
		fmt.Scanln(&month)

	} else if packs == gold {

		d.packs = "Gold"

		montlyprice = pack["Gold Pack: Zee, Sony, Star Plus, Discovery, NatGeo:"]
		fmt.Println("No of months:")
		fmt.Scanln(&month)

	} else {
		fmt.Println("Please provide valid option")
	}

	fmt.Println("Monthly price:", montlyprice, "Rs.")
	fmt.Println("No of months:", month)

	totalAmount = montlyprice * month //Getting amount for subscription

	fmt.Println("Subscription Amount:", totalAmount)

	if month >= 3 {

		discount := montlyprice * month
		value := discount / 10
		totalamount := totalAmount - value //Final amount for subscription

		if totalAmount > d.balance { //validating the balance
			fmt.Println("Please recharge your account, Your current Balance is :", d.balance)
			return
		}

		d.balance = d.balance - totalamount // Updatting the user account balance
		fmt.Println("Discount applied:", value, "Rs.")
		fmt.Println("Final Price after discount:", totalamount, "Rs.")
		fmt.Println("Account balance:", d.balance)
		emailAndPhone()

	} else {

		if totalAmount > d.balance { //validating the balance
			fmt.Println("Please recharge your account, Your current Balance is :", d.balance)
			return
		}

		d.balance = d.balance - totalAmount // Updatting the user account balance
		fmt.Println("Account balance:", d.balance)
		emailAndPhone()

	}

	if montlyprice == montlyprice {

		fmt.Println("You have successfully subscribed the following packs - ", d.packs, "\n")
	}

	return

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
