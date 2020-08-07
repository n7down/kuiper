// +build !unit,integration

package mysql

import (
	"log"
	"os"
	"testing"

	"github.com/n7down/kuiper/internal/settings/persistence"
	"github.com/stretchr/testify/assert"
)

var (
	db *MysqlPersistence
)

func Test_CreateBatCaveSetting(t *testing.T) {
	var (
		setting = persistence.BatCaveSetting{
			DeviceID:       "000000000000",
			DeepSleepDelay: 30,
			CreatedAt:      nil,
			UpdatedAt:      nil,
			DeletedAt:      nil,
		}
		rowsAffectedExpected int64 = 1
	)

	rowsAffectedActual := db.CreateBatCaveSetting(setting)
	assert.Equal(t, rowsAffectedExpected, rowsAffectedActual)
}

func Test_GetBatCaveSetting(t *testing.T) {
	var (
		deviceID        = "000000000011"
		settingExpected = persistence.BatCaveSetting{
			DeviceID:       deviceID,
			DeepSleepDelay: 30,
			CreatedAt:      nil,
			UpdatedAt:      nil,
			DeletedAt:      nil,
		}
		recordNotFoundExpected = false
	)

	recordNotFoundActual, settingActual := db.GetBatCaveSetting(deviceID)
	assert.Equal(t, recordNotFoundExpected, recordNotFoundActual)
	assert.True(t, settingExpected.Equal(settingActual))
}

func Test_UpdateBatCaveSetting(t *testing.T) {
	var (
		deviceID              = "000000001111"
		deepSleepDelay uint32 = 32

		setting = persistence.BatCaveSetting{
			DeviceID:       deviceID,
			DeepSleepDelay: deepSleepDelay,
			CreatedAt:      nil,
			UpdatedAt:      nil,
			DeletedAt:      nil,
		}

		settingExpected = persistence.BatCaveSetting{
			DeviceID:       deviceID,
			DeepSleepDelay: deepSleepDelay,
			CreatedAt:      nil,
			UpdatedAt:      nil,
			DeletedAt:      nil,
		}
		rowsAffectedExpected   int64 = 1
		recordNotFoundExpected       = false
	)

	rowsAffectedActual := db.UpdateBatCaveSetting(setting)
	assert.Equal(t, rowsAffectedExpected, rowsAffectedActual)

	recordNotFoundActual, settingActual := db.GetBatCaveSetting(deviceID)
	assert.Equal(t, recordNotFoundExpected, recordNotFoundActual)
	assert.Equal(t, settingExpected.DeviceID, settingActual.DeviceID)
	assert.Equal(t, settingExpected.DeepSleepDelay, settingActual.DeepSleepDelay)
}

func TestMain(m *testing.M) {
	var (
		err error
	)

	dbConn := os.Getenv("DB_CONN")
	db, err = NewMysqlPersistence(dbConn)
	if err != nil {
		log.Fatal(err.Error())
	}
	code := m.Run()
	os.Exit(code)
}
