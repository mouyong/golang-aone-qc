package models

import (
	"gorm.io/gorm"

	"aone-qc/internal/initialization"
)

type QcTasks struct {
	ID                int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Environment       string `gorm:"size:255;not null" json:"environment"`
	TenantID          string `gorm:"size:255;not null" json:"tenant_id"`
	Slug              string `gorm:"size:255;not null" json:"slug"`
	ProjectName       string `gorm:"size:255;not null" json:"project_name"`
	ExperimentBatchNo string `gorm:"size:255;not null" json:"experiment_batch_no"`
	AnalysesBatchNo   string `gorm:"size:255;null" json:"analysis_batch_number"`
	State             string `gorm:"size:50;null" json:"state"`
	Result            string `gorm:"size:50;null" json:"result"`
	ResultExplain     string `gorm:"type:text;null" json:"result_explain"`
	Remark            string `gorm:"type:text;null" json:"remark"`
}

func (q *QcTasks) Save() (id int64, err error) {
	var result QcTasks

	db := initialization.Db

	err = db.Where(&QcTasks{
		Environment:       q.Environment,
		TenantID:          q.TenantID,
		Slug:              q.Slug,
		ExperimentBatchNo: q.ExperimentBatchNo,
		AnalysesBatchNo:   q.AnalysesBatchNo,
	}).First(&result).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			q.State = "waiting"
			q.Result = ""
			q.ResultExplain = ""

			if createErr := db.Create(&q).Error; createErr != nil {
				return 0, createErr
			}

			return q.ID, nil
		} else {
			return 0, err
		}
	} else {
		updateData := map[string]interface{}{
			"State":         "waiting",
			"Result":        "",
			"ResultExplain": "",
		}

		if updateErr := db.Model(&result).Updates(updateData).Error; updateErr != nil {
			return 0, updateErr
		}
	}

	return result.ID, nil
}

func (q *QcTasks) GetQcTaskListWithPage(offset, limit int, qcTasks QcTasks) (data []QcTasks, count int64, err error) {
	var qcBatchList []QcTasks

	db := initialization.Db
	err = db.Where(qcTasks).Limit(limit).Offset(offset).Find(&qcBatchList).Offset(-1).Limit(-1).Count(&count).Error

	return qcBatchList, count, nil
}
