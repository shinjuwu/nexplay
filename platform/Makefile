# Makefile

# 指定Go編譯器位置
GO_BUILD=go build

# 編譯後的目標文件名
BINARY_NAME=dcc_backend

# 目標目錄，替換為實際路徑
TARGET_DIRECTORY=/usr/local/bin/backend

# 遷移文件目標目錄，替換為實際路徑
MIGRATIONS_TARGET_DIRECTORY=/usr/local/bin/backend/db/migrations

# 默認目標
all: swag_pre swag build move_binary copy_migrations restart_backend

# 
swag_pre:
	chmod 777 backend/swag

# 編譯 swagger 檔案
swag:
	backend/swag init -g backend/main.go -o backend/docs

# 構建應用程序
build:
	$(GO_BUILD) -o $(BINARY_NAME) backend/main.go

# 將構建的程序移動到目標目錄
move_binary:
	mv $(BINARY_NAME) $(TARGET_DIRECTORY)

# 複制遷移文件到指定目錄
copy_migrations:
	cp -r backend/db/migrations/* $(MIGRATIONS_TARGET_DIRECTORY)

# 重啟後台服務
restart_backend:
	systemctl restart backend || exit 1

# 開啟後台服務
start_backend:
	systemctl start backend || exit 1

# 關閉後台服務
stop_backend:
	systemctl stop backend || exit 1

# 清理構建文件
clean:
	rm -f $(TARGET_DIRECTORY)/$(BINARY_NAME)
