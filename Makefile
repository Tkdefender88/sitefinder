GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

BUILD_WIN=env GOOS=windows GOARCH=amd64
BUILD_LIN=env GOOS=linux GOARCH=amd64

TARGETS := sf

.PHONY: all clean

linux:
	$(BUILD_LIN) $(GOBUILD) -o $(TARGETS) -v

clean:
	$(GOCLEAN)
	rm -f $(TARGETS)

windows:
	$(BUILD_WIN) $(GOBUILD) -o $(TARGETS) -v
