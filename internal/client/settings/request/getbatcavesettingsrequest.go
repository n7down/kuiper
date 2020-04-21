package request

import (
	"net/url"
)

type GetBatCaveSettingsRequest struct {
	DeviceID string `json:"deviceID" binding:"required"`
}

func (r *GetBatCaveSettingsRequest) Validate() url.Values {
	errs := url.Values{}

	if r.DeviceID == "" {
		errs.Add("deviceID", "The deviceID field is required!")
	}

	return errs
}
