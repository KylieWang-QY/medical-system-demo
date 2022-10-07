package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	bc "github.com/medical-system/application/blockchain"
	"github.com/medical-system/application/tlib"
	"github.com/medical-system/application/utils"
	"net/http"
)

func Doctor(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "doctor.html", "")
}

func DoctorQueryInformation(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "doctor_query_information.html", "")
}

func ToDoctorQueryInformation(ctx *gin.Context) {
	form := struct {
		Username  string `form:"username"`
		PatientId string `form:"patient_id"`
	}{}
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	args := [][]byte{[]byte("doctor"), []byte(form.Username), []byte(form.PatientId), []byte(utils.GetTime())}
	resp, err := bc.ChannelExecute("queryInformation", args)
	if err != nil {
		if err.Error() == "Transaction processing for endorser [localhost:7051]: Chaincode status Code: (400) UNKNOWN. Description: The patient does not exist" {
			ctx.String(http.StatusBadRequest, "The patient does not exist")
		} else if err.Error() == "Transaction processing for endorser [localhost:7051]: Chaincode status Code: (400) UNKNOWN. Description: Error! Permission denied" {
			ctx.String(http.StatusBadRequest, "Error! Permission denied")
		} else {
			ctx.String(http.StatusBadRequest, err.Error())
		}
		fmt.Println(err.Error())
		return
	}

	var patient tlib.Patient
	if err = json.Unmarshal(resp.Payload, &patient); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "view_patient.html", gin.H{
		"patient": patient,
		"cases":   patient.Cases,
	})
}

func DoctorEnterData(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "enter_data.html", "")

}

func ToDoctorEnterData(ctx *gin.Context) {
	form := struct {
		Username    string `form:"username"`
		PatientId   string `form:"patient_id"`
		TestResults string `form:"test_results"`
		Diagnosis   string `form:"diagnosis"`
		Treatment   string `form:"treatment"`
	}{}
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	args := [][]byte{[]byte(form.Username), []byte(form.PatientId), []byte(form.TestResults), []byte(form.Diagnosis), []byte(form.Treatment), []byte(utils.GetTime())}

	_, err := bc.ChannelExecute("enterData", args)
	if err != nil {
		if err.Error() == "Transaction processing for endorser [localhost:7051]: Chaincode status Code: (400) UNKNOWN. Description: The patient does not exist" {
			ctx.String(http.StatusBadRequest, "The patient does not exist")
		} else if err.Error() == "Transaction processing for endorser [localhost:7051]: Chaincode status Code: (400) UNKNOWN. Description: Error! Permission denied" {
			ctx.String(http.StatusBadRequest, "Error! Permission denied")
		} else {
			ctx.String(http.StatusBadRequest, err.Error())
		}
		fmt.Println(err.Error())
		return
	} else {
		ctx.String(http.StatusOK, "Submission complete")
	}
}
