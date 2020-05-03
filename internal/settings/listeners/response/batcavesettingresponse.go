package response

type BatCaveSettingResponse struct {
	DeepSleepDelay int32 `json:"s"`
}

func GetBatCaveSettingDefault() BatCaveSettingResponse {
	return BatCaveSettingResponse{
		DeepSleepDelay: 15,
	}
}
