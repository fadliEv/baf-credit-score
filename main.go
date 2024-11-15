package main

import (	
	"baf-credit-score/delivery"
)

func main() {
	delivery.NewServer().Run()
}