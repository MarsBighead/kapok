package model

import (
	"database/sql"
	"time"

	"github.com/vmware/govmomi/vim25/types"
)

//VmResponse for API vm
type VmResponse struct {
	DB   *sql.DB
	Data interface{}
}

//NewVmResponse for API vm
func NewVmResponse(db *sql.DB) *VmResponse {
	return &VmResponse{
		DB: db,
	}
}

//Vm with usage meter installed
type Vm struct {
	ID                int64                          `vsql:"column:id;type:SERIAL;primary_key"             json:"id"`
	VcID              int                            `vsql:"column:vcId;type:INT NOT NULL"                 json:"vcId"`
	HostID            int                            `vsql:"column:hostId;type:INT NOT NULL"               json:"hostId"`
	HostName          string                         `vsql:"column:hostName;type:varchar(256)"`
	Moref             string                         `vsql:"column:moref;type:varchar(256) NOT NULL"       json:"moref"`
	CPUReservation    int32                          `vsql:"column:cpuReservation;type:int"`
	GuestFullName     string                         `vsql:"column:guestFullName;type:varchar(256)"`
	InstanceUUID      string                         `vsql:"column:instanceUuid;type:varchar(256)"         json:"instanceUuid"`
	MemoryReservation int32                          `vsql:"column:memoryReservation;type:INT"`
	MemorySizeMB      int32                          `vsql:"column:memorySizeMB;type:INT"`
	Name              string                         `vsql:"column:name;type:varchar(256)"`
	NumCPU            int32                          `vsql:"column:numCpu;type:smallint"`
	Uuid              string                         `vsql:"column:uuid;type:varchar(256)"                 json:"uuid"`
	BootTime          *time.Time                     `vsql:"column:bootTime;type:timestamp with time zone"`
	PowerState        types.VirtualMachinePowerState `vsql:"column:powerState;type:varchar(256)"`
	IPAddress         *string                        `vsql:"column:ipAddress;type:varchar(256)"`
	ChangeTime        *time.Time                     `vsql:"column:changeTime;type:timestamp with time zone NOT NULL"`
}

//Get VCcenter  list
func (r *VmResponse) Get() error {
	stmt, err := r.DB.Prepare(`SELECT id, "vcId", "hostId","moref", "uuid","guestFullName", "ipAddress","instanceUuid", "powerState", "changeTime" FROM "Vm"`)
	if err != nil {
		return err
	}
	rows, err := stmt.Query()
	if err != nil {
		return err
	}
	var vms []*Vm
	for rows.Next() {
		vm := new(Vm)
		err = rows.Scan(&vm.ID, &vm.VcID, &vm.HostID, &vm.Moref, &vm.Uuid, &vm.GuestFullName, &vm.IPAddress, &vm.InstanceUUID, &vm.PowerState, &vm.ChangeTime)
		if err != nil {
			return err

		}
		vms = append(vms, vm)
	}
	r.Data = vms
	return nil
}

//Add vm list
func (r *VmResponse) Add(vms []*Vm) error {
	if len(vms) == 0 {
		return nil
	}
	now := time.Now()
	stmt, err := r.DB.Prepare(`Insert into "Vm" ("vcId", "hostId", "moref", "uuid", "guestFullName", "ipAddress", "instanceUuid", "powerState", "changeTime")
		VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id
	`)
	if err != nil {
		return err
	}
	for _, vm := range vms {
		err = stmt.QueryRow(vm.VcID, vm.HostID, vm.Moref, vm.Uuid, vm.GuestFullName, vm.IPAddress, vm.InstanceUUID, vm.PowerState, now).Scan(&vm.ID)
		if err != nil {
			return err
		}

	}
	return nil
}

//Update vCenter data
func (r *VmResponse) Update(vm *Vm) error {
	return nil
}

//Delete vCenter data
func (r *VmResponse) Delete(vm *Vm) error {
	return nil
}
