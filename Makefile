# shuffler source code
# Author (c) 2023 Belousov Daniil

TEMP_BIN=temp.bin
NAME=shuffler
BINARY_DIRECTORY=/usr/local/bin/

.PHONY: all
.DEFAULT_GOAL: all
all: build install
	@rm -r $(TEMP_BIN) || true

.PHONY: clean
clean:
	@rm -r $(TEMP_BIN) || true

.PHONY: build_app build_getconfig build
build_app:
	@echo "building $(NAME)"
	@mkdir $(TEMP_BIN)
	@go build -o $(TEMP_BIN)/$(NAME) cmd/$(NAME)/*

build_getconfig:
	@echo "building $(NAME).getconfig"
	@go build -o $(TEMP_BIN)/$(NAME).getconfig cmd/$(NAME).getconfig/*

build: clean build_app build_getconfig
	@echo "build successful"


.PHONY: install_app install_getconfig install
install_app:
	@echo "instaling $(NAME)"
	@mv $(TEMP_BIN)/$(NAME) $(BINARY_DIRECTORY)

install_getconfig:
	@echo "instaling $(NAME).getconfig"
	@mv $(TEMP_BIN)/$(NAME).getconfig $(BINARY_DIRECTORY)

install: install_app install_getconfig
	@echo "insallation successful"
	@echo "\n$(NAME) is ready to use, installed: \n - $(NAME) \n - $(NAME).getconfig\n"
