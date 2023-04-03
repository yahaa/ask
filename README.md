# ask
[README](README.md) | [中文文档](README_zh.md)

Ask is a command line tool for ChatGPT that allows you to ask any question.


## Feature list

- [x] Directly ask ChatGPT in the command line, avoiding the need to authenticate as a human every time you open your browser.
- [x] Support for inferring session context
- [x] Support for viewing history records
- [x] Support for sentence translation
- [x] Support for text enhancement
- [x] Support for English grammar/spelling error detection
- [x] Support for creating new sessions

> ask will not record any user data to the cloud. All data will only be saved in the `$HOME/.ask/`.

## golang version
go 1.18+

## install
```bash
go install github.com/yahaa/ask@latest
```
## Usage

Setup env:
```bash
export API_KEY="your openai api key"
```

Run help to see more features:
```bash
ask -h 
```
  ```txt                                                                                
Ask is a command line tool for ChatGPT that allows you to ask any question.

Examples:

$ ask "help me write a hello world demo using golang"
$ ask "Ask is a command line tool for ChatGPT that allows you to ask any question." -t zh
$ ask "Ask is a command line tool for ChatGPT that allows you to ask any question." -p
$ ask "Ask is a command line tool for ChatGPT that allows you to ask any question." -c

Usage:
  ask [flags]

Flags:
  -k, --api-key string     openai api key (default "sk-xxx")
  -c, --check              enable check grammar
  -f, --config string      config save path (default "~/.ask")
  -d, --debug              enable debug mode
  -h, --help               help for ask
      --history            print the history of ask
  -l, --limit int          limit the number of conversation history output (default 3)
      --list               print the list of sessions
  -p, --polish             enable polish sentence
  -n, --session string     new conversation session (default "default")
  -t, --translate string   translate to specify language

  ```



Ask with `--debug(-d)` flag:
```bash
ask "Are you here?" -d
```
  ```txt
Q: Are you here?
A: Yes, I am here. How can I assist you?
  ```


Translate sentence:
```bash
ask "Are you here?" -t zh -d
```
  ```txt
Q: Please help me translate this sentence 'Are you here?' to zh
A: 你在这里吗？(nǐ zài zhèlǐ ma?)
  ```
Polish your sentence:
```bash
ask "Are you here?" -p -d
```
```txt
Q: Please help me polish this sentence 'Are you here?'
A: Is it safe to assume that you are present in this location?
```

Check grammar or spelling:
```bash
ask "Ask is a command line tool for ChatGPT that allows you to ask any question." -c
```
  ```txt
The sentence appears to be grammatically correct and there are no spelling errors. However, there is a minor issue with clarity in using "Ask" as a proper noun. It may be more clear to write "The 'Ask' tool is a command line tool for ChatGPT that allows you to ask any question."
  ```

Print ask history:
```bash
ask --history
```
```txt
T: 2023-03-31 23:59:55 +0800 CST
Q: How to keep context when using openai open api

A: There are a few strategies you can use to keep context when using OpenAI's API:

1. Use the "context" parameter: OpenAI's API allows you to provide context by passing in a string of text as the "context" parameter. This tells the model what information it should use to generate its response. Make sure to provide as much relevant context as possible to get the best results.

2. Use the "stream" parameter: If your context is too long to fit into a single API request, you can use the "stream" parameter to provide it in chunks. This allows you to maintain a continuous context throughout multiple requests.

3. Use the "temperature" parameter: The temperature parameter controls the randomness of the model's responses. By setting a lower temperature, the model will generate responses that are more predictable and closely related to the provided context.

4. Use the "max_tokens" parameter: The max_tokens parameter limits the number of tokens the model will generate for its response. By setting a higher max_tokens value, the model can generate longer, more detailed responses that are still closely related to the provided context.

Keep in mind that OpenAI's API is a powerful tool, but it's still an artificial intelligence system that relies on statistical patterns in large amounts of data. It may not always provide the exact response you're looking for, but by providing relevant context and fine-tuning your parameters, you can increase your chances of getting a useful and accurate response from the model.


Print latest 1 questions were asked to ChatGPT, In the default session.

```
Ask with new session:
```bash
ask "How to learn kernel develop?" -n kernel
```
```txt
To learn kernel development, you will need to have a strong understanding of the C programming language, as well as knowledge of computer architecture and operating system principles.

Here are some steps to guide you in learning kernel development:

1. Start with the basics of operating system principles and computer architecture. You can find online resources, such as books and video tutorials, to start learning this.

2. Learn C programming language. Get comfortable with pointers, structures, memory management, and other fundamental concepts.

3. Study the source code of an operating system kernel, such as Linux or FreeBSD. It will help you understand how operating system components work together.

4. Attend workshops or online courses on kernel development. You can find many resources provided by universities, open-source communities, and technology companies.

5. Start with simple kernel programming exercises, such as developing a device driver or system call. It will help you become more familiar with kernel programming concepts.

6. Contribute to open-source kernel projects, such as Linux or FreeBSD. This will help you learn from experienced developers and gain more practical experience.

With patience, determination, and a lot of practice, you can become an accomplished kernel developer.
```
Print all sessions:
```bash
ask --list
```
```txt
default
kernel

A total of 2 chat sessions were requested from ChatGPT.
```

## Help?


* For more features, please run `ask -h` to view additional information. 
* You can get OpenAI `API_KEY` from here https://platform.openai.com/account/api-keys 
* If you do not have an OpenAI account, simply sign up for one and visit https://sms-activate.org/ to receive assistance.
