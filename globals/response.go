package globals

type (
	API_Resp struct {
		Success bool
		Data    interface{}
	}

	DistanceAPI struct {
		Distance int `json:"distance"`
	}
)
