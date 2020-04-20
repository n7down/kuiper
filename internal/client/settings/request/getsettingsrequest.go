package request

import (
	"net/url"
)

type GetSettingsRequest struct {
	DeviceID string `json:"deviceID" binding:"required"`
}

func (r *GetSettingsRequest) Validate() url.Values {
	errs := url.Values{}

	if r.DeviceID == "" {
		errs.Add("deviceID", "The deviceID field is required!")
	}

	return errs
}
