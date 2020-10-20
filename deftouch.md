Direct to home (D2H) operator DishTV wants you to design a mini software
system for its customers.

DishTV offers multiple predefined channel packages (called as base packs
henceforth) and multiple services. Base packs available to purchase are Gold and
Silver which comes with certain channels.

User has an initial balance of 100 Rs. in his account.

Whenever the base pack is subscribed, the base package price will be deducted
from the account balance.

User will receive a 10% discount on the base pack amount if the duration is at
least 3 months.

User can add the individual channel into the current subscription on which
amount will be deducted from the account balance.

In addition to offering TV channels, DishTV also offers multiple special optional
services - Learn English Service and Learn Cooking service.

If any of the services are purchased by the user, the corresponding amount will
be deducted from the balance.

Whenever user subscribes to any package or service, notifications should be sent
through email and SMS if the user has updated his email and phone number.

Notes:

1. We don't expect from you to develop real Email and SMS notification
   features. You can just show thexx message on the console that Email and SMS
   are sent successfully.

2. The list of packages, channels and prices are shown directly in sample
   inputs.

Sample Inputs
Note: To save space, we haven't shown the input options every time after the
processing is complete for the last selected option. Although you can show them if
you would like :-)

Welcome to DishTV

1. View current balance in the account
2. Recharge Account
3. View available packs, channels and services
4. Subscribe to base packs
5. Add channels to an existing subscription
6. Subscribe to special services
7. View current subscription details
8. Update email and phone number for notifications
9. Exit

Enter the option: 1
View current balance in the account
Current balance is 100 Rs.
======{Here you can show all the options each time if you'd like}==========

Enter the option: 2
Enter the amount to recharge: 500
Recharge completed successfully. Current balance is 600
============================================================

Enter the option: 3
View available packs, channels and services
Available packs for subscription
Silver pack: Zee, Sony, Star Plus: 50 Rs.
Gold Pack: Zee, Sony, Star Plus, Discovery, NatGeo: 100 Rs.
Available channels for subscription
Zee: 10 Rs.
Sony: 15 Rs.
Star Plus: 20 Rs.
Discovery: 10 Rs.
NatGeo: 20 Rs.
Available services for subscription
LearnEnglish Service: 200 Rs.
LearnCooking Service: 100 Rs.
============================================================

Enter the option: 4
Subscribe to channel packs
Enter the Pack you wish to subscribe: (Silver: 'S', Gold: 'G'): G
Enter the months: 3
You have successfully subscribed the following packs - Gold
Monthly price: 100 Rs.
No of months: 3
Subscription Amount: 300 Rs.
Discount applied: 30 Rs.
Final Price after discount: 270 Rs.
Account balance: 330 Rs.
Email notification sent successfully
SMS notification sent successfully
============================================================

Enter the option: 5
Add channels to existing subscription
Enter channel names to add (separated by commas): Discovery, Nat Geo
D2H Coding Problem
Channels added successfully.
Account balance: 300 Rs.
============================================================

Enter the option: 6
Subscribe to special services
Enter the service name: LearnEnglish
Service subscribed successfully
Account balance: 100 Rs.
Email notification sent successfully
SMS notification sent successfully
============================================================

Enter the option: 7
View current subscription details
Currently subscribed packs and channels: Gold + Discovery + Nat Geo
Currently subscribed services: LearnEnglish.
============================================================

Enter the option: 8
Update email and phone number for notifications
Enter the email: john@example.com
Enter phone: 9999999999
Email and Phone updated successfully

Some Instructions/Guidelines

1. Develop simple command-line application.
2. Write object-oriented code.
3. Write simple enough code which is easy to understand.
4. Avoid writing too many lines of production code.
5. Include test cases.
6. Avoid using any framework except testing/mocking frameworks.
7. Include all the dependencies with your solution so that we can open it
