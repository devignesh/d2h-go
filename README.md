# d2h-go

D2H Dish TV Go application

Direct to home (D2H) operator DishTV wants you to design a mini software
system for its customers.

DishTV offers multiple predefined channel packages (called as base packs
henceforth) and multiple services. Base packs available to purchase are Gold and
Silver which comes with certain channels.

simple command-line application
Deftouch.co | Backend Developer

---

## Run the Go application:

Please make sure that you are in inside gopath directory

    cd go/src/

    Clone:
    ----------
    git clone git@github.com:devignesh/d2h-go.git

    cd d2h-go

    go run d2h.go  // Run the file

    go build      //build the go file

    ./d2h-go      // Run the go

##Sample Outputs

    Welcome to DishTV

    1: View current balance in the account
    2: Recharge Account
    3: View available packs, channels and services
    4: Subscribe to base packs
    5: Add channels to an existing subscription
    6: Subscribe to special services
    7: View current subscription details
    8: Update email and phone number for notifications
    9: Exit

    Enter the option

    -----------------------------------------------------------

    1: View current balance in the account
    2: Recharge Account
    3: View available packs, channels and services
    4: Subscribe to base packs
    5: Add channels to an existing subscription
    6: Subscribe to special services
    7: View current subscription details
    8: Update email and phone number for notifications
    9: Exit

    Enter the option
    1

    View current balance in the account
    Current balance is 100 Rs.

    1: View current balance in the account
    2: Recharge Account
    3: View available packs, channels and services
    4: Subscribe to base packs
    5: Add channels to an existing subscription
    6: Subscribe to special services
    7: View current subscription details
    8: Update email and phone number for notifications
    9: Exit

    Enter the option
    2

    Recharge Account
    Enter the amount to recharge:
    5000
    Recharge completed successfully.
    Current balance is 5100 Rs

    1: View current balance in the account
    2: Recharge Account
    3: View available packs, channels and services
    4: Subscribe to base packs
    5: Add channels to an existing subscription
    6: Subscribe to special services
    7: View current subscription details
    8: Update email and phone number for notifications
    9: Exit

    Enter the option
    3

    Available packs for subscription:
    Silver pack:  Zee, Sony, Star Plus: 50
    Gold Pack: Zee, Sony, Star Plus, Discovery, NatGeo: 100
    Available channels for subscription
    Discovery 10
    NatGeo 20
    Zee 10
    Sony 15
    StarPlus 20
    Available services for subscription
    LearnEnglish service: 200

    LearnCooking service: 100

    1: View current balance in the account
    2: Recharge Account
    3: View available packs, channels and services
    4: Subscribe to base packs
    5: Add channels to an existing subscription
    6: Subscribe to special services
    7: View current subscription details
    8: Update email and phone number for notifications
    9: Exit

    Enter the option
    4

    Subscribe to channel packs
    Enter the Pack you wish to subscribe: (Silver: 'S', Gold: 'G'):
    S
    No of months:
    5
    Monthly price: 50 Rs.
    No of months: 5
    Subscription Amount: 250
    Discount applied: 25 Rs.
    Final Price after discount: 225 Rs.
    Account balance: 4875
    Email notification sent successfully
    SMS notification sent successfully

    You have successfully subscribed the following packs -  Silvar

    1: View current balance in the account
    2: Recharge Account
    3: View available packs, channels and services
    4: Subscribe to base packs
    5: Add channels to an existing subscription
    6: Subscribe to special services
    7: View current subscription details
    8: Update email and phone number for notifications
    9: Exit

    Enter the option
    5


    Add channels to existing subscription
    Enter channel names to add (separated by commas):
    Zee,Sony
    Channels added successfully.
    Account balance: 4850 Rs.

    1: View current balance in the account
    2: Recharge Account
    3: View available packs, channels and services
    4: Subscribe to base packs
    5: Add channels to an existing subscription
    6: Subscribe to special services
    7: View current subscription details
    8: Update email and phone number for notifications
    9: Exit

    Enter the option
    6


    Subscribe to special services
    Enter the service name:
    LearnCooking
    Service subscribed successfully
    Account balance: 4750
    Email notification sent successfully
    SMS notification sent successfully

    1: View current balance in the account
    2: Recharge Account
    3: View available packs, channels and services
    4: Subscribe to base packs
    5: Add channels to an existing subscription
    6: Subscribe to special services
    7: View current subscription details
    8: Update email and phone number for notifications
    9: Exit

    Enter the option
    7

    View current subscription details

    Currently subscribed packs and channels: Silvar + Zee,Sony

    Currently subscribed services: LearnCooking

    1: View current balance in the account
    2: Recharge Account
    3: View available packs, channels and services
    4: Subscribe to base packs
    5: Add channels to an existing subscription
    6: Subscribe to special services
    7: View current subscription details
    8: Update email and phone number for notifications
    9: Exit

    Enter the option
    8


    Update email and phone number for notifications
    Enter the email:
    vigneshkumar.mca2016@adhiyamaan.in
    Enter phone:
    9047660920
    Customer Name: Vignesh
    Updated Email Address: vigneshkumar.mca2016@adhiyamaan.in
    Updated Phone Number: 9047660920
    Email and Phone updated successfully

    1: View current balance in the account
    2: Recharge Account
    3: View available packs, channels and services
    4: Subscribe to base packs
    5: Add channels to an existing subscription
    6: Subscribe to special services
    7: View current subscription details
    8: Update email and phone number for notifications
    9: Exit

    Enter the option
    9
    exit status 1

================================================================
