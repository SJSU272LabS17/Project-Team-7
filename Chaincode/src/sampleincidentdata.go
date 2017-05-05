package main

import (
	"time"
)

const INCIDENT_VERIFIED = 1
const INCIDENT_INVALID = 2

//==============================================================================================================================
//	Incident - structure for incident details
//==============================================================================================================================
type Incident struct {
	FirstName          string
	LastName           string
	IncidentDate       time.Time
	PolicyId           string
	VIN                string
	LicencePlateNumber string
	Status             int
}

//==============================================================================================================================
//  GetIncidentsData - Sample data to verify if the incident occured.
//==============================================================================================================================
func GetIncidentsData() [5]Incident {
	var incidentDetails [5]Incident
	var newincident Incident
	newincident.FirstName = "John"
	newincident.LastName = "Doe"
	newincident.PolicyId = "P12345"
	newincident.IncidentDate = time.Date(2017, months["March"], 28, 0, 0, 0, 0, time.UTC)
	newincident.VIN = "Gh564HG445"
	newincident.LicencePlateNumber = "JOHNDOE"
	newincident.Status = 1
	incidentDetails[0] = newincident

	newincident.FirstName = "Amita"
	newincident.LastName = "Kamat"
	newincident.PolicyId = "P28789"
	newincident.IncidentDate = time.Date(2017, months["April"], 2, 0, 0, 0, 0, time.UTC)
	newincident.VIN = "Gh564HY6678"
	newincident.LicencePlateNumber = "AMITAK7"
	newincident.Status = 0
	incidentDetails[1] = newincident

	newincident.FirstName = "Mohammed"
	newincident.LastName = "Haroon"
	newincident.PolicyId = "P13589"
	newincident.IncidentDate = time.Date(2017, months["May"], 1, 0, 0, 0, 0, time.UTC)
	newincident.VIN = "Gh564GR44378"
	newincident.LicencePlateNumber = "MHAROON"
	newincident.Status = 0
	incidentDetails[2] = newincident

	newincident.FirstName = "Nethra"
	newincident.LastName = "Reddy"
	newincident.PolicyId = "P84756"
	newincident.IncidentDate = time.Date(2017, months["February"], 10, 0, 0, 0, 0, time.UTC)
	newincident.VIN = "RT469HY6678"
	newincident.LicencePlateNumber = "N283RED"
	newincident.Status = 1
	incidentDetails[3] = newincident

	newincident.FirstName = "Pavana"
	newincident.LastName = "Achar"
	newincident.PolicyId = "P75985"
	newincident.IncidentDate = time.Date(2016, months["December"], 28, 0, 0, 0, 0, time.UTC)
	newincident.VIN = "RT469HT5567"
	newincident.LicencePlateNumber = "P5ACHAR"
	newincident.Status = 1
	incidentDetails[4] = newincident

	return incidentDetails
}
