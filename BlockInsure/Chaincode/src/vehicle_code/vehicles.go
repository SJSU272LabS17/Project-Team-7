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
type SimpleChaincode struct {
}

//==============================================================================================================================
//	Vehicle - Defines the structure for a car object. JSON on right tells it what JSON fields to map to
//			  that element when reading a JSON object into the struct e.g. JSON make -> Struct Make.
//==============================================================================================================================
type Claim struct {
	incidentdate  string `json:"incidentdate"`
	amount        string `json:"amount"`
	vin           int    `json:"VIN"`
	owner         string `json:"owner"`
	applydate     bool   `json:"applydate"`
	status        int    `json:"status"`
	email         string `json:"email"`
	claimID         string `json:"claimID"`
	policyID 	  string `json:"policyID"`
	licenceplatenumber  string  `json:"licenceplatenumber"`
	settled 		string  `json:"settled"`
}

//==============================================================================================================================
//	 Participant types - Each participant type is mapped to an integer which we use to compare to the value stored in a
//						 user's eCert
//==============================================================================================================================
//CURRENT WORKAROUND USES ROLES CHANGE WHEN OWN USERS CAN BE CREATED SO THAT IT READ 1, 2, 3, 4, 5
const   AUTHORITY      =  "regulator"
const 	IDENTITY_INSPECTION = "Identity Inspector"
const 	VEHICLE_INSPECTION = "Vehicle Inspector"
const 	CLAIM_INSPECTION = "Claim_Inspector"
const 	SETTLEMENT = "Settlement Officer"

//==============================================================================================================================
//	 States within the system
//	 Helps to track the state and perform actions accordingly
//==============================================================================================================================

const   STATE_TEMPLATE  			=  0
const 	STATE_IDENTITY_INSPECTION = 1
const 	STATE_VEHICLE_INSPECTION = 2
const 	STATE_CLAIM_INSPECTION = 3
const 	STATE_SETTLEMENT = 4


func main() {
	// LogDebug, LogInfo, LogNotice, LogWarning, LogError, LogCritical (Default: LogDebug)
	logger.SetLevel(shim.LogInfo)

	logLevel, _ := shim.LogLevel(os.Getenv("SHIM_LOGGING_LEVEL"))
	shim.SetLoggingLevel(logLevel)

	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	} else {
		fmt.Printf("Simple chaincode started...")
	}
}

//==============================================================================================================================
//	V5C Holder - Defines the structure that holds all the v5cIDs for vehicles that have been created.
//				Used as an index when querying all vehicles.
//==============================================================================================================================

type Claim_Holder struct {
	claimIDs 	[]string `json:"v5cs"`
}

//==============================================================================================================================
//	User_and_eCert - Struct for storing the JSON of a user and their ecert
//==============================================================================================================================

type User_and_eCert struct {
	Identity string `json:"identity"`
	eCert string `json:"ecert"`
}


//==============================================================================================================================
//	 Init method for chaincode
//==============================================================================================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Init..")
	var claimIDs Claim_Holder

	bytes, err := json.Marshal(claimIDs)

    if err != nil { return nil, errors.New("Error creating claim record") }

	err = stub.PutState("claimIDs", bytes)

	for i:=0; i < len(args); i=i+2 {
		t.add_ecert(stub, args[i], args[i+1])
	}

	return nil, nil
}

//==============================================================================================================================
//	 General Functions
//==============================================================================================================================
//	 get_ecert - Takes the name passed and calls out to the REST API for HyperLedger to retrieve the ecert
//				 for that user. Returns the ecert as retrived including html encoding.
//==============================================================================================================================
func (t *SimpleChaincode) get_ecert(stub shim.ChaincodeStubInterface, name string) ([]byte, error) {

	ecert, err := stub.GetState(name)

	if err != nil { return nil, errors.New("Couldn't retrieve ecert for user " + name) }

	return ecert, nil
}

//==============================================================================================================================
//	 add_ecert - Adds a new ecert and user pair to the table of ecerts
//==============================================================================================================================

