package models

import (
	"fmt"

	"gorm.io/gorm"

	"aone-qc/internal/initialization"
)

type QcTaskSample struct {
	ID                int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	BatchID           int64  `gorm:"not null" json:"batch_id" comment:"批次ID"`
	ExperimentBatchNo string `gorm:"type:varchar(255);not null" json:"experiment_batch_number" comment:"实验批次编号"`
	AnalysesBatchNo   string `gorm:"type:varchar(255);not null" json:"analysis_batch_number" comment:"分析批次编号"`
	SampleNo          string `gorm:"type:varchar(255);not null" json:"sample_no" comment:"样本编号"`
	Result            string `gorm:"type:varchar(64);not null" json:"result" comment:"质控结果"`
	ResultExplain     string `gorm:"type:text" json:"result_explain" comment:"质控说明"`
	Remark            string `gorm:"type:text" json:"remark" comment:"备注"`
}

func (q *QcTaskSample) Save() (id int64, err error) {
	var result QcTaskSample

	db := initialization.Db

	err = db.Where(&QcTaskSample{
		ExperimentBatchNo: q.ExperimentBatchNo,
		AnalysesBatchNo:   q.AnalysesBatchNo,
		SampleNo:          q.SampleNo,
	}).First(&result).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
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
			"Result":        "",
			"ResultExplain": "",
		}

		if updateErr := db.Model(&result).Updates(updateData).Error; updateErr != nil {
			return 0, updateErr
		}
	}

	return result.ID, nil
}

func (q *QcTaskSample) GetQcTaskSampleListWithPage(offset, limit int, qcTaskSample QcTaskSample) (data []QcTaskSample, count int64, err error) {
	var qcTaskSampleList []QcTaskSample

	db := initialization.Db
	err = db.Where(qcTaskSample).Limit(limit).Offset(offset).Find(&qcTaskSampleList).Offset(-1).Limit(-1).Count(&count).Error
	fmt.Println(limit, offset, qcTaskSampleList)

	return qcTaskSampleList, count, nil
}
