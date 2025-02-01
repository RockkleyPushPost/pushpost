dbconn?="localhost:5432"

migration:
	goose -dir internal/services/user_service/storage/migrations create create_friend_requests_table sql