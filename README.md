# Booker

Go приложение для работы с бронированиями.

## Развертывание

Для настройки развертывания в Yandex Cloud см. [YANDEX_CLOUD_SETUP.md](YANDEX_CLOUD_SETUP.md).

## Переменные окружения

Приложение использует следующие переменные окружения:

| Переменная | Описание | Обязательная |
|------------|----------|--------------|
| `UNSPOT_TOKEN` | Bearer токен с веб-страницы unspot | Да |
| `SPOT_ID` | UUID места из unspot | Да |
| `UNSPOT_URL` | unspot URL | Нет |

### Пример файла .env

```env
UNSPOT_TOKEN=your_bearer_token_here
SPOT_ID=your_spot_uuid_here
UNSPOT_URL=https://your-unspot-instance.com
```

