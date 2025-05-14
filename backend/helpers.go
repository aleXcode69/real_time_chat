// package main

// import (
// 	"log"
// )

// func canUserSubscribe(userID string, channel string) bool {
// 	log.Printf("Checking if user %s can subscribe to channel %s", userID, channel)

// 	return true
// }
// func canUserPublish(userID string, channel string) bool {
// 	log.Printf("Checking if user %s can publish to channel %s", userID, channel)
// 	return true
// }



package main

import (
    "context"
    "log"

    "github.com/google/uuid"
)

var queries *Queries

func InitDB(q *Queries) {
    queries = q
}

func ensureUserAndChannel(uid uuid.UUID, channel string) {
    // Создать пользователя, если нет
    if err := queries.EnsureUserSqlc(context.Background(), uid); err != nil {
        log.Printf("DB error in EnsureUserSqlc: %v", err)
    }
    // Создать канал, если нет
    if err := queries.EnsureChannelSqlc(context.Background(), channel); err != nil {
        log.Printf("DB error in EnsureChannelSqlc: %v", err)
    }
}

func canUserSubscribe(userID string, channel string) bool {
    log.Printf("Checking if user %s can subscribe to channel %s", userID, channel)
    uid, err := uuid.Parse(userID)
    if err != nil {
        log.Printf("Invalid userID: %v", err)
        return false
    }
    ensureUserAndChannel(uid, channel)
    ok, err := queries.CanUserSubscribeSqlc(
        context.Background(),
        CanUserSubscribeSqlcParams{
            UserID: uid,
            Name:   channel,
        },
    )
    if err != nil {
        log.Printf("DB error in canUserSubscribe: %v", err)
        return false
    }
    return ok.Valid && ok.Bool
}

func canUserPublish(userID string, channel string) bool {
    log.Printf("Checking if user %s can publish to channel %s", userID, channel)
    uid, err := uuid.Parse(userID)
    if err != nil {
        log.Printf("Invalid userID: %v", err)
        return false
    }
    ensureUserAndChannel(uid, channel)
    ok, err := queries.CanUserPublishSqlc(
        context.Background(),
        CanUserPublishSqlcParams{
            UserID: uid,
            Name:   channel,
        },
    )
    if err != nil {
        log.Printf("DB error in canUserPublish: %v", err)
        return false
    }
    return ok.Valid && ok.Bool
}