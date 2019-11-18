package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type MessageJson struct {
	Content string `json:"content"`
}

func SendMessage(serverID string, secert string, message string) error {
	messageJSON := &MessageJson{Content: message}
	url := "https://discordapp.com/api/webhooks/" + serverID + "/" + secert
	messageBytes, err := json.Marshal(messageJSON)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(messageBytes))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil

}
