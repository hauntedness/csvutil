package csvutil

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func ReadCsv[T any](filepath string) ([]T, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(records) <= 1 {
		return nil, nil
	}
	header := records[0]
	lines := records[1:]
	objects := make([]T, len(lines))
	var sb strings.Builder
	for i := range lines {
		line := lines[i]
		if len(line) == 0 {
			continue
		}
		sb.WriteRune('{')
		for j := range line {
			sb.WriteRune('"')
			sb.WriteString(header[j])
			sb.WriteRune('"')
			sb.WriteRune(':')
			sb.WriteRune('"')
			sb.WriteString(line[j])
			sb.WriteRune('"')
			if j < len(line)-1 {
				sb.WriteRune(',')
			}
		}
		sb.WriteString("}")
		value := sb.String()
		err = json.Unmarshal([]byte(value), &objects[i])
		if err != nil {
			return nil, fmt.Errorf("error unmarshal json %w at line %d", err, i+1)
		}
		sb.Reset()
	}
	return objects, nil
}