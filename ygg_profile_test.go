package go_mc_profile_api

import (
	"testing"
)

const yggTest = "https://littleskin.cn/api/yggdrasil/"

const wantIdLtSkin = "0298846c1a18483db5df1c3eb71becfd"
const wantNameLtSkin = "Paper99"

func TestYggProfile_FetchByName_Ok(t *testing.T) {
	pro := NewYggProfile(yggTest)

	err := pro.FetchByName(wantNameLtSkin)

	if err != nil {
		t.Error(err)
		return
	}

	if pro.PlayerId != wantIdLtSkin {
		t.Errorf("want %s, got %s", wantIdLtSkin, pro.PlayerId)
		return
	}

	if pro.Name != wantNameLtSkin {
		t.Errorf("want %s, got %s", wantNameLtSkin, pro.Name)
		return
	}
}

func TestYggProfile_FetchById_Ok(t *testing.T) {
	pro := NewYggProfile(yggTest)

	err := pro.FetchById(wantIdLtSkin)

	if err != nil {
		t.Error(err)
		return
	}

	if pro.PlayerId != wantIdLtSkin {
		t.Errorf("want %s, got %s", wantIdLtSkin, pro.PlayerId)
		return
	}

	if pro.Name != wantNameLtSkin {
		t.Errorf("want %s, got %s", wantNameLtSkin, pro.Name)
		return
	}
}
