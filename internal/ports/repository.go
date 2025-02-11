package ports

// PortRepository defines methods for data storage
type PortRepository interface {
	Save(id string, port Port)  // Now `Port` is recognized
	Get(id string) (Port, bool) // Now `Port` is recognized
}
