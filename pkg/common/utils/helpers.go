package utils

import (
	"crypto/rand"
	"fmt"
	"strings"
	"time"
)

// GenerateRequestID generates a unique request ID
func GenerateRequestID() string {
	return fmt.Sprintf("%d-%s", time.Now().UnixNano(), RandomString(8))
}

// RandomString generates a random string of specified length
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[num.Int64()]
	}
	return string(b)
}

// FormatPhoneNumber formats phone number to standard format
func FormatPhoneNumber(phone string) string {
	// Remove all non-digit characters
	digits := strings.NewReplacer(
		"-", "", " ", "", "(", "", ")", "", "+", "",
	).Replace(phone)
	
	if len(digits) < 10 {
		return phone
	}
	
	// Format: +XX (XXX) XXX-XXXX or similar based on length
	if len(digits) == 10 {
		return fmt.Sprintf("(%s) %s-%s", digits[:3], digits[3:6], digits[6:])
	}
	
	return phone
}

// CalculateAge calculates age from date of birth
func CalculateAge(dateOfBirth time.Time) int {
	today := time.Now()
	age := today.Year() - dateOfBirth.Year()
	
	// Check if birthday has occurred this year
	if today.Month() < dateOfBirth.Month() ||
		(today.Month() == dateOfBirth.Month() && today.Day() < dateOfBirth.Day()) {
		age--
	}
	
	return age
}

// GetDaysSince calculates days since a given date
func GetDaysSince(date time.Time) int {
	return int(time.Since(date).Hours() / 24)
}

// GetDaysUntil calculates days until a given date
func GetDaysUntil(date time.Time) int {
	return int(time.Until(date).Hours() / 24)
}

// TruncateString truncates a string to max length with ellipsis
func TruncateString(str string, maxLength int) string {
	if len(str) <= maxLength {
		return str
	}
	if maxLength < 3 {
		maxLength = 3
	}
	return str[:maxLength-3] + "..."
}

// PascalCase converts string to PascalCase
func PascalCase(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, "")
}

// TitleCase converts string to Title Case
func TitleCase(str string) string {
	return strings.Title(strings.ToLower(str))
}

// PointerToString converts *string to string
func PointerToString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

// StringToPointer converts string to *string
func StringToPointer(str string) *string {
	return &str
}

// PointerToInt converts *int to int
func PointerToInt(ptr *int) int {
	if ptr == nil {
		return 0
	}
	return *ptr
}

// IntToPointer converts int to *int
func IntToPointer(num int) *int {
	return &num
}

// PointerToTime converts *time.Time to time.Time
func PointerToTime(ptr *time.Time) time.Time {
	if ptr == nil {
		return time.Time{}
	}
	return *ptr
}

// TimeToPointer converts time.Time to *time.Time
func TimeToPointer(t time.Time) *time.Time {
	return &t
}

// StringInSlice checks if string exists in slice
func StringInSlice(str string, slice []string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

// RemoveDuplicates removes duplicate strings from slice
func RemoveDuplicates(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// Contains checks if slice contains value
func Contains(slice []int, value int) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// Reverse reverses a string
func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsEmpty checks if string is empty or whitespace
func IsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

// GetMonthName returns month name from month number
func GetMonthName(month int) string {
	months := []string{
		"January", "February", "March", "April",
		"May", "June", "July", "August",
		"September", "October", "November", "December",
	}
	if month < 1 || month > 12 {
		return ""
	}
	return months[month-1]
}

// ParseDate parses date string in format YYYY-MM-DD
func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

// FormatDate formats time.Time to YYYY-MM-DD
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatDateTime formats time.Time to RFC3339
func FormatDateTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

// GetWeekNumber returns week number of year
func GetWeekNumber(t time.Time) int {
	_, week := t.ISOWeek()
	return week
}

// GetQuarter returns quarter of year (1-4)
func GetQuarter(t time.Time) int {
	return (int(t.Month())-1)/3 + 1
}
