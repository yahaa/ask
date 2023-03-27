package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var (
	apiKey         string
	translate      string
	polish         bool
	debug          bool
	check          bool
	configSavePath string
)

var contextMessages = make([]openai.ChatCompletionMessage, 0)

var rootCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ask is a command line tool for ChatGPT that allows you to ask any question",
	Long: `Ask is a command line tool for ChatGPT that allows you to ask any question.

Examples:

$ ask "help me write a hello world demo using golang"
$ ask "Ask is a command line tool for ChatGPT that allows you to ask any question." -t zh
$ ask "Ask is a command line tool for ChatGPT that allows you to ask any question." -p
$ ask "Ask is a command line tool for ChatGPT that allows you to ask any question." -c
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalf("args need to specify")
		}

		q := args[0]

		if len(translate) > 0 {
			q = fmt.Sprintf("Could you please help me translate this sentence '%s' to %s", q, translate)
		} else if polish {
			q = fmt.Sprintf("Could you please help me polish this sentence '%s'", q)
		} else if check {
			q = fmt.Sprintf("Could you please assist me in reviewing the grammar and spelling of the sentence '%s', and identify any existing errors within it?", q)
		}

		if debug {
			fmt.Printf("Q: %s\n", q)
		}

		client := openai.NewClient(apiKey)

		req, err := makeChatReq(q)
		if err != nil {
			log.Printf("make chat error: %v", err)
			os.Exit(1)
		}

		stream, err := client.CreateChatCompletionStream(context.TODO(), *req)
		if err != nil {
			log.Printf("ChatCompletionStream error: %v", err)
			os.Exit(1)
		}
		defer stream.Close()

		if debug {
			fmt.Print("A: ")
		}

		respBuffer := bytes.Buffer{}

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
			respBuffer.WriteString(resp.Choices[0].Delta.Content)
		}

		fmt.Println()

		if err := saveContext(q, respBuffer.String()); err != nil {
			log.Printf("save context err: %v", err)
		}
	},
}

func saveContext(ask, ans string) error {
	contextMessages = append(contextMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: ask,
	})
	contextMessages = append(contextMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: ans,
	})

	data, err := json.Marshal(contextMessages)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(configSavePath, 0755); err != nil {
		return err
	}

	saveAs := fmt.Sprintf("%s/context.json", configSavePath)
	f, err := os.Create(saveAs)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	return err
}

func loadContext() {
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/context.json", configSavePath))
	if err != nil {
		return
	}

	_ = json.Unmarshal(data, &contextMessages)
}

func makeChatReq(ask string) (*openai.ChatCompletionRequest, error) {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are a helpful assistant.",
		},
	}

	// append limit histry conversation context
	if len(contextMessages) > 6 {
		messages = append(messages, contextMessages[len(contextMessages)-6:]...)
	} else if len(contextMessages) > 0 {
		messages = append(messages, contextMessages...)
	}

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: ask,
	})

	req := openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: messages,
		Stream:   true,
	}

	return &req, nil
}

func init() {
	curUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().BoolVarP(&check, "check", "c", false, "enable check grammar")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug mode")
	rootCmd.PersistentFlags().BoolVarP(&polish, "polish", "p", false, "enable polish sentence")
	rootCmd.PersistentFlags().StringVarP(&apiKey, "api-key", "k", os.Getenv("API_KEY"), "openai api key")
	rootCmd.PersistentFlags().StringVarP(&translate, "translate", "t", "", "translate to specify language")
	rootCmd.PersistentFlags().StringVarP(&configSavePath, "config", "f", fmt.Sprintf("%v/.ask", curUser.HomeDir), "config save path")

	loadContext()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("run err: %v", err)
		os.Exit(1)
	}
}
