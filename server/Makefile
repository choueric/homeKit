EXEC:=server
SRCS=*.go

all:run

$(EXEC):$(SRCS)
	go build -o $(EXEC)

run:$(SRCS)
	go run $(SRCS)

clean:
	rm -rf $(EXEC)
