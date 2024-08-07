package go_mc_profile_api

import (
	"fmt"
	"testing"
)

const wantName = "jeb_"
const wantId = "853c80ef3c3749fdaa49938b674adae6"

func TestMcProfile_FetchByName_Ok(t *testing.T) {
	p := McProfile{}

	err := p.FetchByName(wantName)

	if err != nil {
		t.Error(err)
		return
	}

	if p.Name != wantName {
		t.Errorf("p.Name is %s, want %s", p.Name, wantName)
		return
	}

	if p.PlayerId != wantId {
		t.Errorf("p.PlayerId is %s, want %s", p.PlayerId, wantId)
		return
	}
}

func TestMcProfile_FetchByName_NotFound(t *testing.T) {

	p := McProfile{}

	name := "jeb_test_1234_xxx"

	err := p.FetchByName(name)

	if err == nil {
		t.Errorf("p.Name is %s, want error", name)
	} else {
		wantErr := fmt.Sprintf("Couldn't find any profile with name %s", name)

		thisError := err.Error()

		if thisError != wantErr {
			t.Errorf("this is %s, want %s", thisError, wantErr)
			return
		}
	}
}

func TestMcProfile_FetchById_Ok(t *testing.T) {

	p := McProfile{}

	err := p.FetchById(wantId)

	if err != nil {
		t.Error(err)
		return
	}

	if p.Name != wantName {
		t.Errorf("p.Name is %s, want %s", p.Name, wantName)
		return
	}

	if p.PlayerId != wantId {
		t.Errorf("p.PlayerId is %s, want %s", p.PlayerId, wantId)
		return
	}
}

func TestMcProfile_FetchById_InvalidId(t *testing.T) {

	p := McProfile{}

	testId := "123123"

	err := p.FetchById(testId)

	if err == nil {
		t.Errorf("want error")
		return
	} else {
		wantErr := fmt.Sprintf("Not a valid UUID: %s", testId)
		eStr := err.Error()
		if eStr != wantErr {
			t.Errorf("want %s, got %s", wantErr, eStr)
			return
		}

		return
	}
}
