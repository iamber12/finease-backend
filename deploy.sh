set -o errexit

cd FinEase/backend
git pull origin release/SIT
go build -a -o ~/FinEase/backend/bin/finease-backend ~/FinEase/backend/cmd/main.go
lsof -i tcp:8000 | awk '{ print $2; }' | grep -v "PID" | xargs kill > /dev/null 2&>1
~/FinEase/backend/bin/finease-backend serve --db-name=$DB_NAME --db-user=$DB_USER --db-password=$DB_PASSWORD --db-host=$DB_HOST --db-port=$DB_PORT &