package handlers

import "github.com/gin-gonic/gin"

type Resp struct {
	Context *gin.Context
}

type Meta struct {
    Page 		int		`json:"page"`
    PerPage     int 	`json:"per_page"`
    Total  		int64 	`json:"total"`
}

func NewResp(c *gin.Context) *Resp {
	return &Resp{
		Context: c,
	}
}

func (resp *Resp) success() {
	resp.Context.JSON(200, gin.H{
		"err_code": 200,
		"err_msg": "success",
		"data": nil,
	})
}

func (resp *Resp) successWithData(data interface{}, meta *Meta) {
	response := gin.H{
        "err_code": 200,
        "err_msg":  "success",
        "data":     data,
    }

    if meta != nil {
        response["meta"] = meta
    }

    resp.Context.JSON(200, response)
}

func (resp *Resp) fail(err_msg string) {
	resp.Context.JSON(200, gin.H{
		"err_code": 400,
		"err_msg": err_msg,
		"data": nil,
	})
}

func (resp *Resp) failWithErrCode(err_msg string, err_code int) {
	resp.Context.JSON(200, gin.H{
		"err_code": err_code,
		"err_msg": err_msg,
		"data": nil,
	})
}