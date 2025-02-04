package hardware

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func handleUserDevices(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	userID := r.URL.Query().Get("user_id")

	devices, err := GetUserDevices(db, userID)
	if err != nil {
		respondWithError(w, "Erro ao buscar dispositivos", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": userID,
		"devices": devices,
	})
}

func handleOrgDevices(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	orgID := r.URL.Query().Get("org_id")

	devices, err := GetOrganizationDevices(db, orgID)
	if err != nil {
		respondWithError(w, "Erro ao buscar dispositivos", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"organization_id": orgID,
		"devices":         devices,
	})
}
