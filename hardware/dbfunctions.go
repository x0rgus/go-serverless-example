package hardware

import (
	"database/sql"
)

// GetUserDevices - Dispositivos associados a um usuário
func GetUserDevices(db *sql.DB, userID string) ([]map[string]interface{}, error) {
	query := `
		SELECT 
			d.id,
			d.name,
			d.type
		FROM 
			users u
		JOIN user_devices ud ON u.id = ud.user_id
		JOIN devices d ON ud.device_id = d.id
		WHERE 
			u.id = ? AND d.status = 1;`

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []map[string]interface{}
	for rows.Next() {
		var id, name, deviceType string
		err := rows.Scan(&id, &name, &deviceType)
		if err != nil {
			continue
		}

		devices = append(devices, map[string]interface{}{
			"id":   id,
			"name": name,
			"type": deviceType,
		})
	}

	return devices, nil
}

// GetOrganizationDevices - Todos dispositivos de uma organização
func GetOrganizationDevices(db *sql.DB, orgID string) ([]map[string]interface{}, error) {
	query := `
		SELECT 
			d.id,
			d.name,
			d.model,
			d.status
		FROM 
			devices d
		WHERE 
			d.organization_id = ?;`

	rows, err := db.Query(query, orgID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var devices []map[string]interface{}
	for rows.Next() {
		var id, name, model string
		var status int
		err := rows.Scan(&id, &name, &model, &status)
		if err != nil {
			continue
		}

		devices = append(devices, map[string]interface{}{
			"id":     id,
			"name":   name,
			"model":  model,
			"status": status,
		})
	}

	return devices, nil
}
