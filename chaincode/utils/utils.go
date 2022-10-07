package utils

import (
	"bytes"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/medical-system/chaincode/lib"
)

// ConstructCaseKey Construct the unique key of the case
func ConstructCaseKey() string {
	lib.NumberOfCases += 1
	return fmt.Sprintf("Case%d",  lib.NumberOfCases+10000)
}

// ConstructRecordKey Construct the unique key of the case
func ConstructRecordKey() string {
	lib.NumberOfRecord += 1
	return fmt.Sprintf("Record%d", lib.NumberOfRecord+10000)
}

// IsContain Determine whether the element is in the array
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// QueryByString Query according to the specified string
func QueryByString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {
	iterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer iterator.Close()
	var buffer bytes.Buffer
	var isSplit bool

	buffer.WriteString("[")
	for iterator.HasNext() {
		result, err := iterator.Next()
		if err != nil {
			return nil, err
		}
		if isSplit {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(result.Value))
		isSplit = true
	}
	buffer.WriteString("]")
	return buffer.Bytes(), nil
}