package main

// Go build first
// If you want to execute the main.go file on the terminal, you need to remove. / application
// In the configpath of sdk.go, follow the static configuration file of this file

import (
	"github.com/gin-gonic/gin"
	bc "github.com/medical-system/application/blockchain"
	ct "github.com/medical-system/application/controller"
)

func setrouter() *gin.Engine {
	// Initialize SDK
	bc.Init()
	router := gin.Default()
	// Load static resource file
	router.LoadHTMLGlob("./application/view/*")

	// system
	router.GET("/", ct.Home)
	router.GET("/showRecords", ct.ShowRecords)
	// patient
	router.GET("/patient", ct.Patient)
	router.GET("/patientQueryInformation", ct.PatientQueryInformation)
	router.POST("/patientQueryInformation", ct.ToPatientQueryInformation)
	router.GET("/authorizeDoctor", ct.AuthorizeDoctor)
	router.POST("/authorizeDoctor", ct.ToAuthorizeDoctor)

	// doctor
	router.GET("/doctor", ct.Doctor)
	router.GET("/doctorQueryInformation", ct.DoctorQueryInformation)
	router.POST("/doctorQueryInformation", ct.ToDoctorQueryInformation)
	router.GET("/doctorEnterData", ct.DoctorEnterData)
	router.POST("/doctorEnterData", ct.ToDoctorEnterData)

	// researcher
	router.GET("/researcher", ct.Researcher)
	router.GET("/researcherQuery", ct.ResearcherQuery)
	router.POST("/researcherQuery", ct.ToResearcherQuery)

	// regulator
	router.GET("/regulator", ct.Regulator)
	router.GET("/regulatorQuery", ct.RegulatorQuery)
	router.POST("/regulatorQuery", ct.ToRegulatorQuery)

	return router
}
func main() {
	router := setrouter()
	router.Run(":8080")
}
