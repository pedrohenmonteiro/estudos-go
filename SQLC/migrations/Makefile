createmigration:
	goose -dir migrations -s create init sql

migrate: 
	goose mysql "root:root@tcp(localhost:3308)/blog?parseTime=true" up

migratedown: 
	goose mysql "root:root@tcp(localhost:3308)/blog?parseTime=true" down


.PHONY: migrate migratedown createmigration