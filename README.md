# Telegram-бот для записи к мастеру

Проект Telegram-бота на Go с использованием [go-telegram/bot](https://github.com/go-telegram/bot).

## Требования

- Go 1.18 или новее
- Токен Telegram-бота, полученный у [@BotFather](https://t.me/BotFather)

## Запуск

1. Клонируйте репозиторий:

   ```bash
   git clone https://github.com/Lane000/tg-bot-nails.git
   cd tg-bot-nails
   ```

2. Установите зависимости:

   ```bash
   go mod tidy
   ```

3. Укажите токен через переменную окружения.

   **macOS / Linux:**

   ```bash
   export TELEGRAM_BOT_TOKEN="ваш_токен_от_BotFather"
   ```

   **Windows PowerShell:**

   ```powershell
   $env:TELEGRAM_BOT_TOKEN="ваш_токен_от_BotFather"
   ```

4. Запустите бота:

   ```bash
   go run .
   ```

5. Откройте своего бота в Telegram и отправьте команду `/start`.

## Текущие возможности

- Запуск бота через Telegram Bot API.
- Ответ приветствием на команду `/start`.
- Получение токена из переменной окружения.

## Безопасность

Не добавляйте токен бота в исходный код и не отправляйте его в GitHub.
