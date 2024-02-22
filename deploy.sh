cd FinEase/backend
git pull origin release/SIT
go build -a -o ~/FinEase/backend/bin/finease-backend ~/FinEase/backend/cmd/main.go
~/FinEase/backend/bin/finease-backend serve &