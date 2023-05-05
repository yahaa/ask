# ask
[README](README.md) | [中文文档](README_zh.md)

ask 是 ChatGPT 的一个命令行工具，它允许你在命令行向 ChatGPT 问任何问题.

## 功能清单
- [x] 命令行直接向 ChatGPT 提问，避免每次打开浏览器都需要认证是不是人类。
- [x] 支持会话上下文推断
- [x] 支持查看历史记录
- [x] 支持语句翻译
- [x] 支持文本润色
- [x] 支持英语语法/拼写错误检测
- [x] 支持新建会话

> ask 不会记录用户任何数据到云端，所有数据只会保存在 `$HOME/.ask/` 目录下。

## golang 版本
go 1.18+

## 安装
```go
$ go install github.com/yahaa/ask@latest
```
## 使用演示

设置 API_KEY 环境变量:
```bash
export API_KEY="your openai api key"
```

设置 MODEL 环境变量(如果你的账号支持 'gpt-4',不设置将使用 'gpt-3.5-trubo'):
```bash
export MODEL="gpt-4"
```


查看帮助信息:
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
  -m, --model string       openai model, using 'gpt-3.5-turbo' if not specified (default "gpt-4")
  -p, --polish             enable polish sentence
  -n, --session string     new conversation session (default "default")
  -t, --translate string   translate to specify language

  ```



使用 `--debug(-d)` 参数:
```bash
ask "Are you here?" -d
```
  ```txt
Q: Are you here?
A: Yes, I am here. How can I assist you?
  ```


翻译句子:
```bash
ask "Are you here?" -t zh -d
```
  ```txt
Q: Please help me translate this sentence 'Are you here?' to zh
A: 你在这里吗？(nǐ zài zhèlǐ ma?)
  ```
 润色句子:
```bash
ask "Are you here?" -p -d
```
```txt
Q: Please help me polish this sentence 'Are you here?'
A: Is it safe to assume that you are present in this location?
```

语法拼写检测:
```bash
ask "Ask is a command line tool for ChatGPT that allows you to ask any question." -c
```
  ```txt
The sentence appears to be grammatically correct and there are no spelling errors. However, there is a minor issue with clarity in using "Ask" as a proper noun. It may be more clear to write "The 'Ask' tool is a command line tool for ChatGPT that allows you to ask any question."
  ```

打印提问历史:
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
创建新会话:
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
打印所有会话:
```bash
ask --list
```
```txt
default
kernel

A total of 2 chat sessions were requested from ChatGPT.
```

## 其他帮助?


* 更多功能请执行 `ask -h` 查看。
* 可以从 https://platform.openai.com/account/api-keys 获取你的 API_KEYS。
* 如果你没有账号，可以按照 chat GPT 提示主持一个账号，注册账号过程中需要提供手机号码和短信验证，你可以从这个网站获取一些帮助 https://sms-activate.org/ to receive assistance。