func (t *SimpleChaincode) add_ecert(stub shim.ChaincodeStubInterface, name string, ecert string) ([]byte, error) {


	err := stub.PutState(name, []byte(ecert))

	if err == nil {
		return nil, errors.New("Error storing eCert for user " + name + " identity: " + ecert)
	}

	return nil, nil

}


//==============================================================================================================================
//	 get_caller - Retrieves the username of the user who invoked the chaincode.
//				  Returns the username as a string.
//==============================================================================================================================

func (t *SimpleChaincode) get_username(stub shim.ChaincodeStubInterface) (string, error) {

    username, err := stub.ReadCertAttribute("username");
	if err != nil { return "", errors.New("Couldn't get attribute 'username'. Error: " + err.Error()) }
	return string(username), nil
}

//==============================================================================================================================
//	 check_affiliation - Takes an ecert as a string, decodes it to remove html encoding then parses it and checks the
// 				  		certificates common name. The affiliation is stored as part of the common name.
//==============================================================================================================================

func (t *SimpleChaincode) check_affiliation(stub shim.ChaincodeStubInterface) (string, error) {
    affiliation, err := stub.ReadCertAttribute("role");
	if err != nil { return "", errors.New("Couldn't get attribute 'role'. Error: " + err.Error()) }
	return string(affiliation), nil

}

//==============================================================================================================================
//	 get_caller_data - Calls the get_ecert and check_role functions and returns the ecert and role for the
//					 name passed.
//==============================================================================================================================

func (t *SimpleChaincode) get_caller_data(stub shim.ChaincodeStubInterface) (string, string, error){

	user, err := t.get_username(stub)

    // if err != nil { return "", "", err }

	// ecert, err := t.get_ecert(stub, user);

    // if err != nil { return "", "", err }

	affiliation, err := t.check_affiliation(stub);

    if err != nil { return "", "", err }

	return user, affiliation, nil
}

//==============================================================================================================================
//	 retrieve_v5c - Gets the state of the data at v5cID in the ledger then converts it from the stored
//					JSON into the Vehicle struct for use in the contract. Returns the Vehcile struct.
//					Returns empty v if it errors.
//==============================================================================================================================
func (t *SimpleChaincode) retrieve_v5c(stub shim.ChaincodeStubInterface, claimID string) (Claim, error) {

	var v Claim

	bytes, err := stub.GetState(claimID);

	if err != nil {	fmt.Printf("RETRIEVE_V5C: Failed to invoke claim_code: %s", err); return v, errors.New("RETRIEVE_V5C: Error retrieving claim with claimID = " + claimID) }

	err = json.Unmarshal(bytes, &v);

    if err != nil {	fmt.Printf("RETRIEVE_V5C: Corrupt claim record "+string(bytes)+": %s", err); return v, errors.New("RETRIEVE_V5C: Corrupt claim record"+string(bytes))	}

	return v, nil
}

//==============================================================================================================================
// save_changes - Writes to the ledger the Vehicle struct passed in a JSON format. Uses the shim file's
//				  method 'PutState'.
//==============================================================================================================================
func (t *SimpleChaincode) save_changes(stub shim.ChaincodeStubInterface, v Claim) (bool, error) {

	bytes, err := json.Marshal(v)

	if err != nil { fmt.Printf("SAVE_CHANGES: Error converting claim record: %s", err); return false, errors.New("Error converting claim record") }

	err = stub.PutState(v.claimID, bytes)

	if err != nil { fmt.Printf("SAVE_CHANGES: Error storing claim record: %s", err); return false, errors.New("Error storing claim record") }

	return true, nil
}

