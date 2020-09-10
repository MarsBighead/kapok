package model

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
)

func TestVcGet(t *testing.T) {
	db, err := sql.Open("postgres", "postgres://postgres:togerme@192.168.198.152:5433/hbu?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	r := NewVcResponse(db)
	vc := &Vc{
		EndPoint:     "10.138.0.218",
		Port:         443,
		Version:      "vCenter 6.5",
		UserName:     "admin",
		Password:     "togerme",
		FullName:     "vsphere65.vc.218",
		InstanceUUID: "528688ac-4b4e-ca2b-a90c-df2c9493a618",
	}
	err = r.Add([]*Vc{vc})
	if err != nil {
		t.Fatal(err)
	}
	err = r.Get()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", r.Data)
}
