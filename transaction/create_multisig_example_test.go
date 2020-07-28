package transaction_test

import (
	"fmt"
	"github.com/MinterTeam/minter-go-sdk/transaction"
)

func ExampleNewCreateMultisigData() {
	data := transaction.NewCreateMultisigData().
		MustAddSigData("Mx08d920c5d93dbf23038fe1a54bbb34f41f77677c", 1).
		MustAddSigData("Mx772fd5bd06356250e5efe572b6ae615860ee0c17", 3).
		MustAddSigData("Mx9c7f68ff71b5819c41e8f87cc99bdf6359da3d75", 5).
		SetThreshold(7)

	tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(data)

	msigAddress := data.Address()
	fmt.Println(msigAddress)
	// Result: Mxd43eef7b9406762aa031b82ed0b1082264a13934

	signedTx, _ := tx.SetNonce(11).SetGasPrice(1).SetGasCoin(1).SetSignatureType(transaction.SignatureTypeSingle).
		SetPayload([]byte(fmt.Sprintf("Multisig Address %s", msigAddress))).Sign("ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63")

	signedTxEncode, _ := signedTx.Encode()
	fmt.Println(signedTxEncode)
	// Result: 0xf901940b02018a4d4e54000000000000000cb848f84607c3010305f83f9408d920c5d93dbf23038fe1a54bbb34f41f77677c94772fd5bd06356250e5efe572b6ae615860ee0c17949c7f68ff71b5819c41e8f87cc99bdf6359da3d75b8f0616530383962333265346530393736636136383838636231303233313438626431613966316363323863356434343265353265353836373534666634386436332c20396437383839356661393534623262303766623366323964326165396635656230646330653932356136386566383336326534306334376261346164623330632c20376534303839633762363833663162386431383332613865393737636637396161343539626631373066663139363335343131323734373132346262643037322c204d78643433656566376239343036373632616130333162383265643062313038323236346131333933348001b845f8431ba03032390eb7457b987b223128b956e514c9847fda62c01b62672712719cca4edfa013d566aef3ba4729a2075ef0db8217e7ae390699be6bc959541d8d4b372219cf

	// Output:
	// Mxd43eef7b9406762aa031b82ed0b1082264a13934
	// 0xf8df0b02018a4d4e54000000000000000cb848f84607c3010305f83f9408d920c5d93dbf23038fe1a54bbb34f41f77677c94772fd5bd06356250e5efe572b6ae615860ee0c17949c7f68ff71b5819c41e8f87cc99bdf6359da3d75b83b4d756c74697369672041646472657373204d78643433656566376239343036373632616130333162383265643062313038323236346131333933348001b845f8431ba00dc74e6b4f1b7e34ff06d106bf1b8d963d1538f55730761a88c4d35784fbccc8a05c758e30e1c535e15688bd1c2b240186485f02199f9e12a159fa06daf462ce15
}
