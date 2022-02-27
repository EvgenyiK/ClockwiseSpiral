# ClockwiseSpiral

Приложение для построения числовой спирали из исходного числа

Запуск сервера docker-compose up

Компиляция,рекомпиляция бинарника клиента  go build -race -ldflags "-s -w" -o bin/client client/main.go


Запуск клиента и передача аргумента флагом коммандной строки bin/client -a=3