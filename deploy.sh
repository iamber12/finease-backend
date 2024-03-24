cd FinEase/backend
git reset --hard HEAD~0
git pull origin dev --rebase
nohup go build -a -o ./bin/finease-backend ./cmd/main.go > build.log 2>&1