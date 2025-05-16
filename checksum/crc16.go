package checksum

import (
	"fmt"
	"github.com/sigurn/crc16"
	"io"
	"os"
)

// checksum: 0x1A88 (replace to 0D 44 and calculate) CRC-16/ARC

func CalculateCRC16(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	all, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("could not read file: %w", err)
	}

	table := crc16.MakeTable(crc16.CRC16_ARC)
	checksum := crc16.Checksum(all, table)

	return fmt.Sprintf("%04x", checksum), nil
}
