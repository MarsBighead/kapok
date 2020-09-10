package model

import (
	"database/sql"
	"time"

	"github.com/vmware/govmomi/vim25/types"
)

//HostResponse for API host
type HostResponse struct {
	DB   *sql.DB
	Data interface{}
}

//NewHostResponse for API host
func NewHostResponse(db *sql.DB) *HostResponse {
	return &HostResponse{
		DB: db,
	}
}

//Host  data structure for table Host
type Host struct {
	ID             int64                      `vsql:"column:id;type:SERIAL;primary_key"`
	VcID           int                        `vsql:"column:vcId;type:INT NOT NULL"`
	Moref          string                     `vsql:"column:moref;type:varchar(256) NOT NULL"`
	Hz             int64                      `vsql:"column:hz;type:BIGINT"`
	NumCPUCores    int16                      `vsql:"column:numCpuCores;type:INT"`
	NumCPUPackages int16                      `vsql:"column:numCpuPackages;type:INT"`
	NumCPUThreads  int16                      `vsql:"column:numCpuThreads;type:INT"`
	MemorySize     int64                      `vsql:"column:memorySize;type:BIGINT"`
	InstanceUUID   string                     `vsql:"column:uuid;type:varchar(256)"`
	Uuid           string                     `vsql:"column:uuid;type:varchar(256)"`
	FullName       string                     `vsql:"column:fullName;type:varchar(256)"`
	Name           string                     `vsql:"column:name;type:varchar(256)"`
	Version        string                     `vsql:"column:version;type:varchar"`
	BootTime       *time.Time                 `vsql:"column:bootTime;type:timestamp with time zone"`
	PowerState     types.HostSystemPowerState `vsql:"column:powerState;type:varchar(256)"`
	ChangeTime     *time.Time                 `vsql:"column:changeTime;type:timestamp with time zone NOT NULL"`
}

//Get VCcenter  list
func (r *HostResponse) Get() error {
	stmt, err := r.DB.Prepare(`SELECT id, "vcId", "moref", "uuid","fullName", "instanceUuid", "powerState", "changeTime" FROM "Host"`)
	if err != nil {
		return err
	}
	rows, err := stmt.Query()
	if err != nil {
		return err
	}
	var hss []*Host
	for rows.Next() {
		h := new(Host)
		err = rows.Scan(&h.ID, &h.VcID, &h.Moref, &h.Uuid, &h.FullName, &h.InstanceUUID, &h.PowerState, &h.ChangeTime)
		if err != nil {
			return err

		}
		hss = append(hss, h)
	}
	r.Data = hss
	return nil
}

//Add host list
func (r *HostResponse) Add(hss []*Host) error {
	if len(hss) == 0 {
		return nil
	}
	now := time.Now()
	stmt, err := r.DB.Prepare(`Insert into "Host" ("vcId", "moref", "uuid", "fullName", "instanceUuid", "powerState", "changeTime")
		VALUES ( $1, $2, $3, $4, $5, $6, $7) RETURNING id
	`)
	if err != nil {
		return err
	}
	for _, h := range hss {
		err = stmt.QueryRow(h.VcID, h.Moref, h.Uuid, h.FullName, h.InstanceUUID, h.PowerState, now).Scan(&h.ID)
		if err != nil {
			return err
		}

	}
	return nil
}

//Update vCenter data
func (r *HostResponse) Update(host *Vc) error {
	return nil
}

//Delete vCenter data
func (r *HostResponse) Delete(host *Vc) error {
	return nil
}
