package main

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v9"
	"strconv"
	"strings"
)

var ctx = context.Background()

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example, but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	authorId := m.Author.ID

	if getUserScore(authorId) < 100 {
		//s.ChannelMessageEdit(m.ChannelID, m.ID, "[censored due to low social points]")
		s.ChannelMessageDelete(m.ChannelID, m.ID)
		s.ChannelMessageSend(m.ChannelID, "[censored]")
		return
	}
	if getUserScore(authorId) < 200 {
		s.ChannelMessageSend(m.ChannelID, m.Author.Username+" you have less than 200 social points. Be careful")
		return
	}
	if getUserScore(authorId) < 300 {
		s.ChannelMessageSend(m.ChannelID, "This user has low social credit, be careful who you trust")
		return
	}

	message := m.Content
	if strings.Contains(message, "bad") {
		users := m.Mentions
		for _, user := range users {
			userId := user.ID
			if userId == m.Author.ID {
				continue
			}
			oldScore := getUserScore(userId)
			newScore := oldScore - 30
			setUserScore(userId, newScore)
			s.ChannelMessageSend(m.ChannelID, user.Username+" has now "+strconv.Itoa(newScore)+" points")
		}
	}
	if strings.Contains(message, "good") {
		users := m.Mentions
		for _, user := range users {
			userId := user.ID
			if userId == m.Author.ID {
				continue
			}
			oldScore := getUserScore(userId)
			newScore := oldScore + 10
			setUserScore(userId, newScore)
			s.ChannelMessageSend(m.ChannelID, user.Username+" has now "+strconv.Itoa(newScore)+" points")
		}
	}
	if strings.Contains(message, "spy") {
		users := m.Mentions
		for _, user := range users {
			userId := user.ID
			if userId == m.Author.ID {
				continue
			}
			oldScore := getUserScore(userId)
			newScore := oldScore - 200
			setUserScore(userId, newScore)
			s.ChannelMessageSend(m.ChannelID, user.Username+" has now "+strconv.Itoa(newScore)+" points")
		}
	}
	if strings.Contains(message, "communist") {
		users := m.Mentions
		for _, user := range users {
			userId := user.ID
			if userId == m.Author.ID {
				continue
			}
			oldScore := getUserScore(userId)
			newScore := oldScore + 100
			setUserScore(userId, newScore)
			s.ChannelMessageSend(m.ChannelID, user.Username+" has now "+strconv.Itoa(newScore)+" points")
		}
	}
}

func setUserScore(userId string, score int) {
	Rdb.Set(ctx, userId, score, 0)
}

func getUserScore(userId string) int {
	val, err := Rdb.Get(ctx, userId).Result()
	if err == redis.Nil {
		setUserScore(userId, 500)
		return 500
	}
	intVal, _ := strconv.Atoi(val)
	return intVal
}
