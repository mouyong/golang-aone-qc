package models

import (
	"aone-qc/internal/initialization"
)

type QcTasks struct {
	ID                 int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Environment        string `gorm:"size:255;not null" json:"environment"`
	TenantID           string `gorm:"size:255;not null" json:"tenantId"`
	Slug               string `gorm:"size:255;not null" json:"slug"`
	ProjectName        string `gorm:"size:255;not null" json:"projectName"`
	ExperimentBatchNo  string `gorm:"size:255;not null" json:"experimentBatchNo"`
	AnalysesBatchNo    string `gorm:"size:255;null" json:"analysesBatchNo"`
	State              string `gorm:"size:50;null" json:"state"`
	Result             string `gorm:"size:50;null" json:"result"`
	ResultExplain      string `gorm:"type:text;null" json:"resultExplain"`
	Remark             string `gorm:"type:text;null" json:"remark"`
}

func (q *QcTasks) Save() error {
	db := initialization.Db

	if err := db.Create(q).Error; err != nil {
		return err
	}
	return nil
}
