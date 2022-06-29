package main

import (
	"os"

	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot/handlers"
	"github.com/PaulSonOfLars/gotgbot/handlers/Filters"

	"github.com/PaulSonOfLars/gotgbot"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	log := zap.NewProductionEncoderConfig()
	log.EncodeLevel = zapcore.CapitalLevelEncoder
	log.EncodeTime = zapcore.RFC3339TimeEncoder

	logger := zap.New(zap.NewCore(zapcore.NewConsoleEncoder(log), os.Stdout, zap.InfoLevel))
	updater, err := gotgbot.NewUpdater(logger, "5570548452:AAFT5hAdoBAYodEA23BFMRscbR5dQOOaucA")
	if err != nil {
		logger.Panic("UPDATER FAILED TO START")
		return
	}
	logger.Sugar().Info("UPDATER STARTED SUCCESFULLY")
	updater.StartCleanPolling()
	updater.Dispatcher.AddHandler(handlers.NewsMessage(Filters.Text, echo))
}

func echo(b ext.Bot, u *gotgbot.Update) error {
	b.SendMessage(u.EffectiveChat.Id, u, EffectiveMessage.Text)
	return nil
}
