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

func Regulator(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "researcher.html", "")
}

func RegulatorQuery(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "researcher_query_cases.html", "")
}

func ToRegulatorQuery(ctx *gin.Context) {
	form := struct {
		Username string `form:"username"`
		KeyWord  string `form:"key_word"`
	}{}
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	args := [][]byte{[]byte("regulator"), []byte(form.Username), []byte(form.KeyWord), []byte(utils.GetTime())}
	resp, err := bc.ChannelExecute("queryCases", args)
	if err != nil {
		if err.Error() == "Transaction processing for endorser [localhost:7051]: Chaincode status Code: (400) UNKNOWN. Description: No data that meets the conditions" {
			ctx.String(http.StatusBadRequest, "No data that meets the conditions")
		} else {
			ctx.String(http.StatusBadRequest, err.Error())
			fmt.Println(err.Error())
			return
		}
	}

	var cases []tlib.Case
	if err = json.Unmarshal(resp.Payload, &cases); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "researcher_query_cases.html", gin.H{
		"cases": cases,
	})
}
