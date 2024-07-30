package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/karokojnr/GoBuzz/internal/store"
)

var usernames = []string{
	"ken", "kelvin", "grace", "emmanuel", "alice", "bob", "charlie", "david",
	"eve", "frank", "george", "hannah", "isabella", "jack", "karen", "liam",
	"michael", "nina", "oliver", "paula", "quinn", "rachel", "steve", "tina",
	"ursula", "victor", "wendy", "xavier", "yara", "zane", "aaron", "betty",
	"carl", "diana", "ethan", "fiona", "gary", "helen", "ian", "julia", "kyle",
	"linda", "matt", "nora", "oscar", "pam", "quincy", "roger", "susan", "tony",
}

var titles = []string{
	"Understanding Go Routines",
	"Mastering Interfaces in Go",
	"Intro to Go Modules",
	"Building APIs with Go",
	"Concurrency in Go",
	"Go Error Handling",
	"Using Context in Go",
	"Effective Go Testing",
	"Go Data Structures",
	"Writing Middleware in Go",
	"Go for Web Development",
	"Deploying Go Apps",
	"Go and Docker",
	"Go and Microservices",
	"Go Best Practices",
	"Optimizing Go Performance",
	"Go Patterns and Practices",
	"Building CLI Tools in Go",
	"Go for Beginners",
	"Advanced Go Techniques",
}

var contents = []string{
	"In this post, we explore the concept of goroutines and how they enable concurrent programming in Go.",
	"Learn how to use interfaces in Go to write more flexible and modular code.",
	"This guide covers the basics of Go modules, including how to create and manage them.",
	"Discover the steps to build robust APIs with Go, including routing, middleware, and handlers.",
	"Concurrency is a powerful feature in Go. This article delves into concurrent programming patterns and best practices.",
	"Error handling is a critical part of any application. Here, we discuss Go's approach to managing errors.",
	"Understanding and using the context package is crucial for managing timeouts and cancellations in Go.",
	"Learn how to write effective tests for your Go code, including unit tests, integration tests, and benchmarks.",
	"This post covers common data structures in Go, including slices, maps, and structs, and how to use them.",
	"Middleware can help you manage cross-cutting concerns in your web applications. Learn how to create and use middleware in Go.",
	"Explore the tools and techniques for building web applications with Go, including templating and routing.",
	"Deploying Go applications can be straightforward. This article covers various deployment strategies and best practices.",
	"Learn how to containerize your Go applications using Docker and why it’s beneficial.",
	"Microservices architecture is popular for scalable applications. See how Go fits into this paradigm.",
	"Follow these best practices to write clean, maintainable, and efficient Go code.",
	"Performance optimization can make a significant difference in your Go applications. Discover tips and tricks for optimizing Go code.",
	"Go has many design patterns that can help you solve common problems. Learn about these patterns and how to implement them.",
	"Creating command-line interfaces (CLI) can be powerful. This guide shows you how to build CLI tools using Go.",
	"This beginner’s guide to Go will help you get started with the language and understand its basic concepts.",
	"Advanced Go techniques can help you push the boundaries of what’s possible with the language. Explore these advanced topics in this post.",
}

var tags = []string{
	"golang", "concurrency", "api", "web development", "docker", "microservices",
	"testing", "error handling", "data structures", "middleware", "performance",
	"best practices", "cli tools", "deployment", "go routines", "interfaces",
	"context", "modules", "patterns", "advanced go",
}

var comments = []string{
	"Great article! This really helped me understand Go routines.",
	"Thanks for the detailed explanation on interfaces. Very useful!",
	"I was struggling with Go modules, but this post made it much clearer.",
	"Building APIs with Go is so much easier now. Appreciate the guide!",
	"Concurrency was always confusing to me. This article simplified it.",
	"Error handling in Go makes more sense now. Thanks for the insights!",
	"The context package was a mystery to me until I read this. Great job!",
	"Effective testing strategies in Go were exactly what I needed. Thanks!",
	"Your explanation of Go data structures was spot on!",
	"Middleware in Go is something I needed to understand better. Great post!",
	"This guide to web development with Go is very comprehensive.",
	"Deploying Go apps seemed daunting, but this article helped a lot.",
	"Containerizing with Docker is much clearer now. Thanks!",
	"Microservices in Go is a game-changer. Excellent post!",
	"These best practices will definitely improve my Go code.",
	"Optimizing performance in Go was a challenge before. Thanks for the tips!",
	"Design patterns in Go are easier to implement with this guide.",
	"Building CLI tools in Go is now part of my skill set thanks to this post.",
	"As a beginner, this guide to Go was extremely helpful. Thanks!",
	"Advanced Go techniques were well explained. Great article!",
}

func Seed(store store.Storage) {

	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("Error creating user: ", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post: ", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment: ", err)
			return
		}
	}

	log.Println("Database seeded successfully!")
}

func generateUsers(n int) []*store.User {
	users := make([]*store.User, n)
	for i := 0; i < n; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "changeme",
		}
	}
	return users
}

func generatePosts(n int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, n)
	for i := 0; i < n; i++ {
		user := users[rand.Intn(len(users))]
		posts[i] = &store.Post{
			UserID:  user.ID,
			Version: 0,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}
	return posts
}

func generateComments(n int, users []*store.User, posts []*store.Post) []*store.Comment {
	cmts := make([]*store.Comment, n)
	for i := 0; i < n; i++ {
		cmts[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}
	return cmts
}
