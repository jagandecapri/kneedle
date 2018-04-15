package kneedle

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	//Test data based on that used in:
	//"Finding a Kneedle in a Haystack: Detecting Knee Points in System Behavior"

	testData := [][]float64{
	{0,0},
	{0.1, 0.55},
	{0.2, 0.75},
	{0.35, 0.825},
	{0.45, 0.875},
	{0.55, 0.9},
	{0.675, 0.925},
	{0.775, 0.95},
	{0.875, 0.975},
	{1,1},
	}

	kneePoints, _ := Run(testData, 1, 1, false)

	for _, kneePoint := range kneePoints{
		fmt.Println("Knee point:", kneePoint)
	}

	assert.Equal(t, 1, len(kneePoints))
	assert.InDeltaSlice(t, []float64{0.2, 0.75}, kneePoints[0], 1e-05)
}