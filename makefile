dbconn?="localhost:5432"

migrate-all:
		goose -dir internal/services/user_service/storage/migrations pgx "postgres://postgres:pushword@$(dbconn)/pushpost_user" up
		goose -dir internal/services/message_service/storage/migrations pgx "postgres://postgres:pushword@$(dbconn)/pushpost_message" up
		goose -dir internal/services/notification_service/storage/migrations pgx "postgres://postgres:pushword@$(dbconn)/pushpost_notification" up

migration:
		goose -dir internal/services/user_service/storage/migrations create create_user_service_schema sql
