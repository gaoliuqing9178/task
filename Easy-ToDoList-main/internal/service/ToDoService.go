package service

import (
	"todolist/internal/dto"
	"todolist/internal/model"
	"todolist/internal/repository"
	"todolist/internal/vo"
)

type ToDoService struct {
	toDoRepository *repository.ToDoRepository
}

func NewToDoService(toDoRepository *repository.ToDoRepository) *ToDoService {
	return &ToDoService{
		toDoRepository: toDoRepository,
	}
}

func (s *ToDoService) Create(req dto.ToDoAddReq, userId uint) error {
	toDo := &model.ToDo{
		Title:       req.Title,
		Description: req.Description,
		Status:      0,
		EndTime:     req.EndTime,
		UserId:      userId,
	}
	return s.toDoRepository.Create(toDo)
}

func (s *ToDoService) UpdateByID(req dto.ToDoUpdateReq, id uint) error {
	toDo := &model.ToDo{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		EndTime:     req.EndTime,
	}
	return s.toDoRepository.UpdateById(toDo, id)
}

func (s *ToDoService) Delete(id uint) error {
	return s.toDoRepository.Delete(id)
}

func (s *ToDoService) GetByID(id uint) (*model.ToDo, error) {
	return s.toDoRepository.GetByID(id)
}

func (s *ToDoService) ListByUser(userId uint) ([]vo.ToDoVO, error) {
	toDos, err := s.toDoRepository.ListByUser(userId)
	if err != nil {
		return nil, err
	}

	var toDoVOs []vo.ToDoVO

	for _, toDo := range toDos {
		toDoVO := vo.ToDoVO{
			ID:          toDo.ID,
			Title:       toDo.Title,
			Description: toDo.Description,
			Status:      toDo.Status,
			EndTime:     toDo.EndTime,
		}

		toDoVOs = append(toDoVOs, toDoVO)
	}

	return toDoVOs, nil
}
