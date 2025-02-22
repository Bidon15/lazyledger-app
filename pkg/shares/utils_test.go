package shares

import (
	"reflect"
	"testing"

	"github.com/celestiaorg/celestia-app/testutil/testfactory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/types"
)

func FuzzBlobSharesUsed(f *testing.F) {
	f.Add(uint32(1))
	f.Fuzz(func(t *testing.T, a uint32) {
		if a < 1 {
			t.Skip()
		}
		ml := SparseSharesNeeded(a)
		blob := testfactory.GenerateRandomBlob(int(a))
		rawShares, err := SplitBlobs(0, nil, []types.Blob{blob}, false)
		require.NoError(t, err)
		require.Equal(t, len(rawShares), ml)
	})
}

func Test_zeroPadIfNecessary(t *testing.T) {
	type args struct {
		share []byte
		width int
	}
	tests := []struct {
		name               string
		args               args
		wantPadded         []byte
		wantBytesOfPadding int
	}{
		{"pad", args{[]byte{1, 2, 3}, 6}, []byte{1, 2, 3, 0, 0, 0}, 3},
		{"not necessary (equal to shareSize)", args{[]byte{1, 2, 3}, 3}, []byte{1, 2, 3}, 0},
		{"not necessary (greater shareSize)", args{[]byte{1, 2, 3}, 2}, []byte{1, 2, 3}, 0},
	}
	for _, tt := range tests {
		tt := tt // stupid scopelint :-/
		t.Run(tt.name, func(t *testing.T) {
			gotPadded, gotBytesOfPadding := zeroPadIfNecessary(tt.args.share, tt.args.width)
			if !reflect.DeepEqual(gotPadded, tt.wantPadded) {
				t.Errorf("zeroPadIfNecessary gotPadded %v, wantPadded %v", gotPadded, tt.wantPadded)
			}
			if gotBytesOfPadding != tt.wantBytesOfPadding {
				t.Errorf("zeroPadIfNecessary gotBytesOfPadding %v, wantBytesOfPadding %v", gotBytesOfPadding, tt.wantBytesOfPadding)
			}
		})
	}
}

func TestParseDelimiter(t *testing.T) {
	for i := uint64(0); i < 100; i++ {
		tx := testfactory.GenerateRandomTxs(1, int(i))[0]
		input, err := MarshalDelimitedTx(tx)
		if err != nil {
			panic(err)
		}
		res, txLen, err := ParseDelimiter(input)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, i, txLen)
		assert.Equal(t, []byte(tx), res)
	}
}
