# ClockwiseSpiral

Приложение для построения числовой спирали из исходного числа

Компиляция,рекомпиляция бинарника сервера go build -race -ldflags "-s -w" -o bin/server server/main.go

Компиляция,рекомпиляция бинарника клиента  go build -race -ldflags "-s -w" -o bin/client client/main.go

Запуск сервера bin/server
Запуск клиента и передача аргумента флагом коммандной строки bin/client -a=3