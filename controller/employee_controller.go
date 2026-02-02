package controller

import (
	"encoding/json"
	"go-mysql-crud/model"
	"go-mysql-crud/repository"
	"net/http"
	"strconv"
)

type EmployeeController struct {
	Repo *repository.EmployeeRepository
}

/* POST /employees */
func (c *EmployeeController) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp model.Employee
	json.NewDecoder(r.Body).Decode(&emp)

	err := c.Repo.Save(emp)
	if err != nil {
		http.Error(w, "Failed to create employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Employee created successfully"))
}

/* GET /employees */
func (c *EmployeeController) GetEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := c.Repo.FindAll()
	if err != nil {
		http.Error(w, "Failed to fetch employees", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

/* PUT /employees?id=1 */
func (c *EmployeeController) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	var emp model.Employee
	json.NewDecoder(r.Body).Decode(&emp)
	emp.ID = id

	err := c.Repo.Update(emp)
	if err != nil {
		http.Error(w, "Failed to update employee", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Employee updated successfully"))
}

/* DELETE /employees?id=1 */
func (c *EmployeeController) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	err := c.Repo.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete employee", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Employee deleted successfully"))
}
