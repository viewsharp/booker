# Booker

Go приложение для работы с бронированиями.

## Автоматическая сборка Docker образа

Приложение настроено для автоматической сборки и публикации Docker образа в Docker Hub при пуше в master/main ветку.

### Настройка

1. **Создайте аккаунт в Docker Hub** (если еще нет): https://hub.docker.com/

2. **Создайте репозиторий в Docker Hub** с именем, соответствующим вашему GitHub репозиторию (например, `username/booker`)

3. **Добавьте секреты в GitHub репозиторий**:
   - Перейдите в Settings → Secrets and variables → Actions
   - Добавьте следующие секреты:
     - `DOCKER_USERNAME` - ваше имя пользователя в Docker Hub
     - `DOCKER_PASSWORD` - ваш токен доступа Docker Hub (не пароль!)

### Получение Docker Hub токена

1. Войдите в Docker Hub
2. Перейдите в Account Settings → Security
3. Создайте новый Access Token
4. Скопируйте токен и используйте его как `DOCKER_PASSWORD`

### Использование

После настройки, при каждом пуше в master/main ветку:
- Автоматически собирается Docker образ
- Образ публикуется в Docker Hub с тегами:
  - `latest` - для master ветки
  - `main` - для main ветки
  - `v1.0.0` - для тегов релизов
  - `1.0` - для мажорных версий
  - `master-abc123` - для конкретных коммитов

### Локальная сборка

```bash
# Сборка образа
docker build -t booker .

# Запуск контейнера
docker run booker
```

### Платформы

Образ собирается для следующих платформ:
- linux/amd64
- linux/arm64

