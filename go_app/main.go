// Created by: Mr Coxall
// Created on: Sep 2020
//
// This program finds the area and perimeter of a rectangle

package main

import "fmt"

func main() {
	// This function finds the area of trapezoid
	var Base1 int
	var Base2 int
	var Height int
	var Area int

	// input
	fmt.Println("This program finds the varea of trapezoid.")
	fmt.Println()
	fmt.Print("Enter the Base1 (mm): ")
	fmt.Scanln(&Base1)
	fmt.Print("Enter the Base2 (mm): ")
	fmt.Scanln(&Base2)
	fmt.Print("Enter the Height (mm): ")
	fmt.Scanln(&Height)

	// process
	Area = (Base1 + Base2) /2 * Height

	// output
	fmt.Println()
	fmt.Println("The Area is:", Area, "mm2.")
	fmt.Println("\nDone.")
}
