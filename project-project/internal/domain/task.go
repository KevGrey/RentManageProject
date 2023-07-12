package domain

import (
	"context"
	"pomfret.cn/project-project/internal/dao"
	"pomfret.cn/project-project/internal/repo"
	"pomfret.cn/project-project/pkg/model"
	"pomfret.cn/project_common/errs"
)

type TaskDomain struct {
	taskRepo repo.TaskRepo
}

func NewTaskDomain() *TaskDomain {
	return &TaskDomain{
		taskRepo: dao.NewTaskDao(),
	}
}

func (d *TaskDomain) FindProjectIdByTaskId(taskId int64) (int64, bool, *errs.BError) {
	//config.SendLog(kk.Info("Find", "TaskDomain.FindProjectIdByTaskId", kk.FieldMap{
	//	"taskId": taskId,
	//}))
	task, err := d.taskRepo.FindTaskById(context.Background(), taskId)
	if err != nil {
		//config.SendLog(kk.Error(err, "TaskDomain.FindProjectIdByTaskId.taskRepo.FindTaskById", kk.FieldMap{
		//	"taskId": taskId,
		//}))
		return 0, false, model.DBError
	}
	if task == nil {
		return 0, false, nil
	}
	return task.ProjectCode, true, nil
}
