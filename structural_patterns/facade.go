package structural_patterns

import (
	"fmt"
	"log"
)

type account struct {
	name string
}

func newAccount(name string) *account {
	return &account{name}
}

func (act *account) checkAccount(id string) error {
	if act.name != id {
		return fmt.Errorf("account name is incorrect")
	}

	fmt.Println("Account verified")
	return nil
}

type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{code}
}

func (sc *securityCode) checkCode(code int) error {
	if sc.code != code {
		return fmt.Errorf("invalid security code")
	}

	fmt.Println("Security code verified")
	return nil
}

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{0}
}

func (w *wallet) creditBalance(amt int) {
	w.balance += amt
	fmt.Println("Balance credited successfully.")
}

func (w *wallet) debitBalance(amt int) error {
	if w.balance < amt {
		return fmt.Errorf("insufficient balance in wallet")
	}

	w.balance -= amt
	fmt.Println("Balance debited successfully.")
	return nil
}

type ledger struct {
	records []string
}

func newLedger() *ledger {
	return &ledger{}
}

func (l *ledger) makeEntry(accountId string, amt int, txnType string) {
	record := fmt.Sprintf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountId, txnType, amt)

	l.records = append(l.records, record)
}

type notification struct {}

func (n *notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}

type walletFacade struct {
	account *account
	securityCode *securityCode
	wallet *wallet
	ledger *ledger
	notification *notification
}

func newWalletFacade(accountId string, code int) *walletFacade {
	fmt.Println("Start creating account.")

	walletFacade := &walletFacade{
		account: newAccount(accountId),
		securityCode: newSecurityCode(code),
		wallet: newWallet(),
		ledger: newLedger(),
		notification: &notification{},
	}

	fmt.Println("Account created successfully.")
	
	return walletFacade
}

func (wf *walletFacade) addMoneyToWallet(account string, code, amt int) error {
	fmt.Println("Start adding money to wallet")

	err := wf.account.checkAccount(account)
	if err != nil {
		return err
	}

	err = wf.securityCode.checkCode(code)
	if err != nil {
		return err
	}

	wf.wallet.creditBalance(amt)
	wf.notification.sendWalletCreditNotification()
	wf.ledger.makeEntry(account, amt, "credit")
	return nil
}

func (wf *walletFacade) deductMoneyFromWallet(account string, code, amt int) error {
	fmt.Println("Start deduction money from wallet")

	err := wf.account.checkAccount(account)
	if err != nil {
		return err
	}

	err = wf.securityCode.checkCode(code)
	if err != nil {
		return err
	}

	err = wf.wallet.debitBalance(amt)
	if err != nil {
		return err
	}
	wf.notification.sendWalletDebitNotification()
	wf.ledger.makeEntry(account, amt, "debit")
	return nil
}

func FacadePattern() {
	wallet := newWalletFacade("abc", 1234)
	fmt.Println()

	err := wallet.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	fmt.Println()

	err = wallet.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
	
}