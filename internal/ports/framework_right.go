package ports

type DBPort interface {
	CloseDBConnection()
	Create(answer int32, operation string) error
	// Read()
	// Update()
	// Delete()
}
