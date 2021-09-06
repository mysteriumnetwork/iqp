package energy

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func TestHalfLife(t *testing.T) {
	type args struct {
		InitialValue             string
		GapHalvingPeriod, T0, T1 int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{
				InitialValue:     "1152921504606846976000",
				GapHalvingPeriod: 20,
				T0:               100,
				T1:               120,
			},
			wantErr: false,
			want:    "576460752303423488000",
		},
		{
			args: args{
				InitialValue:     "1152921504606846976000",
				GapHalvingPeriod: 20,
				T0:               100,
				T1:               140,
			},
			wantErr: false,
			want:    "288230376151711744000",
		},
		{
			args: args{
				InitialValue:     "1152921504606846976000",
				GapHalvingPeriod: 20,
				T0:               100,
				T1:               110,
			},
			wantErr: false,
			want:    "815238614083298888273",
		},
		{
			args: args{
				InitialValue:     "2302672475076025122816",
				GapHalvingPeriod: 20,
				T0:               100,
				T1:               110,
			},
			wantErr: false,
			want:    "1628235321977868704599",
		},
		{
			args: args{
				InitialValue:     "2302672475076025122816",
				GapHalvingPeriod: 2373046875,
				T0:               0,
				T1:               2373046874,
			},
			wantErr: false,
			want:    "1151336237874308264392",
		},
		{
			args: args{
				InitialValue:     "5192296858534774017680532110835712",
				GapHalvingPeriod: 2373046875,
				T0:               0,
				T1:               2373046874,
			},
			wantErr: false,
			want:    "2596148430025700290926036596036491",
		},
		{
			args: args{
				InitialValue:     "5192296858534827052069744025796608",
				GapHalvingPeriod: 2373046875,
				T0:               0,
				T1:               2373046874,
			},
			wantErr: false,
			want:    "2596148430025726808120650298968105",
		},
		{
			name: "throws err when period is negative",
			args: args{
				InitialValue:     "1000",
				GapHalvingPeriod: 50,
				T0:               50,
				T1:               20,
			},
			wantErr: true,
		},
	}
	for i, tt := range tests {
		if tt.name == "" {
			tt.name = fmt.Sprintf("test case %v", i)
		}
		t.Run(tt.name, func(t *testing.T) {
			iv, _ := big.NewInt(0).SetString(tt.args.InitialValue, 10)
			in := HalfLifeParams{
				InitialValue: iv,
				GapHalvingParams: GapHalvingParams{
					GapHalvingPeriod: tt.args.GapHalvingPeriod,
					T0:               tt.args.T0,
					T1:               tt.args.T1,
				},
			}
			got, err := HalfLife(in)
			fmt.Println(err != nil)
			fmt.Println(!tt.wantErr)
			if (err != nil) && !tt.wantErr {
				t.Errorf("HalfLife() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got.String() != tt.want {
				t.Errorf("HalfLife() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}

func TestCalculateLinearEnergy(t *testing.T) {
	type args struct {
		InitialValue             string
		GapHalvingPeriod, T0, T1 int64
		Power                    string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{
				InitialValue:     "200",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               50,
				Power:            "0",
			},
			wantErr: false,
			want:    "200",
		},
		{
			args: args{
				InitialValue:     "200",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               50,
				Power:            "5",
			},
			wantErr: false,
			want:    "201",
		},
		{
			args: args{
				InitialValue:     "200",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               25,
				Power:            "100",
			},
			wantErr: false,
			want:    "212",
		},
		{
			args: args{
				InitialValue:     "200",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               50,
				Power:            "100",
			},
			wantErr: false,
			want:    "225",
		},
		{
			args: args{
				InitialValue:     "200",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               100,
				Power:            "100",
			},
			wantErr: false,
			want:    "250",
		},
		{
			args: args{
				InitialValue:     "200",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               50,
				Power:            "-100",
			},
			wantErr: false,
			want:    "175",
		},
		{
			args: args{
				InitialValue:     "200",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               25,
				Power:            "-100",
			},
			wantErr: false,
			// this does not strictly match the js implementation, as they're using floats
			// on their side this returns 188
			want: "187",
		},
		{
			args: args{
				InitialValue:     "200",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               100,
				Power:            "-100",
			},
			wantErr: false,
			want:    "150",
		},
		{
			args: args{
				InitialValue:     "-200",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               100,
				Power:            "-100",
			},
			wantErr: false,
			want:    "-250",
		},
		{
			args: args{
				InitialValue:     "-100",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               100,
				Power:            "-200",
			},
			wantErr: false,
			want:    "-200",
		},
		{
			args: args{
				InitialValue:     "0",
				GapHalvingPeriod: 9007199254740991,
				T0:               0,
				T1:               9007199254740991,
				Power:            "300",
			},
			wantErr: false,
			want:    "75",
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "" {
				tt.name = fmt.Sprintf("test case %v", i)
			}
			iv, _ := big.NewInt(0).SetString(tt.args.InitialValue, 10)
			p, _ := big.NewInt(0).SetString(tt.args.Power, 10)
			in := EnergyCalculationParams{
				HalfLifeParams: HalfLifeParams{
					InitialValue: iv,
					GapHalvingParams: GapHalvingParams{
						GapHalvingPeriod: tt.args.GapHalvingPeriod,
						T0:               tt.args.T0,
						T1:               tt.args.T1,
					},
				},
				Power: p,
			}
			got, err := CalculateLinearEnergy(in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateLinearEnergy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.String() != tt.want {
				t.Errorf("CalculateLinearEnergy() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}

func TestCalculateEnergyCap(t *testing.T) {
	type args struct {
		InitialValue             string
		GapHalvingPeriod, T0, T1 int64
		Power                    string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{
				InitialValue:     "0",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               0,
				Power:            "100",
			},
			wantErr: false,
			want:    "0",
		},
		{
			args: args{
				InitialValue:     "0",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               25,
				Power:            "100",
			},
			wantErr: false,
			want:    "29",
		},
		{
			args: args{
				InitialValue:     "0",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               50,
				Power:            "100",
			},
			wantErr: false,
			want:    "50",
		},
		{
			args: args{
				InitialValue:     "10",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               50,
				Power:            "100",
			},
			wantErr: false,
			want:    "55",
		},
		{
			args: args{
				InitialValue:     "0",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               100,
				Power:            "100",
			},
			wantErr: false,
			want:    "75",
		},
		{
			args: args{
				InitialValue:     "0",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               150,
				Power:            "100",
			},
			wantErr: false,
			want:    "88",
		},
		{
			args: args{
				InitialValue:     "0",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               200,
				Power:            "100",
			},
			wantErr: false,
			want:    "94",
		},
		{
			args: args{
				InitialValue:     "200",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               25,
				Power:            "100",
			},
			wantErr: false,
			want:    "171",
		},
		{
			args: args{
				InitialValue:     "200",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               50,
				Power:            "100",
			},
			wantErr: false,
			want:    "150",
		},
		{
			args: args{
				InitialValue:     "200",
				GapHalvingPeriod: 50,
				T0:               0,
				T1:               100,
				Power:            "100",
			},
			wantErr: false,
			want:    "125",
		},
		{
			args: args{
				InitialValue:     "0",
				GapHalvingPeriod: 9007199254740991,
				T0:               0,
				T1:               1,
				Power:            "1",
			},
			wantErr: false,
			want:    "0",
		},
	}
	for i, tt := range tests {
		if tt.name == "" {
			tt.name = fmt.Sprintf("test case %v", i)
		}
		t.Run(tt.name, func(t *testing.T) {
			iv, _ := big.NewInt(0).SetString(tt.args.InitialValue, 10)
			p, _ := big.NewInt(0).SetString(tt.args.Power, 10)
			in := EnergyCalculationParams{
				HalfLifeParams: HalfLifeParams{
					InitialValue: iv,
					GapHalvingParams: GapHalvingParams{
						GapHalvingPeriod: tt.args.GapHalvingPeriod,
						T0:               tt.args.T0,
						T1:               tt.args.T1,
					},
				},
				Power: p,
			}
			got, err := CalculateEnergyCap(in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateEnergyCap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.String() != tt.want {
				t.Errorf("CalculateEnergyCap() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}

func TestCalculateEffectiveEnergy(t *testing.T) {
	type args struct {
		GapHalvingPeriod, T0, T1 int64
		Power                    string
		Energy                   string
		EnergyCap                string
	}
	tests := []struct {
		name    string
		args    args
		want    EffectiveEnergyResult
		wantErr bool
	}{
		{
			args: args{
				GapHalvingPeriod: 50,
				Power:            "100",
				Energy:           "200",
				EnergyCap:        "200",
				T0:               0,
				T1:               50,
			},
			want: EffectiveEnergyResult{
				EnergyCap:    big.NewInt(150),
				LinearEnergy: big.NewInt(225),
				Energy:       big.NewInt(150),
			},
		},
	}
	for i, tt := range tests {
		if tt.name == "" {
			tt.name = fmt.Sprintf("test case %v", i)
		}
		t.Run(tt.name, func(t *testing.T) {
			p, _ := big.NewInt(0).SetString(tt.args.Power, 10)
			e, _ := big.NewInt(0).SetString(tt.args.Energy, 10)
			ec, _ := big.NewInt(0).SetString(tt.args.EnergyCap, 10)
			in := EffectiveEnergyCalculationParams{
				GapHalvingParams: GapHalvingParams{
					GapHalvingPeriod: tt.args.GapHalvingPeriod,
					T0:               tt.args.T0,
					T1:               tt.args.T1,
				},
				Power:     p,
				Energy:    e,
				EnergyCap: ec,
			}
			got, err := CalculateEffectiveEnergy(in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateEffectiveEnergy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateEffectiveEnergy() = %v, want %v", got, tt.want)
			}
		})
	}
}
