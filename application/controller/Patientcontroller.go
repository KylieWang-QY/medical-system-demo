package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	bc "github.com/medical-system/application/blockchain"
	"github.com/medical-system/application/tlib"
	"github.com/medical-system/application/utils"
	"net/http"
)

func Patient(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "patient.html", "")
}

func PatientQueryInformation(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "patient_query_information.html", "")
}

func ToPatientQueryInformation(ctx *gin.Context) {
	form := struct {
		Username string `form:"username" binding:"required"`
	}{}
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	args := [][]byte{[]byte("patient"), []byte(form.Username), []byte(utils.GetTime())}
	resp, err := bc.ChannelExecute("queryInformation", args)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
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

func AuthorizeDoctor(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "authorize_doctor.html", "")
}

func ToAuthorizeDoctor(ctx *gin.Context) {
	form := struct {
		Username string `form:"username"`
		DoctorId string `form:"doctor_id"`
	}{}
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	args := [][]byte{[]byte(form.Username), []byte(form.DoctorId)}

	_, err := bc.ChannelExecute("authorizeDoctor", args)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	} else {
		ctx.String(http.StatusOK, "Authorization to the doctor successfully")
	}
}
