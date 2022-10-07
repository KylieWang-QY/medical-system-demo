package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/medical-system/chaincode/lib"
	"github.com/medical-system/chaincode/utils"
	"github.com/rs/xid"
	"time"
)

type MedicalSystem struct {
}

func (t *MedicalSystem) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// Initialize five patient information
	patient1 := lib.Patient{ObjectType: lib.PatientObj, Username: "maryyy", Name: "Mary", Address: "2501 W Colorado Ave, Colorado Springs, CO 80904", Telephone: "626-780-7552", Id: xid.New().String()}
	informationAsBytes, err := json.Marshal(patient1)
	if err != nil {
		errMes := fmt.Sprintf("Marshal error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}
	err = stub.PutState(patient1.Username, informationAsBytes)
	if err != nil {
		errMes := fmt.Sprintf("PutState error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}
	record1 := lib.UsageRecord{ObjectType: lib.UsageRecordObj, Id: patient1.Id, Operations: lib.Add, Roles: "patient", Username: patient1.Username, Time: time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")}
	recordAsBytes, err := json.Marshal(record1)
	if err != nil {
		errMes := fmt.Sprintf("Marshal error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}
	err = stub.PutState(utils.ConstructRecordKey(), recordAsBytes)
	if err != nil {
		errMes := fmt.Sprintf("PutState error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}

	patient2 := lib.Patient{ObjectType: lib.PatientObj, Username: "helenlen", Name: "Helen", Address: "15 Independence Blvd. 1st Fl. Warren , NJ 07059", Telephone: "646-952-7115", Id: xid.New().String()}
	informationAsBytes, err = json.Marshal(patient2)
	if err != nil {
		errMes := fmt.Sprintf("Marshal error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}
	err = stub.PutState(patient2.Username, informationAsBytes)
	if err != nil {
		errMes := fmt.Sprintf("PutState error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}
	record2 := lib.UsageRecord{ObjectType: lib.UsageRecordObj, Id: patient2.Id, Operations: lib.Add, Roles: "patient", Username: patient2.Username, Time: time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")}
	recordAsBytes2, err := json.Marshal(record2)
	if err != nil {
		errMes := fmt.Sprintf("Marshal error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}
	err = stub.PutState(utils.ConstructRecordKey(), recordAsBytes2)
	if err != nil {
		errMes := fmt.Sprintf("PutState error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}

	patient3 := lib.Patient{ObjectType: lib.PatientObj, Username: "akanancy", Name: "Nancy", Address: "10815 Ranch Road 2222. Austin, TX 78730", Telephone: "658-152-1545", Id: xid.New().String()}
	informationAsBytes, err = json.Marshal(patient3)
	if err != nil {
		errMes := fmt.Sprintf("Marshal error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}
	err = stub.PutState(patient3.Username, informationAsBytes)
	if err != nil {
		errMes := fmt.Sprintf("PutState error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}
	record3 := lib.UsageRecord{ObjectType: lib.UsageRecordObj, Id: patient3.Id, Operations: lib.Add, Roles: "patient", Username: patient3.Username, Time: time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")}
	recordAsBytes3, err := json.Marshal(record3)
	if err != nil {
		errMes := fmt.Sprintf("Marshal error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}
	err = stub.PutState(utils.ConstructRecordKey(), recordAsBytes3)
	if err != nil {
		errMes := fmt.Sprintf("PutState error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}

	patient4 := lib.Patient{ObjectType: lib.PatientObj, Username: "evayam", Name: "Eva", Address: "4263 City Terrace Dr , Los Angeles, CA 90063", Telephone: "917-154-1685", Id: xid.New().String()}
	informationAsBytes, err = json.Marshal(patient4)
	if err != nil {
		errMes := fmt.Sprintf("Marshal error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}
	err = stub.PutState(patient4.Username, informationAsBytes)
	if err != nil {
		errMes := fmt.Sprintf("PutState error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}
	record4 := lib.UsageRecord{ObjectType: lib.UsageRecordObj, Id: patient4.Id, Operations: lib.Add, Roles: "patient", Username: patient4.Username, Time: time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")}
	recordAsBytes4, err := json.Marshal(record4)
	if err != nil {
		errMes := fmt.Sprintf("Marshal error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}
	err = stub.PutState(utils.ConstructRecordKey(), recordAsBytes4)
	if err != nil {
		errMes := fmt.Sprintf("PutState error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}

	patient5 := lib.Patient{ObjectType: lib.PatientObj, Username: "francesaa", Name: "Frances", Address: "1055 Thomas Jefferson St NW, Washington, DC 20007", Telephone: "541-312-1548", Id: xid.New().String()}
	informationAsBytes, err = json.Marshal(patient5)
	if err != nil {
		errMes := fmt.Sprintf("Marshal error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}
	err = stub.PutState(patient5.Username, informationAsBytes)
	if err != nil {
		errMes := fmt.Sprintf("PutState error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}
	record5 := lib.UsageRecord{ObjectType: lib.UsageRecordObj, Id: patient5.Id, Operations: lib.Add, Roles: "patient", Username: patient5.Username, Time: time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")}
	recordAsBytes5, err := json.Marshal(record5)
	if err != nil {
		errMes := fmt.Sprintf("Marshal error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}
	err = stub.PutState(utils.ConstructRecordKey(), recordAsBytes5)
	if err != nil {
		errMes := fmt.Sprintf("PutState error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}

	return pb.Response{Status: 200, Message: "Initialize successful", Payload: nil}
}

func (t *MedicalSystem) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "queryInformation" {
		// Doctor or patient query information
		return t.queryInformation(stub, args)
	} else if function == "authorizeDoctor" {
		// Authorized to the doctor
		return t.authorizeDoctor(stub, args)
	} else if function == "enterData" {
		// Enter the patient's medical data
		return t.enterData(stub, args)
	} else if function == "queryCases" {
		// researcher or regulator query cases
		return t.queryCases(stub, args)
	} else if function == "showRecords" {
		// Show all ID usage records
		return t.showRecords(stub, args)
	}

	return shim.Error("Invalid invoke function name.")
}

// Doctor or patient query information
func (t *MedicalSystem) queryInformation(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// check the number of parameters
	if len(args) != 3 && len(args) != 4 {
		errMes := "Incorrect number of arguments. Expecting 3 or 4"
		return pb.Response{Status: 400, Message: errMes, Payload: []byte(errMes)}
	}

	// judge if this is a doctor or a patient
	if args[0] == "patient" {
		username := args[1]
		ntime := args[2]
		// Get the state from the ledger
		informationAsBytes, err := stub.GetState(username)
		if err != nil {
			errMes := fmt.Sprintf("GetState error: %s", err)
			return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
		} else if informationAsBytes == nil {
			errMes := "The patient does not exist"
			return pb.Response{Status: 400, Message: errMes, Payload: []byte(errMes)}
		}

		var patient lib.Patient
		err = json.Unmarshal(informationAsBytes, &patient)
		if err != nil {
			errMes := fmt.Sprintf("Unmarshal error: %s", err)
			return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
		}

		// Record the reading operation
		record := lib.UsageRecord{ObjectType: lib.UsageRecordObj, Id: patient.Id, Operations: lib.Read, Roles: args[0], Username: username, Time: ntime}
		recordAsBytes, err := json.Marshal(record)
		if err != nil {
			errMes := fmt.Sprintf("Marshal error: %s", err)
			return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
		}
		// Write the state to the ledger
		err = stub.PutState(utils.ConstructRecordKey(), recordAsBytes)
		if err != nil {
			errMes := fmt.Sprintf("PutState error: %s", err)
			return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
		}

		return pb.Response{Status: 200, Message: "query successful", Payload: informationAsBytes}
	} else if args[0] == "doctor" {
		username := args[1]
		patientId := args[2]
		ntime := args[3]
		// Get the state from the ledger
		informationAsBytes, err := stub.GetState(patientId)
		if err != nil {
			errMes := fmt.Sprintf("GetState error: %s", err)
			return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
		} else if informationAsBytes == nil {
			errMes := "The patient does not exist"
			return pb.Response{Status: 400, Message: errMes, Payload: []byte(errMes)}
		}

		var patient lib.Patient
		err = json.Unmarshal(informationAsBytes, &patient)
		if err != nil {
			errMes := fmt.Sprintf("Unmarshal error: %s", err)
			return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
		}
		// Verify doctor's permissions
		if utils.IsContain(patient.AuthorizedDoctors, username) {
			// Record the reading operation
			record := lib.UsageRecord{ObjectType: lib.UsageRecordObj, Id: patient.Id, Operations: lib.Read, Roles: args[0], Username: username, Time: ntime}
			recordAsBytes, err := json.Marshal(record)
			if err != nil {
				errMes := fmt.Sprintf("Marshal error: %s", err)
				return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
			}
			// Write the state to the ledger
			err = stub.PutState(utils.ConstructRecordKey(), recordAsBytes)
			if err != nil {
				errMes := fmt.Sprintf("PutState error: %s", err)
				return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
			}

			return pb.Response{Status: 200, Message: "query successful", Payload: informationAsBytes}
		} else {
			errMes := fmt.Sprintf("Error! Permission denied")
			return pb.Response{Status: 400, Message: errMes, Payload: nil}
		}
	}

	errMes := "The role must be 'patient' or 'doctor'"
	return pb.Response{Status: 400, Message: errMes, Payload: []byte(errMes)}
}

// Authorized to the doctor
func (t *MedicalSystem) authorizeDoctor(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// check the number of parameters
	if len(args) != 2 {
		errMes := "Incorrect number of arguments. Expecting 2"
		return pb.Response{Status: 400, Message: errMes, Payload: []byte(errMes)}
	}
	username := args[0]
	doctorId := args[1]

	// Get the state from the ledger
	informationAsBytes, err := stub.GetState(username)
	if err != nil {
		errMes := fmt.Sprintf("GetState error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	} else if informationAsBytes == nil {
		errMes := "The patient does not exist"
		return pb.Response{Status: 400, Message: errMes, Payload: []byte(errMes)}
	}

	var patient lib.Patient
	err = json.Unmarshal(informationAsBytes, &patient)
	if err != nil {
		errMes := fmt.Sprintf("Unmarshal error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}

	// Add permissions
	patient.AuthorizedDoctors = append(patient.AuthorizedDoctors, doctorId)
	informationAsBytes, err = json.Marshal(patient)
	if err != nil {
		errMes := fmt.Sprintf("Marshal error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}
	// Write the state back to the ledger
	err = stub.PutState(username, informationAsBytes)
	if err != nil {
		errMes := fmt.Sprintf("PutState error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}
	return pb.Response{Status: 200, Message: "Authorize doctor successful", Payload: informationAsBytes}
}

// Enter the patient's medical data
func (t *MedicalSystem) enterData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// check the number of parameters
	if len(args) != 6 {
		errMes := "Incorrect number of arguments. Expecting 6"
		return pb.Response{Status: 400, Message: errMes, Payload: []byte(errMes)}
	}
	username := args[0]
	patientId := args[1]
	testResults := args[2]
	diagnosis := args[3]
	treatment := args[4]
	wtime := args[5]

	// Get the state from the ledger
	informationAsBytes, err := stub.GetState(patientId)
	if err != nil {
		errMes := fmt.Sprintf("GetState error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	} else if informationAsBytes == nil {
		errMes := "The patient does not exist"
		return pb.Response{Status: 400, Message: errMes, Payload: []byte(errMes)}
	}

	var patient lib.Patient
	err = json.Unmarshal(informationAsBytes, &patient)
	if err != nil {
		errMes := fmt.Sprintf("Unmarshal error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}
	// Verify doctor's permissions
	if utils.IsContain(patient.AuthorizedDoctors, username) {
		// append the case
		ncase := lib.Case{ObjectType: lib.CaseObj, Id: patient.Id, TestResults: testResults, Diagnosis: diagnosis, Treatment: treatment}
		patient.Cases = append(patient.Cases, ncase)
		informationAsBytes, err = json.Marshal(patient)
		if err != nil {
			errMes := fmt.Sprintf("Marshal error: %s", err)
			return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
		}
		// Write the state back to the ledger
		err = stub.PutState(patientId, informationAsBytes)
		if err != nil {
			errMes := fmt.Sprintf("PutState error: %s", err)
			return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
		}

		caseAsBytes, err := json.Marshal(ncase)
		if err != nil {
			errMes := fmt.Sprintf("Marshal error: %s", err)
			return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
		}
		// Write the state to the ledger
		err = stub.PutState(utils.ConstructCaseKey(), caseAsBytes)
		if err != nil {
			errMes := fmt.Sprintf("PutState error: %s", err)
			return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
		}

		// Record the reading operation
		record := lib.UsageRecord{ObjectType: lib.UsageRecordObj, Id: patient.Id, Operations: lib.Append, Roles: "doctor", Username: username, Time: wtime}
		recordAsBytes, err := json.Marshal(record)
		if err != nil {
			errMes := fmt.Sprintf("Marshal error: %s", err)
			return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
		}
		// Write the state to the ledger
		err = stub.PutState(utils.ConstructRecordKey(), recordAsBytes)
		if err != nil {
			errMes := fmt.Sprintf("PutState error: %s", err)
			return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
		}

		return pb.Response{Status: 200, Message: "Enter data successful", Payload: informationAsBytes}
	} else {
		errMes := fmt.Sprintf("Error! Permission denied")
		return pb.Response{Status: 400, Message: errMes, Payload: []byte(errMes)}
	}
}

// researcher or regulator query cases
func (t *MedicalSystem) queryCases(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// check the number of parameters
	//   0      1       2     3
	// role,username,keyword,ntime
	if len(args) != 4 {
		errMes := "Incorrect number of arguments. Expecting 4"
		return pb.Response{Status: 400, Message: errMes, Payload: []byte(errMes)}
	}
	if args[0] != "researcher" && args[0] != "regulator" {
		errMes := "The role must be 'researcher' or 'regulator'"
		return pb.Response{Status: 400, Message: errMes, Payload: []byte(errMes)}
	}
	roles := args[0]
	username := args[1]
	keyword := args[2]
	ntime := args[3]

	// Splicing query statement
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"caseObj\",\"diagnosis\":{\"$regex\": \".*%s.*\"}}}", keyword)
	// query data by statement
	result, err := utils.QueryByString(stub, queryString)
	if err != nil {
		errMes := fmt.Sprintf("GetQueryResult error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	} else if string(result) == "[]" {
		errMes := "No data that meets the conditions"
		return pb.Response{Status: 400, Message: errMes, Payload: []byte(errMes)}
	}

	var cases []lib.Case
	err = json.Unmarshal(result, &cases)
	if err != nil {
		errMes := fmt.Sprintf("Unmarshal error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	}

	// Record the reading operation
	for _, tcase := range cases {
		record := lib.UsageRecord{ObjectType: lib.UsageRecordObj, Id: tcase.Id, Operations: lib.Read, Roles: roles, Username: username, Time: ntime}
		recordAsBytes, err := json.Marshal(record)
		if err != nil {
			errMes := fmt.Sprintf("Marshal error: %s", err)
			return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
		}
		// Write the state to the ledger
		err = stub.PutState(utils.ConstructRecordKey(), recordAsBytes)
		if err != nil {
			errMes := fmt.Sprintf("PutState error: %s", err)
			return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
		}
	}
	return pb.Response{Status: 200, Message: "query cases successful", Payload: result}
}

// Show all ID usage records
func (t *MedicalSystem) showRecords(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// check the number of parameters
	if len(args) != 0 {
		errMes := "Incorrect number of arguments. Expecting 0"
		return pb.Response{Status: 400, Message: errMes, Payload: []byte(errMes)}
	}

	// Splicing query statement
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"usageRecordObj\"}}")
	// query data by statement
	result, err := utils.QueryByString(stub, queryString)
	if err != nil {
		errMes := fmt.Sprintf("GetQueryResult error: %s", err)
		return pb.Response{Status: 500, Message: errMes, Payload: []byte(errMes)}
	} else if string(result) == "[]" {
		errMes := "No data that meets the conditions"
		return pb.Response{Status: 400, Message: errMes, Payload: []byte(errMes)}
	}
	return pb.Response{Status: 200, Message: "show records successful", Payload: result}
}

func main() {
	err := shim.Start(new(MedicalSystem))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
