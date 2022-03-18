# Телеграм бот для автоматизації репорту пропагандистьских каналів

Телеграм бот використовується для автоматизації репорту пропагандистьских каналів. 

Це телеграм бот, якому ви можете надати доступ до будь-якої кількості ваших телеграм аккаунтів ввівши номер телефону, код, 2FA пароль якщо необхіден. Та бот буде надсилати репорти із авторизованих в ньому акаунтів. Список пропагандистьских каналів наразі задається через відправлення повідомлень із посиланнями (https://t.me/channel_name) адмінами боту.

## Вимоги

1. Заводимо телеграм бота за допомогою https://t.me/BotFather  та отримуємо `TOKEN`
2. Реєсруємо нову телеграм аплікацію https://my.telegram.org та отримуємо `api_id`, `api_hash`, конфігурацію продакшин датасерверів телеграм:`Public keys`, `Ip`, `Id`

## Конфігурація

```yaml
Usage:
  tg-bot [OPTIONS] bot [bot-OPTIONS]

Application Options:
  --debug              Is debug mode? [$DEBUG]

Help Options:
  -h, --help           Show this help message

  [bot command options]
  --token=             Bot token [$TOKEN]
  --db=                Database filepath (default: bbolt.db) [$DB]
  -a, --admin=         Bot admin telegram usernames [$ADMIN]
  --pattern=           Bot server handler pattern (default: /) [$PATTERN]
  --cert=              Bot server tls cert file [$CERT_FILE]
  --key=               Bot server tls key file [$KEY_FILE]
  --webhook_url=       Bot server webhook url [$WEBHOOK_URL]
  --address=           Bot server bind address (default: 0.0.0.0:443) [$ADDRESS]

reporter:
  --reporter.msg=      A report message (default: The channel undermines the integrity of the Ukrainian state. Spreading fake news, misleading people. There are a lot of posts with threats against Ukrainians and Ukrainian soldiers. Block him ASAP) [$REPORTER_MESSAGE]
  --reporter.interval= Interval between sending reports (default: 40m) [$REPORTER_INTERVAL]
  --reporter.max_reps= Max number of sent reports from a telegram client (default: 25) [$REPORTER_INTERVAL_MAX_REPORTS]

hub:
  --hub.app_hash=      Telegram API app hash [$HUB_APP_HASH]
  --hub.pk=            Telegram API public key [$HUB_PUBLIC_KEY]
  --hub.device=        Telegram API device model (default: Dmitry Nev) [$HUB_DEVICE]
  --hub.client_ttl=    A telegram API client TTL (default: 3m) [$HUB_CLIENT_TTL]
  --hub.app_id=        Telegram API app id [$HUB_APP_ID]

dc:
  --hub.dc.ip=         DC ip address [$HUB_DC_IP]
  --hub.dc.id=         DC id (default: 2) [$HUB_DC_ID]
  --hub.dc.port=       DC port (default: 443) [$HUB_DC_PORT]
```

За допомогою `docker-compose.yml`:

```yaml
version: '3'
services:

  tg-stand-for-ukraine:
    build: .
    container_name: tg-stand-with-ukraine
    hostname: tg-stand-with-ukraine
    environment:
      - TOKEN=<bot token>
      - DB=/app/db/bbolt.db
      - ADMIN=
      - DEBUG=true
      - PATTERN=/
      - CERT_FILE=
      - KEY_FILE=
      - WEBHOOK_URL=
      - ADDRESS=

      - HUB_APP_ID=<app_id>
      - HUB_APP_HASH=<app_hash>
      - HUB_PUBLIC_KEY=/app/tg_app_public_key.pem
      - HUB_DEVICE=Dmitry Nev
      - HUB_CLIENT_TTL=7m

      - HUB_DC_PORT=
      - HUB_DC_IP=
      - HUB_DC_ID=

      - REPORTER_INTERVAL=50m
      - REPORTER_INTERVAL_MAX_REPORTS=20
      - REPORTER_MESSAGE=The channel undermines the integrity of the Ukrainian state. Spreading fake news, misleading people. There are a lot of posts with threats against Ukrainians and Ukrainian soldiers. Block him ASAP
    volumes:
      - <PUBLIC KEY FILE>:/app/publicKey
      - <DATABASE FOLDER>::/app/db
```

## Запуск

### Docker

Обовʼязково маунтіть `publicKey` та `database folder` до контейнеру для того щоб клієнтська база зберігалася від запуску до запуску `-v ./publicKey:/app/publicKey -v db:/app/db`

```shell
make docker-build

docker run -v ./publicKey:/app/publicKey -v db:/app/db  tg-stand-with-ukraine app bot --token==<token> <other config...>
```

### Compose

Обовʼязково маунтіть `publicKey` та `database folder` до сервіс для того щоб клієнтська база зберігалася від запуску до запуску:

```yaml
    volumes:
      - <PUBLIC KEY FILE>:/app/publicKey
      - <DATABASE FOLDER>:/app/db
```

```shell
docker-compose up
```

## Використання

Після успішного запуску бота переходимо до нього в чат натискаемо `/login` команду, проходимо авторизацію. Після успішної авторизації клієнтів адміни бота, ті що зазначені в `-a, --admin= Bot admin telegram username [$ADMIN]`,  можут надсилати повідомлення із посиланнями до пропагандистьских каналів. Бот автоматично розбере їх та збереже до своєї бази, після чого авторизовані телеграм клієнти будуть відпраляти одноразово репорти із інтервалом `--hub.rep_interval= Interval between sending reports (default: 40m) [$HUB_SEND_REPORTS_INTERVAL]`

## Рекомендації 

- НІКОМУ НЕ ПОВІДОМЛЯЙТЕ ДАНІ отримані із https://t.me/BotFather та https://my.telegram.org

- Не задавайте `--hub.rep_interval= Interval between sending reports (default: 40m) [$HUB_SEND_REPORTS_INTERVAL]` меншим ніж 5-10хв для запобігання бану.

- Після успішної авторизації клієнта видаліть повідомлення із приватними даними номер, код, 2fa пароль.

- Не змінюйте `--hub.device= Telegram API device model (default: Dmitry Nev) [$HUB_DEVICE]`  може призвести до бану. Якщо бажаєте змінити, то спочатку розлогіньтесь із всіх кліентів.

## Ящко ваш акаунт заблокували

Пишіть на пошу `abuse@telegram.org`, `recover@telegram.org`, `login@stel.com` з проханням до розблокування та опишіть ситуацію для чого використовували аккаунт і що намагались допомогти українскьому народу! Вас повинні розблокувати!

Ініші дискусії:
 - https://github.com/gotd/td/blob/main/.github/SUPPORT.md#how-to-not-get-banned

 - https://github.com/lonamiwebs/telethon/issues/824#issuecomment-432182634

## TODO

- [ ] підписка телеграм клієнтів на канал з оновленнями щодо бази пропагандистьских каналів. Замість того щоб адміни надсилали списки у повідомленнях до бота

- [ ] можливість `loguot` з певного клієнта

- [ ] публічний `docker image`

- [ ] поповнення спику пропагандистьских каналів через файл

- [ ] автоматичне видалення повідомлення із номером, кодом, 2fa паролем 

- [x] розбирати посилання на пропагандистьскі канали у форматі `@channel_name`

- [ ] розбирати посилання на пропагандистьскі канали у форматі посилання на запрошення


## Дякую за натхнення

- [Antcating/telegram_report_bot_ua](https://github.com/Antcating/telegram_report_bot_ua)
