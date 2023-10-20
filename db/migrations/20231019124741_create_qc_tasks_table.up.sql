CREATE TABLE IF NOT EXISTS qc_tasks (
    id bigint AUTO_INCREMENT PRIMARY KEY COMMENT 'id',
    environment VARCHAR(255) NOT NULL COMMENT '系统环境',
    tenant_id VARCHAR(255) NOT NULL COMMENT '租户ID',
    slug VARCHAR(255) NOT NULL COMMENT '项目',
    project_name VARCHAR(255) NOT NULL COMMENT '项目名',
    experiment_batch_no VARCHAR(255) NOT NULL COMMENT '实验批次编号',
    analyses_batch_no VARCHAR(255) NOT NULL COMMENT '分析批次编号',
    state VARCHAR(50) NOT NULL COMMENT '质控进度',
    result VARCHAR(50) NOT NULL COMMENT '质控结果',
    result_explain TEXT COMMENT '质控说明',
    remark TEXT COMMENT '备注'
) COMMENT='质控任务表';