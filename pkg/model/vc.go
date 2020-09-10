package model

import (
	"database/sql"
	"time"
)

//VcResponse for API vc
type VcResponse struct {
	DB   *sql.DB
	Data interface{}
}

//NewVcResponse for API vc
func NewVcResponse(db *sql.DB) *VcResponse {
	return &VcResponse{
		DB: db,
	}
}

//Vc  data structure for table Vc
//    active: vCenter is active in usage meter or not
type Vc struct {
	ID           int        `vsql:"column:id;type:SERIAL;primary_key"                       json:"-"`
	EndPoint     string     `vsql:"column:endPoint;type:varchar(256) NOT NULL"              json:"endPoint"`
	Port         int        `vsql:"column:port;type:integer NOT NULL"                       json:"port"`
	UserName     string     `vsql:"column:userName;type:varchar(256) NOT NULL"              json:"-"`
	Password     string     `vsql:"column:password;type:varchar(256) NOT NULL"              json:"-"`
	Sso          int        `vsql:"column:sso;type:integer NULL;default:0"                  json:"sso"`
	FullName     string     `vsql:"column:fullName;type:varchar(256) NOT NULL"              json:"fullName"`
	Version      string     `vsql:"column:version;type:varchar(256) NOT NULL"               json:"version"`
	InstanceUUID string     `vsql:"column:instanceUuid;type:varchar(256) NOT NULL"            json:"instanceUuid"`
	ChangeTime   *time.Time `vsql:"column:changeTime;type:timestamp with time zone NOT NULL"  json:"changeTime"`
}

//Get VCcenter  list
func (r *VcResponse) Get() error {
	stmt, err := r.DB.Prepare(`SELECT id,"endPoint", "port", "fullName", "version", "instanceUuid", "changeTime" FROM "Host"`)
	if err != nil {
		return err
	}
	rows, err := stmt.Query()
	if err != nil {
		return err
	}
	var vcs []*Vc
	for rows.Next() {
		vc := new(Vc)
		err = rows.Scan(&vc.ID, &vc.EndPoint, &vc.Port, &vc.FullName, &vc.Version, &vc.InstanceUUID, &vc.ChangeTime)
		if err != nil {
			return err

		}
		vcs = append(vcs, vc)
	}
	r.Data = vcs
	return nil
}

//Add vCenter list
func (r *VcResponse) Add(vcs []*Vc) error {
	if len(vcs) == 0 {
		return nil
	}
	now := time.Now()
	stmt, err := r.DB.Prepare(`Insert into "Vc" ("endPoint", "port", "userName", "password", "fullName", "version", "instanceUuid", "changeTime")
		VALUES ( $1, $2, $3, $4, $5, $6, $7, $8) RETURNING id
	`)
	if err != nil {
		return err
	}
	for _, vc := range vcs {
		err = stmt.QueryRow(vc.EndPoint, vc.Port, vc.UserName, vc.Password, vc.FullName, vc.Version, vc.InstanceUUID, now).Scan(&vc.ID)
		if err != nil {
			return err
		}

	}
	return nil
}

//Update vCenter data
func (r *VcResponse) Update(vc *Vc) error {
	return nil
}

//Delete vCenter data
func (r *VcResponse) Delete(vc *Vc) error {
	return nil
}
