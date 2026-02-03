package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"inventory_management_system/database/dbhelper"
	"inventory_management_system/middlewares"
	"inventory_management_system/models"
	"inventory_management_system/utils"
	"net/http"
)

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	managerId, roles, err := middlewares.GetUserAndRolesFromContext(r)
	if err != nil {
		utils.RespondError(w, http.StatusUnauthorized, err, "unauthorized")
		return
	}

	role := roles[0]
	if role != "admin" && role != "employee_manager" {
		utils.RespondError(w, http.StatusForbidden, nil, "only admin and asset are allowd")
		return
	}

	managerUUID, err := uuid.Parse(managerId)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid manager id")
		return
	}

	var req models.UpdateEmployeeReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid request body")
		return
	}
	err = validator.New().Struct(req)
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, err, "invalid input")
		return
	}
	if req.Username == "" && req.Email == "" && req.ContactNo == "" {
		utils.RespondError(w, http.StatusBadRequest, nil, "at least one field must be provided for update")
		return
	}
	err = dbhelper.UpdateEmployeeInfo(req, managerUUID)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err, "failed to update employee")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "employee updated successfully",
	})
}
