APP_NAME="pass-genie"
rm -rf dist/
mkdir dist

ANDROID_FILE_NAME="${APP_NAME}-android-arm64"
GOOS=android GOARCH=arm64 go build -o ${ANDROID_FILE_NAME} ./main.go
mv ${ANDROID_FILE_NAME} dist

LINUX_FILE_NAME="${APP_NAME}-linux-arm64"
GOOS=linux GOARCH=arm64 go build -o ${LINUX_FILE_NAME} ./main.go
mv ${LINUX_FILE_NAME} dist

LINUX_FILE_NAME="${APP_NAME}-linux-x86"
GOOS=linux GOARCH=386 go build -o ${LINUX_FILE_NAME} ./main.go
mv ${LINUX_FILE_NAME} dist

MAC_FILE_NAME="${APP_NAME}-mac-arm64"
GOOS=darwin GOARCH=arm64 go build -o ${MAC_FILE_NAME} ./main.go
mv ${MAC_FILE_NAME} dist

WINDOWS_FILE_NAME="${APP_NAME}-windows-x86.exe"
GOOS=windows GOARCH=386 go build -o ${WINDOWS_FILE_NAME} ./main.go
mv ${WINDOWS_FILE_NAME} dist
