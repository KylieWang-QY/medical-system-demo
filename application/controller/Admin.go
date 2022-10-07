package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	bc "github.com/medical-system/application/blockchain"
	"github.com/medical-system/application/tlib"
	"net/http"
)

func Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home.html", "")
}

func ShowRecords(ctx *gin.Context) {
	var args [][]byte
	resp, err := bc.ChannelQuery("showRecords", args)
	if err != nil {
		fmt.Println(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	var records []tlib.UsageRecord
	if err = json.Unmarshal(resp.Payload, &records); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.HTML(http.StatusOK, "show_records.html", gin.H{
		"records": records,
	})
}
