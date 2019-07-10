package repository

import (
	"fmt"

	"reflect"
	"testing"
)

func TestLoadEventTable(t *testing.T) {
	db := Connect()
	r := LoadEventTable["logurl"](2, 0, db)
	fmt.Println(reflect.TypeOf(r[0]))
	if reflect.TypeOf(r[0]).String() != "models.LogURL" {
		t.Errorf("TestLoadEventTable Error : %s.", reflect.TypeOf(r[0]))
	}
	r = LoadEventTable["guardianreward"](2, 0, db)
	fmt.Println(reflect.TypeOf(r[0]))
	if reflect.TypeOf(r[0]).String() != "models.GuardianReward" {
		t.Errorf("TestLoadEventTable Error : %s.", reflect.TypeOf(r[0]))
	}
}

func TestSearchTx(t *testing.T) {
	db := Connect()
	r := SearchRelatedEvents(200, "hash", "385f702381b6a19b0c5e636b", db)
	for _, e := range r {
		fmt.Println(reflect.TypeOf(e))
	}
	fmt.Println(len(r))
}

func TestSearchMethod(t *testing.T) {
	db := Connect()
	r := SearchRelatedEvents(200, "method", "signalGroupFormation", db)
	for _, e := range r {
		fmt.Println(reflect.TypeOf(e))
	}
	fmt.Println(len(r))
}

func TestSearchAddr(t *testing.T) {
	db := Connect()
	r := SearchRelatedEvents(200, "sender", "0xCdD9759439dF580FF183414C491F27E852Ac6240", db)
	for _, e := range r {
		fmt.Println(reflect.TypeOf(e))
	}
	fmt.Println(len(r))
}
