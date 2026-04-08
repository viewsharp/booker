# Настройка Yandex Cloud Function

## Предварительные требования

1. **Аккаунт в Yandex Cloud** с активным платежным аккаунтом

## Пошаговая настройка

### 1. Скачивание ZIP архива с GitHub

1. Перейдите на страницу репозитория booker
2. Нажмите кнопку "Code"
3. Выберите "Download ZIP"
4. Сохраните архив на вашем компьютере

### 2. Создание Cloud Function

1. Перейдите в [Yandex Cloud Console](https://console.cloud.yandex.ru/)
2. Выберите ваш каталог
3. Перейдите в раздел "Cloud Functions"
4. Нажмите "Создать функцию"
5. Заполните параметры:
   - **Имя**: `booker`
   - **Описание**: `Booker - бронирование места в Unspot`
   - **Среда выполнения**: `go123`
   - **Точка входа**: `index.Handler`
   - **Размер памяти**: `128 МБ`
   - **Таймаут**: `30 секунд`
6. В разделе "Способ развертывания" выберите "ZIP-архив"
7. Загрузите скачанный ZIP архив
8. Нажмите "Создать функцию"

### 3. Настройка переменных окружения

В созданной функции перейдите в раздел "Переменные окружения" и добавьте:

| Переменная | Значение | Описание |
|------------|----------|----------|
| `UNSPOT_TOKEN` | `your_token_here` | Bearer токен с веб-страницы unspot |
| `SPOT_ID` | `your_spot_id_here` | UUID места из unspot |
| `UNSPOT_URL` | `https://your-unspot-instance.com` | Unspot URL |

### 4. Создание триггера

1. В функции перейдите в раздел "Триггеры"
2. Нажмите "Создать триггер"
3. Выберите тип "Таймер"
4. Заполните параметры:
   - **Имя**: `booker-timer-trigger`
   - **Описание**: `Timer trigger for booker application`
   - **Часовой пояс**: `Europe/Moscow` (или ваш часовой пояс)

5. Выберите вариант расписания, например:
```
0 18 * * *  # Каждый день в 18:00
```
6. Нажмите "Создать триггер"

## Полезные ссылки

- [Документация Yandex Cloud Functions](https://cloud.yandex.ru/docs/functions/)
- [Триггеры](https://cloud.yandex.ru/docs/functions/concepts/trigger/)
- [Cron синтаксис](https://cloud.yandex.ru/docs/functions/concepts/trigger/timer#cron-expression)