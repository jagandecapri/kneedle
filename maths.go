package kneedle

import (
	"math"
	"github.com/pkg/errors"
)

//gaussian calculates the gaussian of an input
//where height is the height of the center of the curve (sometimes called 'a'),
//center is the center of the curve (sometimes called 'b'),
//and width is the standard deviation, i.e ~68% of the data will be contained in center ± the width. /
func gaussian(x float64, height float64, center float64, width float64) float64{
	return height * math.Exp(-(x-center)*(x-center)/(2.0*width*width) )
}

//gaussianSmooth2d smooths the data using a gaussian kernel
//where w is the size of sliding window (i.e number of indices either side to sample).
func gaussianSmooth2d(data [][]float64, w int) (smoothed [][]float64, err error){
	dataSize := len(data)

	if(dataSize == 0){
		err = errors.New("Cannot smooth empty data.")
		return
	}

	nDims := len(data[0])

	if(nDims == 0){
		err = errors.New("Cannot smooth a data point with no values. Uniformly populate every entry in your data with 1 or more dimensions.")
	}

	smoothed = make([][]float64, dataSize)

	for i := 0; i < dataSize; i++{
		var startIdx, endIdx int

		if 0 < i -w {
			startIdx = i - w
		}

		if dataSize - 1 < i + w{
			endIdx = dataSize - 1
		} else {
			endIdx = i + w
		}

		sumWeights := make([]float64, nDims)
		var sumIndexWeight float64

		for j := startIdx; j < endIdx + 1; j++{
			indexScore := math.Abs(float64(j - i))/float64(w)
			indexWeight := gaussian(indexScore, 1, 0, 1);

			for n := 0; n < nDims; n++{
				sumWeights[n] += (indexWeight * data[j][n])
			}
			sumIndexWeight += indexWeight
		}

		tmp := make([]float64, nDims)
		for n := 0; n < nDims; n++{
			tmp[n] = sumWeights[n]/sumIndexWeight
		}

		smoothed[i] = tmp
	}

	return
}

// minmaxNormalise performs min-max normalisation on n-dimensional data (as long as the dimensionality is uniform, that is, all data is 2d or all 3d etc.).
// For reference, refer to <a href="https://en.wikipedia.org/wiki/Feature_scaling#Rescaling">Wikipedia article about feature re-scaling.</a>
func minmaxNormalise(data [][]float64) (outputNormalised [][]float64, err error){

	dataSize := len(data)

	if(dataSize == 0){
		err = errors.New("Cannot smooth empty data.")
		return
	}

	nDims := len(data[0])

	if(nDims == 0){
		err = errors.New("Cannot smooth a data point with no values. " +
		"Uniformly populate every entry in your data with 1 or more dimensions.")
		return
	}

//1) get min and max for each dimension of the data

	minEachDim := make([]float64, nDims)
	maxEachDim := make([]float64, nDims)
	for i := 0; i < nDims; i++ {
		minEachDim[i] = math.MaxFloat64
		maxEachDim[i] = math.SmallestNonzeroFloat64
	}

	for _, coords := range data{
		for n := 0; n < nDims; n++ {
			v := coords[n]
			if (v < minEachDim[n]) {
				minEachDim[n] = v
			}
			if (v > maxEachDim[n]) {
				maxEachDim[n] = v
			}
		}
	}

	//2) normalise the data using the min and max
	rangeEachDim := make([]float64, nDims)
	for n := 0; n < nDims; n++{
		rangeEachDim[n] = maxEachDim[n] - minEachDim[n]
	}

	outputNormalised = make([][]float64, dataSize)
	for i := 0; i < dataSize; i++{
		tmp := make([]float64, nDims)
		for n := 0; n < nDims; n++{
			//normalising step
			tmp[n] = (data[i][n] - minEachDim[n]) / rangeEachDim[n]
		}
		outputNormalised[i] = tmp
	}
	return
}