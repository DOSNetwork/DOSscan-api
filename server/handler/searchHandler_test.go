package handler

import (
	"fmt"

	"reflect"
	"testing"

	"github.com/DOSNetwork/DOSscan-api/server/repository"
	//"github.com/DOSNetwork/DOSscan-api/models"
	//"github.com/jinzhu/gorm"
)

const (
	abiPath = "../../abi/DOSProxy.abi"
)

func TestGetEventsAndMethodFromABI(t *testing.T) {
	events, methods, err := getEventsAndMethodFromABI(abiPath)
	if err != nil {
		t.Errorf("TestGetEventsAndMethodFromABI Error : %s.", err.Error())
	}
	fmt.Println(events)
	fmt.Println(methods)
}

func TestSearchEvents(t *testing.T) {
	db := repository.Connect("postgres", "postgres", "postgres")
	repo := repository.NewDBEventsRepository(db)
	repo.SetTxRelatedEvents(events)
	//r, err := searchEvents(repo, "0xa49a38aa1c69090e9d4927535b3be2dfe027eb47190dd7809511e6e26a317934", 100, 0)
	eventMap, _, err := getEventsAndMethodFromABI(abiPath)
	if err != nil {
		t.Errorf("TestGetEventsAndMethodFromABI Error : %s.", err.Error())
	}

	r, err := searchEventsByEventName(repo, eventMap, events, "log", 100, 0)
	if err != nil {
		t.Errorf("TestSearchEvents Error : %s.", err.Error())
	}
	for _, event := range r {
		fmt.Println("event ", reflect.TypeOf(event).String())
	}
}
