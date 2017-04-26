package main

//==============================================================================================================================
//  GetUserDate - Sample data to verify user identity
//==============================================================================================================================
func GetMultipleUserData() []User {
	var userDetails []User
	var user User
	user.FirstName = "John"
	user.LastName = "Doe"
	user.Email = "johndoe@gmail.com"
	user.SSN = "12FG254HG"
	user.BirthDate = "04/23/1989"
	user.PolicyId = "P12345"
	user.VIN = "Gh564HG445"
	user.LicencePlateNumber = "JOHNDOE"
	userDetails[0] = user

	user.FirstName = "Amita"
	user.LastName = "Kamat"
	user.Email = "amitakamat@gmail.com"
	user.SSN = "12F76HG4HG"
	user.BirthDate = "08/07/1990"
	user.PolicyId = "P28789"
	user.VIN = "Gh564HY6678"
	user.LicencePlateNumber = "AMITAK7"
	userDetails[1] = user

	user.FirstName = "Mohammed"
	user.LastName = "Haroon"
	user.Email = "mohammedh@gmail.com"
	user.SSN = "34T56HG4HG"
	user.BirthDate = "02/28/1989"
	user.PolicyId = "P13589"
	user.VIN = "Gh564GR44378"
	user.LicencePlateNumber = "MHAROON"
	userDetails[2] = user

	user.FirstName = "Nethra"
	user.LastName = "Reddy"
	user.Email = "nethrar@gmail.com"
	user.SSN = "90T56HGTY34"
	user.BirthDate = "03/28/1988"
	user.PolicyId = "P84756"
	user.VIN = "RT469HY6678"
	user.LicencePlateNumber = "N283RED"
	userDetails[3] = user

	user.FirstName = "Pavana"
	user.LastName = "Achar"
	user.Email = "pachar@gmail.com"
	user.SSN = "87TR6HGTY34"
	user.BirthDate = "05/05/1991"
	user.PolicyId = "P75985"
	user.VIN = "RT469HT5567"
	user.LicencePlateNumber = "P5ACHAR"
	userDetails[4] = user

	return userDetails
}

func GetSingleUserData() User {
	var newuser User
	newuser.FirstName = "John"
	newuser.LastName = "Doe"
	newuser.Email = "johndoe@gmail.com"
	newuser.SSN = "12FG254HG"
	newuser.BirthDate = "04/23/1989"
	newuser.PolicyId = "P12345"
	newuser.VIN = "Gh564HG445"
	newuser.LicencePlateNumber = "JOHNDOE"

	return newuser
}
