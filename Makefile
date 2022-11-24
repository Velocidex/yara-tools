all:  linux windows

linux:
	go build -o yara_tool -ldflags="-s -w " .

windows:
	GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o yara_tool.exe  -ldflags="-s -w " .

test:
	./yara_tool clean ./testing/yara.txt --verify
