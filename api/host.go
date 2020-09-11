package api

import (
	"database/sql"
	"fmt"
	"kapok/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type hostRequest struct {
	Response model.HostResponse
}

func NewHostRequest(db *sql.DB) *hostRequest {
	resp := model.HostResponse{
		DB: db,
	}
	return &hostRequest{
		Response: resp,
	}
}
func (r *hostRequest) Query(c *gin.Context) {

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
func (r *hostRequest) Add(c *gin.Context) {
	h := &model.Host{
		VcID:         2,
		Moref:        "host-21",
		PowerState:   "poweredOn",
		Version:      "vCenter 6.5",
		FullName:     "vsphere65.host.218",
		Uuid:         "423a511a-9db6-fccb-6877-040db7f9572a",
		InstanceUUID: "528688ac-4b4e-ca2b-a90c-df2c9493a618",
	}

	err := r.Response.Add([]*model.Host{h})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "添加成功",
		"data":    h.ID,
	})
}

func (r *hostRequest) Del(c *gin.Context) {
	host := new(model.Vc)
	id, err := strconv.Atoi(c.Param("id"))
	host.ID = id
	err = r.Response.Delete(host)
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
func (r *hostRequest) Update(c *gin.Context) {
	host := new(model.Vc)
	id, err := strconv.Atoi(c.Param("id"))
	host.ID = id
	host.UserName = c.Request.FormValue("username")
	host.Password = c.Request.FormValue("password")
	err = r.Response.Update(host)
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
