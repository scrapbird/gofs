# makefile for gofs

CC=go

SOURCES=gofs.go
BIN=gofs

all : $(SOURCES) $(BIN)

$(BIN) : 
	$(CC) build $(SOURCES)

