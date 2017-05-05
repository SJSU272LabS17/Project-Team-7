package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var logger = shim.NewLogger("InsuranceLogger")

// CarInsuranceChaincode
type CarInsuranceChaincode struct {
}

//==============================================================================================================================
//	 Participants within the system
//==============================================================================================================================
const POLICY_HOLDER = "policy_holder"
const IDENTITY_INSPECTION = "identity_inspection"
const VEHICLE_INSPECTION = "vehicle_inspection"
const CLAIM_INSPECTION = "claim_inspection"
const SETTLEMENT = "settlement"

//==============================================================================================================================
//	 States within the system
//	 Helps to track the state and perform actions accordingly
//==============================================================================================================================
const STATE_INIT_CLAIM = 0
const STATE_IDENTITY_INSPECTION = 1
const STATE_VEHICLE_INSPECTION = 2
const STATE_CLAIM_INSPECTION = 3
const STATE_SETTLEMENT = 4
const STATE_CANCELLED = 5

func main() {
	// LogDebug, LogInfo, LogNotice, LogWarning, LogError, LogCritical (Default: LogDebug)
	logger.SetLevel(shim.LogInfo)

	logLevel, _ := shim.LogLevel(os.Getenv("SHIM_LOGGING_LEVEL"))
	shim.SetLoggingLevel(logLevel)

	err := shim.Start(new(CarInsuranceChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	} else {
		fmt.Printf("Simple chaincode started...")
	}
}

//==============================================================================================================================
//	 Init method for chaincode
//==============================================================================================================================
func (t *CarInsuranceChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Init..")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1 argument.")
	}

	return nil, nil
}

//==============================================================================================================================
//	 Invoke method to invoke a chaincode function
//==============================================================================================================================
func (t *CarInsuranceChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Invoke..")

	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "createClaim" {
		return t.createClaim(stub, args)
	} else if function == "verifyIdentity" {
		return t.verifyUserIdentity(stub, args[0])
	} else if function == "inspectVehicle" {
		return t.doVehicleInspection(stub, args[0])
	}

	return nil, nil
}

//==============================================================================================================================
//	 Query method for queries on blockchain
//==============================================================================================================================
func (t *CarInsuranceChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query..")

	if function == "getClaim" {
		return t.getClaim(stub, args[0])
	}

	return nil, errors.New("Received unknown function query: " + function)
}

//=================================================================================================================================
//	 createClaim - Creates a new Claim object and saves it.
//   args - ID, IncidentDay, IncidentMonth, IncidentYear, Amount,FirstName, LastName, Email, SSN, BirthDate, PolicyId, VIN, LicencePlateNumber
//=================================================================================================================================
func (t *CarInsuranceChaincode) createClaim(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 13 {
		return nil, errors.New("Incorrect number of arguments. ID, IncidentDate, Amount, FirstName, LastName, Email, SSN, BirthDate, PolicyId, VIN, LicencePlateNumber required.")
	}

	if len(args[4]) == 0 {
		return nil, errors.New("Invalid Amount.")
	}
	year, err := strconv.ParseInt(args[3], 10, 32)
	day, err := strconv.ParseInt(args[1], 10, 32)
	var incidentdate = time.Date(int(year), months[args[2]], int(day), 0, 0, 0, 0, time.UTC)

	var newUser = NewUser(args[5], args[6], args[7], args[8], args[9], args[10], args[11], args[12])

	data, err := strconv.ParseFloat(args[4], 32)
	if err != nil {
		return nil, errors.New("Error getting amount.")
	}
	var newClaim = NewClaim(args[0], incidentdate, data, newUser)

	bytes, err := json.Marshal(newClaim)

	if err != nil {
		return nil, errors.New("Error creating new claim")
	}

	err = stub.PutState(args[0], bytes)

	/*bytes, err = json.Marshal(STATE_INIT_CLAIM)

	if err != nil {
		return nil, errors.New("Error setting init claim state.")
	}*/

	return nil, nil
}

//=================================================================================================================================
//	 getClaim - Gets claim details.
//   args - key
//=================================================================================================================================
func (t *CarInsuranceChaincode) getClaim(stub shim.ChaincodeStubInterface, id string) ([]byte, error) {
	var jsonResp string
	//var err error

	if len(id) == 0 {
		return nil, errors.New("No Id specified. Expecting id to query")
	}

	valAsbytes, err := stub.GetState(id)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get value of " + id + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}

//=================================================================================================================================
//	 getClaimStatus - Gets current state of claim process.
//=================================================================================================================================
func (t *CarInsuranceChaincode) getClaimStatus(stub shim.ChaincodeStubInterface, id string) (int, error) {
	var jsonResp string
	var claimData Claim

	valAsbytes, err := stub.GetState(id)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get claim status\"}"
		return -1, errors.New(jsonResp)
	}

	err = json.Unmarshal(valAsbytes, &claimData)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to UnMarshal claim data\"}"
		return -1, errors.New(jsonResp)
	}

	return claimData.Status, nil
}

//=================================================================================================================================
//	 updateClaimStatus - Updates current state of claim process.
//=================================================================================================================================
func (t *CarInsuranceChaincode) updateClaimStatus(stub shim.ChaincodeStubInterface, claimData Claim) ([]byte, error) {

	bytes, err := json.Marshal(claimData)

	if err != nil {
		return nil, errors.New("Error marshalling claim data")
	}

	err = stub.PutState(claimData.Id, bytes)

	if err != nil {
		return nil, errors.New("Error updating claim status " + string(claimData.Status) + " to the ledger.")
	}

	return nil, nil
}

