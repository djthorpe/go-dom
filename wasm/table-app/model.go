package main

import (
	"syscall/js"
)

// Employee represents an employee record
type Employee struct {
	Name     string
	Position string
	Salary   float64
	Location string
}

// EmployeeModel manages employee data
type EmployeeModel struct {
	employees []Employee
	loaded    bool
	offset    int
	limit     int
}

// NewEmployeeModel creates a new employee model
func NewEmployeeModel() *EmployeeModel {
	return &EmployeeModel{
		employees: []Employee{},
		loaded:    false,
		offset:    0,
		limit:     10,
	}
}

// LoadFromWindow loads employee data from window.testData
// Returns true if data was loaded successfully
func (m *EmployeeModel) LoadFromWindow() bool {
	testData := js.Global().Get("testData")

	if testData.IsUndefined() || testData.IsNull() {
		return false
	}

	length := testData.Length()
	m.employees = make([]Employee, length)

	for i := 0; i < length; i++ {
		item := testData.Index(i)
		m.employees[i] = Employee{
			Name:     item.Get("name").String(),
			Position: item.Get("position").String(),
			Salary:   item.Get("salary").Float(),
			Location: item.Get("location").String(),
		}
	}

	m.loaded = true
	return true
}

// WaitForData waits for window.testData to be available and loads it
// Calls the callback function when data is loaded
func (m *EmployeeModel) WaitForData(callback func()) {
	var intervalID js.Value
	var checkFunc js.Func

	checkFunc = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if m.LoadFromWindow() {
			js.Global().Call("clearInterval", intervalID)
			checkFunc.Release()
			callback()
		}
		return nil
	})

	intervalID = js.Global().Call("setInterval", checkFunc, 100)
}

// Offset returns the current offset within the data
func (m *EmployeeModel) Offset() int {
	return m.offset
}

// Limit returns the limit of number of rows shown
func (m *EmployeeModel) Limit() int {
	return m.limit
}

// Count returns the total number of rows, regardless of offset and limit
func (m *EmployeeModel) Count() int {
	return len(m.employees)
}

// SetOffsetLimit changes the offset and limit
func (m *EmployeeModel) SetOffsetLimit(offset, limit int) {
	if offset < 0 {
		offset = 0
	}
	if limit < 1 {
		limit = 1
	}
	m.offset = offset
	m.limit = limit
}

// GetAll returns employees starting at the offset and no more than limit rows
func (m *EmployeeModel) GetAll() []Employee {
	if m.offset >= len(m.employees) {
		return []Employee{}
	}

	end := m.offset + m.limit
	if end > len(m.employees) {
		end = len(m.employees)
	}

	return m.employees[m.offset:end]
}

// GetRandom returns a random employee from the data
func (m *EmployeeModel) GetRandom() Employee {
	if len(m.employees) == 0 {
		return Employee{}
	}

	randomIndex := js.Global().Get("Math").Call("floor",
		js.Global().Get("Math").Call("random").Float()*float64(len(m.employees))).Int()

	return m.employees[randomIndex]
}

// InsertBefore inserts an employee before the specified index
// If index is 0 or negative, inserts at the beginning
// If index is >= length, appends to the end
func (m *EmployeeModel) InsertBefore(index int, employee Employee) {
	if index <= 0 {
		// Insert at the beginning
		m.employees = append([]Employee{employee}, m.employees...)
		return
	}

	if index >= len(m.employees) {
		// Append to the end
		m.employees = append(m.employees, employee)
		return
	}

	// Insert at the specified position
	m.employees = append(m.employees[:index], append([]Employee{employee}, m.employees[index:]...)...)
}

// Delete removes the employee at the specified index
// Returns true if the delete was successful, false if index is out of bounds
func (m *EmployeeModel) Delete(index int) bool {
	if index < 0 || index >= len(m.employees) {
		return false
	}

	m.employees = append(m.employees[:index], m.employees[index+1:]...)
	return true
}

// Update replaces the employee at the specified index with a new employee
// Returns true if the update was successful, false if index is out of bounds
func (m *EmployeeModel) Update(index int, employee Employee) bool {
	if index < 0 || index >= len(m.employees) {
		return false
	}

	m.employees[index] = employee
	return true
}

// IsLoaded returns true if data has been loaded
func (m *EmployeeModel) IsLoaded() bool {
	return m.loaded
}

// FindByName returns the employee with the given name, or nil if not found
func (m *EmployeeModel) FindByName(name string) *Employee {
	for i := range m.employees {
		if m.employees[i].Name == name {
			return &m.employees[i]
		}
	}
	return nil
}

// GetByIndex returns the employee at the given table row index (offset + index)
// Returns nil if the index is out of bounds
func (m *EmployeeModel) GetByIndex(index int) *Employee {
	actualIndex := m.offset + index
	if actualIndex < 0 || actualIndex >= len(m.employees) {
		return nil
	}
	return &m.employees[actualIndex]
}
