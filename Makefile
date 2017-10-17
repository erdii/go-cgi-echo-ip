BUILD_DIR = .build

.PHONY: all clean

all: | clean build
build: $(BUILD_DIR)

clean:
	rm -rf $(BUILD_DIR)

$(BUILD_DIR): $(BUILD_DIR)/cgi-echo-ip

$(BUILD_DIR)/cgi-echo-ip:
	go build -o $(BUILD_DIR)/cgi-echo-ip main.go