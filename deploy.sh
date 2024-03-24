cd FinEase/backend
git reset --hard HEAD~0
git pull origin release/SIT --rebase
go build -v -o ~/FinEase/backend/bin/finease-backend ~/FinEase/backend/cmd/main.go
~/FinEase/backend/bin/finease-backend serve --db-name=$DB_NAME --db-user=$DB_USER --db-password=$DB_PASSWORD --db-host=$DB_HOST --db-port=$DB_PORT &