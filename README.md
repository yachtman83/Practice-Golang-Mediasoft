# Practice-Golang-Mediasoft
Задание по производственной практике в ООО "MEDIASOFT"

#  Задание

##  Хабибуллин Хабиб Решатович

## УлГТУ, ФИСТ, ИВТАСбд-32

## Задание - [https://docs.google.com/document/d/1-laS0wKfca9m3r0FOBkMI1GuZ6HSyC73/edit](https://docs.google.com/document/d/1-laS0wKfca9m3r0FOBkMI1GuZ6HSyC73/edit)

## Описание проекта

Создать REST API сервис
Сервер должен хранить в файле
в формате JSON сущности на выбор:
1. Автомобили
Перечисление полей:
1. Идентификатор (назначать самому)
2. Название марки
3. Название модели
4. Пробег
5. Количество владельцев

#### Запуск проекта:
Чтобы установить Go в VScode - [https://learn.microsoft.com/ru-ru/azure/developer/go/configure-visual-studio-code](https://learn.microsoft.com/ru-ru/azure/developer/go/configure-visual-studio-code)

#### Тестирование:

Примечание: 

Так как я использовал командную строку (cmd) в VScode это Command Promt, чтобы запросы работали корректно, нужно экранизировать кавычки

Пример запроса:

curl -X POST -d "{\"brand\":\"Subaru\", \"model\":\"Legacy\", \"mileage\":25700, \"owners\":1}" -H "Content-Type: application/json" http://localhost:8000/cars


Создание сущности:
До:
![Снимок экрана 2024-07-07 161833](https://github.com/yachtman83/Practice-Golang-Mediasoft/assets/80544566/a36d352e-7b46-4db0-b9c8-492b156a67d9)

После:
![Снимок экрана 2024-07-07 163926](https://github.com/yachtman83/Practice-Golang-Mediasoft/assets/80544566/33ca22b0-ff47-479d-9e5a-019772f23e66)  ![Снимок экрана 2024-07-07 163943](https://github.com/yachtman83/Practice-Golang-Mediasoft/assets/80544566/ca061154-8906-4e08-84a1-0c84187ea86f)


Получение списка сущностей:

![Снимок экрана 2024-07-07 164352](https://github.com/yachtman83/Practice-Golang-Mediasoft/assets/80544566/ee5ee08d-77b9-4eda-9c91-1a26de2e6fd5)


Получение одной сущности:

![Снимок экрана 2024-07-07 164443](https://github.com/yachtman83/Practice-Golang-Mediasoft/assets/80544566/4aca0278-c12e-4ee4-8ec4-2a0edb8bf5c4)
![Снимок экрана 2024-07-07 164520](https://github.com/yachtman83/Practice-Golang-Mediasoft/assets/80544566/92d8bde5-131b-40e6-b1ee-3c7103eaf2eb)


Обновление всех полей сущности:

До обновления:

![image](https://github.com/yachtman83/Practice-Golang-Mediasoft/assets/80544566/145343ac-dbb4-4789-b566-8486d7a59179)

После обновления:

![image](https://github.com/yachtman83/Practice-Golang-Mediasoft/assets/80544566/09ca695e-a939-41ad-9316-5d7f042cfa4a)

![image](https://github.com/yachtman83/Practice-Golang-Mediasoft/assets/80544566/9ae6f466-d5a2-4690-afe1-0a1903cbad7b)


Обновление одного или несколько полей одной сущности:

До обновления:

![image](https://github.com/yachtman83/Practice-Golang-Mediasoft/assets/80544566/44119c22-7453-4167-a157-0d42d4ef3744)

После обновления:

![image](https://github.com/yachtman83/Practice-Golang-Mediasoft/assets/80544566/3ccbb8d8-6f91-4624-aa62-7ff24e770d4b)

![image](https://github.com/yachtman83/Practice-Golang-Mediasoft/assets/80544566/8cd00286-8d0d-48f3-bcf7-385aeec4a38e)

![image](https://github.com/yachtman83/Practice-Golang-Mediasoft/assets/80544566/9ff735e0-3391-45c3-84fc-90a116920847)

![image](https://github.com/yachtman83/Practice-Golang-Mediasoft/assets/80544566/ecb0eba8-2ae1-4206-859a-54a38c6bfd56)


Удаление сущности:

Список до удаления:

![image](https://github.com/yachtman83/Practice-Golang-Mediasoft/assets/80544566/63a1da90-9db5-47c2-8035-c09dddd3ca3b)

Список после удаления:

![image](https://github.com/yachtman83/Practice-Golang-Mediasoft/assets/80544566/889d9dfe-4486-4a0d-982a-dabd51e4a941)

![image](https://github.com/yachtman83/Practice-Golang-Mediasoft/assets/80544566/aaf7d0b2-6fd2-4485-b07b-dc0ae4036973)


















