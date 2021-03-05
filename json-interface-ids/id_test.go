package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIdStruct(t *testing.T) {
	var fname = "../../data/ids.json"
    var ids []id

	ids = make([]id,0)

	ids,err := loadIDs(fname)

	if err != nil {
		t.Error("load failed",err)
	}
	if ids == nil {
		t.Error("load failed (2)")
	}

	// use reflect to print the struct values
	v := reflect.ValueOf(ids[0])
	u := v.Type()

	for i:=0;i<v.NumField();i++ {
		fmt.Printf("%v : %v\n",u.Field(i).Name,v.Field(i).Interface())
	}
}