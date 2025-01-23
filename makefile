dbconn?="localhost:5432"

migration:
	goose -dir internal/services/message_service/storage/migrations create create_message_table sql