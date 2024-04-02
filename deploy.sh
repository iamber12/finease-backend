sudo lsof -i :8000 | grep LISTEN | awk '{print $2}' | xargs -r kill
cd FinEase/backend
git reset --hard HEAD~0
git pull origin dev --rebase
go build -a -o ./bin/finease-backend ./cmd/main.go > build.log 2>&1
~/FinEase/backend/bin/finease-backend serve --db-name=$DB_NAME --db-user=$DB_USER --db-password=$DB_PASSWORD --db-host=$DB_HOST --db-port=$DB_PORT &