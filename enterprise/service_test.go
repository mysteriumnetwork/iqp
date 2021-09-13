package enterprise

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/mysteriumnetwork/iqp/blockchain/eip155"
	"github.com/mysteriumnetwork/iqp/enterprise/mocks/mocks"
	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	t.Run("returns correct id", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockBc := mocks.NewMockServiceBlockchain(mockCtrl)

		mockBc.EXPECT().GetChainID().Return(int64(1), nil).Times(1)

		s := NewService(common.Address{}, mockBc)
		acc, err := s.GetID()
		assert.NoError(t, err)

		assert.Equal(t, "eip155:1:0x0000000000000000000000000000000000000000", acc.String())
	})
	t.Run("retrieves on-chain data via blockchain provider", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockBc := mocks.NewMockServiceBlockchain(mockCtrl)

		mockBc.EXPECT().GetServiceInfo(common.Address{}).Return(eip155.ServiceInfo{}, nil).Times(1)

		s := NewService(common.Address{}, mockBc)
		_, err := s.GetInfo()
		assert.NoError(t, err)
	})
	t.Run("retrieves on-chain state via blockchain provider", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockBc := mocks.NewMockServiceBlockchain(mockCtrl)
		address := common.HexToAddress("0x52De41D6a2104812f84ef596BE15B84d1d846ee5")

		mockBc.EXPECT().GetAccountState(common.Address{}, address).Return(eip155.AccountState{}, nil).Times(1)
		s := NewService(common.Address{}, mockBc)
		_, err := s.GetAccountState(address)
		assert.NoError(t, err)
	})
}
