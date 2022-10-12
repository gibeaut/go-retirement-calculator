package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func Get_Info() (string, int, int, float64, int) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter birthday mm/dd/yyyy: ")
	input, err := reader.ReadString('\n')

	birthdayString := strings.Split(input, "\n")[0]

	fmt.Print("At what age do you expect to die (you can check actuarial tables online): ")
	input, err = reader.ReadString('\n')

	deathString := strings.Split(input, "\n")[0]

	var death int64
	death, err = strconv.ParseInt(deathString, int(0), int(32))

	fmt.Print("How much do you have saved: ")
	input, err = reader.ReadString('\n')

	savingsString := strings.Split(input, "\n")[0]

	var savings int64
	savings, err = strconv.ParseInt(savingsString, int(0), int(32))

	fmt.Print("What is your expected interest rate (EX. 0.04 if 4%): ")
	input, err = reader.ReadString('\n')

	rateString := strings.Split(input, "\n")[0]

	var rate float64
	rate, err = strconv.ParseFloat(rateString, int(64))

	fmt.Print("What will be your daily retirement budget: ")
	input, err = reader.ReadString('\n')

	budgetString := strings.Split(input, "\n")[0]

	var budget int64
	budget, err = strconv.ParseInt(budgetString, int(0), int(32))

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return "09/09/1999", 99, 1000, 0.04, 100
	}

	return birthdayString, int(death), int(savings), rate, int(budget)
}

func Get_Retirement_Date(birthday time.Time, death int, dailyrate float64, budget int, savings int) (int, float64) {

	ageInDays := math.Round(time.Since(birthday).Hours() / 24)

	daysToDeath := (death * 365) - int(ageInDays)

	totalAmountNeeded := daysToDeath * budget

	numberOfDaysFromToday := 0

	requiredInSavings := float64(savings)

	for (totalAmountNeeded-int(requiredInSavings)) > 0 && daysToDeath > 1 {
		numberOfDaysFromToday++
		daysToDeath--
		requiredInSavings = requiredInSavings * dailyrate
		totalAmountNeeded = totalAmountNeeded - budget
	}

	return numberOfDaysFromToday, requiredInSavings
}

func main() {
	timeFormat := "01/02/2006"

	birthdayString, death, savings, rate, budget := Get_Info()

	birthday, _ := time.Parse(timeFormat, birthdayString)

	var powerRate float64 = 0.00273972602
	dailyRate := math.Pow((rate + 1.0), powerRate)

	daysleft, totalSavings := Get_Retirement_Date(birthday, death, dailyRate, int(budget), int(savings))

	retirementDate := time.Now().AddDate(0, 0, daysleft)

	fmt.Printf("You can retire on %s with %d\n", retirementDate.Format(timeFormat), int(totalSavings))

}
