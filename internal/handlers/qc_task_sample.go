package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"aone-qc/internal/models"
)

type QcTaskSampleParamsDTO struct {
	Environment       string `json:"environment"`
	TenantId          string `json:"tenant_id"`
	Slug              string `json:"slug"`
	ProjectName       string `json:"project_name"`
	ExperimentBatchNo string `json:"experiment_batch_no"`
	AnalysesBatchNo   string `json:"analyses_batch_no"`
}

type QcTaskSampleResp struct {
	*models.QcTaskSample

	IsNeedAlert bool `json:"is_need_alert"`
}

func GetQcTaskSampleListWithPage(c *gin.Context) {
	resp := NewResp(c)

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	per_page, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	offset := (page * per_page) - per_page

	qcTaskSample := models.QcTaskSample{
		ExperimentBatchNo: c.DefaultQuery("experiment_batch_no", ""),
		AnalysesBatchNo:   c.DefaultQuery("analyses_batch_no", ""),
		SampleNo:          c.DefaultQuery("sample_no", ""),
		Result:            c.DefaultQuery("result", ""),
	}

	qcTaskSampleModel := models.QcTaskSample{}
	list, count, err := qcTaskSampleModel.GetQcTaskSampleListWithPage(offset, per_page, qcTaskSample)
	if err != nil {
		resp.fail("保存错误: " + err.Error())
		return
	}

	var data []*QcTaskSampleResp
	for _, v := range list {
		data = append(data, getQcTaskSampleInfo(&v))
	}

	metaStruct := &Meta{
		page,
		per_page,
		count,
	}

	resp.successWithData(data, metaStruct)
}

func getQcTaskSampleInfo(qcTaskSample *models.QcTaskSample) *QcTaskSampleResp {
	qcTaskSampleListDetail := &QcTaskSampleResp{
		QcTaskSample: qcTaskSample,
		IsNeedAlert:  false,
	}

	if qcTaskSampleListDetail.Result == "fail" {
		qcTaskSampleListDetail.IsNeedAlert = true
	}

	return qcTaskSampleListDetail
}
