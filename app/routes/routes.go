package routes

import (
	"fmt"

	"github.com/haleyrc/phir/app/models"
)

func IndexPath() string                   { return "/" }
func PatientsPath() string                { return "/patients" }
func NewPatientPath() string              { return "/patients/new" }
func PatientPath(p models.Patient) string { return fmt.Sprintf("/patients/%d", p.ID) }
