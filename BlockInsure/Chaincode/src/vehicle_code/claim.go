package main

import (
	"time"
)

//==============================================================================================================================
//	User - structure for policy holder details
//==============================================================================================================================
type User struct {
	FirstName          string
	LastName           string
	Email              string
	SSN                string
	BirthDate          string
	PolicyId           string
	VIN                string
	LicencePlateNumber string
}

//==============================================================================================================================
//	Claim - structure for claim
//==============================================================================================================================
type Claim struct {
	Id           string
	IncidentDate time.Time
	Status       int
	Amount       float64
	UserDetails  User
	ApplyDate    time.Time
}

//==============================================================================================================================
//	 Constants for Month
//==============================================================================================================================
var months = map[string]time.Month{
	"January":   time.January,
	"February":  time.February,
	"March":     time.March,
	"April":     time.April,
	"May":       time.May,
	"June":      time.June,
	"July":      time.July,
	"August":    time.August,
	"September": time.September,
	"October":   time.October,
	"November":  time.November,
	"December":  time.December,
}

//==============================================================================================================================
//	 NewClaim - created new claim with the parameters passed.
//==============================================================================================================================
func NewClaim(id string, incident_date time.Time, amount float64, user_details User) Claim {

	var newClaim Claim

	newClaim.Id = id
	newClaim.IncidentDate = incident_date
	newClaim.UserDetails = user_details
	newClaim.Amount = amount
	newClaim.Status = STATE_INIT_CLAIM
	newClaim.ApplyDate = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

	return newClaim
}

//==============================================================================================================================
//	 NewClaim - created new claim with the parameters passed.
//==============================================================================================================================
func NewClaimWithState(id string, incident_date time.Time, amount float64, user_details User, state int) Claim {

	var newClaim Claim

	newClaim.Id = id
	newClaim.IncidentDate = incident_date
	newClaim.UserDetails = user_details
	newClaim.Amount = amount
	newClaim.Status = state
	newClaim.ApplyDate = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

	return newClaim
}

//==============================================================================================================================
//	 NewUser - created new user with the parameters passed.
//==============================================================================================================================
func NewUser(first_name string, last_name string, email string, ssn string, birth_date string, policy_id string, vin string, lpn string) User {

	var newUser User

	newUser.FirstName = first_name
	newUser.LastName = last_name
	newUser.Email = email
	newUser.SSN = ssn
	newUser.BirthDate = birth_date
	newUser.PolicyId = policy_id
	newUser.VIN = vin
	newUser.LicencePlateNumber = lpn

	return newUser
}