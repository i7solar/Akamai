package main

import (
	"fmt"
	os "github.com/ojrac/opensimplex-go"
	"math/rand"
	"strconv"
	"time"
)

// MactEntity MactEntity Object */
type MactEntity struct {
	indexValue []int64
	eventType  []int64
	eventTime  []int64
	xPos       []int64
	yPos       []int64
}

// Init Allocates the Mact arrays and allocates the index value array based on the number (99.)
func Init() MactEntity {
	/* System Design Idea: Use a random number rather than the 0-99 values. */
	//indexAmount := int(bmak.GenerateRandomNumber(62, 98))
	indexAmount := 99
	indexValue := make([]int64, indexAmount)
	eventType := make([]int64, indexAmount)
	eventTime := make([]int64, indexAmount)
	xPos := make([]int64, indexAmount)
	yPos := make([]int64, indexAmount)

	for i := 0; i < indexAmount; i++ {
		indexValue[i] = int64(i + 1)
	}

	for i := 0; i < indexAmount; i++ {
		eventType[i] = 1
	}

	return MactEntity{indexValue, eventType, eventTime, xPos, yPos}
}

/* buildMact - Returns the string and values all computed for bmak.ta. */
func buildMact(entity MactEntity, it0Flag bool) (string, int64) {

	var (
		mactEntity, mactString string
		mactTotal              int64
		i                      int
	)

	if it0Flag == true {
		for i = 1; i < len(entity.indexValue); i++ {

			if i == len(entity.indexValue) {
				mactEntity = strconv.Itoa(i) + "," + strconv.FormatInt(entity.eventType[i], 10) + "," + strconv.FormatInt(entity.eventTime[i], 10) + "," + strconv.FormatInt(entity.xPos[i], 10) + "," + strconv.FormatInt(entity.yPos[i], 10)
				mactTotal += int64(i) + entity.eventType[i] + entity.eventTime[i] + entity.xPos[i] + entity.yPos[i]
			} else {
				mactEntity = strconv.Itoa(i) + "," + strconv.FormatInt(entity.eventType[i], 10) + "," + strconv.FormatInt(entity.eventTime[i], 10) + "," + strconv.FormatInt(entity.xPos[i], 10) + "," + strconv.FormatInt(entity.yPos[i], 10) + ";"
				mactTotal += int64(i) + entity.eventType[i] + entity.eventTime[i] + entity.xPos[i] + entity.yPos[i]
			}

			mactString += mactEntity
		}
	} else {
		for i = 0; i < len(entity.indexValue); i++ {

			if i == len(entity.indexValue) {
				mactEntity = strconv.Itoa(i) + "," + strconv.FormatInt(entity.eventType[i], 10) + "," + strconv.FormatInt(entity.eventTime[i], 10) + "," + strconv.FormatInt(entity.xPos[i], 10) + "," + strconv.FormatInt(entity.yPos[i], 10)
				mactTotal += int64(i) + entity.eventType[i] + entity.eventTime[i] + entity.xPos[i] + entity.yPos[i]
			} else {
				mactEntity = strconv.Itoa(i) + "," + strconv.FormatInt(entity.eventType[i], 10) + "," + strconv.FormatInt(entity.eventTime[i], 10) + "," + strconv.FormatInt(entity.xPos[i], 10) + "," + strconv.FormatInt(entity.yPos[i], 10) + ";"
				mactTotal += int64(i) + entity.eventType[i] + entity.eventTime[i] + entity.xPos[i] + entity.yPos[i]
			}

			mactString += mactEntity
		}
	}

	return mactString, mactTotal
}

/* timestampPivotsGen - Creates an array of pivot numbers, based on the number. */
func timestampPivotsGen(num int) []int {
	tspivs := make([]int, num)
	for i := 0; i < num; i++ {
		tspivs[i] = randNumGen(100)
	}
	return tspivs
}

/* genMact - Returns the string, total of all values, and total of timestamp values... Uses the it0 flag to start at 1 with it0 as true and false to start at 0. */
func simplexNoiseGen(it0Flag bool) (string, int64, int64) {
	Mact := Init()
	noise := os.NewNormalized(rand.Int63())
	rand.Seed(time.Now().UnixNano())

	w, h := 100, 100
	heightmap := make([]float64, w*h)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			xFloat := float64(x) / float64(w)
			yFloat := float64(y) / float64(h)
			heightmap[(y*w)+x] = noise.Eval2(xFloat, yFloat)
		}
	}
	staticrandnum := randNumGen(1500)
	for i := 0; i < len(heightmap); i++ {
		heightmap[i] = (heightmap[i] * 1000) + float64(staticrandnum)
	}

	ystart := randNumGen(800)
	ts := randNumGen(2000)
	totaltime := 0

	for i := 1; i < len(Mact.indexValue); i++ {
		// set the ts
		ts += randNumGen(14)
		totaltime += ts
		Mact.eventTime[i] = int64(ts)

		Mact.xPos[i] = int64(heightmap[i])

		ystart += randNumGen(17)
		Mact.yPos[i] = int64(ystart)
	}

	mactString, mactCompleteTotal := buildMact(Mact, it0Flag)

	return mactString, int64(totaltime), mactCompleteTotal
}

/* Obsolete. */
func mathGen(it0Flag bool) (string, int64, int64) {

	var bmak Bmak

	Mact := Init()

	tspivots := timestampPivotsGen(int(bmak.GenerateRandomNumber(2, 5)))
	xstart := bmak.GenerateRandomNumber(800, 1600)
	ystart := bmak.GenerateRandomNumber(500, 800)
	ts := randNumGen(2000)

	xpivot := bmak.GenerateRandomNumber(45, 65)

	totaltime := 0

	for i := 1; i < len(Mact.indexValue); i++ {
		// set the ts
		ts += randNumGen(14)
		totaltime += ts
		Mact.eventTime[i] = int64(ts)

		// make proper adjustment to timestamp
		for _, tspivot := range tspivots {
			if i == tspivot {
				ts += randNumGen(1000)
			}
		}

		// set x pos
		if int64(i) < xpivot {
			xstart -= int64(randNumGen(20))
			Mact.xPos[i] = xstart
		} else if int64(i) > xpivot {
			xstart += int64(randNumGen(30))
			Mact.xPos[i] = int64(xstart)
		} else if int64(i) == xpivot {
			xstart = bmak.GenerateRandomNumber(800, 1600)
			Mact.xPos[i] = xstart
		}
		ystart += int64(randNumGen(17))
		// set y pos
		Mact.yPos[i] = ystart

	}

	mactString, mactCompleteTotal := buildMact(Mact, it0Flag)

	return mactString, int64(totaltime), mactCompleteTotal
}

/* getMact - [DEBUG] :: Display MACT strings. */
func getMact() {
	for i := 0; i < 1; i++ {
		//mactgen, timestamp_total, mact_total  := mathGen(true)
		mactgen, timestamp_total, mact_total := simplexNoiseGen(true)
		fmt.Println("Mact " + strconv.Itoa(i) + ":\n" + mactgen)
		fmt.Println("Mact " + strconv.Itoa(i) + " timestamp total:" + strconv.FormatInt(timestamp_total, 10))
		fmt.Println("Mact " + strconv.Itoa(i) + " mact total:" + strconv.FormatInt(mact_total, 10))
	}
}
