package account

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestAccountID(t *testing.T) {
	t.Run("formats to string correctly", func(t *testing.T) {
		acc, err := NewAccountID(ChainEthGoerli, common.HexToAddress("0x5c3f649ffdbc91a247ac45fc2c4c63f9319e5132"))
		assert.NoError(t, err)

		assert.Equal(t, "eip155:5:0x5c3f649ffdbc91a247ac45fc2c4c63f9319e5132", acc.String())
	})
}

func Test_extractChainID(t *testing.T) {
	type args struct {
		chid ChainIdentitifier
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "extracts chain id",
			args: args{
				chid: ChainEthGoerli,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "errs on invalid integer",
			args: args{
				chid: "asdasd:asdasd:asdasd",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractChainID(tt.args.chid)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractChainID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractChainID() = %v, want %v", got, tt.want)
			}
		})
	}
}
