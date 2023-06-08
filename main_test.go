package main

import (
    "testing"
    "strings"
)

func TestHandleDraws(t *testing.T){
    name = "DiegoAndriano"
    sgame := `[Event "Rated Rapid game"]
    [Site "https://lichess.org/h6wlmISG"]
    [Date "2023.06.06"]
    [White "VaheMovsisyan"]
    [Black "DiegoAndriano"]
    [Result "1/2-1/2"]
    [UTCDate "2023.06.06"]
    [UTCTime "15:48:37"]
    [WhiteElo "1516"]
    [BlackElo "1591"]
    [WhiteRatingDiff "-4"]
    [BlackRatingDiff "+4"]
    [Variant "Standard"]
    [TimeControl "600+0"]
    [ECO "B12"]
    [Termination "Time forfeit"]

    1. e4 c6 2. d4 d5 3. e5 Bf5 0-1`
    game := strings.Split(sgame, "[")
    handleDraws(game)

    if(totalDraws[15] == 0){
        t.Errorf("Result was incorrect, got: 0, want: 1.")
    }
    if(draws[15] == 0){
        t.Errorf("Result was incorrect, got: 0, want: 1.")
    }

    totalDraws[15] = 0
    draws[15] = 0
    sgame = `[Event "Rated Rapid game"]
    [Site "https://lichess.org/h6wlmISG"]
    [Date "2023.06.06"]
    [White "VaheMovsisyan"]
    [Black "DiegoAndriano"]
    [Result "1-0"]
    [UTCDate "2023.06.06"]
    [UTCTime "15:48:37"]
    [WhiteElo "1516"]
    [BlackElo "1591"]
    [WhiteRatingDiff "-4"]
    [BlackRatingDiff "+4"]
    [Variant "Standard"]
    [TimeControl "600+0"]
    [ECO "B12"]
    [Termination "Time forfeit"]

    1. e4 c6 2. d4 d5 3. e5 Bf5 0-1`
    game = strings.Split(sgame, "[")
    handleDraws(game)

    if(totalDraws[15] == 1){
        t.Errorf("Result was incorrect, got: 1, want: 0.")
    }
    if(draws[15] == 1){
        t.Errorf("Result was incorrect, got: 1, want: 0.")
    }
}

func TestHandleWhiteWon(t *testing.T){
    name = "DiegoAndriano"

    sgame := `[Event "Rated Rapid game"]
    [Site "https://lichess.org/h6wlmISG"]
    [Date "2023.06.06"]
    [White "DiegoAndriano"]
    [Black "VaheMovsisyan"]
    [Result "1-0"]
    [UTCDate "2023.06.06"]
    [UTCTime "15:48:37"]
    [WhiteElo "1516"]
    [BlackElo "1591"]
    [WhiteRatingDiff "-4"]
    [BlackRatingDiff "+4"]
    [Variant "Standard"]
    [TimeControl "600+0"]
    [ECO "B12"]
    [Termination "Time forfeit"]

    1. e4 c6 2. d4 d5 3. e5 Bf5 0-1`
    game := strings.Split(sgame, "[")
    handleWhiteWon(game)

    if(totalWins[15] == 0){
        t.Errorf("Result was incorrect, got: 0, want: 1.")
    }
    if(wins[15] == 0){
        t.Errorf("Result was incorrect, got: 0, want: 1.")
    }

    totalWins[15] = 0
    wins[15] = 0

    sgame = `[Event "Rated Rapid game"]
    [Site "https://lichess.org/h6wlmISG"]
    [Date "2023.06.06"]
    [White "VaheMovsisyan"]
    [Black "DiegoAndriano"]
    [Result "1-0"]
    [UTCDate "2023.06.06"]
    [UTCTime "15:48:37"]
    [WhiteElo "1516"]
    [BlackElo "1591"]
    [WhiteRatingDiff "-4"]
    [BlackRatingDiff "+4"]
    [Variant "Standard"]
    [TimeControl "600+0"]
    [ECO "B12"]
    [Termination "Time forfeit"]

    1. e4 c6 2. d4 d5 3. e5 Bf5 0-1`
    game = strings.Split(sgame, "[")
    handleWhiteWon(game)

    if(totalLosses[15] == 0){
        t.Errorf("Result was incorrect, got: 0, want: 1.")
    }
    if(losses[15] == 0){
        t.Errorf("Result was incorrect, got: 0, want: 1.")
    }
}

func TestHandleBlackWon(t *testing.T){
    name = "DiegoAndriano"

    sgame := `[Event "Rated Rapid game"]
    [Site "https://lichess.org/h6wlmISG"]
    [Date "2023.06.06"]
    [White "VaheMovsisyan"]
    [Black "DiegoAndriano"]
    [Result "0-1"]
    [UTCDate "2023.06.06"]
    [UTCTime "15:48:37"]
    [WhiteElo "1516"]
    [BlackElo "1591"]
    [WhiteRatingDiff "-4"]
    [BlackRatingDiff "+4"]
    [Variant "Standard"]
    [TimeControl "600+0"]
    [ECO "B12"]
    [Termination "Time forfeit"]

    1. e4 c6 2. d4 d5 3. e5 Bf5 0-1`
    game := strings.Split(sgame, "[")
    handleBlackWon(game)

    if(totalWins[15] == 0){
        t.Errorf("Result was incorrect, got: 0, want: 1.")
    }
    if(wins[15] == 0){
        t.Errorf("Result was incorrect, got: 0, want: 1.")
    }

    totalWins[15] = 0
    wins[15] = 0

    sgame = `[Event "Rated Rapid game"]
    [Site "https://lichess.org/h6wlmISG"]
    [Date "2023.06.06"]
    [White "DiegoAndriano"]
    [Black "VaheMovsisyan"]
    [Result "0-1"]
    [UTCDate "2023.06.06"]
    [UTCTime "15:48:37"]
    [WhiteElo "1516"]
    [BlackElo "1591"]
    [WhiteRatingDiff "-4"]
    [BlackRatingDiff "+4"]
    [Variant "Standard"]
    [TimeControl "600+0"]
    [ECO "B12"]
    [Termination "Time forfeit"]

    1. e4 c6 2. d4 d5 3. e5 Bf5 0-1`
    game = strings.Split(sgame, "[")
    handleBlackWon(game)

    if(totalLosses[15] == 0){
        t.Errorf("Result was incorrect, got: 0, want: 1.")
    }
    if(losses[15] == 0){
        t.Errorf("Result was incorrect, got: 0, want: 1.")
    }
}

func TestGetRates(t *testing.T){
    totalWins[0] = 50.0
    totalDraws[0] = 5.0
    totalLosses[0] = 50.0
    wins[0] = 50.0
    totalDraws[0] = 5.0
    totalLosses[0] = 50.0

    getRates()

    if(wins[0] > 48){
        t.Errorf("Result was incorrect, got: %f, want: ~47", wins[0])
    }
    if(losses[0] > 48){
        t.Errorf("Result was incorrect, got: %f, want: ~47", losses[0])
    }
    if(draws[0] > 5){
        t.Errorf("Result was incorrect, got: %f, want: ~47", draws[0])
    }
}

func TestGetGames(t *testing.T){


}
