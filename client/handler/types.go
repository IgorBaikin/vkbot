package handler

import "encoding/json"

type MessageResponse struct {
	Message Message `json:"message"`
}

type Message struct {
	Payload string `json:"payload"`
	PeerID  int    `json:"peer_id"`
	Text    string `json:"text"`
}

type LongPoolServerResponse struct {
	Response LongPoolServer `json:"response"`
}

type LongPoolServer struct {
	Key    string `json:"key"`
	Server string `json:"server"`
	Ts     string `json:"ts"`
}

type Response struct {
	Ts      string   `json:"ts"`
	Updates []Events `json:"updates"`
	Failed  int      `json:"failed"`
}

type Events struct {
	Type    string          `json:"type"`
	Object  json.RawMessage `json:"object"`
	GroupID int             `json:"group_id"`
	EventID string          `json:"event_id"`
	V       string          `json:"v"`
}
