CREATE TABLE IF NOT EXISTS qc_task_samples (
    id bigint AUTO_INCREMENT PRIMARY KEY COMMENT 'id',
    batch_id bigint NOT NULL COMMENT '批次ID',
    experiment_batch_no VARCHAR(255) NOT NULL COMMENT '实验批次编号',
    analyses_batch_no VARCHAR(255) NOT NULL COMMENT '分析批次编号',
    sample_no VARCHAR(255) NOT NULL COMMENT '样本编号',
    result VARCHAR(64) NOT NULL COMMENT '质控结果',
    result_explain TEXT COMMENT '质控说明',
    remark TEXT COMMENT '备注'
) COMMENT='质控任务样本表';
