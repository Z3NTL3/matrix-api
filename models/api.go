package models

type (
	API_Resp struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
	}

	DistanceAPI struct {
		Distance int `json:"distance"`
	}
)
