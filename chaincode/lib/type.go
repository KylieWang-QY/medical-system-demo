package lib

const (
	PatientObj     = "patientObj"
	CaseObj        = "caseObj"
	UsageRecordObj = "usageRecordObj"

	Add    = "add"
	Read   = "read"
	Append = "append"
)

var NumberOfCases = 0
var NumberOfRecord = 0

// Patient The patient's personal information and medical data
type Patient struct {
	ObjectType        string   `json:"docType"`            // patientObj
	Username          string   `json:"username"`           // username
	Name              string   `json:"name"`               // full name
	Address           string   `json:"address"`            // address
	Telephone         string   `json:"telephone"`          // telephone number
	Id                string   `json:"id"`                 // UUid for the patient's medical data
	Cases             []Case   `json:"cases"`              // medical data
	AuthorizedDoctors []string `json:"authorized_doctors"` // username of the authorized doctor
}

// Case The medical data
type Case struct {
	ObjectType  string `json:"docType"`      // caseObj
	Id          string `json:"id"`           // UUid for the patient's medical data
	TestResults string `json:"test_results"` // test results
	Diagnosis   string `json:"diagnosis"`    // doctor's diagnosis
	Treatment   string `json:"treatment"`    // treatment plan
}

// UsageRecord Usage record of id
type UsageRecord struct {
	ObjectType string `json:"docType"`    // usageRecordObj
	Id         string `json:"id"`         // UUid for the patient's medical data
	Operations string `json:"operations"` // operation type, like: add, read, append
	Roles      string `json:"roles"`      // role of operator
	Username   string `json:"username"`   // username of the operator
	Time       string `json:"time"`       // time of operation
}

