# Begin

## This is a demo for wm-wallet-sdk

## How to enjoy it

```shell
cd test
go test -v -run TestDepositCreate   ./deposit_test.go             # create a order
go test -v -run TestDepositDetail   ./deposit_test.go             # get this order detail, you need modify code, the param source_id.
go test -v -run TestDepositCancel   ./deposit_test.go             # cancel a order, you need modify code, the param source_id.
go test -v -run TestWithdrawCreate  ./deposit_test.go             # withdraw from wm-wallet, you need modify code, the param source_id.
go test -v -run TestWithdrawDetail  ./deposit_test.go             # get this withdraw order detail, you need modify code, the param source_id.
```

# End
