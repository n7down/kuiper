package response

type UpdateBatCaveSettingResponse struct {
	DeviceID       string `json:"deviceID"`
	DeepSleepDelay int32  `json:"deepSleepDelay"`
}
