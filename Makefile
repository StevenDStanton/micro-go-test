default: flash

flash:
	@echo "Building firmware..."
	@tinygo build -opt=2 -target=arduino-mega2560 -o firmware.hex main.go
	@echo "Flashing firmware..."
	@avrdude -c wiring -p m2560 -P /dev/ttyACM0 -b 115200 -U flash:w:firmware.hex -D
	@echo "Cleaning up..."
	@rm -f $(TARGET)

.PHONY: default flash
