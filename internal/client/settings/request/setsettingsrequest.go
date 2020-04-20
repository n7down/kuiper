package request

import (
	"net/url"
)

type SetSettingsRequest struct {
	DeviceID       string `json:"deviceID" binding:"required"`
	DeepSleepDelay int32  `json:"deepSleepDelay" binding:"required"`
}

func (r *SetSettingsRequest) Validate() url.Values {
	errs := url.Values{}

	if r.DeviceID == "" {
		errs.Add("deviceID", "The deviceID field is required!")
	}

	if r.DeepSleepDelay == 0 {
		errs.Add("deepSleepDelay", "The deepSleepDelay field is required!")
	}

	return errs
}
