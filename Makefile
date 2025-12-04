binary := gss
source := $(addsuffix .go, $(binary))

default: build

.PHONY: build
build: $(binary)

$(binary): $(source)
	go build -o $@ $<

.PHONY: clean
clean:
	@rm -f $(binary)
