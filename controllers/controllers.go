package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jonathas/home-automation-helpers/models"
)

func CallVoiceMonkeyPost(c *gin.Context) {
	var voiceMonkey models.VoiceMonkey

	if err := c.ShouldBindJSON(&voiceMonkey); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	callVoiceMonkey(c, voiceMonkey)
}

func CallVoiceMonkeyGet(c *gin.Context) {
	var voiceMonkey models.VoiceMonkey
	device := c.Query("device")

	if device == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Device is required"})
		return
	}

	voiceMonkey.Device = device

	callVoiceMonkey(c, voiceMonkey)
}

func callVoiceMonkey(c *gin.Context, voiceMonkey models.VoiceMonkey) {
	url := os.Getenv("VOICEMONKEY_URL")
	token := os.Getenv("VOICEMONKEY_TOKEN")
	test := url + "/trigger?token=" + token + "&device=" + voiceMonkey.Device

	resp, err := http.Get(test)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resp.StatusCode != 200 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error calling Voice Monkey"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "VoiceMonkey called successfully"})
}

func EndpointNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Endpoint not found!"})
}
