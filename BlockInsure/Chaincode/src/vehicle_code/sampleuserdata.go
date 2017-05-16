package main

//==============================================================================================================================
//  GetUserDate - Sample data to verify user identity
//==============================================================================================================================
func GetMultipleUserData() [5]User {
	var userDetails [5]User
	var newuser User
	newuser.FirstName = "John"
	newuser.LastName = "Doe"
	newuser.Email = "johndoe@gmail.com"
	newuser.SSN = "12FG254HG"
	newuser.BirthDate = "04/23/1989"
	newuser.PolicyId = "P12345"
	newuser.VIN = "Gh564HG445"
	newuser.LicencePlateNumber = "JOHNDOE"
	userDetails[0] = newuser

	newuser.FirstName = "Amita"
	newuser.LastName = "Kamat"
	newuser.Email = "amitakamat@gmail.com"
	newuser.SSN = "12F76HG4HG"
	newuser.BirthDate = "08/07/1990"
	newuser.PolicyId = "P28789"
	newuser.VIN = "Gh564HY6678"
	newuser.LicencePlateNumber = "AMITAK7"
	userDetails[1] = newuser

	newuser.FirstName = "Mohammed"
	newuser.LastName = "Haroon"
	newuser.Email = "mohammedh@gmail.com"
	newuser.SSN = "34T56HG4HG"
	newuser.BirthDate = "02/28/1989"
	newuser.PolicyId = "P13589"
	newuser.VIN = "Gh564GR44378"
	newuser.LicencePlateNumber = "MHAROON"
	userDetails[2] = newuser

	newuser.FirstName = "Nethra"
	newuser.LastName = "Reddy"
	newuser.Email = "nethrar@gmail.com"
	newuser.SSN = "90T56HGTY34"
	newuser.BirthDate = "03/28/1988"
	newuser.PolicyId = "P84756"
	newuser.VIN = "RT469HY6678"
	newuser.LicencePlateNumber = "N283RED"
	userDetails[3] = newuser

	newuser.FirstName = "Pavana"
	newuser.LastName = "Achar"
	newuser.Email = "pachar@gmail.com"
	newuser.SSN = "87TR6HGTY34"
	newuser.BirthDate = "05/05/1991"
	newuser.PolicyId = "P75985"
	newuser.VIN = "RT469HT5567"
	newuser.LicencePlateNumber = "P5ACHAR"
	userDetails[4] = newuser

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