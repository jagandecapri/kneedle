package kneedle

import (
	"math"
	"github.com/pkg/errors"
)

/*
  Given set of values, look for the elbow/knee points.
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

//Run takes in a 2D slice containing data where knee or elbow needs to be found.
//The function also takes in the number of "flat" points that is required before considering
//a point as knee or elbow. The smootingWindow parameter is used to indicate the avarage used for
//the Gaussian kernel average smoother (you can try with 3 to begin with). The findElbows parameter indicates
//whether to find an elbow or a knee when the value of parameter is true or false respectively
func Run(data [][]float64, s float64, smoothingWindow int, findElbows bool) (localMinMaxPts [][]float64, err error){

	if(len(data) == 0){
		err = errors.New("Cannot find elbow or knee points in empty data.")
		return
	}

	if(len(data[0]) != 2){
		err = errors.New("Cannot run Kneedle, this method expects all data to be 2d.")
		return
	}

	//do steps 1,2,3 of the paper in the prepare method
	normalisedData := prepare(data, smoothingWindow)
	//find candidate indices (this is step 4 in the paper)

	candidateIndices := findCandidateIndices(normalisedData, findElbows)
	//go through each candidate index, i, and see if the indices after i are satisfy the threshold requirement
	//(this is step 5 in the paper)
	step := computeAverageVarianceX(normalisedData)

	if findElbows{
		step = step * s
	} else {
		step = step * -s
	}

	//check each candidate to see if it is a real elbow/knee
	//(this is step 6 in the paper)
	for i := 0; i < len(candidateIndices); i++{
		candidateIdx := candidateIndices[i]
		var endIdx int

		if i + 1 < len(candidateIndices){
			endIdx = candidateIndices[i+1]
		} else {
			endIdx = len(data)
		}

		threshold := normalisedData[candidateIdx][1] + step

		for j := candidateIdx + 1; j < endIdx; j++{
			var isRealElbowOrKnee bool
			if findElbows{
				isRealElbowOrKnee = normalisedData[j][1] > threshold
			} else {
				isRealElbowOrKnee = normalisedData[j][1] < threshold
			}

			if isRealElbowOrKnee {
				localMinMaxPts = append(localMinMaxPts, data[candidateIdx])
				break
			}
		}
	}

	return
}