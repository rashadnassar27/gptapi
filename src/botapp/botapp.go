package botapp

import (
	"fmt"
	"gptapi/internal/storage/redis"
	"gptapi/internal/tbot"
	"gptapi/pkg/api/httpserver"
	"gptapi/pkg/utils"
	"os"
)

const rule3 = `you are a music and song suggestions machine,
 on any user quetion your answer must be only a list of 5 songs (excepts user requested specific number of songs) 
 related to question or text that user has sent to you.
 each song or music in list descibed by name and artist only. 
 Don't explain anything or `

type BotApp struct {
	server *httpserver.HttpServer
	cache  *redis.RedisClient
}

func NewBotAPP() *BotApp {
	b := &BotApp{}
	b.init()
	b.initRestAPI()
	return b
}

func (b *BotApp) init() {
	utils.LoadEnv("")
	redisHost := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	b.cache = redis.NewRedisClient(fmt.Sprintf(`%s:%s`, redisHost, port))
	b.server = httpserver.NewHttpServer()
}

func (h *BotApp) initRestAPI() {
	h.server.RegisterAction("GET", "/generate", h.generate)
}

func (h *BotApp) StartAPI(port string) {
	h.server.Start(port)
}

func (b *BotApp) Start() {
	botKey := os.Getenv("TELEGRAM_TEST_TOKEN")
	bot := tbot.NewTelegramBot(botKey, b.cache)
	bot.SetPrompt(rule3)
	bot.Start()
}

func (h *BotApp) generate(params map[string]string, queryString map[string][]string, bodyJson map[string]interface{}) (string, error) {
	return "", nil
}
