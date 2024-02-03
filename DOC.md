# Архитектура:

C:.\
│   DOC.md \
│   go.mod \
│   go.sum \
│   main.go \
│   README.md \
│ \
├───bot \
│       bot.go \
│ \
└───db \
│       db.txt \
│       saveUserState.go \

# команды 

!info \
!remind \
!remindList \
!remindDelete \

Вся их логика написано в ReadMe файле

# БД

мой код парсит сообщение от пользователя и добавляет данные о напоминании в текстовый файл db.txt в папке db. \
затем для вывода списка команд также читает этот текстовый файл и выводит список записей по channelID

