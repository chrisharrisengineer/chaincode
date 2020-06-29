package main

import (
	"time"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


type RawResource struct {
	PropertyAddress 			string 		'json:"property_address"'
	City 						string 		'json:"city"'
	State 						string 		'json:"state"'
	ZipCode 					string 		'json:"zip_code"'
}		


// Heloc implements a simple chaincode to manage an asset
type RawResource struct {
	PropertyAddress 			string 		'json:"property_address"'
	City 						string 		'json:"city"'
	State 						string 		'json:"state"'
	ZipCode 					string 		'json:"zip_code"'
	PropertyType 				string 		'json:"property_type"'
	FinancingPurpose 			string 		'json:"financing_purpose"'
	LegalFirstName 				string 		'json:"legal_first_name"'
	LegalLastName 				string 		'json:"legal_last_name"'
	Suffix 						string 		'json:"suffix"'
	DateOfBirth 				string 		'json:"date_of_birth"'
	AnnualGrossHouseholdIncome 	int 		'json:"annual_gross_household_income"'
	OtherIncome 				int 		'json:"other_income"'
	Email 						string 		'json:"email"'
	ArrivalTime 				*time.Time 	'json:"arrival_time"'
	Timestamp 					*time.Time 	'json:"timestamp"'

}

type Chaincode struct{
	
}

func (c *Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response{
	return shim.Sucess(nil)
}

func (c *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response{

	function, args := stub.GetFunctionAndParameters()

	switch function {
	case "store":
		return Store(stub, args)
	case "index"
		return Index(stub, args)
	default:
		return shim.Error("Available Functions: ")
	}

}

func main() {
	if err := shim.Start(new(Chaincode)); err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}

}