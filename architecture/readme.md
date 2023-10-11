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

---