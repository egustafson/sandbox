package main

import (
	"fmt"
	"os"

	"github.com/aclindsa/ofxgo"
)

func main() {
	f, err := os.Open("/tmp/demo.qfx")
	if err != nil {
		fmt.Printf("can't open file: %v\n", err)
		return
	}
	defer f.Close()

	resp, err := ofxgo.ParseResponse(f)
	if err != nil {
		fmt.Printf("can't parse qfx file: %v\n", err)
		return
	}

	if stmt, ok := resp.Bank[0].(*ofxgo.StatementResponse); ok {
		fmt.Printf("Balance: %s %s (as of %s)\n", stmt.BalAmt, stmt.CurDef, stmt.DtAsOf)
		fmt.Println("Transactions:")
		for _, tran := range stmt.BankTranList.Transactions {
			currency := stmt.CurDef
			if tran.Currency != nil {
				if ok, _ := tran.Currency.Valid(); ok {
					currency = tran.Currency.CurSym
				}
			}
			fmt.Printf("%s %-15s %-11s %s | %s\n",
				tran.DtPosted,
				tran.TrnAmt.String()+" "+currency.String(),
				tran.TrnType,
				tran.Name,
				//tran.Payee.Name,
				tran.Memo,
			)
		}
	}
	_ = resp

	fmt.Println("done.")
}
