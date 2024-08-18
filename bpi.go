package bpi

import "math"

// Calculates the "Beat Performance Index" of an IIDX score.
func calculate(exScore, kaidenAverage, worldRecord, maxScore int, powCoef float64) float64 {
	powCoefficient := powCoef

	if powCoef == 0 || powCoef < 0 {
		powCoefficient = 1.175
	}

	yourPGF := pikaGreatFunction(exScore, maxScore)
	kaidenPGF := pikaGreatFunction(kaidenAverage, maxScore)
	wrPGF := pikaGreatFunction(worldRecord, maxScore)

	yourScorePrime := yourPGF / kaidenPGF
	wrScorePrime := wrPGF / kaidenPGF

	isWorseThanKavg := exScore < kaidenAverage

	logExPrime := logToBase(yourScorePrime, wrScorePrime)

	if isWorseThanKavg {
		negativeRaisedValue := math.Pow((-1 * logExPrime), powCoefficient)

		bpi := 100 * -1 * negativeRaisedValue

		if bpi < -15 {
			return float64(-15)
		}

		return bpi
	} else {
		return 100 * math.Pow(logExPrime, powCoefficient)
	}

}

/*
Calculates the "PGF" of an IIDX score. This returns a number that indicates how
many pgreats you are expected to get for every great.
*/
func pikaGreatFunction(exScore, maxScore int) float64 {
	if exScore == maxScore {
		return float64(maxScore) * 0.8
	}

	scorePercent := float64(exScore) / float64(maxScore)

	return 0.5 / (1 - scorePercent)
}

/*
Returns the EX Score necessary to achieve the provided BPI.
*/
func inverse() {

}

func logToBase(number float64, base float64) float64 {
	return math.Log(number) / math.Log(base)
}
