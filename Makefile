all:  linux windows

linux:
	go build -o yara_tool -ldflags="-s -w " .

windows:
	GOOS=windows go build -o yara_tool.exe  -ldflags="-s -w " .
