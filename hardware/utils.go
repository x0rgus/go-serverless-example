package hardware

import (
	"database/sql"
)

// Helper para converter campos nullable
func ConvertNullString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
