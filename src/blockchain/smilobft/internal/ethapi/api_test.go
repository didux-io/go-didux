package ethapi

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
	"go-didux/src/blockchain/smilobft/core/types"
	"testing"
)

func TestMe(t *testing.T) {
	src, _ := hexutil.Decode("0xf8b501800194000000000000000000000000000000000000000080b85874412b473435326e564f35315a6c766f366a34565a6c70452f645472565666306a3135696948367477324870657032736143444965655039566435663479334161523133756a753272734c5a7a3464694e6e4e5754673d3d26a061674284fe07bd6554931892297b30e0912983dbd45ef9ad223178662df0551d9f3e5f267d4e6565b88b655bcb77b6e467c322dfc75378cc7df38d916a628514")
	fmt.Println("Encoded: ", src)
	tx := new(types.Transaction)

	fmt.Println("BeforeDecode: ", tx.String())
	if err := rlp.DecodeBytes(src, tx); err != nil {
		fmt.Println(common.Hash{}, err)
	}

	sender, _ := types.Sender(types.HomesteadSigner{}, tx)
	fmt.Println(sender.String())

	fmt.Println("Incoming: ", tx.String())
}