//==============================================================================================================================
//	 Router Functions
//==============================================================================================================================
//	Invoke - Called on chaincode invoke. Takes a function name passed and calls that function. Converts some
//		  initial arguments passed to other things for use in the called function e.g. name -> ecert
//==============================================================================================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	caller, caller_affiliation, err := t.get_caller_data(stub)

	if err != nil { return nil, errors.New("Error retrieving caller information")}


	if function == "create_claim" {
        return t.create_vehicle(stub, caller, caller_affiliation, args[0])
	} else if function == "ping" {
        return t.ping(stub)
    } else { 																				// If the function is not a create then there must be a car so we need to retrieve the car.
		argPos := 1

		if function == "settle_claim" {																// If its a scrap vehicle then only two arguments are passed (no update value) all others have three arguments and the v5cID is expected in the last argument
			argPos = 0
		}

		v, err := t.retrieve_v5c(stub, args[argPos])

        if err != nil { fmt.Printf("INVOKE: Error retrieving v5c: %s", err); return nil, errors.New("Error retrieving v5c") }


        if strings.Contains(function, "update") == false && function != "settle_claim"    { 									// If the function is not an update or a scrappage it must be a transfer so we need to get the ecert of the recipient.


				if 		   function == "authority_to_identityinspector" { return t.authority_to_identityinspector(stub, v, caller, caller_affiliation, args[0], "IDENTITY_INSPECTION ")
				} else if  function == "identityinspector_to_vehicleinspector"   { return t.identityinspector_to_vehicleinspector(stub, v, caller, caller_affiliation, args[0], "VEHICLE_INSPECTION ")
				} else if  function == "vehicleinspector_to_vehicleinspector" 	   { return t.vehicleinspector_to_vehicleinspector(stub, v, caller, caller_affiliation, args[0], "VEHICLE_INSPECTION ")
				} else if  function == "vehicleinspector_to_claiminspector"  { return t.vehicleinspector_to_claiminspector(stub, v, caller, caller_affiliation, args[0], "CLAIM_INSPECTION")
				} else if  function == "claiminspector_to_vehicleinspector"  { return t.claiminspector_to_vehicleinspector(stub, v, caller, caller_affiliation, args[0], "VEHICLE_INSPECTION")
				} else if  function == "vehicleinspector_to_settlement" { return t.vehicleinspector_to_settlement(stub, v, caller, caller_affiliation, args[0], "SETTLEMENT ")
				}

		} else if function == "update_amount"  	    { return t.update_amount(stub, v, caller, caller_affiliation, args[0])
		} else if function == "update_email"        { return t.update_email(stub, v, caller, caller_affiliation, args[0])
		} else if function == "update_licenceplatenumber" { return t.update_licenceplatenumber(stub, v, caller, caller_affiliation, args[0])
		} else if function == "update_incidentdate" 			{ return t.update_incidentdate(stub, v, caller, caller_affiliation, args[0])
        } else if function == "update_vin" 		{ return t.update_vin(stub, v, caller, caller_affiliation, args[0])
		} else if function == "settle_claim" 		{ return t.settle_claim(stub, v, caller, caller_affiliation) }

		return nil, errors.New("Function of the name "+ function +" doesn't exist.")

	}
}
//=================================================================================================================================
//	Query - Called on chaincode query. Takes a function name passed and calls that function. Passes the
//  		initial arguments passed are passed on to the called function.
//=================================================================================================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	caller, caller_affiliation, err := t.get_caller_data(stub)
	if err != nil { fmt.Printf("QUERY: Error retrieving caller details", err); return nil, errors.New("QUERY: Error retrieving caller details: "+err.Error()) }

    logger.Debug("function: ", function)
    logger.Debug("caller: ", caller)
    logger.Debug("affiliation: ", caller_affiliation)

	if function == "get_claim_details" {
		if len(args) != 1 { fmt.Printf("Incorrect number of arguments passed"); return nil, errors.New("QUERY: Incorrect number of arguments passed") }
		v, err := t.retrieve_v5c(stub, args[0])
		if err != nil { fmt.Printf("QUERY: Error retrieving v5c: %s", err); return nil, errors.New("QUERY: Error retrieving v5c "+err.Error()) }
		return t.get_claim_details(stub, v, caller, caller_affiliation)
	} else if function == "check_unique_v5c" {
		return t.check_unique_v5c(stub, args[0], caller, caller_affiliation)
	} else if function == "get_claims" {
		return t.get_claims(stub, caller, caller_affiliation)
	} else if function == "get_ecert" {
		return t.get_ecert(stub, args[0])
	} else if function == "ping" {
		return t.ping(stub)
	}

	return nil, errors.New("Received unknown function invocation " + function)

}

