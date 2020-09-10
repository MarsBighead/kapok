package api

import (
	"database/sql"
	"fmt"
	"kapok/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type vcRequest struct {
	Response model.VcResponse
}

func NewVcRequest(db *sql.DB) *vcRequest {
	resp := model.VcResponse{
		DB: db,
	}
	return &vcRequest{
		Response: resp,
	}
}
func (r *vcRequest) Query(c *gin.Context) {

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
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": data,
	})
}

//Add Vc data to database
func (r *vcRequest) Add(c *gin.Context) {
	vc := &model.Vc{
		EndPoint:     "10.138.0.218",
		Port:         443,
		Version:      "vCenter 6.5",
		UserName:     "admin",
		Password:     "togerme",
		FullName:     "vsphere65.vc.218",
		InstanceUUID: "528688ac-4b4e-ca2b-a90c-df2c9493a618",
	}
	err := r.Response.Add([]*model.Vc{vc})
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
		"data":    vc.ID,
	})
}

func (r *vcRequest) Del(c *gin.Context) {
	vc := new(model.Vc)
	id, err := strconv.Atoi(c.Param("id"))
	vc.ID = id
	err = r.Response.Delete(vc)
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
func (r *vcRequest) Update(c *gin.Context) {
	vc := new(model.Vc)
	id, err := strconv.Atoi(c.Param("id"))
	vc.ID = id
	vc.UserName = c.Request.FormValue("username")
	vc.Password = c.Request.FormValue("password")
	err = r.Response.Update(vc)
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
