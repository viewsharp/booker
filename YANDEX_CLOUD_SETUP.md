# Настройка Yandex Cloud Container Registry

## Предварительные требования

1. **Аккаунт в Yandex Cloud** с активным платежным аккаунтом
2. **Созданный Container Registry** в Yandex Cloud
3. **OAuth токен** для доступа к Yandex Cloud

## Пошаговая настройка

### 1. Создание Container Registry

1. Перейдите в [Yandex Cloud Console](https://console.cloud.yandex.ru/)
2. Выберите ваш каталог
3. Перейдите в раздел "Container Registry"
4. Нажмите "Создать реестр"
5. Укажите имя реестра и нажмите "Создать"
6. Запомните ID реестра (понадобится для настройки GitHub Actions)

### 2. Получение OAuth токена

1. Получние OAuth токена описанов в [документации](https://yandex.cloud/ru/docs/container-registry/operations/authentication#user)

### 3. Настройка GitHub Secrets

В вашем GitHub репозитории перейдите в **Settings → Secrets and variables → Actions** и добавьте следующие секреты:

| Секрет | Значение | Описание |
|--------|----------|----------|
| `YC_OAUTH_TOKEN` | Ваш OAuth токен | Токен для аутентификации в Yandex Cloud |
| `YC_REGISTRY_ID` | ID вашего реестра | ID Container Registry (например: `crp1234567890abcdef`) |

### 4. Проверка настройки

После настройки секретов, при пуше в master/main ветку:
- GitHub Actions автоматически соберет Docker образ
- Образ будет опубликован в ваш Yandex Cloud Container Registry
- Доступные теги:
  - `latest` - для master ветки
  - `main` - для main ветки
  - `v1.0.0` - для тегов релизов
  - `1.0` - для мажорных версий
  - `master-abc123` - для конкретных коммитов

## Использование образа

После успешной публикации вы можете использовать образ:

```bash
# Аутентификация в Yandex Cloud
yc container registry configure-docker

# Запуск контейнера
docker run cr.yandex/YOUR_REGISTRY_ID/booker:latest
```

## Настройка Serverless Container

### 1. Создание Serverless Container

1. Перейдите в [Yandex Cloud Console](https://console.cloud.yandex.ru/)
2. Выберите ваш каталог
3. Перейдите в раздел "Serverless Containers"
4. Нажмите "Создать контейнер"
5. Заполните параметры:
   - **Имя**: `booker-container`
   - **Описание**: `Booker application container`
   - **Образ**: `cr.yandex/YOUR_REGISTRY_ID/booker:latest`
   - **Память**: `128 МБ` (или больше при необходимости)
   - **Время выполнения**: `30 секунд`
   - **Количество экземпляров**: `1`

### 2. Настройка переменных окружения

В разделе "Переменные окружения" добавьте:

| Переменная | Значение | Описание |
|------------|----------|----------|
| `UNSPOT_TOKEN` | `your_token_here` | Bearer токен с веб-страницы unspot |
| `SPOT_ID` | `your_spot_id_here` | UUID места из unspot |
| `UNSPOT_URL` | `https://your-unspot-instance.com` | Unspot URL |

### 3. Настройка прав доступа

1. В разделе "Права доступа" создайте новый сервисный аккаунт или выберите существующий
2. Назначьте роль `serverless.containers.invoker`

## Настройка триггера по таймеру

### 1. Создание триггера

1. Перейдите в раздел "Триггеры"
2. Нажмите "Создать триггер"
3. Выберите тип "Таймер"
4. Заполните параметры:
   - **Имя**: `booker-timer-trigger`
   - **Описание**: `Timer trigger for booker application`
   - **Часовой пояс**: `Europe/Moscow` (или ваш часовой пояс)

### 2. Настройка расписания

Выберите один из вариантов расписания:

#### Ежедневно в определенное время
```
0 9 * * *  # Каждый день в 9:00
```

#### В рабочие дни
```
0 9 * * 1-5  # Понедельник-пятница в 9:00
```

#### Каждые N минут
```
*/15 * * * *  # Каждые 15 минут
```

#### Каждый час
```
0 * * * *  # Каждый час в 0 минут
```

### 3. Привязка к контейнеру

1. В разделе "Действие" выберите "Вызвать контейнер"
2. Выберите созданный контейнер `booker-container`
3. Нажмите "Создать триггер"

## Мониторинг и логи

### Просмотр логов

1. Перейдите в раздел "Логи" вашего контейнера
2. Выберите временной интервал
3. Фильтруйте по уровню логирования (INFO, ERROR, DEBUG)

### Настройка алертов

1. Перейдите в раздел "Мониторинг"
2. Создайте алерт на основе метрик:
   - Количество ошибок выполнения
   - Время выполнения
   - Количество вызовов

## Полезные ссылки

- [Документация Yandex Cloud Container Registry](https://cloud.yandex.ru/docs/container-registry/)
- [Настройка Docker для работы с Container Registry](https://cloud.yandex.ru/docs/container-registry/operations/authentication)
- [GitHub Actions с Yandex Cloud](https://cloud.yandex.ru/docs/container-registry/operations/ci-cd/github-actions)
- [Serverless Containers](https://cloud.yandex.ru/docs/serverless-containers/)
- [Триггеры](https://cloud.yandex.ru/docs/functions/concepts/trigger/)
- [Cron синтаксис](https://cloud.yandex.ru/docs/functions/concepts/trigger/timer#cron-expression)
