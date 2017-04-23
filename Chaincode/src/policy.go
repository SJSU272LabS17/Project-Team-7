package main

//==============================================================================================================================
//	Policy - structure for policy details
//==============================================================================================================================

type Policy struct {
	Id        string
	Type      string
	OwnerId   string
	Insurer   string
	CarNumber string
	StartDate string
	EndDate   string
}
