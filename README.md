# Kneedle-Go

## Description

A Go package for the Kneedle algorithm which can be used to detect a knee in a list of values. This package is adapted into Go from the Java implementation at [github.com/lukehb/137-stopmove](https://github.com/lukehb/137-stopmove/blob/master/src/main/java/onethreeseven/stopmove/algorithm/Kneedle.java).

## Installing

Use `go get github.com/jagandecapri/kneedle`

## Example use

```Go
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
	
    	kneePoints, _ := kneedle.Run(testData, 1, 1, false)
    
    	for _, kneePoint := range kneePoints{
        	fmt.Println("Knee point:", kneePoint)
    	}
```

## References

Ville Satopää, Jeannie Albrecht, David Irwin, Barath Raghavan. [Finding a "Kneedle" in a Haystack: Detecting Knee Points in System Behavior](http://ieeexplore.ieee.org/xpl/login.jsp?tp=&arnumber=5961514&url=http%3A%2F%2Fieeexplore.ieee.org%2Fxpls%2Fabs_all.jsp%3Farnumber%3D5961514). 31st International Conference on Distributed Computing Systems Workshops, pp. 166-171, Minneapolis, Minnesota, USA, June 2011.

## Authors

* **[Jagatheesan Jack](https://github.com/jagandecapri)** - *Initial work*

See also the list of [contributors](https://github.com/jagandecapri/kneedle/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
