package task

type Service interface {
	Store(input InputTask) (Task, error)
	FetchAll() ([]Task, error)
	FetchById(id int) (Task, error)
	Update(id int, inputTask InputTask) (Task, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) Store(input InputTask) (Task, error) {
	var task Task
	task.Name = input.Name
	task.Description = input.Description

	newTask, err := s.repository.Store(task)
	if err != nil {
		return task, err
	}

	return newTask, nil
}

func (s *service) FetchAll() ([]Task, error) {
	tasks, err := s.repository.FetchAll()
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (s *service) FetchById(id int) (Task, error) {
	task, err := s.repository.FetchById(id)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (s *service) Update(id int, input InputTask) (Task, error) {
	uTask, err := s.repository.Update(id, input)
	if err != nil {
		return uTask, err
	}

	return uTask, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
