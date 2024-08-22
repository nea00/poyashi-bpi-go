package bpi

import "math"

// Calculates the "Beat Performance Index" of an IIDX score.
func Calculate(exScore, kaidenAverage, worldRecord, maxScore int, powCoef float64) float64 {
	powCoefficient := powCoef

	if powCoef == 0 || powCoef < 0 {
		powCoefficient = 1.175
	}

	yourPGF := PikaGreatFunction(exScore, maxScore)
	kaidenPGF := PikaGreatFunction(kaidenAverage, maxScore)
	wrPGF := PikaGreatFunction(worldRecord, maxScore)

	yourScorePrime := yourPGF / kaidenPGF
	wrScorePrime := wrPGF / kaidenPGF

	isWorseThanKavg := exScore < kaidenAverage

	logExPrime := LogToBase(yourScorePrime, wrScorePrime)

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
func PikaGreatFunction(exScore, maxScore int) float64 {
	if exScore == maxScore {
		return float64(maxScore) * 0.8
	}

	scorePercent := float64(exScore) / float64(maxScore)

	return 0.5 / (1 - scorePercent)
}

// Inverts the PGF by returning the EXScore necessary to get a given PGF.
func InversePikaGreatFunction(pgf, maxScore float64) float64 {
	return (pgf*maxScore - 0.5*maxScore) / pgf
}

// Returns the EX Score necessary to achieve the provided BPI.
func Inverse(bpi float64, kaidenAverage, worldRecord, maxScore int, powCoef float64) int {
	powCoefficient := powCoef

	if powCoef == 0 || powCoef < 0 {
		powCoefficient = 1.175
	}

	isWorseThanKavg := bpi < 0

	var logExPrime float64

	if isWorseThanKavg {
		logExPrime = -1 * math.Pow((bpi/float64(-100)), (1/powCoefficient))

	} else {
		logExPrime = math.Pow((bpi / float64(100)), (1 / powCoefficient))
	}

	kaidenPGF := PikaGreatFunction(kaidenAverage, maxScore)
	wrPGF := PikaGreatFunction(worldRecord, maxScore)

	wrScorePrime := wrPGF / kaidenPGF
	exScorePrime := math.Pow(wrScorePrime, logExPrime)

	exScorePGF := exScorePrime * kaidenPGF

	exScore := InversePikaGreatFunction(exScorePGF, float64(maxScore))

	if bpi > 100 && math.Ceil(exScore) == float64(maxScore) {
		return maxScore
	}

	return int(math.Round(exScore))
}

func LogToBase(number float64, base float64) float64 {
	return math.Log(number) / math.Log(base)
}
