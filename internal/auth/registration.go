package auth

import (
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	repository "github.com/gpunker/rss-telegram-bot/internal/repositories"

	types "github.com/gpunker/rss-telegram-bot/internal"
)


func RegisterUser(telegramUser *tgbotapi.User) (bool, error) {
	// через репозиторий дергаем пользователя по id и username
	params := repository.UserParams{
		TelegramID: types.Ptr(telegramUser.ID),
		UserName: types.Ptr(telegramUser.UserName),
	}
	_, err := repository.FindUser(params)

	// если такой не найден
	if err == sql.ErrNoRows {
		//   то записываем в БД через тот же репозиторий
		_, err := repository.InsertUser(params)
		//   если из репозитория вернулась ошибка то возвращаем false, err
		if err != nil {
			return false, err
		}
	}
	// возвращаем true (это будет значить, что пользователь уже зареган или зареган сейчас)
	return true, nil	
}
