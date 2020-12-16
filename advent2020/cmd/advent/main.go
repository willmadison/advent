package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/willmadison/advent/advent2020"
	"github.com/willmadison/advent/internal/problems"
)

func main() {
	response, err := problems.Fetch(time.Now().Year(), os.Getenv("DAY"), os.Getenv("SESSION_ID"))

	if err != nil {
		fmt.Printf("encountered an error fetching Day %s input: %v", os.Getenv("DAY"), err)
		os.Exit(1)
	}

	defer response.Close()

	rules, tickets := advent2020.ParseTicketRules(response)
	errorRate, validTickets := rules.FindErrorScanRate(tickets[1:])
	fmt.Println(errorRate)

	fieldLocale := rules.DetermineFieldLocale(validTickets)
	fmt.Println("Locale:", fieldLocale)

	var departureFields []int

	for field, index := range fieldLocale {
		if strings.HasPrefix(field, "departure") {
			departureFields = append(departureFields, index)
		}
	}

	departureProduct := 1

	myTicket := tickets[0]

	for _, f := range departureFields {
		departureProduct *= myTicket[f]
	}

	fmt.Println(departureProduct)
}
