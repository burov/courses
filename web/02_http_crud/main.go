package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

var (
	ErrRecordNotFound     = fmt.Errorf("record not found")
	ErrRecordAlreadyExist = fmt.Errorf("already exist")
	ErrEmptyIdentifier    = fmt.Errorf("empty identifier")
)

var employeesStorage = NewEmployeesStorage()

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/users").Subrouter()

	s.HandleFunc("/", ListEmployee).Methods(http.MethodGet)
	s.HandleFunc("/", CreateEmployee).Methods(http.MethodPost)

	s.HandleFunc("/{id}", ReadEmployee).Methods(http.MethodGet)
	s.HandleFunc("/{id}", ReplaceEmployee).Methods(http.MethodPut)
	s.HandleFunc("/{id}", UpdateEmployee).Methods(http.MethodPatch)
	s.HandleFunc("/{id}", DeleteEmployee).Methods(http.MethodDelete)

	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Employee

	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
		return
	}

	result, err := employeesStorage.Save(employee)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
		return
	}

	marshaled, _ := json.MarshalIndent(result, "", "  ")
	w.Write(marshaled)
}

func ReplaceEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Employee

	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
		return
	}
	id := mux.Vars(r)["id"]
	employee.ID = id

	result, err := employeesStorage.Update(employee)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
		return
	}

	marshaled, _ := json.MarshalIndent(result, "", "  ")
	w.Write(marshaled)

}

//TODO: now UpdateEmployee works with patch behavior, add field mask
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Employee

	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
		return
	}
	id := mux.Vars(r)["id"]
	employee.ID = id

	result, err := employeesStorage.Update(employee)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
		return
	}

	marshaled, _ := json.MarshalIndent(result, "", "  ")
	w.Write(marshaled)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID mustn't be empty")
		return
	}

	err := employeesStorage.Delete(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal server error")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func ListEmployee(w http.ResponseWriter, r *http.Request) {
	employees, err := employeesStorage.GetAll()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Bad request")
		return
	}

	marshaled, _ := json.MarshalIndent(employees, "", " ")
	w.Write(marshaled)
}

func ReadEmployee(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID mustn't be empty")
		return
	}

	employee, err := employeesStorage.GetByID(id)
	if err == ErrRecordNotFound {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Record Not Found")
		return
	} else if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal server error")
		return
	}

	marshaled, _ := json.MarshalIndent(employee, "", " ")
	w.Write(marshaled)
}

type EmployeesStorage struct {
	employees map[string]Employee
	sync.RWMutex
}

func (s *EmployeesStorage) GetAll() ([]Employee, error) {
	result := make([]Employee, 0, len(s.employees))

	s.RLock()
	defer s.RUnlock()

	for _, e := range s.employees {
		result = append(result, e)
	}

	return result, nil
}

func (s *EmployeesStorage) GetByID(id string) (Employee, error) {
	s.RLock()
	defer s.RUnlock()

	employee, ok := s.employees[id]
	if !ok {
		return Employee{}, ErrRecordNotFound
	}

	return employee, nil
}

func (s *EmployeesStorage) Save(employee Employee) (Employee, error) {
	s.Lock()
	defer s.Unlock()

	employee.ID = uuid.New().String()
	s.employees[employee.ID] = employee

	return employee, nil
}

func (s *EmployeesStorage) Update(employee Employee) (Employee, error) {
	s.Lock()
	defer s.Unlock()

	if employee.ID == "" {
		return Employee{}, ErrEmptyIdentifier
	}

	s.employees[employee.ID] = employee

	return employee, nil
}

func (s *EmployeesStorage) Delete(id string) error {
	s.Lock()
	defer s.Unlock()

	delete(s.employees, id)
	return nil
}

func NewEmployeesStorage() *EmployeesStorage {
	storage := EmployeesStorage{
		employees: make(map[string]Employee, 0),
	}

	return &storage
}
