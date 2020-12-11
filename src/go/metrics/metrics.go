package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const dataLines = 100

func writeMetrics(in, out string) error {
	data, err := readData(in)
	if err != nil {
		return err
	}

	f, err := os.Create(out)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(f)
	writer.Write([]string{"promedio", "desviacion", "tamano muestra"})
	mean := strconv.FormatFloat(calcMean(data), 'f', 3, 64)
	deviation := strconv.FormatFloat(calcDeviation(data), 'f', 3, 64)
	writer.Write([]string{mean, deviation, "0"})
	writer.Flush()

	return nil
}

func readData(in string) ([]float64, error) {
	f, err := os.Open(in)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	lines := make([]string, dataLines)

	for i := 0; i < dataLines; i++ {
		line, err := reader.ReadString('\n')
		if line == "" {
			break
		}

		line = strings.TrimSuffix(line, "\n")

		if err != nil && err != io.EOF {
			return nil, nil
		}

		lines[i] = line
	}
	parsedLine, err := strToFloat(lines)
	if err != nil {
		return nil, err
	}

	return parsedLine, nil
}

func strToFloat(strData []string) ([]float64, error) {
	data := make([]float64, len(strData))
	for i := 0; i < len(strData); i++ {
		if strData[i] == "" {
			i++
			continue
		}

		n, err := strconv.ParseFloat(strData[i], 64)
		if err != nil {
			return nil, err
		}

		data[i] = n
	}
	return data, nil
}

func calcMean(data []float64) float64 {
	sum := 0.000
	for _, n := range data {
		sum += n
	}

	return sum / float64(len(data))
}

func calcDeviation(data []float64) float64 {
	var sd float64
	mean := calcMean(data)

	for i := 0; i < len(data); i++ {
		sd += math.Pow(data[i]-mean, 2)
	}

	return math.Sqrt(sd / float64(len(data)))
}

func calcSampleSize(data []float64) float64 {
	return 0
}

func main() {
	err := writeMetrics("data/pc2-go-700-version2-tratamiento1.csv", "test")
	if err != nil {
		fmt.Println(err)
	}
}
