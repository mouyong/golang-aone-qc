package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"aone-qc/internal/models"
)

type QcTaskDTO struct {
	Environment       string `json:"environment"`
	TenantId          string `json:"tenant_id"`
	Slug              string `json:"slug"`
	ProjectName       string `json:"project_name"`
	ExperimentBatchNo string `json:"experiment_batch_number"`
	AnalysesBatchNo   string `json:"analysis_batch_number"`
}

type QcTaskResp struct {
	*models.QcTasks

	IsNeedAlert bool `json:"is_need_alert"`
}

func CreateOrRetryQcTask(c *gin.Context) {
	resp := NewResp(c)

	var qcTaskDTO QcTaskDTO
	if err := c.ShouldBind(&qcTaskDTO); err != nil {
		resp.fail("请求错误: " + err.Error())
		return
	}

	qcTaskModel := models.QcTasks{
		Environment:       qcTaskDTO.Environment,
		TenantID:          qcTaskDTO.TenantId,
		Slug:              qcTaskDTO.Slug,
		ProjectName:       qcTaskDTO.ProjectName,
		ExperimentBatchNo: qcTaskDTO.ExperimentBatchNo,
		AnalysesBatchNo:   qcTaskDTO.AnalysesBatchNo,
	}

	// 以下是一个示意性的数据库保存操作，具体实现可能会有所不同
	_, err := qcTaskModel.Save()
	if err != nil {
		resp.fail("保存错误: " + err.Error())
		return
	}

	resp.success()
}

func GetQcTaskListWithPage(c *gin.Context) {
	resp := NewResp(c)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per_page, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	offset := (page * per_page) - per_page

	qcTasks := models.QcTasks{
		Environment:       c.DefaultQuery("environment", ""),
		TenantID:          c.DefaultQuery("tenant_id", ""),
		Slug:              c.DefaultQuery("slug", ""),
		ExperimentBatchNo: c.DefaultQuery("experiment_batch_no", ""),
		AnalysesBatchNo:   c.DefaultQuery("analyses_batch_no", ""),
		State:             c.DefaultQuery("state", ""),
	}

	// 以下是一个示意性的数据库保存操作，具体实现可能会有所不同
	qcTaskModel := models.QcTasks{}
	list, count, err := qcTaskModel.GetQcTaskListWithPage(offset, per_page, qcTasks)
	if err != nil {
		resp.fail("保存错误: " + err.Error())
		return
	}

	var data []*QcTaskResp
	for _, v := range list {
		data = append(data, getQcTaskInfo(&v))
	}

	metaStruct := &Meta{
		page,
		per_page,
		count,
	}

	resp.successWithData(data, metaStruct)
}

func getQcTaskInfo(q *models.QcTasks) *QcTaskResp {
	qcDataDetail := &QcTaskResp{
		QcTasks:     q,
		IsNeedAlert: false,
	}

	if qcDataDetail.State == "batch_fail" {
		qcDataDetail.IsNeedAlert = true
	}

	if qcDataDetail.Result == "part_pass" {
		qcDataDetail.IsNeedAlert = true
	}

	if qcDataDetail.Result == "batch_fail" {
		qcDataDetail.IsNeedAlert = true
	}

	return qcDataDetail
}
