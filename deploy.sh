sudo systemctl stop finease-backend.service
cd FinEase/backend
git reset --hard HEAD~0
git pull origin release/SIT --rebase
/usr/local/go/bin/go build -o ./bin/finease-backend ./cmd/main.go
sudo systemctl daemon-reload
sudo systemctl start finease-backend.service