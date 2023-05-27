package discord

type Discord interface {
	AddCommand(command Command)error
}
