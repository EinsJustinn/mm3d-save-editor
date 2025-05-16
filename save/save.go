package save

import (
	"bytes"
	"fmt"
	"mm3d-save-editor/checksum"
	"mm3d-save-editor/fileutils"
	"os"
)

// money: 0x140 (4 bit)
// bank: 0x1230 (4 bit)
// name: 0x128 (16 bit) 0x21A 0x22C 0x23E

type info struct {
	Name  string
	Money int
	Bank  int
}

type SaveFile struct {
	*info
}

func Load(filePath string) (*SaveFile, error) {
	var info info

	money, err := fileutils.ReadIntFromOffset(filePath, 0x140)
	if err != nil {
		return &SaveFile{}, fmt.Errorf("could not read money: %w", err)
	}
	info.Money = money

	bank, err := fileutils.ReadIntFromOffset(filePath, 0x1230)
	if err != nil {
		return &SaveFile{}, fmt.Errorf("could not read bank: %w", err)
	}
	info.Bank = bank

	name, err := fileutils.ReadString(filePath, 0x128, 16)
	if err != nil {
		return &SaveFile{}, fmt.Errorf("could not read name: %w", err)
	}
	name = string(bytes.ReplaceAll([]byte(name), []byte("\x00"), []byte("")))
	info.Name = name

	return &SaveFile{&info}, nil
}

func (s *SaveFile) SetName(name string) error {
	maxLength := 16

	if len(name) > maxLength {
		return fmt.Errorf("name is too long; maximum allowed length is %d", maxLength)
	}

	s.Name = name
	return nil
}

func (s *SaveFile) SetMoney(value int) error {
	if value > 999 {
		return fmt.Errorf("money must be between 0 and 999")
	}
	s.Money = value

	return nil
}

func (s *SaveFile) SetBank(value int) error {

	// uint16
	if value > 65535 {
		return fmt.Errorf("bank must be between 0 and 65535")
	}
	s.Bank = value
	return nil
}

func (s *SaveFile) SaveToFile(filePath string) error {

	// set name
	offset := int64(0x128)
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	_, err = file.Seek(offset, 0)
	if err != nil {
		return fmt.Errorf("could not seek to offset %d: %w", offset, err)
	}

	_, err = file.Write([]byte(s.fixName(s.Name)))
	if err != nil {
		return fmt.Errorf("could not write name to offset %d: %w", offset, err)
	}

	// set money
	offset = int64(0x140)
	err = fileutils.WriteIntToOffset(filePath, offset, s.Money)
	if err != nil {
		return fmt.Errorf("failed to set money: %w", err)
	}

	// set bank
	offset = int64(0x1230)
	err = fileutils.WriteIntToOffset(filePath, offset, s.Bank)
	if err != nil {
		return fmt.Errorf("failed to set bank: %w", err)
	}

	err = s.fixChecksum(filePath)
	if err != nil {
		return err
	}
	return nil
}

func (s *SaveFile) fixName(name string) string {
	maxLength := 16
	paddedName := make([]byte, 0, maxLength)
	for i := 0; i < len(name); i++ {
		paddedName = append(paddedName, name[i])
		if i < len(name)-1 {
			paddedName = append(paddedName, 0x00)
		}
	}

	for len(paddedName) < maxLength {
		paddedName = append(paddedName, 0x00)
	}
	return string(paddedName)
}

func (s *SaveFile) fixChecksum(filePath string) error {

	err := fileutils.WriteByteToOffset(filePath, 0x1A88, 0xD)
	if err != nil {
		return fmt.Errorf("could not write 0xD to 0x1A88: %w", err)
	}
	err = fileutils.WriteByteToOffset(filePath, 0x1A89, 0x44)
	if err != nil {
		return fmt.Errorf("could not write 0x44 to 0x1A89: %w", err)
	}

	checksumString, err := checksum.CalculateCRC16(filePath)
	if err != nil {
		return fmt.Errorf("could not calculate checksum: %w", err)
	}

	var parsedChecksum uint16
	_, err = fmt.Sscanf(checksumString, "%04x", &parsedChecksum)
	if err != nil {
		return fmt.Errorf("could not parse checksum: %w", err)
	}

	leftByte := byte((parsedChecksum >> 8) & 0xFF)
	rightByte := byte(parsedChecksum & 0xFF)

	err = fileutils.WriteByteToOffset(filePath, 0x1A88, rightByte)
	if err != nil {
		return fmt.Errorf("could not write checksum to 0x1A88: %w", err)
	}

	err = fileutils.WriteByteToOffset(filePath, 0x1A89, leftByte)
	if err != nil {
		return fmt.Errorf("could not write checksum to 0x1A89: %w", err)
	}

	return nil
}
