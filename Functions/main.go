package main

import (
	///"compress/flate"
	"fmt"
)

// func sum(number int, number2 int, name string){
// 	summation := number + number2
// 	fmt.Printf("%s,your answer is %d\n",name,summation)
// }

// func main() {
// sum(50,890,"Kelvin")

//func calculateLoyaltyPoints(amountSpent float64) int {
//	loyaltyPoints := int(amountSpent * 2)
//	return loyaltyPoints
//}

func updatedPoints(currentPoints int, oldPoints int) int{
	 finalPoints := currentPoints + oldPoints

	 return finalPoints
}

func twoValues(vaule1 string, value2 float64) (string,float64){
	return vaule1, value2
}


func main() {
	totalPoints := 120
	calculateLoyaltyPoints := func (amountSpent float64) int {
	loyaltyPoints := int(amountSpent * 2)
	return loyaltyPoints
}

	var newlyEarnedPoints int = calculateLoyaltyPoints(9.30)

	fmt.Println("Earned points today:", newlyEarnedPoints)

	totalPoints = updatedPoints(totalPoints,newlyEarnedPoints)

	val1,val2 := twoValues("garri", 50.6)
	fmt.Printf("%s %.2f\n",val1,val2)

	fmt.Println("Updated loyalty points:", totalPoints) // 120 + 18 = 138
}