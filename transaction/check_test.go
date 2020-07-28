package transaction

import (
	"encoding/hex"
	"math/big"
	"strings"
	"testing"
)

func TestCheck_Sign(t *testing.T) {
	check := NewCheck(
		480,
		TestNetChainID,
		999999,
		1,
		big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)),
		1,
	).SetPassphrase("pass")

	sign, err := check.Sign("64e27afaab363f21eec05291084367f6f1297a7b280d69d672febecda94a09ea")
	if err != nil {
		t.Fatal(err)
	}

	encode, err := sign.Encode()
	if err != nil {
		t.Fatal(err)
	}
	validCheck := "Mcf89a8334383002830f423f01888ac7230489e8000001b841ea3d022c8326965556f1b651b14d3124947b8683f7b3ab56fca06e0b4204757b2a11dace85d0139ce4e8fdb18369d07905e733683b8229f41bc216c784b4d714011ca017bffff4b3f431dc938239cd2727f0c1dfa61ccdc98727fa8e9baf608b3755f5a05b768c53d09c5e9517487820df439f496e16e459862e7d449360ce69a2ccc4d6"
	checkString := encode
	if checkString != validCheck {
		t.Errorf("check want %s,\ngot %s", validCheck, checkString)
	}
}

func TestCheck_Sign1(t *testing.T) {
	check := NewCheck(
		1,
		MainNetChainID,
		999999,
		1,
		big.NewInt(0).Mul(big.NewInt(10), big.NewInt(0).Exp(big.NewInt(10), big.NewInt(18), nil)),
		1,
	).SetPassphrase("pass")

	sign, err := check.Sign("2919c43d5c712cae66f869a524d9523999998d51157dc40ac4d8d80a7602ce02")
	if err != nil {
		t.Fatal(err)
	}

	checkString, err := sign.Encode()
	if err != nil {
		t.Fatal(err)
	}

	validCheck := "Mcf8963101830f423f01888ac7230489e8000001b8416f6efe02689cacf2b169de53e35eb006f3b82c3119f092857d9a30e5f1bafe440f5ae5f617861a00908767ab13db9c08a07e1d207084c62b441cd3ea6eeb5423011ca09f5eee367f22a982766fa6b60b84512f6801714720dd2f36c8a594d3c24452ca9f49175aea5b59ee4162b91e286ba77810a34368c8240b0c96da23507d9c92f6"
	if checkString != validCheck {
		t.Errorf("check want %s,\ngot %s", validCheck, checkString)
	}
}

func TestCheckAddress_Proof(t *testing.T) {
	check, err := NewCheckAddress("Mxa7bc33954f1ce855ed1a8c768fdd32ed927def47", "pass")
	if err != nil {
		t.Fatal(err)
	}

	proofString, err := check.Proof()
	if err != nil {
		t.Fatal(err)
	}

	validProof := "da021d4f84728e0d3d312a18ec84c21768e0caa12a53cb0a1452771f72b0d1a91770ae139fd6c23bcf8cec50f5f2e733eabb8482cf29ee540e56c6639aac469600"
	if proofString != validProof {
		t.Errorf("check want %s,\ngot %s", validProof, proofString)
	}
}

func TestDecodeCheck(t *testing.T) {
	data, err := DecodeCheck("+K6DNDgwAoMPQj+KTU5UAAAAAAAAAIiKxyMEiegAAIpNTlQAAAAAAAAAuEFJfF8+b8GC/Rp5FSKp73V2cQvfvIb9vxZUdu8iDon5/xOA+T8tmi+S/asO3B4mBcwsabcHzUBLLLFSK3q6Te/VABugg8mUUWnwp7vllpc7MtyIdgh4BYCx07x7GIvts704VZSgR7LVNFlG7VSY9b7nE/hidqrARqX++CC+rud6m2+bwd8=")
	if err != nil {
		t.Fatal(err)
	}

	if string(data.Nonce) != "480" {
		t.Errorf("Nonce want %s, got %s", string(data.Nonce), "480")
	}
	if data.ChainID != TestNetChainID {
		t.Errorf("ChainID want %d, got %d", data.ChainID, TestNetChainID)
	}
	if data.Coin != 1 {
		t.Errorf("CoinID want %s, got %d", data.Coin.String(), 1)
	}
	if data.DueBlock != 999999 {
		t.Errorf("DueBlock want %d, got %d", data.DueBlock, 999999)
	}
	if data.Value.String() != "10"+strings.Repeat("0", 18) {
		t.Errorf("Value want %s, got %s", data.Value.String(), "10"+strings.Repeat("0", 18))
	}
	if data.Lock.String() != "985283871505876742053353384809055983203325304659217336382177168476233943196356552072809135181493031263037291805582875144952622907359402361966594748392133888" {
		t.Errorf("Lock want %s, got %s", data.Lock.String(), "985283871505876742053353384809055983203325304659217336382177168476233943196356552072809135181493031263037291805582875144952622907359402361966594748392133888")
	}
	if data.V.Int64() != 27 {
		t.Errorf("V want %d, got %d", data.V.Int64(), 28)
	}
	if hex.EncodeToString(data.R.Bytes()) != "83c9945169f0a7bbe596973b32dc887608780580b1d3bc7b188bedb3bd385594" {
		t.Errorf("R want %s, got %s", hex.EncodeToString(data.R.Bytes()), "83c9945169f0a7bbe596973b32dc887608780580b1d3bc7b188bedb3bd385594")
	}
	if hex.EncodeToString(data.S.Bytes()) != "47b2d5345946ed5498f5bee713f86276aac046a5fef820beaee77a9b6f9bc1df" {
		t.Errorf("S want %s, got %s", hex.EncodeToString(data.S.Bytes()), "47b2d5345946ed5498f5bee713f86276aac046a5fef820beaee77a9b6f9bc1df")
	}

	sender, err := data.Sender()
	if err != nil {
		t.Fatal(err)
	}

	address := "Mxce931863b9c94a526d94acd8090c1c5955a6eb4b"
	if sender != address {
		t.Errorf("Sender want %s, got %s", address, sender)
	}
}

func TestDecodeCheck_Sender(t *testing.T) {
	data, err := DecodeCheck("+K6DNDgwAoMPQj+KTU5UAAAAAAAAAIiKxyMEiegAAIpNTlQAAAAAAAAAuEFJfF8+b8GC/Rp5FSKp73V2cQvfvIb9vxZUdu8iDon5/xOA+T8tmi+S/asO3B4mBcwsabcHzUBLLLFSK3q6Te/VABugg8mUUWnwp7vllpc7MtyIdgh4BYCx07x7GIvts704VZSgR7LVNFlG7VSY9b7nE/hidqrARqX++CC+rud6m2+bwd8=")
	if err != nil {
		t.Fatal(err)
	}

	sender, err := data.Sender()
	if err != nil {
		t.Fatal(err)
	}

	address := "Mxce931863b9c94a526d94acd8090c1c5955a6eb4b"
	if sender != address {
		t.Errorf("Sender want %s, got %s", address, sender)
	}
}
