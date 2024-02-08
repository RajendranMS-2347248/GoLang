/* You're developing an online store application in GoLang. 
As part of the application, you need to keep track of various 
product details such as name, price, and quantity in stock. 
Design a set of variables and assign values to represent a 
specific product in the inventory. Ensure you use appropriate 
data types for each variable to accurately capture the information.*/

package main

import "fmt"

func main()  {
	var prodname string
	var prodprice float32
	var qtyinstock int

	fmt.Println("Enter the product name:")
	fmt.Scan(&prodname)
	fmt.Println("Enter the product price:")
	fmt.Scan(&prodprice)
	fmt.Println("Enter the product quantity:")
	fmt.Scan(&qtyinstock)
	if prodname!="" ||  prodprice !=0.0 || qtyinstock > 0{
		fmt.Printf("\nProduct Name : %s \nPrice : Rs.%f \nQuantity In Stock : %d",prodname,prodprice,qtyinstock)
	}else{
	fmt.Println("Enter proper values")
}
}