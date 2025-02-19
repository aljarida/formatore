package utils

// Stores a single column's name and type (e.g., "INTEGER").
type ColumnBlueprint struct {
	Name string
	Type string
	//	Constraints []string
}

// Stores a table's name and its column's column blueprints.
type TableBlueprint struct {
	Name string
	ColumnBlueprints []ColumnBlueprint 
}
