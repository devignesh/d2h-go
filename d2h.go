package main

import (
	"fmt"
	"os"
	"strings"
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

	fmt.Println("Add channels to existing subscription")
	var chanels string

	fmt.Println("Enter channel names to add (separated by commas):")
	fmt.Scanln(&chanels)

	chanalInput := strings.Split(chanels, ",")

	var totalprice int

	for _, value := range chanalInput {

		chanelPrice, ok := chanalMap[value]
		if !ok {
			fmt.Println("The channels are not valid", value)
			continue
		}
		totalprice = totalprice + chanelPrice

		if totalprice > d.balance {
			fmt.Println("Please recharge your account, Your current Balance is :", d.balance, "\n")
			return
		}

		d.balance = d.balance - totalprice // updating account balance
		d.chanels = chanels                // updating account channels
	}

	fmt.Println("Channels added successfully.")
	fmt.Println("Account balance:", d.balance, "Rs.\n")

	return
}

//Subscribe New services
func (d *Dth) subscribeNewService() {

	fmt.Println("Subscribe to special services")

	var servicename string
	var servicePrice int

	fmt.Println("Enter the service name:")
	fmt.Scanln(&servicename)

	servicePrice, ok := service[servicename]

	if !ok {
		fmt.Println("Please provide valid service:", servicename)
		return
	}

	if servicePrice > d.balance {
		fmt.Println("Please recharge your account, Your current Balance is :", d.balance, "\n")
		return
	}

	d.services = servicename
	fmt.Println("Service subscribed successfully")
	d.balance = d.balance - servicePrice
	fmt.Println("Account balance:", d.balance)
	emailAndPhone()

	return
}

//Listing account Details
func (d *Dth) getSubscriptionDetails() {

	fmt.Println("View current subscription details\n")
	fmt.Println("Currently subscribed packs and channels:", d.packs, "+", d.chanels, "\n")
	fmt.Println("Currently subscribed services:", d.services, "\n")

	return
}

//Updating email and phone nuber
func (d *Dth) updateEmailandPhone() {

	fmt.Println("Update email and phone number for notifications")

	var email string
	var phone int
	fmt.Println("Enter the email:")
	fmt.Scanln(&email)
	fmt.Println("Enter phone:")

	fmt.Scanln(&phone)

	d.email = email
	d.phone = phone

	fmt.Println("Customer Name:", d.name)
	fmt.Println("Updated Email Address:", d.email)
	fmt.Println("Updated Phone Number:", d.phone)
	fmt.Println("Email and Phone updated successfully\n")

	return
}

func emailAndPhone() {

	fmt.Println("Email notification sent successfully")
	fmt.Println("SMS notification sent successfully\n")
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

	dt := &Dth{name: "Vignesh", email: "vinesh1865@gmail.com", phone: 9047660920, balance: 100, packs: "Gold", chanels: "Zee,Sony", services: "LearnEnglish"}
	// fmt.Println(dt)
	help()

	for {
		var option int
		_, err := fmt.Scanf("%d", &option) //Getting user input for options

		if err != nil {
			fmt.Println(err)
		}

		if option > 9 {

			fmt.Println("Please Enter valid option: 1 - 9 \n")

		}

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
