package repository

import (
	"context"
	"database/sql"
	"go-mysql-crud/model"
)

type EmployeeRepository struct {
	DB *sql.DB
}

/* CREATE */
func (r *EmployeeRepository) Save(emp model.Employee) error {
	query := "INSERT INTO employees(name, email, salary) VALUES (?,?,?)"
	_, err := r.DB.ExecContext(context.Background(),
		query, emp.Name, emp.Email, emp.Salary)
	return err
}

/* READ ALL */
func (r *EmployeeRepository) FindAll() ([]model.Employee, error) {
	rows, err := r.DB.Query("SELECT id, name, email, salary FROM employees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []model.Employee

	for rows.Next() {
		var emp model.Employee
		rows.Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Salary)
		employees = append(employees, emp)
	}

	return employees, nil
}

/* UPDATE */
func (r *EmployeeRepository) Update(emp model.Employee) error {
	query := "UPDATE employees SET name=?, email=?, salary=? WHERE id=?"
	_, err := r.DB.Exec(query,
		emp.Name, emp.Email, emp.Salary, emp.ID)
	return err
}

/* DELETE */
func (r *EmployeeRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM employees WHERE id=?", id)
	return err
}
