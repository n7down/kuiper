package request

import (
	"net/url"
	"regexp"
)

type CreateBatCaveSettingsRequest struct {
	DeviceID       string `json:"deviceID" binding:"required"`
	DeepSleepDelay int32  `json:"deepSleepDelay" binding:"required"`
}

func (r *CreateBatCaveSettingsRequest) Validate() url.Values {
	errs := url.Values{}

	if r.DeviceID == "" {
		errs.Add("deviceID", "The deviceID field is required!")
	}

	if len(r.DeviceID) != 12 {
		errs.Add("deviceID", "The deviceID field needs to be a valid mac with 12 characters!")
	}

	regex, _ := regexp.Compile("[a-f0-9]{12}")
	isMacAddress := regex.MatchString(r.DeviceID)
	if !isMacAddress {
		errs.Add("deviceID", "The deviceID field needs to be a valid mac!")
	}

	if r.DeepSleepDelay < 1 {
		errs.Add("deepSleepDelay", "The deepSleepDelay field is required!")
	}

	return errs
}
