package payload

type RegisterVehicleRequest struct {
	PlateNumber string `json:"plate_number" binding:"required"`
	Type        string `json:"type" binding:"required"`
}