//=================================================================================================================================
//	 Ping Function
//=================================================================================================================================
//	 Pings the peer to keep the connection alive
//=================================================================================================================================
func (t *SimpleChaincode) ping(stub shim.ChaincodeStubInterface) ([]byte, error) {
	return []byte("Hello, world!"), nil
}

//=================================================================================================================================
//	 Create Function
//=================================================================================================================================
//	 Create Vehicle - Creates the initial JSON for the vehcile and then saves it to the ledger.
//=================================================================================================================================
func (t *SimpleChaincode) create_claim(stub shim.ChaincodeStubInterface, caller string, caller_affiliation string, claimID string) ([]byte, error) {
	var v Claim

	claimID         := "\"claimID\":\""+ claimID +"\", "							// Variables to define the JSON
	vin            := "\"VIN\":0, "
	amount           := "\"Amount\":\"UNDEFINED\", "
	incidentdate          := "\"Incidentdate\":\"UNDEFINED\", "
	applydate          := "\"Applydate\":\"UNDEFINED\", "
	owner          := "\"Owner\":\""+caller+"\", "
	email         := "\"Email\":\"UNDEFINED\", "
	policyID  := "\"PolicyID\":\"UNDEFINED\", "
	licenceplatenumber  := "\"Licenceplatenumber\":\"UNDEFINED\", "
	status         := "\"Status\":0, "
	settled       := "\"Settled\":false"

	claim_json := "{"+claimID+vin+amount+incidentdate+applydate+owner+email+policyID+licenceplatenumber+status+settled+"}" 	// Concatenates the variables to create the total JSON object

	matched, err := regexp.Match("^[A-z][A-z][0-9]{7}", []byte(claimID))  				// matched = true if the v5cID passed fits format of two letters followed by seven digits

												if err != nil { fmt.Printf("CREATE_CLAIM: Invalid claimID: %s", err); return nil, errors.New("Invalid claimID") }

	if 				claimID  == "" 	 ||
					matched == false    {
																		fmt.Printf("CREATE_VEHICLE: Invalid v5cID provided");
																		return nil, errors.New("Invalid v5cID provided")
	}

	err = json.Unmarshal([]byte(claim_json), &v)							// Convert the JSON defined above into a vehicle object for go

																		if err != nil { return nil, errors.New("Invalid JSON object") }

	record, err := stub.GetState(v.claimID) 								// If not an error then a record exists so cant create a new car with this V5cID as it must be unique

																		if record != nil { return nil, errors.New("Vehicle already exists") }

	if 	caller_affiliation != AUTHORITY {							// Only the regulator can create a new v5c

		return nil, errors.New(fmt.Sprintf("Permission Denied. create_claim. %v === %v", caller_affiliation, AUTHORITY))

	}

	_, err  = t.save_changes(stub, v)

																		if err != nil { fmt.Printf("CREATE_CLAIM: Error saving changes: %s", err); return nil, errors.New("Error saving changes") }

	bytes, err := stub.GetState("claimIDs")

																		if err != nil { return nil, errors.New("Unable to get claimIDs") }

	var claimIDs claim_Holder

	err = json.Unmarshal(bytes, &claimIDs)

																		if err != nil {	return nil, errors.New("Corrupt claim_Holder record") }

	claimIDs.V5Cs = append(claimIDs.V5Cs, claimID)


	bytes, err = json.Marshal(claimIDs)

															if err != nil { fmt.Print("Error creating claim_Holder record") }

	err = stub.PutState("claimIDs", bytes)

															if err != nil { return nil, errors.New("Unable to put the state") }

	return nil, nil

}

