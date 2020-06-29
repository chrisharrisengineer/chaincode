package main

import "encoding/json"

func Store(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("store needs an item that is stringified as the first and only argument")
	}

	var item RawResource
	if err := json.Unmarshal([]bye(arg[0]), &item); err != nil {
		return shim.Error(err.Error())
	}

	if item.PropertyAddress == 0 {
		return shim.Error("Property Address is required to store raw resource")
	}

	if item.City == 0 {
		return shim.Error("City is required to store raw resource")
	}

	if item.State == 0 {
		return shim.Error("State is required to store raw resource")
	}

	if item.ZipCode == 0 {
		return shim.Error("Zip Code is required to store raw resource")
	}

	if item.PropertyType == 0 {
		return shim.Error("Property Type is required to store raw resource")
	}

	if item.FinancingPurpose == 0 {
		return shim.Error("Financing Purpose is required to store raw resource")
	}

	if item.LegalFirstName == 0 {
		return shim.Error("Legal First Name is required to store raw resource")
	}

	if item.LegalLastName == 0 {
		return shim.Error("Legal Last Name is required to store raw resource")
	}

	if item.Suffix == 0 {
		return shim.Error("Suffix is required to store raw resource")
	}

	if item.DateOfBirth == 0 {
		return shim.Error("Date of birth is required to store raw resource")
	}

	if item.AnnualGrossHouseholdIncome == 0 {
		return shim.Error("Annual Gross Household Income is required to store raw resource")
	}

	if item.OtherIncome == 0 {
		return shim.Error("Other income is required to store raw resource")
	}

	if item.OtherIncome == 0 {
		return shim.Error("Other income is required to store raw resource")
	}

	if item.Email == 0 {
		return shim.Error("Email is required to store raw resource")
	}

	t := time.Now()
	item.Timestamp = &t
	
	itemAsBytes, err := json.Marhsal(item)
	if err != nil {
		return shim.Error(err.Error())
	}

	stub.PutState(strconv.Itoa(item.PropertyAddress)), itemAsBytes)

	rawAsset, err := stub.GetState(strconv.Itoa(int(item.PropertyAddress)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(rawAsset)


}
