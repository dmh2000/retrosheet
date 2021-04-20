package jsontypes

import (
	"os"
	"testing"
)

func TestPersonnel(t *testing.T) {
	var fname = os.Getenv("RETROSHEET_DATA") + "/personnel.json"
    var ids []Person

	ids = make([]Person,0)

	ids,err := ReadPersonnel(fname)

	if err != nil {
		t.Error("load failed",err)
	}
	if ids == nil {
		t.Error("load failed (2)")
	}

	// // uncomment this to print and verify the fields of this instance are correct
	// // use reflect to print the struct values
	// v := reflect.ValueOf(ids[0])
	// u := v.Type()

	// // print the person data
	// for i:=0;i<v.NumField();i++ {
	// 	fmt.Printf("%v : %v\n",u.Field(i).Name,v.Field(i).Interface())
	// }
}