ssh -tt -o StrictHostKeyChecking=no ubuntu@ec2-3-15-22-52.us-east-2.compute.amazonaws.com
cd FinEase/backend
git reset --hard HEAD~0
git pull origin dev --rebase
go build -v -o ./bin/finease-backend ./cmd/main.go
~/FinEase/backend/bin/finease-backend serve --db-name=$DB_NAME --db-user=$DB_USER --db-password=$DB_PASSWORD --db-host=$DB_HOST --db-port=$DB_PORT &