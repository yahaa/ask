# ask
Ask is a command line tool for ChatGPT that allows you to ask any question.

## golang version
go 1.18+

## install
```go
$ go install github.com/yahaa/ask@latest

```
## Usage

```bash
$ export API_KEY="your openai api key"

$ ask -h                                                                                 
Ask is a command line tool for ChatGPT that allows you to ask any question.

Examples:
$ ask "help me write a hello world demo using golang"
$ ask "Ask is a command line tool for ChatGPT that allows you to ask any question." -t zh
$ ask "Ask is a command line tool for ChatGPT that allows you to ask any question." -p

Usage:
  ask [flags]

Flags:
  -k, --api-key string     openai api key
  -h, --help               help for ask
  -p, --polish             polishing sentence
  -t, --translate string   translate to specify language


$ ask "Are you here?"
Q: Are you here?
A: Yes, I am here. How can I assist you?


$ ask "Are you here?" -t zh
Q: Please help me translate this sentence 'Are you here?' to zh
A: 你在这里吗？(nǐ zài zhèlǐ ma?)

$ ask "Are you here?" -p
Q: Please help me polish this sentence 'Are you here?'
A: Is it safe to assume that you are present in this location?

```

## Help?


* For more features, please run `ask -h` to view additional information. 
* You can get OpenAI `API_KEY` from here https://platform.openai.com/account/api-keys 
* If you do not have an OpenAI account, simply sign up for one and visit https://sms-activate.org/ to receive assistance.
