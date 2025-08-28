# Booker

Go приложение для работы с бронированиями.

## Переменные окружения

Приложение использует следующие переменные окружения:

| Переменная | Описание | Обязательная |
|------------|----------|--------------|
| `UNSPOT_TOKEN` | Bearer токен с веб-страницы unspot | Да |
| `SPOT_ID` | UUID места из unspot | Да |
| `UNSPOT_URL` | unspot URL | Нет |

## Использование с Docker

### Публичный образ

Доступен публичный Docker образ: `atamanvll/booker`

```bash
# Запуск контейнера с переменными окружения
docker run -e UNSPOT_TOKEN="your_token_here" \
           -e SPOT_ID="your_spot_id_here" \
           -e UNSPOT_URL="https://your-unspot-instance.com" \
           atamanvll/booker

# Или с файлом .env
docker run --env-file .env atamanvll/booker
```

### Локальная сборка

```bash
# Сборка образа
docker build -t booker .

# Запуск контейнера с переменными окружения
docker run -e UNSPOT_TOKEN="your_token_here" \
           -e SPOT_ID="your_spot_id_here" \
           booker
```

### Пример файла .env

```env
UNSPOT_TOKEN=your_bearer_token_here
SPOT_ID=your_spot_uuid_here
UNSPOT_URL=https://your-unspot-instance.com
```

