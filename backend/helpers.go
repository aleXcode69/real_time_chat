package main

import (
	"log"
)

func canUserSubscribe(userID string, channel string) bool {
	log.Printf("Checking if user %s can subscribe to channel %s", userID, channel)

	return true
}
func canUserPublish(userID string, channel string) bool {
	log.Printf("Checking if user %s can publish to channel %s", userID, channel)
	return true
}
