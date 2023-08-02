package main

import (
	"testing"
)

// Test the 'validateUserInput' function with valid and invalid input.
func TestValidateUserInput(t *testing.T) {
	// Test with valid input
	if !validateUserInput("John", "Doe", "john@example.com", 2, 10) {
		t.Errorf("Expected 'true', but got 'false' for valid input")
	}

	// Test with invalid first name
	if validateUserInput("J", "Doe", "john@example.com", 2, 10) {
		t.Errorf("Expected 'false', but got 'true' for invalid first name")
	}

	// Test with invalid email address
	if validateUserInput("John", "Doe", "john.example.com", 2, 10) {
		t.Errorf("Expected 'false', but got 'true' for invalid email address")
	}

	// Test with invalid ticket count
	if validateUserInput("John", "Doe", "john@example.com", 11, 10) {
		t.Errorf("Expected 'false', but got 'true' for invalid ticket count")
	}
}

// Test the 'GetfirstNames' function with some test data.
func TestGetfirstNames(t *testing.T) {
	// Create some test data for bookings
	bookings := []userData{
		{firstName: "John", lastName: "Doe", email: "john@example.com", userTickets: 2},
		{firstName: "Jane", lastName: "Smith", email: "jane@example.com", userTickets: 1},
		{firstName: "Mike", lastName: "Johnson", email: "mike@example.com", userTickets: 3},
	}

	// Call the function
	firstNames := GetfirstNames(bookings)

	// Check the results
	expected := []string{"John", "Jane", "Mike"}
	if len(firstNames) != len(expected) {
		t.Errorf("Expected %v first names, but got %v", len(expected), len(firstNames))
	}

	for i, name := range expected {
		if firstNames[i] != name {
			t.Errorf("Expected %v, but got %v", name, firstNames[i])
		}
	}
}

// Add more unit tests for other functions if needed.

// Note: In order to test functions that involve user input/output (e.g., getUserInput, sendTicket),
// you may need to use interfaces and mocking libraries like 'testify/mock' to mock the user input/output.
// For simplicity, we are not testing those functions in this example.

