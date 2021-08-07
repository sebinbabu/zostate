package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sebinbabu/zostate"
)

const (
	Start           zostate.StateType = "Start"
	EnterPIN        zostate.StateType = "EnterPIN"
	SelectOperation zostate.StateType = "SelectOperation"
	WaitForDeposit  zostate.StateType = "WaitForDeposit"
	EnterAmount     zostate.StateType = "EnterAmount"
	ConfirmAmount   zostate.StateType = "ConfirmAmount"
	WithdrawSuccess zostate.StateType = "WithdrawSuccess"
	DepositSuccess  zostate.StateType = "DepositSuccess"
)

const (
	CardInserted       zostate.EventType = "CardInserted"
	BadPINEntered      zostate.EventType = "BadPINEntered"
	CorrectPINEntered  zostate.EventType = "CorrectPINEntered"
	DepositSelected    zostate.EventType = "DepositSelected"
	WithdrawalSelected zostate.EventType = "WithdrawalSelected"
	AmountEntered      zostate.EventType = "AmountEntered"
	AmountDeposited    zostate.EventType = "AmountDeposited"
	ConfirmedYes       zostate.EventType = "ConfirmedYes"
	ConfirmedNo        zostate.EventType = "ConfirmedNo"
	CardWithdrawn      zostate.EventType = "CardWithdrawn"
)

func main() {
	machine, err := zostate.NewMachine(
		"ATM",
		Start,
		zostate.States{
			{
				Name: Start,
				Transitions: zostate.Transitions{
					{Event: CardInserted, Dst: EnterPIN},
				},
			},
			{
				Name: EnterPIN,
				Transitions: zostate.Transitions{
					{Event: BadPINEntered, Dst: Start},
					{Event: CorrectPINEntered, Dst: SelectOperation},
				},
			},
			{
				Name: SelectOperation,
				Transitions: zostate.Transitions{
					{Event: DepositSelected, Dst: WaitForDeposit},
					{Event: WithdrawalSelected, Dst: EnterAmount},
				},
			},
			{
				Name: WaitForDeposit,
				Transitions: zostate.Transitions{
					{Event: AmountDeposited, Dst: DepositSuccess},
				},
			},
			{
				Name: EnterAmount,
				Transitions: zostate.Transitions{
					{Event: AmountEntered, Dst: ConfirmAmount},
				},
			},
			{
				Name: ConfirmAmount,
				Transitions: zostate.Transitions{
					{Event: ConfirmedYes, Dst: WithdrawSuccess},
					{Event: ConfirmedNo, Dst: EnterAmount},
				},
			},
			{
				Name: WithdrawSuccess,
				Transitions: zostate.Transitions{
					{Event: CardWithdrawn, Dst: Start},
				},
			},
			{
				Name: DepositSuccess,
				Transitions: zostate.Transitions{
					{Event: CardWithdrawn, Dst: Start},
				},
			},
		},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintln(os.Stderr, "current state: ", machine.Current())

	dot := zostate.DrawMachine(machine)
	io.Copy(os.Stdout, strings.NewReader(dot))
}