//=================================================================================================================================
//	 Transfer Functions
//=================================================================================================================================
//	 authority_to_manufacturer
//=================================================================================================================================
func (t *SimpleChaincode) authority_to_identityinspector(stub shim.ChaincodeStubInterface, v Claim, caller string, caller_affiliation string, recipient_name string, recipient_affiliation string) ([]byte, error) {

	if     	v.Status				== STATE_TEMPLATE	&&
			v.Owner					== caller			&&
			caller_affiliation		== AUTHORITY		&&
			recipient_affiliation	== IDENTITY_INSPECTION 		&&
			v.Settled				== false			{		// If the roles and users are ok

					v.Owner  = recipient_name		// then make the owner the new owner
					v.Status = STATE_IDENTITY_INSPECTION 			// and mark it in the state of manufacture

	} else {									// Otherwise if there is an error
															fmt.Printf("AUTHORITY_TO_IDENITY_INSPECTION: Permission Denied");
                                                            return nil, errors.New(fmt.Sprintf("Permission Denied. authority_to_identity_inspection. %v %v === %v, %v === %v, %v === %v, %v === %v, %v === %v", v, v.Status, STATE_VEHICLE_INSPECTION , v.Owner, caller, caller_affiliation, VEHICLE_INSPECTION , recipient_affiliation, SETTLEMENT , v.Settled, false))


	}

	_, err := t.save_changes(stub, v)						// Write new state

															if err != nil {	fmt.Printf("AUTHORITY_TO_IDENITY_INSPECTION: Error saving changes: %s", err); return nil, errors.New("Error saving changes")	}

	return nil, nil									// We are Done

}

//=================================================================================================================================
//	 manufacturer_to_private
//=================================================================================================================================
func (t *SimpleChaincode) identityinspector_to_vehicleinspector(stub shim.ChaincodeStubInterface, v Claim, caller string, caller_affiliation string, recipient_name string, recipient_affiliation string) ([]byte, error) {

	if 		v.Amount 	 == "UNDEFINED" ||
			v.Incidentdate  == "UNDEFINED" ||
			v.Licenceplatenumber 	 == "UNDEFINED" ||
			v.Email == "UNDEFINED" ||
			v.VIN == 0				{					//If any part of the claim is undefined it has not bene fully manufacturered so cannot be sent
															fmt.Printf("MANUFACTURER_TO_VEHICLE_INSPECTION : Claim not fully defined")
															return nil, errors.New(fmt.Sprintf("Claim not fully defined. %v", v))
	}

	if 		v.Status				== STATE_IDENTITY_INSPECTION 	&&
			v.Owner					== caller				&&
			caller_affiliation		== IDENTITY_INSPECTION 			&&
			recipient_affiliation	== VEHICLE_INSPECTION 		&&
			v.Settled    == false							{

					v.Owner = recipient_name
					v.Status = STATE_VEHICLE_INSPECTION 

	} else {
        return nil, errors.New(fmt.Sprintf("Permission Denied. identityinspector_to_vehicleinspector. %v %v === %v, %v === %v, %v === %v, %v === %v, %v === %v", v, v.Status, STATE_VEHICLE_INSPECTION, v.Owner, caller, caller_affiliation, VEHICLE_INSPECTION , recipient_affiliation, SETTLEMENT , v.Settled, false))
    }

	_, err := t.save_changes(stub, v)

	if err != nil { fmt.Printf("IDENTITY_INSPECTION _TO_VEHICLE_INSPECTION: Error saving changes: %s", err); return nil, errors.New("Error saving changes") }

	return nil, nil

}

//=================================================================================================================================
//	 private_to_private
//=================================================================================================================================
func (t *SimpleChaincode) vehicleinspector_to_vehicleinspector(stub shim.ChaincodeStubInterface, v Claim, caller string, caller_affiliation string, recipient_name string, recipient_affiliation string) ([]byte, error) {

	if 		v.Status				== STATE_VEHICLE_INSPECTION 	&&
			v.Owner					== caller					&&
			caller_affiliation		== VEHICLE_INSPECTION 			&&
			recipient_affiliation	== VEHICLE_INSPECTION 			&&
			v.Settled				== false					{

					v.Owner = recipient_name

	} else {
        return nil, errors.New(fmt.Sprintf("Permission Denied. vehicleinspector_to_vehicleinspector. %v %v === %v, %v === %v, %v === %v, %v === %v, %v === %v", v, v.Status, STATE_VEHICLE_INSPECTION , v.Owner, caller, caller_affiliation, VEHICLE_INSPECTION , recipient_affiliation, SETTLEMENT , v.Settled, false))
	}

	_, err := t.save_changes(stub, v)

															if err != nil { fmt.Printf("VEHICLE_INSPECTION _TO_VEHICLE_INSPECTION : Error saving changes: %s", err); return nil, errors.New("Error saving changes") }

	return nil, nil

}

