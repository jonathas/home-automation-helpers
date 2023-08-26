package models

type VoiceMonkey struct {
	Device string `json:"device" binding:"required"`
}
