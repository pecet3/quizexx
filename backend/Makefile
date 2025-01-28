run:
	go run ./cmd

rm:
	rm -rf store.db 


reset_db:
	rm -rf store.db 
	touch store.db
	sqlite3 store.db < ./data/migrations/0001_initial.up.sql