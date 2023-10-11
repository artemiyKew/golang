# Структура проекта на Golang

*Пакеты* - В go пакет (package) - это коллекция связанных между собой функций, типов, констант и перемнных, которые могут быть использованы в других программах.

---

## Плохой пакет
![Alt text](assets/bad_pkg.png)

---

## Хороший пакет
![Alt text](assets/good_pkg.png)

---

## Плохая структура
![Alt text](assets/bad_structure.png)

### Папка API
- REST
- POST-Like
- GRPC

![Alt text](assets/api.png)

---

### Папка CMD
Содержит: 
- main файлы приложений 
- конфиги
- internal приложения

![Alt text](assets/cmd.png)

#### Папка CMD/INTERNAL
Содержит:
- Внутренние пакеты отдельного приложения

![Alt text](assets/internal.png)

#### Папка CONFIGS
Содержит: 
- Дефолтные конфигурации сервера

![Alt text](assets/config.png)

#### Папка DEPLOYMENTS
Содержит: 
- Конфигурации для локального деплоя и разработки
- Docker, Docker Compose
- CI\CD

![Alt text](assets/DEPLOYMENTS.png)

#### Папка DOCS
Содержит: 
- Документацию по проекту
- Документацию по API
- Документацию по CI

![Alt text](assets/DOCS.png)

---

### Папка INTERNAL
Содержит:
- Внутрениие пакеты утилит и сервисов
- Метрики
- Кэши
- Пакет для работы с JWT и т.д.

![Alt text](assets/INTERNAL1.png)

---

### Папка PKG
Содержит:
- Пакеты, которые могут импортироваться другими проектами

![Alt text](assets/pkg.png)

---

### Папка MIGRATIONS
Содержит:
- Скрипты миграций для баз данных

![Alt text](assets/MIGRATIONS.png)

---

### Папка TEST
Содержит:
- Интеграционные тесты
- E2E тесты
- Spanshot тесты

![Alt text](assets/TEST.png)

---

### OTHER
Содержит:
- go.mod и go.sum
- MAKEfile (TaskFile)
- README
- License

![Alt text](assets/OTHER.png)


## Конфигурация проекта
_Конфигурация_ _проекта_ является одной из важнейших частей архитектры. Если проект легко натсраивать и конфигурировать, то это существенно влияет на скорость разработки и гибкость, расширяемость, тестируемость кода.

---

## Dependency Injection (DI)
_Dependency Injection (DI)_ - это паттерн проектирования, который позволяет управлять зависимостями между объектами в приложении. Он заклячается в том, что объекты не создают свои зависимости самостоятельно, а получают из извне.

![Alt text](assets/plain.png)


# Паттерн Репозиторий
_Репозиторий_ (Repository) - это популярный подход к организации доступа к данным в приложении. Он позволяет абстрагироваться от деталей хранения данных и предоставляет единый интерфейс для работа с данными.

```golang
    // repository.go
   type User struct {
        ID       int
        Name     string
        Email    string
        Password string
    }

    type UserInterface interface {
        GetByID(id int) (*User, error)
        GetAll() ([]*User, error)
        Create(user *User) error
        Update(user *User) error
        Delete(id int) error
    }
```