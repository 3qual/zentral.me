#!/bin/bash

# Каталоги для данных и логов
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

# Удаляем каталоги data и logs, если они существуют
rm -rf ./data ./logs

# Создаем каталоги
for dir in "${dirs[@]}"; do
  mkdir -p "$dir"
done

# Создаем пустые лог-файлы, если их нет
touch ./logs/go/go.log
touch ./logs/java/java.log
touch ./logs/python/python.log
touch ./logs/uptimekuma/uptimekuma.log

echo "Каталоги и файлы логов подготовлены."