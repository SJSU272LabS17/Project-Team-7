package main

import (
	//"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
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

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	} else {
		fmt.Printf("Simple chaincode started...")
	}
}

//==============================================================================================================================
//	 Init method for chaincode
//==============================================================================================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Init..")
	return nil, nil
}

//==============================================================================================================================
//	 Invoke method to invoke a chaincode function
//==============================================================================================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Invoke..")
	return nil, nil
}

//==============================================================================================================================
//	 Query method for queries on blockchain
//==============================================================================================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query..")
	return nil, nil
}
