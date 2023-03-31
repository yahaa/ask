package cmd

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"time"

	"github.com/yahaa/ask/kvdb"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var opt Option

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
		kv, err := kvdb.New(opt.DBPath())
		if err != nil {
			log.Fatalf("new kvdb err: %v", err)
		}

		defer kv.Close()

		if opt.History {
			history(kv, opt)
			return
		}

		if opt.SessionList {
			sessions(kv, opt)
			return
		}

		if len(args) <= 0 {
			log.Fatalf("args need to specify")
		}

		q := args[0]

		if len(opt.Translate) > 0 {
			q = fmt.Sprintf("Could you please help me translate this sentence '%s' to %s", q, opt.Translate)
		} else if opt.Polish {
			q = fmt.Sprintf("Could you please help me polish this sentence '%s'", q)
		} else if opt.Check {
			q = fmt.Sprintf("Could you please assist me in reviewing the grammar and spelling of the sentence '%s', and identify any existing errors within it?", q)
		}

		if opt.Debug {
			fmt.Printf("Q: %s\n", q)
		}

		client := openai.NewClient(opt.APIKey)

		req, err := makeChatReq(q, kv.Query(kvdb.QueryParams{Bucket: opt.Session}))
		if err != nil {
			log.Fatalf("make chat error: %v", err)
		}

		stream, err := client.CreateChatCompletionStream(context.TODO(), *req)
		if err != nil {
			log.Fatalf("ChatCompletionStream error: %v", err)
		}
		defer stream.Close()

		if opt.Debug {
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

		conver := kvdb.SaveParmas{
			Bucket: opt.Session,
			ChatContext: kvdb.ChatContext{
				Time: time.Now(),
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: q,
					},
					{
						Role:    openai.ChatMessageRoleAssistant,
						Content: respBuffer.String(),
					},
				},
			},
		}

		if err := kv.Save(conver); err != nil {
			log.Fatalf("kvdb save context err: %v", err)
		}

		fmt.Println()
	},
}

func makeChatReq(ask string, chatCtxs []kvdb.ChatContext) (*openai.ChatCompletionRequest, error) {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are a helpful assistant.",
		},
	}

	// append limit histry conversation context
	for _, item := range chatCtxs {
		messages = append(messages, item.Messages...)
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

func sessions(kv kvdb.Interface, _ Option) {
	keys := kv.Keys()
	for _, item := range keys {
		fmt.Printf("%v\n", item)
	}

	fmt.Printf("\nA total of %d chat sessions were requested from ChatGPT.\n", len(keys))
}

func history(kv kvdb.Interface, opt Option) {
	chatContexts := kv.Query(kvdb.QueryParams{
		Bucket: opt.Session,
		Limit:  opt.Limit,
	})

	for _, item := range chatContexts {
		if len(item.Messages) < 2 {
			continue
		}

		fmt.Printf("T: %v\n", item.Time.Local())
		fmt.Printf("Q: %v\n\n", item.Messages[0].Content)
		fmt.Printf("A: %v\n\n", item.Messages[1].Content)
	}

	fmt.Printf("\nPrint latest %d questions were asked to ChatGPT, In the %v session.\n", len(chatContexts), opt.Session)
}

func init() {
	curUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().BoolVar(&opt.History, "history", false, "print the history of ask")
	rootCmd.PersistentFlags().BoolVar(&opt.SessionList, "list", false, "print the list of sessions")
	rootCmd.PersistentFlags().BoolVarP(&opt.Check, "check", "c", false, "enable check grammar")
	rootCmd.PersistentFlags().BoolVarP(&opt.Debug, "debug", "d", false, "enable debug mode")
	rootCmd.PersistentFlags().BoolVarP(&opt.Polish, "polish", "p", false, "enable polish sentence")
	rootCmd.PersistentFlags().StringVarP(&opt.APIKey, "api-key", "k", os.Getenv("API_KEY"), "openai api key")
	rootCmd.PersistentFlags().StringVarP(&opt.Translate, "translate", "t", "", "translate to specify language")
	rootCmd.PersistentFlags().StringVarP(&opt.ConfigSavePath, "config", "f", fmt.Sprintf("%v/.ask", curUser.HomeDir), "config save path")
	rootCmd.PersistentFlags().StringVarP(&opt.Session, "session", "n", "default", "new conversation session")
	rootCmd.PersistentFlags().IntVarP(&opt.Limit, "limit", "l", 3, "limit the number of conversation history output")
}

func Execute() error {
	return rootCmd.Execute()
}
