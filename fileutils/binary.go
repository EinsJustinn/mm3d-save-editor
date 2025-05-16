package fileutils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func WriteIntToOffset(filePath string, offset int64, value int) error {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	_, err = file.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("could not seek to offset %d: %w", offset, err)
	}

	intValue := int32(value)
	err = binary.Write(file, binary.LittleEndian, intValue)
	if err != nil {
		return fmt.Errorf("could not write value at offset %d: %w", offset, err)
	}

	return nil
}

func WriteByteToOffset(filePath string, offset int64, bytes byte) error {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	_, err = file.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("could not seek to offset %d: %w", offset, err)
	}

	err = binary.Write(file, binary.LittleEndian, bytes)
	if err != nil {
		return fmt.Errorf("could not write value at offset %d: %w", offset, err)
	}

	return nil
}

func ReadIntFromOffset(filePath string, offset int64) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	_, err = file.Seek(offset, 0)
	if err != nil {
		return 0, fmt.Errorf("could not seek to offset %d: %w", offset, err)
	}

	var value int32
	err = binary.Read(file, binary.LittleEndian, &value)
	if err != nil {
		return 0, fmt.Errorf("could not read value at offset %d: %w", offset, err)
	}

	return int(value), nil
}

func ReadString(filePath string, offset int64, length int) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	_, err = file.Seek(offset, 0)
	if err != nil {
		return "", fmt.Errorf("could not seek to offset %d: %w", offset, err)
	}

	buffer := make([]byte, length)
	_, err = file.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("could not read string from offset %d: %w", offset, err)
	}

	buffer = bytes.Trim(buffer, "\x00")

	return string(buffer), nil
}
