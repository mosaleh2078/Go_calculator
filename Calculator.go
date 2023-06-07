package main

import(
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)

func main() {
	// Simple calculator
	fmt.Print("Enter number 1: ")
	reader := bufio.NewReader(os.Stdin)
	var1, err1 := reader.ReadString('\n')
	fmt.Print("Enter number 2: ")
	reader = bufio.NewReader(os.Stdin)
	var2, err2 := reader.ReadString('\n')
	if err1 != nil {
		log.Fatal(err1)
	} else if err2 != nil {
		log.Fatal(err2)
	} else {
		var1 = strings.TrimSpace(var1)
		var2 = strings.TrimSpace(var2)
		num1, err_num_1 := strconv.ParseFloat(var1, 64)
		num2, err_num_2 := strconv.ParseFloat(var2, 64)
		if err_num_1 != nil {
			log.Fatal(err_num_1)
		} else if err_num_2 != nil {
			log.Fatal(err_num_2)
		} else {
			fmt.Print("Choose operator (+ - * /): ")
			reader = bufio.NewReader(os.Stdin)
			op, err_op := reader.ReadString('\n')
			if err_op != nil {
				log.Fatal(err_op)
			} else {
				op = strings.TrimSpace(op)
				if op == "+" {
					fmt.Printf("%.2f %s %.2f = %.2f", num1, op, num2, num1 + num2)
				} else if op == "-" {
					fmt.Printf("%.2f %s %.2f = %.2f", num1, op, num2, num1 - num2)
				} else if op == "*" {
					fmt.Printf("%.2f %s %.2f = %.2f", num1, op, num2, num1 * num2)
				} else if op == "/" {
					fmt.Printf("%.2f %s %.2f = %.2f", num1, op, num2, num1 / num2)
				} else {
					fmt.Println("Wrong operator!")
				}
			}
		}
	}
}