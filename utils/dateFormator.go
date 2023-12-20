package utils

import (
	"fmt"
	"strings"
	"time"
)

func BleagueDate(date string) string {
	parsedDate, err := time.Parse("2006.01.02", date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}

	formattedDate := FormattedDate(parsedDate)

	// fmt.Println("Formatted Date:", formattedDate)

	return formattedDate
}

func BnxtDate(date string) string {
	parsedDate, err := time.Parse("02-Jan-2006", date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}

	formattedDate := FormattedDate(parsedDate)
	// fmt.Println("Formatted Date:", formattedDate)

	return formattedDate
}

func BritishBasketBallDate(date string) string {
	parsedDate, err := time.Parse("Jan 2, 2006, 3:04 PM", date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}

	formattedDate := FormattedDate(parsedDate)
	// fmt.Println("Formatted Date:", formattedDate)

	return formattedDate
}

func EspnDate(date string) string {
	parts := strings.Split(date, " ")

	formattedDate := parts[1]

	return formattedDate
}

func EuroBasketDate(date string) string {
	parsedDate, err := time.Parse("1/2/2006", date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}

	formattedDate := FormattedDate(parsedDate)
	// fmt.Println("Formatted Date:", formattedDate)

	return formattedDate
}

func NblDate(date string) string {
	parts := strings.Split(date, " ")

	if len(parts) != 2 {
		fmt.Println("Invalid date format")
		return ""
	}

	month, err := time.Parse("Jan", parts[0])
	if err != nil {
		return ""
	}

	day := strings.TrimSuffix(parts[1], "th")

	currentYear := time.Now().Year()
	dateStringWithYear := fmt.Sprintf("%s %s %d", month.Format("01"), day, currentYear)

	parsedDate, err := time.Parse("01 2 2006", dateStringWithYear)
	if err != nil {
		return ""
	}

	formattedDate := FormattedDate(parsedDate)
	// fmt.Println("Formatted Date:", formattedDate)

	return formattedDate
}

func FormattedDate(parsedDate time.Time) string {
	formattedDate := parsedDate.Format("01/02")
	return formattedDate
}