//=================================================================================================================================
//	 private_to_lease_company
//=================================================================================================================================
func (t *SimpleChaincode) vehicleinspector_to_claiminspector(stub shim.ChaincodeStubInterface, v Claim, caller string, caller_affiliation string, recipient_name string, recipient_affiliation string) ([]byte, error) {

	if 		v.Status				== STATE_VEHICLE_INSPECTION	&&
			v.Owner					== caller					&&
			caller_affiliation		== VEHICLE_INSPECTION			&&
			recipient_affiliation	== CLAIM_INSPECTION 			&&
            v.Settled    			== false					{

					v.Owner = recipient_name

	} else {
        return nil, errors.New(fmt.Sprintf("Permission denied. vehicleinspector_to_claiminspector. %v === %v, %v === %v, %v === %v, %v === %v, %v === %v", v.Status, STATE_VEHICLE_INSPECTION, v.Owner, caller, caller_affiliation, VEHICLE_INSPECTION, recipient_affiliation, SETTLEMENT, v.Settled, false))

	}

	_, err := t.save_changes(stub, v)
															if err != nil { fmt.Printf("VEHICLE_INSPECTION_TO_CLAIM_INSPECTION: Error saving changes: %s", err); return nil, errors.New("Error saving changes") }

	return nil, nil

}

//=================================================================================================================================
//	 lease_company_to_private
//=================================================================================================================================
func (t *SimpleChaincode) claiminspector_to_vehicleinspector(stub shim.ChaincodeStubInterface, v Claim, caller string, caller_affiliation string, recipient_name string, recipient_affiliation string) ([]byte, error) {

	if		v.Status				== STATE_VEHICLE_INSPECTION	&&
			v.Owner  				== caller					&&
			caller_affiliation		== CLAIM_INSPECTION 			&&
			recipient_affiliation	== VEHICLE_INSPECTION 			&&
			v.Scrapped				== false					{

				v.Owner = recipient_name

	} else {
		return nil, errors.New(fmt.Sprintf("Permission Denied. claiminspector_to_vehicleinspector. %v %v === %v, %v === %v, %v === %v, %v === %v, %v === %v", v, v.Status, STATE_VEHICLE_INSPECTION , v.Owner, caller, caller_affiliation, VEHICLE_INSPECTION , recipient_affiliation, SETTLEMENT , v.Settled, false))
	}

	_, err := t.save_changes(stub, v)
															if err != nil { fmt.Printf("CLAIM_INSPECTION_TO_VEHICLE_INSPECTION : Error saving changes: %s", err); return nil, errors.New("Error saving changes") }

	return nil, nil

}

//=================================================================================================================================
//	 private_to_scrap_merchant
//=================================================================================================================================
func (t *SimpleChaincode) vehicleinspector_to_settlement(stub shim.ChaincodeStubInterface, v Claim, caller string, caller_affiliation string, recipient_name string, recipient_affiliation string) ([]byte, error) {

	if		v.Status				== STATE_VEHICLE_INSPECTION	&&
			v.Owner					== caller					&&
			caller_affiliation		== VEHICLE_INSPECTION 			&&
			recipient_affiliation	== SETTLEMENT 			&&
			v.Settled				== false					{

					v.Owner = recipient_name
					v.Status = STATE_SETTLEMENT 

	} else {
        return nil, errors.New(fmt.Sprintf("Permission Denied. vehicleinspector_to_settlement. %v %v === %v, %v === %v, %v === %v, %v === %v, %v === %v", v, v.Status, STATE_VEHICLE_INSPECTION , v.Owner, caller, caller_affiliation, VEHICLE_INSPECTION , recipient_affiliation, SETTLEMENT, v.Settled, false))
	}

	_, err := t.save_changes(stub, v)

															if err != nil { fmt.Printf("VEHICLE_INSPECTION_TO_SETTLEMENT: Error saving changes: %s", err); return nil, errors.New("Error saving changes") }

	return nil, nil

}

