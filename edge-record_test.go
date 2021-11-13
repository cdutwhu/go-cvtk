package main

import (
	"testing"
)

func TestEdgeRecord(t *testing.T) {
	record := NewEdgeRecord("AB", "./in/sample/1.jpg")
	record.AddPtInfo(
		100,
		100,
		[]byte{100, 200},
		[]byte{34, 5, 6},
		[]byte{12},
		[]byte{44, 23},
	)
	record.AddPtInfo(
		200,
		200,
		[]byte{100, 200},
		[]byte{34, 5, 6},
		[]byte{12},
		[]byte{44, 23},
	)
	record.Log("./out/AB-edge.json")
}

// func TestLoadEdgeRecord(t *testing.T) {
// 	edge := LoadLastRecord("./out/AB-edge.json")
// 	fmt.Println(edge)
// 	fmt.Println(edge.Pts[0].GreenPeaks[0])
// }
