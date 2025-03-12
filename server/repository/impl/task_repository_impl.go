package impl

import (
	"fmt"

	"github.com/Kanatanagano/gorm-golang-migrate-todo/entity"
	"github.com/Kanatanagano/gorm-golang-migrate-todo/repository"
	"gorm.io/gorm"
)

type task_repository_impl struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) repository.Task_repository {
	return &task_repository_impl{db: db}
}

func (r *task_repository_impl) FindAll() ([]entity.Task, error) {
	var tasks []entity.Task
	if err := r.db.Preload("Labels").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *task_repository_impl) FindById(id int) (entity.Task, error) {
	var task entity.Task
	if err := r.db.Preload("Labels").First(&task, id).Error; err != nil {
		return entity.Task{}, err
	}
	return task, nil
}

func (r *task_repository_impl) Create(task entity.Task) (string, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// タスクを保存
		if err := tx.Create(&task).Error; err != nil {
			return err
		}

		// ラベルIDが指定されている場合
		if len(task.LabelIDs) > 0 {
			// 指定されたIDのラベルを取得
			var labels []*entity.Label
			if err := tx.Where("id IN ?", task.LabelIDs).Find(&labels).Error; err != nil {
				return err
			}

			// すべてのラベルが存在することを確認
			if len(labels) != len(task.LabelIDs) {
				return fmt.Errorf("some labels do not exist")
			}

			// 中間テーブルに関連を追加
			for _, labelID := range task.LabelIDs {
				if err := tx.Exec("INSERT INTO label_tasks (task_id, label_id) VALUES (?, ?)", task.ID, labelID).Error; err != nil {
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return "task created successfully", nil
}
