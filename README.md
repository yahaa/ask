# ask
Ask is a command line tool for ChatGPT that allows you to ask any question.

## golang version
go 1.18+

## install
```go
$ go install github.com/yahaa/ask@v0.1.0

```
## Usage

```bash
$ export API_KEY="your openai api key"

$ ask -h                                                                                 
Ask is a command line tool for ChatGPT that allows you to ask any question.

Examples:

$ ask "help write a hello world demo using golang"
$ ask "Ask is a command line tool for ChatGPT that allows you to ask any question." -t zh
$ ask "Ask is a command line tool for ChatGPT that allows you to ask any question." -p
$ ask "Ask is a command line tool for ChatGPT that allows you to ask any question." -c

Usage:
  ask [flags]

Flags:
  -k, --api-key string     openai api key
  -c, --check              enable check grammar
  -d, --debug              enable debug mode
  -h, --help               help for ask
  -p, --polish             enable polish sentence
  -t, --translate string   translate to specify language




$ ask "Are you here?" -d
Q: Are you here?
A: Yes, I am here. How can I assist you?


$ ask "Are you here?" -t zh -d
Q: Please help me translate this sentence 'Are you here?' to zh
A: 你在这里吗？(nǐ zài zhèlǐ ma?)

$ ask "Are you here?" -p -d
Q: Please help me polish this sentence 'Are you here?'
A: Is it safe to assume that you are present in this location?

$ ask "Ask is a command line tool for ChatGPT that allows you to ask any question." -c

The sentence appears to be grammatically correct and there are no spelling errors. However, there is a minor issue with clarity in using "Ask" as a proper noun. It may be more clear to write "The 'Ask' tool is a command line tool for ChatGPT that allows you to ask any question."


```

## Help?


* For more features, please run `ask -h` to view additional information. 
* You can get OpenAI `API_KEY` from here https://platform.openai.com/account/api-keys 
* If you do not have an OpenAI account, simply sign up for one and visit https://sms-activate.org/ to receive assistance.
