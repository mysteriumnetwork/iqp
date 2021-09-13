package enterprise

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/mysteriumnetwork/iqp/enterprise/mocks/mocks"
	"github.com/stretchr/testify/assert"
)

func TestEnterprise(t *testing.T) {
	enterpriseAdress := common.HexToAddress("0x34437589B4DC1EAcBe08824645164F93E5d989E1")
	t.Run("returns correct ID", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockBc := mocks.NewMockEnterpriseBlockchain(mockCtrl)

		mockBc.EXPECT().GetChainID().Return(int64(1), nil).Times(1)

		s := NewEnterprise(enterpriseAdress, mockBc)
		acc, err := s.GetID()
		assert.NoError(t, err)

		assert.Equal(t, fmt.Sprintf("eip155:1:%v", strings.ToLower(enterpriseAdress.Hex())), acc.String())
	})
	t.Run("returns enterprise info", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockBc := mocks.NewMockEnterpriseBlockchain(mockCtrl)
		mockBc.EXPECT().GetEnterpriseInfo(enterpriseAdress).Times(1)

		s := NewEnterprise(enterpriseAdress, mockBc)
		_, err := s.GetInfo()
		assert.NoError(t, err)
	})
	t.Run("lists registered services", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockBc := mocks.NewMockEnterpriseBlockchain(mockCtrl)
		addr1 := common.HexToAddress("0x52De41D6a2104812f84ef596BE15B84d1d846ee5")
		addr2 := common.HexToAddress("0x2C368A2E9Bd1bf16eb3DfCd924CA7eF4969CBBD9")
		mockBc.EXPECT().GetServices(enterpriseAdress).Times(1).Return([]common.Address{addr1, addr2}, nil)

		s := NewEnterprise(enterpriseAdress, mockBc)
		res, err := s.GetServices()
		assert.NoError(t, err)

		assert.Len(t, res, 2)
		assert.Equal(t, addr1, res[0].GetAddress())
		assert.Equal(t, addr2, res[1].GetAddress())
	})
	t.Run("allows to estimate loan", func(t *testing.T) {
		estimation := big.NewInt(250)
		serviceAddress := common.HexToAddress("0x52De41D6a2104812f84ef596BE15B84d1d846ee5")
		paymentTokenAddress := common.HexToAddress("0x2C368A2E9Bd1bf16eb3DfCd924CA7eF4969CBBD9")
		amount := big.NewInt(1000)
		var duration uint32 = 86400
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockBc := mocks.NewMockEnterpriseBlockchain(mockCtrl)
		mockBc.EXPECT().EstimateLoan(enterpriseAdress, serviceAddress, paymentTokenAddress, amount, duration).Times(1).Return(estimation, nil)

		s := NewEnterprise(enterpriseAdress, mockBc)
		res, err := s.EstimateLoan(serviceAddress, paymentTokenAddress, amount, duration)
		assert.NoError(t, err)
		assert.Equal(t, estimation, res)
	})
	t.Run("retrieves accrued interest", func(t *testing.T) {
		interestTokenID := big.NewInt(1)
		mockInterest := big.NewInt(125)
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockBc := mocks.NewMockEnterpriseBlockchain(mockCtrl)
		mockBc.EXPECT().GetAccruedInterest(enterpriseAdress, interestTokenID).Times(1).Return(mockInterest, nil)

		s := NewEnterprise(enterpriseAdress, mockBc)
		res, err := s.GetAccruedInterest(interestTokenID)
		assert.NoError(t, err)
		assert.Equal(t, mockInterest, res)
	})
	t.Run("allows to withdraw interest", func(t *testing.T) {
		interestTokenID := big.NewInt(1)
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockBc := mocks.NewMockEnterpriseBlockchain(mockCtrl)
		mockBc.EXPECT().WithdrawInterest(enterpriseAdress, interestTokenID).Times(1)

		s := NewEnterprise(enterpriseAdress, mockBc)
		_, err := s.WithdrawInterest(interestTokenID)
		assert.NoError(t, err)
	})
	t.Run("allows to return loan", func(t *testing.T) {
		borrowTokenID := big.NewInt(1)
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockBc := mocks.NewMockEnterpriseBlockchain(mockCtrl)
		mockBc.EXPECT().ReturnLoan(enterpriseAdress, borrowTokenID).Times(1)

		s := NewEnterprise(enterpriseAdress, mockBc)
		_, err := s.ReturnLoan(borrowTokenID)
		assert.NoError(t, err)
	})
}
