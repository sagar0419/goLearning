package mocktest

import "fmt"

type Database interface {
	Save(data string)
	Fetch() string
}

type RealDatabase struct {
	Data string
}

type MockDatabase struct {
	Data string
}

func (r *RealDatabase) Save(data string) {
	fmt.Println("Saving the data to the real database", data)
	r.Data = data
}

func (m *MockDatabase) Save(data string) {
	fmt.Println("Mock saving data:", data)
	m.Data = data
}

func (r *RealDatabase) Fetch() string {
	fmt.Println("fetching the data from the real database", r.Data)
	return r.Data
}

func (r *MockDatabase) Fetch() string {
	fmt.Println("fetching the data from the real database", r.Data)
	return r.Data
}

func InteractWithDatabase(db Database, data string) string {
	db.Save(data)
	return db.Fetch()
}

func MockTest() {

	fmt.Println("\nMock Database")
	mockDb := &MockDatabase{}
	result := InteractWithDatabase(mockDb, "Test data")
	fmt.Println("Result from mock database:", result)

	fmt.Println("\nReal database")
	realDb := &RealDatabase{}
	result = InteractWithDatabase(realDb, "real data")
	fmt.Println("Result from mock database:", result)
}
