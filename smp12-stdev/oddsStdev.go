package utils

import (
	"fmt"
	"math"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

type HorseDataType struct {
	Index  int
	RaceId string
	Odds   float64
	Pops   int
}

func MySort(c *gin.Context) {
	// サンプルデータの投入
	horsesData := []HorseDataType{
		{0, "2023082701020601", 0.46, 9},
		{1, "2023082701020601", 15.69, 3},
		{2, "2023082701020601", 36.36, 1},
		{3, "2023082701020601", 2.90, 7},
		{4, "2023082701020601", 3.69, 6},
		{5, "2023082701020601", 6.61, 5},
		{6, "2023082701020601", 18.18, 2},
		{7, "2023082701020601", 0.53, 8},
		{8, "2023082701020601", 11.27, 4},
	}
	// horsesData := []HorseDataType{
	// 	{12, "2023082701020602", 11.94, 2},
	// 	{14, "2023082701020602", 2.80, 6},
	// 	{15, "2023082701020602", 61.54, 1},
	// 	{16, "2023082701020602", 6.84, 4},
	// 	{17, "2023082701020602", 2.05, 5},
	// 	{20, "2023082701020602", 1.19, 7},
	// 	{21, "2023082701020602", 8.33, 3},
	// }
	// fmt.Println(horsesData)

	// 上位馬半数を取得
	_horsesData := PickupUpperHorses(horsesData)
	// ピックアップした(対象の)頭数
	dataSize := len(_horsesData)
	fmt.Printf("[1レースの上位馬のオッズ] 集計頭数: %d\n", dataSize)

	// 上位馬半数の「単勝シェアの標準偏差」を求める
	stdev := CalcStdev(_horsesData)

	c.IndentedJSON(http.StatusOK, gin.H{"標準偏差": stdev, "集計頭数": dataSize})
}

func PickupUpperHorses(horsesData []HorseDataType) []HorseDataType {
	// 人気で昇順ソート後のスライスを渡す
	// インデックスを値に基づいてソート
	sort.Slice(horsesData, func(i, j int) bool {
		return horsesData[i].Pops < horsesData[j].Pops
	})
	// 1レースの頭数
	_member_count := len(horsesData)
	// 集計する対象の頭数
	dataSize := math.Ceil(float64(_member_count) / 2.0)

	return horsesData[:int(dataSize)]
}

func CalcStdev(horsesData []HorseDataType) float64 {
	var oddsSum float64

	// 1レースの頭数
	// _member_count := len(horsesData)
	dataSize := len(horsesData)
	// 集計する対象の頭数
	// dataSize := math.Ceil(float64(_member_count) / 2.0)
	// 集計対象とする馬たちのオッズを取得する
	for i := 0; i < int(dataSize); i++ {
		oddsSum += horsesData[i].Odds
	}
	// オッズの平均
	oddsMu := oddsSum / float64(dataSize)
	// オッズの分散
	var muDiffTotal float64
	for i := 0; i < int(dataSize); i++ {
		_mu_diff := horsesData[i].Odds - oddsMu
		muDiffTotal += math.Pow(_mu_diff, 2)
	}
	oddsVariance := muDiffTotal / float64(dataSize)
	oddsStdev := math.Pow(oddsVariance, 0.5)

	fmt.Printf("  合計: %g, 平均: %g, 二乗誤差計: %g, 分散: %g, 標準偏差: %g\n",
		oddsSum, oddsMu, muDiffTotal, oddsVariance, oddsStdev)

	return oddsStdev
}
