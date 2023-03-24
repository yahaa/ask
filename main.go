package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var (
	apiKey    string
	translate string
	polish    bool
)

var rootCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ask is a command line tool for ChatGPT that allows you to ask any question",
	Long: `Ask is a command line tool for ChatGPT that allows you to ask any question.

Examples:

$ ask "help write a hello world demo using golang"
$ ask "Ask is a command line tool for ChatGPT that allows you to ask any question." -t zh
$ ask "Ask is a command line tool for ChatGPT that allows you to ask any question." -p
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalf("args need to specify")
		}

		q := args[0]

		if len(translate) > 0 {
			q = fmt.Sprintf("Please help me translate this sentence '%s' to %s", q, translate)
		} else if polish {
			q = fmt.Sprintf("Please help me polish this sentence '%s'", q)
		}

		fmt.Printf("Q: %s\n", q)

		client := openai.NewClient(apiKey)

		req := openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: q,
				},
			},
			Stream: true,
		}
		stream, err := client.CreateChatCompletionStream(context.TODO(), req)
		if err != nil {
			log.Printf("ChatCompletionStream error: %v", err)
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
				log.Printf("stream response err: %v", err)
				return
			}

			fmt.Print(resp.Choices[0].Delta.Content)
		}

		fmt.Println()
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&polish, "polish", "p", false, "polishing sentence")
	rootCmd.PersistentFlags().StringVarP(&apiKey, "api-key", "k", os.Getenv("API_KEY"), "openai api key")
	rootCmd.PersistentFlags().StringVarP(&translate, "translate", "t", "", "translate to specify language")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("run err: %v", err)
		os.Exit(1)
	}
}