//=================================================================================================================================
//	 Update Functions
//=================================================================================================================================
//	 update_vin
//=================================================================================================================================
func (t *SimpleChaincode) update_vin(stub shim.ChaincodeStubInterface, v Claim, caller string, caller_affiliation string, new_value string) ([]byte, error) {

	new_vin, err := strconv.Atoi(string(new_value)) 		                // will return an error if the new vin contains non numerical chars

															if err != nil || len(string(new_value)) != 15 { return nil, errors.New("Invalid value passed for new VIN") }

	if 		//v.Status			== STATE_MANUFACTURE	&&
			v.Owner				== caller				&&
			//caller_affiliation	== MANUFACTURER			&&
			//v.VIN				== 0					&&			// Can't change the VIN after its initial assignment
			v.Settled			== false				{

					v.VIN = new_vin					// Update to the new value
	} else {

        return nil, errors.New(fmt.Sprintf("Permission denied. update_vin %v %v %v %v %v", v.Status, STATE_IDENTITY_INSPECTION , v.Owner, caller, v.VIN, v.Settled))

	}

	_, err  = t.save_changes(stub, v)						// Save the changes in the blockchain

															if err != nil { fmt.Printf("UPDATE_VIN: Error saving changes: %s", err); return nil, errors.New("Error saving changes") }

	return nil, nil

}


//=================================================================================================================================
//	 update_registration
//=================================================================================================================================
func (t *SimpleChaincode) update_email(stub shim.ChaincodeStubInterface, v Vehicle, caller string, caller_affiliation string, new_value string) ([]byte, error) {


	if		v.Owner				== caller			&&
			//caller_affiliation	!= SCRAP_MERCHANT	&&
			v.Settled			== false			{

					v.Email = new_value

	} else {
        return nil, errors.New(fmt.Sprint("Permission denied. update_email"))
	}

	_, err := t.save_changes(stub, v)

															if err != nil { fmt.Printf("UPDATE_EMAIL: Error saving changes: %s", err); return nil, errors.New("Error saving changes") }

	return nil, nil

}

//=================================================================================================================================
//	 update_colour
//=================================================================================================================================
func (t *SimpleChaincode) update_amount(stub shim.ChaincodeStubInterface, v Vehicle, caller string, caller_affiliation string, new_value string) ([]byte, error) {

	if 		v.Owner				== caller				&&
			//caller_affiliation	== MANUFACTURER			&&/*((v.Owner				== caller			&&
			//caller_affiliation	== MANUFACTURER)		||
			//caller_affiliation	== AUTHORITY)			&&*/
			v.Settled		== false				{

					v.Amount = new_value
	} else {

		return nil, errors.New(fmt.Sprint("Permission denied. update_amount" ))
	}

	_, err := t.save_changes(stub, v)

		if err != nil { fmt.Printf("UPDATE_AMOUNT: Error saving changes: %s", err); return nil, errors.New("Error saving changes") }

	return nil, nil

}

//=================================================================================================================================
//	 update_make
//=================================================================================================================================
func (t *SimpleChaincode) update_licenceplatenumber(stub shim.ChaincodeStubInterface, v Claim, caller string, caller_affiliation string, new_value string) ([]byte, error) {

	if 		//v.Status			== STATE_MANUFACTURE	&&
			v.Owner				== caller				&&
			//caller_affiliation	== MANUFACTURER			&&
			v.Settled			== false				{

					v.LicencePlateNumber = new_value
	} else {

        return nil, errors.New(fmt.Sprint("Permission denied. update_licenceplatenumber"))


	}

	_, err := t.save_changes(stub, v)

															if err != nil { fmt.Printf("UPDATE_MAKE: Error saving changes: %s", err); return nil, errors.New("Error saving changes") }

	return nil, nil

}

