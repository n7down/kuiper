package request

import (
	"net/url"
	"regexp"
)

type UpdateBatCaveSettingRequest struct {
	DeviceID       string `json:"-"`
	DeepSleepDelay int32  `json:"deepSleepDelay" binding:"required"`
}

func (r *UpdateBatCaveSettingRequest) Validate() url.Values {
	errs := url.Values{}

	regex, _ := regexp.Compile("[a-f0-9]{12}")
	isMacAddress := regex.MatchString(r.DeviceID)
	if !isMacAddress {
		errs.Add("deviceID", "The deviceID field needs to be a valid mac!")
	}

	if r.DeepSleepDelay < 1 {
		errs.Add("deepSleepDelay", "The deepSleepDelay field should be a positive non-zero value!")
	}

	return errs
}
