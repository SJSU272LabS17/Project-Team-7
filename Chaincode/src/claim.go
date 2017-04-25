package main

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
	IncidentDate string
	Status       int
	Amount       string
	UserDetails  User
}

//==============================================================================================================================
//	 NewClaim - created new claim with the parameters passed.
//==============================================================================================================================
func NewClaim(id string, incident_date string, amount string, user_details User) Claim {

	var newClaim Claim

	newClaim.Id = id
	newClaim.IncidentDate = incident_date
	newClaim.UserDetails = user_details
	newClaim.Amount = amount
	newClaim.Status = STATE_INIT_CLAIM

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
