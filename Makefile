APP=atlas
BIN=$(PWD)/$(APP)
GOOS=linux
GOARM=7
GOARCH=arm
GO ?= go


build b: clean
	@echo "[build] Building atlas..."
	@cd cmd/peripheral && GOOS=${GOOS} GOARM=${GOARM} GOARCH=${GOARCH} $(GO) build -o $(BIN)

clean c:
	@echo "[clean] Cleaning files..."
	@rm -f $(BIN)

.PHONY: build b clean c

