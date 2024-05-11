package payload

type RegisterVehicleRequest struct {
	PlateNumber string `json:"plate_number" binding:"required"`
	Type        string `json:"type" binding:"required"`
}

type VehicleActionRequest struct {
	PlateNumber string `json:"plate_number" binding:"required"`
	ImageUrl    string `json:"image_url" binding:"required"`
}
