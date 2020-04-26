package response

type GetBatCaveSettingResponse struct {
	DeviceID       string `json:"deviceID"`
	DeepSleepDelay int32  `json:"deepSleepDelay"`
}
