package core

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func WritePointsToFile(method, fileName string, path Paths) {
	f, err := os.Create(fileName)
	w := bufio.NewWriter(f)

	if err != nil {
		log.Panicln(err.Error())
	}

	w.WriteString(fmt.Sprintf("# %s Method - Path Length: %f\n", method, path.Score))

	for idx, point := range path.Points {
		w.WriteString(fmt.Sprintf("P%d %d %d\n", idx, point.X, point.Y))
	}
	w.Flush()
}

func GetPointsFromFile(fileName string) []PointD {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	result := make([]PointD, 0)

	for scanner.Scan() {

		text := scanner.Text()
		textSplit := strings.Split(text, " ")

		if len(textSplit) > 3 {
			continue
		}

		x, err := strconv.Atoi(textSplit[1])
		if err != nil {
			log.Panicln("Error casting X", err.Error())
		}
		y, err := strconv.Atoi(textSplit[2])
		if err != nil {
			log.Panicln("Error casting X", err.Error())
		}

		result = append(result,
			PointD{
				X: int16(x),
				Y: int16(y),
			},
		)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func Output(name, version string, path Paths) {
	file := fmt.Sprintf("output/hw10_tsp_points_%s-%s_%d.txt", name, version, len(path.Points))
	WritePointsToFile(name, file, path)

	cmd := exec.Command("python3", "./core/image_gen.py", file, fmt.Sprintf("%s-%s", name, version))

	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal("python3 response: ", string(out))
		log.Fatal("python3 Err: ", err.Error())
	}

	log.Println("python3 response: ", string(out))
}
