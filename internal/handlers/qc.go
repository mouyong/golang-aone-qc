package handlers

import (
	"github.com/gin-gonic/gin"
	"aone-qc/internal/models"
)

type QcTaskParamsDTO struct {
	Environment       string `json:"environment"`
	TenantId          string `json:"tenant_id"`
	Slug              string `json:"slug"`
	ProjectName       string `json:"project_name"`
	ExperimentBatchNo string `json:"experiment_batch_no"`
	AnalysesBatchNo   string `json:"analyses_batch_no"`
}

func CreateOrRetryQcTask(c *gin.Context) {
	resp := NewResp(c)

	var qcTaskParamsDTO QcTaskParamsDTO
	if err := c.ShouldBind(&qcTaskParamsDTO); err != nil {
		resp.fail("请求错误: " + err.Error())
		return
	}

	qcTask := models.QcTasks{
		Environment:       qcTaskParamsDTO.Environment,
		TenantID:          qcTaskParamsDTO.TenantId,
		Slug:              qcTaskParamsDTO.Slug,
		ProjectName:       qcTaskParamsDTO.ProjectName,
		ExperimentBatchNo: qcTaskParamsDTO.ExperimentBatchNo,
		AnalysesBatchNo:   qcTaskParamsDTO.AnalysesBatchNo,
	}

	// 以下是一个示意性的数据库保存操作，具体实现可能会有所不同
	if err := qcTask.Save(); err != nil {
		resp.fail("保存错误: " + err.Error())
		return
	}

	resp.success()
}
