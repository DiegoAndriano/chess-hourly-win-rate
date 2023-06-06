package main

import (
  "fmt"
  "strings"
  "regexp"
  "strconv"
  "net/http"
  "io"
  "github.com/gin-gonic/gin"
)

var name string = ""

var draws [24]float32
var wins [24]float32
var losses [24]float32
var totalDraws [24]float32
var totalWins [24]float32
var totalLosses [24]float32

func handleDraws(game []string){
  if(strings.Contains(game[6], "1/2-1/2")){
    re := regexp.MustCompile(`"[^"]+"`)
    newStrs := re.FindAllString(game[8], -1)
    hour, _ := strconv.Atoi(string(newStrs[0])[1:3])
    draws[hour]++
    totalDraws[hour]++
  }
}

func handleWhiteWon(game []string) {
 if(strings.Contains(game[6], "1-0")){
  if(strings.Contains(game[4], name)){
    re := regexp.MustCompile(`"[^"]+"`)
    newStrs := re.FindAllString(game[8], -1)
    hour, _ := strconv.Atoi(string(newStrs[0])[1:3])
    totalWins[hour]++
    wins[hour]++
  } else {
    re := regexp.MustCompile(`"[^"]+"`)
    newStrs := re.FindAllString(game[8], -1)
    hour, _ := strconv.Atoi(string(newStrs[0])[1:3])
    totalLosses[hour]++
    losses[hour]++
  }
}
}
func handleBlackWon(game []string) {
  if(strings.Contains(game[6], "0-1")){
    if(strings.Contains(game[5], name)){
      re := regexp.MustCompile(`"[^"]+"`)
      newStrs := re.FindAllString(game[8], -1)
      hour, _ := strconv.Atoi(string(newStrs[0])[1:3])

      totalWins[hour]++
      wins[hour]++
    } else {
      re := regexp.MustCompile(`"[^"]+"`)
      newStrs := re.FindAllString(game[8], -1)
      hour, _ := strconv.Atoi(string(newStrs[0])[1:3])

      totalLosses[hour]++
      losses[hour]++
    }
  }
}

func getRates() {
  for i, v := range draws {
    if(v > 0){
      drawrate := ((v) / (totalWins[i] + totalDraws[i] + totalLosses[i])) * 100.0
      draws[i] = drawrate
    }
    if(wins[i] > 0){
      winrate := ((wins[i]) / (totalWins[i] + totalDraws[i] + totalLosses[i])) * 100.0
      wins[i] = winrate
    }
    if(losses[i] > 0){
      lossrate := ((losses[i]) / (totalWins[i] + totalDraws[i] + totalLosses[i])) * 100.0
      losses[i] = lossrate
    }
  }
}

func getGames(username string, format string) string {
  url := fmt.Sprintf("https://lichess.org/api/games/user/%s?analysed=false&evals=false&perfType=%s", username, format)

    var client http.Client
    resp, err := client.Get(url)
    if err != nil {
      fmt.Printf("%v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
      bodyBytes, err := io.ReadAll(resp.Body)
      if err != nil {
        fmt.Printf("%v", err)
      }    
      return string(bodyBytes)
    }
    return ""
  }

  func main() {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
      username, format := c.Query("username"), c.Query("format")
      name = username
      games := getGames(username, format)

      arr := strings.Split(games, "\n\n\n")[0:]
      for _, v := range arr {
        game := strings.Split(v, "[")
        if(len(game) > 1){
          handleDraws(game)
          handleWhiteWon(game)
          handleBlackWon(game)
        }
      }

      getRates()
      
      c.JSON(http.StatusOK, gin.H{
        "00": fmt.Sprintf("w%f l%f d%f", wins[0], losses[0], draws[0]),
        "01": fmt.Sprintf("w%f l%f d%f", wins[1], losses[1], draws[1]),
        "02": fmt.Sprintf("w%f l%f d%f", wins[2], losses[2], draws[2]),
        "03": fmt.Sprintf("w%f l%f d%f", wins[3], losses[3], draws[3]),
        "04": fmt.Sprintf("w%f l%f d%f", wins[4], losses[4], draws[4]),
        "05": fmt.Sprintf("w%f l%f d%f", wins[5], losses[5], draws[5]),
        "06": fmt.Sprintf("w%f l%f d%f", wins[6], losses[6], draws[6]),
        "07": fmt.Sprintf("w%f l%f d%f", wins[7], losses[7], draws[7]),
        "08": fmt.Sprintf("w%f l%f d%f", wins[8], losses[8], draws[8]),
        "09": fmt.Sprintf("w%f l%f d%f", wins[9], losses[9], draws[9]),
        "10": fmt.Sprintf("w%f l%f d%f", wins[10], losses[10], draws[10]),
        "11": fmt.Sprintf("w%f l%f d%f", wins[11], losses[11], draws[11]),
        "12": fmt.Sprintf("w%f l%f d%f", wins[12], losses[12], draws[12]),
        "13": fmt.Sprintf("w%f l%f d%f", wins[13], losses[13], draws[13]),
        "14": fmt.Sprintf("w%f l%f d%f", wins[14], losses[14], draws[14]),
        "15": fmt.Sprintf("w%f l%f d%f", wins[15], losses[15], draws[15]),
        "16": fmt.Sprintf("w%f l%f d%f", wins[16], losses[16], draws[16]),
        "17": fmt.Sprintf("w%f l%f d%f", wins[17], losses[17], draws[17]),
        "18": fmt.Sprintf("w%f l%f d%f", wins[18], losses[18], draws[18]),
        "19": fmt.Sprintf("w%f l%f d%f", wins[19], losses[19], draws[19]),
        "20": fmt.Sprintf("w%f l%f d%f", wins[20], losses[20], draws[20]),
        "21": fmt.Sprintf("w%f l%f d%f", wins[21], losses[21], draws[21]),
        "22": fmt.Sprintf("w%f l%f d%f", wins[22], losses[22], draws[22]),
        "23": fmt.Sprintf("w%f l%f d%f", wins[23], losses[23], draws[23]),
      })
    })

    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  }
