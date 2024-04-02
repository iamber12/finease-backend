export DB_NAME=$DB_NAME
export DB_USER=$DB_USER
export DB_PASSWORD=$DB_PASSWORD
export DB_HOST=$DB_HOST
export DB_PORT=$DB_PORT
sudo lsof -i :8000 | grep LISTEN | awk '{print $2}' | xargs -r kill
cd FinEase/backend
git reset --hard HEAD~0
git pull origin release/SIT --rebase
/usr/local/go/bin/go build -o ./bin/finease-backend ./cmd/main.go
/usr/bin/screen -dmS finease-backend bash -c 'export DB_NAME="your_db_name"; export DB_USER="your_db_user"; export DB_PASSWORD="your_password"; export DB_HOST="your_db_host"; export DB_PORT="your_db_port"; exec ~/FinEase/backend/bin/finease-backend serve --db-name=$DB_NAME --db-user=$DB_USER --db-password=$DB_PASSWORD --db-host=$DB_HOST --db-port=$DB_PORT'