//=================================================================================================================================
//	 update_model
//=================================================================================================================================
func (t *SimpleChaincode) update_incidentdate(stub shim.ChaincodeStubInterface, v Claim, caller string, caller_affiliation string, new_value string) ([]byte, error) {

	if 		//v.Status			== STATE_MANUFACTURE	&&
			v.Owner				== caller				&&
			//caller_affiliation	== MANUFACTURER			&&
			v.Settled		== false				{

					v.Incidentdate = new_value

	} else {
        return nil, errors.New(fmt.Sprint("Permission denied. update_incidentdate")

	}

	_, err := t.save_changes(stub, v)

															if err != nil { fmt.Printf("UPDATE_MODEL: Error saving changes: %s", err); return nil, errors.New("Error saving changes") }

	return nil, nil

}

//=================================================================================================================================
//	 scrap_vehicle
//=================================================================================================================================
func (t *SimpleChaincode) settle_claim(stub shim.ChaincodeStubInterface, v Claim, caller string, caller_affiliation string) ([]byte, error) {

	if		v.Status			== STATE_SETTLEMENT 	&&
			v.Owner				== caller				&&
			caller_affiliation	== SETTLEMENT 		&&
			v.Settled			== false				{

					v.Settled = true

	} else {
		return nil, errors.New("Permission denied. settle_claim")
	}

	_, err := t.save_changes(stub, v)

															if err != nil { fmt.Printf("SETTLE_CLAIM: Error saving changes: %s", err); return nil, errors.New("SETTLE CLAIM Error saving changes") }

	return nil, nil

}

//=================================================================================================================================
//	 Read Functions
//=================================================================================================================================
//	 get_vehicle_details
//=================================================================================================================================
func (t *SimpleChaincode) get_claim_details(stub shim.ChaincodeStubInterface, v Claim, caller string, caller_affiliation string) ([]byte, error) {

	bytes, err := json.Marshal(v)

																if err != nil { return nil, errors.New("GET_CLAIM_DETAILS: Invalid claim object") }

	if 		v.Owner				== caller		||
			caller_affiliation	== AUTHORITY	{

					return bytes, nil
	} else {
																return nil, errors.New("Permission Denied. get_claim_details")
	}

}

//=================================================================================================================================
//	 get_vehicles
//=================================================================================================================================

func (t *SimpleChaincode) get_claims(stub shim.ChaincodeStubInterface, caller string, caller_affiliation string) ([]byte, error) {
	bytes, err := stub.GetState("claimIDs")

																			if err != nil { return nil, errors.New("Unable to get claimIDs") }

	var claimIDs claim_Holder

	err = json.Unmarshal(bytes, &claimIDs)

																			if err != nil {	return nil, errors.New("Corrupt claim_Holder") }

	result := "["

	var temp []byte
	var v Claim

	for _, v5c := range claimIDs.V5Cs {

		v, err = t.retrieve_v5c(stub, v5c)

		if err != nil {return nil, errors.New("Failed to retrieve V5C")}

		temp, err = t.get_claim_details(stub, v, caller, caller_affiliation)

		if err == nil {
			result += string(temp) + ","
		}
	}

	if len(result) == 1 {
		result = "[]"
	} else {
		result = result[:len(result)-1] + "]"
	}

	return []byte(result), nil
}

//=================================================================================================================================
//	 check_unique_v5c
//=================================================================================================================================
func (t *SimpleChaincode) check_unique_v5c(stub shim.ChaincodeStubInterface, v5c string, caller string, caller_affiliation string) ([]byte, error) {
	_, err := t.retrieve_v5c(stub, v5c)
	if err == nil {
		return []byte("false"), errors.New("V5C is not unique")
	} else {
		return []byte("true"), nil
	}
}
