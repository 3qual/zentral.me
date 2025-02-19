#!/bin/bash

# Каталоги, которые нужно создать
dirs=(
  "./data/go"
  "./data/java"
  "./data/python"
  "./data/mongodb"
  "./data/uptimekuma"
  "./logs/go"
  "./logs/java"
  "./logs/python"
  # "./logs/mongodb"
  "./logs/uptimekuma"
)

# Создание каталогов
for dir in "${dirs[@]}"; do
  if [ ! -d "$dir" ]; then
    echo "Создаю каталог: $dir"
    mkdir -p "$dir"
  else
    echo "Каталог уже существует: $dir"
  fi
done

# Создаем пустые лог-файлы, если их нет
touch ./logs/go/go.log
touch ./logs/java/java.log
touch ./logs/python/python.log
touch ./logs/uptimekuma/uptimekuma.log

echo "Все каталоги созданы."