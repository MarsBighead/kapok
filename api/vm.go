package api

import (
	"database/sql"
	"fmt"
	"kapok/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type vmRequest struct {
	Response model.VmResponse
}

func NewVmRequest(db *sql.DB) *vmRequest {
	resp := model.VmResponse{
		DB: db,
	}
	return &vmRequest{
		Response: resp,
	}
}
func (r *vmRequest) Query(c *gin.Context) {

	err := r.Response.Get()
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}
	data := r.Response.Data
	c.IndentedJSON(http.StatusOK, gin.H{
		"code": 1,
		"data": data,
	})
}

//Add Vc data to database
func (r *vmRequest) Add(c *gin.Context) {
	ipAddress := "10.138.1.108"
	h := &model.Vm{
		VcID:          2,
		HostID:        1,
		Moref:         "vm-21",
		PowerState:    "poweredOn",
		GuestFullName: "UsageMeter361.vm.129",
		Uuid:          "423a511a-9db6-fccb-6877-040db7f9572a",
		InstanceUUID:  "528688ac-4b4e-ca2b-a90c-df2c9493a618",
		IPAddress:     &ipAddress,
	}

	err := r.Response.Add([]*model.Vm{h})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "添加成功",
		"data":    h.ID,
	})
}

func (r *vmRequest) Del(c *gin.Context) {
	vm := new(model.Vm)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	vm.ID = id
	err = r.Response.Delete(vm)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "删除失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "删除成功",
	})
}

//Put update data with method put
func (r *vmRequest) Update(c *gin.Context) {
	vm := new(model.Vm)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	vm.ID = id
	err = r.Response.Update(vm)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "Success",
	})
}
