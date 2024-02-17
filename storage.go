package main

type Storage interface {
	Create(*Todo) error
	List() ([]*Todo, error)
	Get(int) (*Todo, error)
	Delete(int) error
	Update(int, *UpdateTodoRequest) error
}

func (s *PostgresStore) Create(t *Todo) error {
	query := `INSERT INTO todo
	(title, status, created_at)
	VALUES ($1, $2, $3)`

	_, err := s.db.Query(
		query,
		t.Title,
		t.Status,
		t.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) List() ([]*Todo, error) {
	query := `select * from todo order by id`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	todos := []*Todo{}
	for rows.Next() {
		task := new(Todo)
		err := rows.Scan(&task.ID, &task.Title, &task.Status, &task.CreatedAt)
		if err != nil {
			return nil, err
		}

		todos = append(todos, task)

	}
	return todos, nil
}

func (s *PostgresStore) Get(id int) (*Todo, error) {
	query := `select * from todo where id = $1`
	row, err := s.db.Query(query, id)
	if err != nil {
		return nil, err
	}

	task := new(Todo)
	if row.Next() {
		err := row.Scan(&task.ID, &task.Title, &task.Status, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return task, nil
}

func (s *PostgresStore) Delete(id int) error {
	query := `delete from todo where id = $1`
	_, err := s.db.Query(query, id)

	return err
}

func (s *PostgresStore) Update(id int, todo *UpdateTodoRequest) error {
	query := `update todo set title = $1, status = $2 where id = $3`

	_, err := s.db.Query(query, todo.Title, todo.Status, id)

	return err
}
