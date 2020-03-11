package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	g1 := getPlayersFromFile("g1.csv")
	g2 := getPlayersFromFile("g2.csv")

	g3 := append(g2, convertByZValue(g1, g2)...)

	var csv string
	if len(csv) == 0 {
		csv = "player,DKP\n"
	}
	for _, p := range g3 {
		csv += fmt.Sprintf("%s,%.2f\n", p.Name, p.Points)
	}

	ioutil.WriteFile("g3.csv", []byte(csv), 0644)
}

func convertByZValue(g1 []player, g2 []player) []player {
	meang1 := getMeanFromPlayers(g1)
	meang2 := getMeanFromPlayers(g2)
	stddev1 := getStdDev(g1)
	stddev2 := getStdDev(g2)
	var out []player

	for _, p := range g1 {
		out = append(out, player{Name: p.Name, Points: math.Round(stddev2*((p.Points-meang1)/stddev1) + meang2)})
	}

	return out
}

func getStdDev(players []player) float64 {
	mean := getMeanFromPlayers(players)
	var subsqr []float64
	for _, p := range players {
		subsqr = append(subsqr, (p.Points-mean)*(p.Points-mean))
	}
	variance := getMean(subsqr)
	return math.Sqrt(variance)
}

func getMeanFromPlayers(players []player) float64 {
	var sum float64
	for _, p := range players {
		sum += p.Points
	}

	return sum / float64(len(players))
}

func getMean(nums []float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}

	return sum / float64(len(nums))
}

type player struct {
	Name   string
	Points float64
}

func getPlayersFromFile(filePath string) []player {
	file, err := os.Open(filePath)
	if err != nil {
		panic("File read error.")
	}
	defer file.Close()

	var players []player
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ",")
		if s[0] == "player" && s[1] == "DKP" {
			continue
		}
		p := player{Name: s[0], Points: toFloat(s[1])}
		players = append(players, p)
	}

	return players
}

func toFloat(s string) float64 {
	out, _ := strconv.ParseFloat(s, 64)
	return out
}
