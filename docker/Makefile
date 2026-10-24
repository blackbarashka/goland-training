# Запуск только PostgreSQL через docker compose
postgres-up:
	docker compose up -d postgres

# Запуск HTTP приложения через docker compose (с предварительной сборкой)
run-http-app:
	docker compose up --build -d company-app

# Дополнительно: запуск сразу всего стека
compose-up:
	docker compose up --build -d

# Остановка всех контейнеров
compose-down:
	docker compose down