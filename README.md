![Minter Logo](https://github.com/MinterTeam/minter-go-sdk/raw/v2/minter-logo.svg?sanitize=true)

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/mod/github.com/MinterTeam/minter-go-sdk/v2)

## Overview

This is a pure Go SDK for working with **Minter** blockchain

## Table of contents

* [Installing](#installing)
* [Minter API](#minter-api)
    - [Using API v2](#using-api-v2)
    - [Using gRPC](#using-grpc)
* [Using Transactions](#using-transactions)
    - [Sign transaction](#sign-transaction)
        - [Single signature](#single-signature)
        - [Multi signature](#multi-signatures)
* [Minter Wallet](#minter-wallet)

### Installing

```bash
go get github.com/MinterTeam/minter-go-sdk/v2
```

### Minter API

```go

```

#### Using API v2

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/MinterTeam/minter-go-sdk/v2/api/http_client?tab=doc)

#### Using gRPC

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/MinterTeam/minter-go-sdk/v2/api/grpc_client?tab=doc)

### Using Transactions

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/MinterTeam/minter-go-sdk/v2/transaction?tab=doc)

#### Sign transaction

Returns a signed tx.

⚠️ After sending the transaction, to make sure that the transaction was successfully committed on the blockchain, you
need to find the transaction by hash and make sure that the status code is 0.

##### Single signature

###### Example

```go
tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(
    transaction.NewSendData().
        SetCoin(0).
        SetValue(transaction.BipToPip(big.NewInt(1))).
        MustSetTo("Mx1b685a7c1e78726c48f619c497a07ed75fe00483"),
)

signedTransaction, _ := tx.
    SetGasPrice(1).
    SetGasCoin(0).
    SetNonce(1).
    Sign("07bc17abdcee8b971bb8723e36fe9d2523306d5ab2d683631693238e0f9df142")

encode, _ := signedTransaction.Encode()
```

##### Multi signatures

###### Example

```go
tx, _ := transaction.NewBuilder(transaction.TestNetChainID).NewTransaction(
    transaction.NewSendData().
        SetCoin(0).
        SetValue(transaction.BipToPip(big.NewInt(1))).
        MustSetTo("Mx1b685a7c1e78726c48f619c497a07ed75fe00483")
)

signedTx, _ := tx.SetNonce(1).SetGasPrice(1).SetGasCoin(0).SetSignatureType(transaction.SignatureTypeMulti).Sign(
    multisigAddress,
    "ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63",
    "b0a65cd84d57189b70d80fe0b3d5fa3ea6e02fa48041314a587a1f8fdba703d7",
    "4c8dbfb3258f383adf656c2131e5ed77ec482a36125db71fb49d29e0528ff2ba",
)

encode, _ := signedTx.Encode()
```

You can transfer the transaction to the remaining addresses

```go
signedTx1, _ := tx.Sign(msigAddress, privateKey1)
encode, _ := signedTx.Encode()
// transfer encode transaction
signedTx1, _ = transaction.Decode(encode)
// and continue its signature by the remaining participants
signedTx12, _ := signedTx1.Sign(msigAddress, privateKey2)
signedTx123, _ := signedTx12.Sign(msigAddress, privateKey3)
encode, _ := signedTx123.Encode()
```

You can collect all signatures in one place without revealing the private key

```go
signedTx1, _ := tx.Clone().Sign(msigAddress, "ae089b32e4e0976ca6888cb1023148bd1a9f1cc28c5d442e52e586754ff48d63")
signedTx2, _ := tx.Clone().Sign(msigAddress, "b0a65cd84d57189b70d80fe0b3d5fa3ea6e02fa48041314a587a1f8fdba703d7")
signedTx3, _ := tx.Clone().Sign(msigAddress, "4c8dbfb3258f383adf656c2131e5ed77ec482a36125db71fb49d29e0528ff2ba")
simpleSignatureData1, _ := signedTx1.SingleSignatureData()
simpleSignatureData2, _ := signedTx2.SingleSignatureData()
simpleSignatureData3, _ := signedTx3.SingleSignatureData()
signedTransaction, _ := tx.Clone().Sign(msigAddress)
signedTx123, _ := signedTransaction.AddSignature(simpleSignatureData1, simpleSignatureData2, simpleSignatureData3)

encode, _ := signedTx123.Encode()
```

### Minter Wallet

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/MinterTeam/minter-go-sdk/v2/wallet?tab=doc)
