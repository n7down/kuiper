//+build unit

package settings

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	settings_pb "github.com/n7down/kuiper/internal/pb/settings"
)

// TODO: add test cases
// func TestGetResourceById(t *testing.T) {
//     w := httptest.NewRecorder()
//     c, _ := gin.CreateTestContext(w)
//     GetResourceById(c)
//     assert.Equal(t, 200, w.Code) // or what value you need it to be

//     var got gin.H
//     err := json.Unmarshal(&got, w.Body().Bytes())
//     if err != nil {
//         t.Fatal(err)
//     }
//     assert.Equal(t, want, got) // want is a gin.H that contains the wanted map.
// }

type MockSettingsServiceClient struct {
	MockCreateBatCaveSetting func(ctx context.Context, in *settings_pb.CreateBatCaveSettingRequest, opts ...grpc.CallOption) (*settings_pb.CreateBatCaveSettingResponse, error)
	MockGetBatCaveSetting    func(ctx context.Context, in *settings_pb.GetBatCaveSettingRequest, opts ...grpc.CallOption) (*settings_pb.GetBatCaveSettingResponse, error)
	MockUpdateBatCaveSetting func(ctx context.Context, in *settings_pb.UpdateBatCaveSettingRequest, opts ...grpc.CallOption) (*settings_pb.UpdateBatCaveSettingResponse, error)
}

func (s *MockSettingsServiceClient) CreateBatCaveSetting(ctx context.Context, in *settings_pb.CreateBatCaveSettingRequest, opts ...grpc.CallOption) (*settings_pb.CreateBatCaveSettingResponse, error) {
	return s.MockCreateBatCaveSetting(ctx, in, opts...)
}

func (s *MockSettingsServiceClient) GetBatCaveSetting(ctx context.Context, in *settings_pb.GetBatCaveSettingRequest, opts ...grpc.CallOption) (*settings_pb.GetBatCaveSettingResponse, error) {
	return s.MockGetBatCaveSetting(ctx, in, opts...)
}

func (s *MockSettingsServiceClient) UpdateBatCaveSetting(ctx context.Context, in *settings_pb.UpdateBatCaveSettingRequest, opts ...grpc.CallOption) (*settings_pb.UpdateBatCaveSettingResponse, error) {
	return s.UpdateBatCaveSetting(ctx, in, opts...)
}

func Test_CreateBatCaveSetting(t *testing.T) {
	assert.Fail(t, "not implemented")
}

func Test_GetBatCaveSetting(t *testing.T) {
	assert.Fail(t, "not implemented")
}

func Test_UpdateBatCaveSetting(t *testing.T) {
	assert.Fail(t, "not implemented")
}
