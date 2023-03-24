package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/sashabaranov/go-openai"
	"k8s.io/klog"
)

var (
	apiKey string
	help   bool
)

var helpMsg = `
Ask is a command line tool for ChatGPT that allows you to ask any question.

Usage:
$ ask "help write a hello world demo using golang"
`

func init() {
	flag.BoolVar(&help, "help", false, "help")
	flag.StringVar(&apiKey, "api-key", os.Getenv("API_KEY"), "openai api key")
	flag.Parse()
}

func main() {
	if help {
		fmt.Println(helpMsg)
		os.Exit(0)
	}

	if len(os.Args) <= 1 {
		klog.Fatalf("args need to specify")
	}

	args := os.Args[1]

	fmt.Printf("Q: %s\n", args)

	client := openai.NewClient(apiKey)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: args,
			},
		},
		Stream: true,
	}
	stream, err := client.CreateChatCompletionStream(context.TODO(), req)
	if err != nil {
		klog.Errorf("ChatCompletionStream error: %v", err)
		os.Exit(1)
	}
	defer stream.Close()

	fmt.Print("A: ")
	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			klog.Errorf("stream response err: %v", err)
			return
		}

		fmt.Print(resp.Choices[0].Delta.Content)
	}

	fmt.Println()
}
