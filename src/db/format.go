package db

import (
	"formatore/src/structs"
	"strings"
)

// Return name without spaces in lowercase.
func formatCBName(cbName string) string {
	return strings.ReplaceAll(strings.ToLower(cbName), " ", "_")
}

// Return type in all uppercase.
func formatCBType(cbType string) string {
	return strings.ToUpper(cbType)
}

// Return formatted string versions of ColumnBlueprints' name and type.
func formatColumnBlueprints(cbs []structs.ColumnBlueprint) {
	for i := range cbs {
		cbs[i].Name = formatCBName(cbs[i].Name)
		cbs[i].Type = formatCBType(cbs[i].Type)
	}
}
