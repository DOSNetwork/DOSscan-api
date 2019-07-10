package handler

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/DOSNetwork/DOSscan-api/models"
	"github.com/ethereum/go-ethereum/accounts/abi"
	//"github.com/jinzhu/gorm"
)

const (
	abiPath = "../../abi/DOSProxy.abi"
)

func TestReadABI(t *testing.T) {
	var proxyAbi abi.ABI
	jsonFile, err := os.Open(abiPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		t.Errorf("TestReadABI Error : %s.", err.Error())
	}
	abiJsonByte, _ := ioutil.ReadAll(jsonFile)
	proxyAbi, err = abi.JSON(strings.NewReader(string(abiJsonByte)))
	if err != nil {
		t.Errorf("TestReadABI Error : %s.", err.Error())
	}
	for key, _ := range proxyAbi.Methods {
		fmt.Println("Method: ", key)
	}
	for key, _ := range proxyAbi.Events {
		fmt.Println("Event: ", key)
	}
}

func TestLoadEventTable(t *testing.T) {
	db := models.Connect()
	r := models.LoadEventTable["logurl"](2, 0, db)
	fmt.Println(reflect.TypeOf(r[0]))
	if reflect.TypeOf(r[0]).String() != "models.LogURL" {
		t.Errorf("TestLoadEventTable Error : %s.", reflect.TypeOf(r[0]))
	}
	r = models.LoadEventTable["guardianreward"](2, 0, db)
	fmt.Println(reflect.TypeOf(r[0]))
	if reflect.TypeOf(r[0]).String() != "models.GuardianReward" {
		t.Errorf("TestLoadEventTable Error : %s.", reflect.TypeOf(r[0]))
	}
}
