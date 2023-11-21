package models

type (
	QueryCtx struct {
		Origin string `json:"origin" validate:"gte=2,lte=300,required"`
		Dest   string `json:"dest"   validate:"gte=2,lte=300,required"`
		Token  string `json:"token"  validate:"gte=0,lte=300"`
	}
)
