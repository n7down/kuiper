package response

type CreateBatCaveSettingResponse struct {
	DeviceID       string `json:"deviceID"`
	DeepSleepDelay int32  `json:"deepSleepDelay"`
}
