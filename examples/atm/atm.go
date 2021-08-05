package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sebinbabu/zostate"
)

const (
	Start           zostate.StateType = "START"
	EnterPIN        zostate.StateType = "ENTER_PIN"
	SelectOperation zostate.StateType = "SELECT_OPERATION"
	WaitForDeposit  zostate.StateType = "WAIT_FOR_DEPOSIT"
	EnterAmount     zostate.StateType = "ENTER_AMOUNT"
	ConfirmAmount   zostate.StateType = "CONFIRM_AMOUNT"
	WithdrawSuccess zostate.StateType = "WITHDRAW_SUCCESS"
	DepositSuccess  zostate.StateType = "DEPOSIT_SUCCESS"
)

const (
	CardInserted       zostate.EventType = "CARD_INSERTED"
	BadPINEntered      zostate.EventType = "BAD_PIN_ENTERED"
	CorrectPINEntered  zostate.EventType = "CORRECT_PIN_ENTERED"
	DepositSelected    zostate.EventType = "DEPOSIT_SELECTED"
	WithdrawalSelected zostate.EventType = "WITHDRAWAL_SELECTED"
	AmountEntered      zostate.EventType = "AMOUNT_ENTERED"
	AmountDeposited    zostate.EventType = "AMOUNT_DEPOSITED"
	ConfirmedYes       zostate.EventType = "CONFIRMED_YES"
	ConfirmedNo        zostate.EventType = "CONFIRMED_NO"
	CardWithdrawn      zostate.EventType = "CARD_WITHDRAWN"
)

func main() {
	machine, err := zostate.NewMachine(
		"ATM",
		Start,
		zostate.States{
			{
				Name: Start,
				Transitions: zostate.Transitions{
					{Name: CardInserted, Dst: EnterPIN},
				},
			},
			{
				Name: EnterPIN,
				Transitions: zostate.Transitions{
					{Name: BadPINEntered, Dst: Start},
					{Name: CorrectPINEntered, Dst: SelectOperation},
				},
			},
			{
				Name: SelectOperation,
				Transitions: zostate.Transitions{
					{Name: DepositSelected, Dst: WaitForDeposit},
					{Name: WithdrawalSelected, Dst: EnterAmount},
				},
			},
			{
				Name: WaitForDeposit,
				Transitions: zostate.Transitions{
					{Name: AmountDeposited, Dst: DepositSuccess},
				},
			},
			{
				Name: EnterAmount,
				Transitions: zostate.Transitions{
					{Name: AmountEntered, Dst: ConfirmAmount},
				},
			},
			{
				Name: ConfirmAmount,
				Transitions: zostate.Transitions{
					{Name: ConfirmedYes, Dst: WithdrawSuccess},
					{Name: ConfirmedNo, Dst: EnterAmount},
				},
			},
			{
				Name: WithdrawSuccess,
				Transitions: zostate.Transitions{
					{Name: CardWithdrawn, Dst: Start},
				},
			},
			{
				Name: DepositSuccess,
				Transitions: zostate.Transitions{
					{Name: CardWithdrawn, Dst: Start},
				},
			},
		},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintln(os.Stderr, "current", machine.Current())

	dot := zostate.DrawMachine(machine)
	io.Copy(os.Stdout, strings.NewReader(dot))
}
