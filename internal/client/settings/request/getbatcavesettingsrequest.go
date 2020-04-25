package request

import (
	"net/url"
	"regexp"
)

type GetBatCaveSettingsRequest struct {
	DeviceID string `json:"-"`
}

func (r *GetBatCaveSettingsRequest) Validate() url.Values {
	errs := url.Values{}

	regex, _ := regexp.Compile("[a-f0-9]{12}")
	isMacAddress := regex.MatchString(r.DeviceID)
	if !isMacAddress {
		errs.Add("deviceID", "The deviceID field needs to be a valid mac!")
	}

	return errs
}
