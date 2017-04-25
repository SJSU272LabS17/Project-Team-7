package main

//==============================================================================================================================
//  GetUserDate - Sample data to verify user identity
//==============================================================================================================================
func GetUserData() User {
	var user User
	user.FirstName = "John"
	user.LastName = "Doe"
	user.Email = "johndoe@gmail.com"
	user.SSN = "12FG254HG"
	user.BirthDate = "04/23/1989"
	user.PolicyId = "P12345"
	user.VIN = "Gh564HG445"
	user.LicencePlateNumber = "JOHNDOE"

	return user
}
