package kneedle

import (
	"math"
	"github.com/jagandecapri/vision/data"
)

/*
  Given set of values look for the elbow/knee points.
  See paper: "Finding a Kneedle in a Haystack: Detecting Knee Points in System Behavior"
  @author Jagatheesan
*/

// findCancidateIndices finds the indices of all local minimum or local maximum values
// where findMinima is to indicate whether to find local minimums or local maximums.
func findCandidateIndices(data [][]float64, findMinima bool) (candidates []int){
	//a coordinate is considered a candidate if both of its adjacent points have y-values
	//that are greater or less (depending on whether we want local minima or local maxima)
	for i := 1; i < len(data) - 1; i++{
		prev := data[i-1][1]
		cur := data[i][1]
		next := data[i+1][1]
		var isCandidate bool
		if findMinima == true{
			isCandidate = prev > cur && next > cur
		} else {
			isCandidate = prev < cur && next < cur
		}
		if(isCandidate){
			candidates = append(candidates, i)
		}
	}
	return
}

//findElbowIndex fings the index in the data the represents a most exaggerated elbow point.
func findElbowIndex(data []float64) (bestIdx int){
	var bestScore float64

	for i := 0; i < len(data); i++{
		score := math.Abs(data[i])
		if score > bestScore{
			bestScore = score
			bestIdx = i
		}
	}
	return bestIdx
}

//Prepare prepares the data by smoothing, then normalising into unit range 0-1,
//and finally, subtracting the y-value from the x-value where
//smoothingWindow is the size of the smoothing window.
func prepare(data [][]float64, smoothingWindow int) (normalisedData [][]float64){
	//smooth the data to make local minimum/maximum easier to find (this is Step 1 in the paper)
	smoothedData, _ := gaussianSmooth2d(data, smoothingWindow)

	//prepare the data into the unit range (step 2 of paper)
	normalisedData, _ = minmaxNormalise(smoothedData)

	//subtract normalised x from normalised y (this is step 3 in the paper)
	for i := 0; i < len(normalisedData); i++{
		normalisedData[i][1] = normalisedData[i][1] - normalisedData[i][0]
	}
	return
}

func computeAverageVarianceX(data [][]float64) float64{
	var sumVariance float64
	for i := 0; i < len(data) - 1; i++{
		sumVariance += data[i + 1][0] - data[i][0]
	}
	return sumVariance / float64((len(data) - 1))
}