package discord

import "os"

type discordImpl struct {
	AppID        string
	AppPublicKey string
	Token        string
	URL          string
}

func New() Discord {
	return &discordImpl{
		AppID:        os.Getenv("DISCORD_APP_ID"),
		AppPublicKey: os.Getenv("DISCORD_APP_PUBLIC_KEY"),
		Token:        os.Getenv("DISCORD_APP_TOKEN"),
		URL:          "https://discord.com/api/v10/applications",
	}
}