//=================================================================================================================================
//	 verifyUserIdentity - Verifies user identity (first stage).
//=================================================================================================================================
func (t *CarInsuranceChaincode) verifyUserIdentity(stub shim.ChaincodeStubInterface, id string) ([]byte, error) {
	var claimData Claim
	var jsonResp string
	var userData [5]User
	var claimUser User
	var log string = ""
	userData = GetMultipleUserData()
	data, err := t.getClaim(stub, id)
	var userMatched bool = false

	if err != nil {
		jsonResp = "{\"Error\":\"Failed to retrieve claim details\"}"
		return nil, errors.New(jsonResp)
	}

	err = json.Unmarshal(data, &claimData)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to UnMarshal claim data\"}"
		return nil, errors.New(jsonResp)
	}

	claimUser = claimData.UserDetails

	if claimData.Status == STATE_INIT_CLAIM {
		for i := 0; i < len(userData); i++ {
			if userData[i].FirstName == claimUser.FirstName && userData[i].LastName == claimUser.LastName && userData[i].BirthDate == claimUser.BirthDate && userData[i].Email == claimUser.Email && userData[i].LicencePlateNumber == claimUser.LicencePlateNumber && userData[i].PolicyId == claimUser.PolicyId && userData[i].SSN == claimUser.SSN && userData[i].VIN == claimUser.VIN {
				log = log + "User Details Verified!"
				claimData.Status = STATE_IDENTITY_INSPECTION
				t.updateClaimStatus(stub, claimData)
				userMatched = true
				break
			}
		}

		if userMatched != true {
			claimData.Status = STATE_CANCELLED
			data, err = t.updateClaimStatus(stub, claimData)
			if err != nil {
				jsonResp = "{\"Error\":\"User Identity authentication failed. Status could not be updated.\"}"
				return nil, errors.New(jsonResp)
			} else {
				jsonResp = "{\"Error\":\"User Identity authentication failed. Claim cancelled.\"}"
				return nil, errors.New(jsonResp)
			}
			logger.Infof("User Identity authentication failed")
		}

		data, err = json.Marshal(log)

		if err != nil {
			return nil, errors.New("Error creating log")
		}
	} else {
		jsonResp = "{\"Error\":\"User Identity authentication cannot be done. Claim is not in required state.\"}"
		return nil, errors.New(jsonResp)
	}

	return data, nil

}

//=================================================================================================================================
//	 doVehicleInspection - Inspects Vehicle and updates state (first stage).
//	 id - claim ID
//=================================================================================================================================
func (t *CarInsuranceChaincode) doVehicleInspection(stub shim.ChaincodeStubInterface, id string) ([]byte, error) {
	var claimData Claim
	var jsonResp string
	var incidentData [5]Incident
	var claimUser User
	var log string = ""
	incidentData = GetIncidentsData()
	data, err := t.getClaim(stub, id)
	var userMatched bool = false

	if err != nil {
		jsonResp = "{\"Error\":\"Failed to retrieve claim details for Vehicle Inspection\"}"
		return nil, errors.New(jsonResp)
	}

	err = json.Unmarshal(data, &claimData)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to UnMarshal claim data for Vehicle Inspection\"}"
		return nil, errors.New(jsonResp)
	}

	claimUser = claimData.UserDetails

	if claimData.Status == STATE_IDENTITY_INSPECTION {
		for i := 0; i < len(incidentData); i++ {
			if incidentData[i].FirstName == claimUser.FirstName && incidentData[i].LastName == claimUser.LastName && incidentData[i].LicencePlateNumber == claimUser.LicencePlateNumber && incidentData[i].PolicyId == claimUser.PolicyId && incidentData[i].VIN == claimUser.VIN {
				if incidentData[i].Status == 1 {
					log = log + "Inspection Verified!"
					claimData.Status = STATE_VEHICLE_INSPECTION
					data, err = t.updateClaimStatus(stub, claimData)
					if err != nil {
						jsonResp = "{\"Error\":\"Vehicle Inspection successful. Status could not be updated.\"}"
						return nil, errors.New(jsonResp)
					}
					userMatched = true
					break
				} else {
					log = log + "Inspection Verification Failed. Cancelling claim transaction!"
					claimData.Status = STATE_CANCELLED
					data, err = t.updateClaimStatus(stub, claimData)
					if err != nil {
						jsonResp = "{\"Error\":\"Vehicle Inspection Failed. Status could not be updated.\"}"
						return nil, errors.New(jsonResp)
					}
					userMatched = true
					break
				}
			}
		}

		if userMatched != true {
			claimData.Status = STATE_CANCELLED
			data, err = t.updateClaimStatus(stub, claimData)
			if err != nil {
				jsonResp = "{\"Error\":\"Vehicle Inspection failed. Record not found. Status could not be updated.\"}"
				return nil, errors.New(jsonResp)
			}
			logger.Infof("Vehicle Inspection failed")
		}

	} else {
		jsonResp = "{\"Error\":\"Vehicle Inspection cannot be done. Claim is not in required state.\"}"
		return nil, errors.New(jsonResp)
	}

	return data, nil
}
