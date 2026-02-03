package handler

import (
	"encoding/json"
	"inventory_management_system/database/dbhelper"
	"inventory_management_system/models"
	"inventory_management_system/utils"
	"net/http"
)

func UpdateAssetWithConfigHandler(w http.ResponseWriter, r *http.Request) {
	var req models.UpdateAssetReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid request")
		return
	}

	err := dbhelper.UpdateAssetWithConfig(req)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to update asset")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "asset updated successfully",
	})
}
