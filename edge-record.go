package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	fd "github.com/digisan/gotk/filedir"
)

type EdgeRecord struct {
	Name     string
	Tm       string
	FilePath string
	Pts      []struct {
		X          int
		Y          int
		GrayPeaks  []byte
		RedPeaks   []byte
		GreenPeaks []byte
		BluePeaks  []byte
	}
}

func (r *EdgeRecord) String() (s string) {
	s += fmt.Sprintln("Name:      ", r.Name)
	s += fmt.Sprintln("Created:   ", r.Tm)
	s += fmt.Sprintln("Image File:", r.FilePath)
	for i, pt := range r.Pts {
		s += fmt.Sprintf("%d:", i)
		s += fmt.Sprintln(" X:", pt.X, " Y:", pt.Y)
		s += fmt.Sprintln(" -- GrayPeaks: ", pt.GrayPeaks)
		s += fmt.Sprintln(" -- RedPeaks:  ", pt.RedPeaks)
		s += fmt.Sprintln(" -- GreenPeaks:", pt.GreenPeaks)
		s += fmt.Sprintln(" -- BluePeaks: ", pt.BluePeaks)
	}
	return
}

func NewEdgeRecord(name, filePath string) *EdgeRecord {
	abspath, err := fd.AbsPath(filePath, true)
	if err != nil {
		log.Fatalln(err)
	}
	return &EdgeRecord{
		Name:     name,
		Tm:       time.Now().Format(time.RFC3339),
		FilePath: abspath,
	}
}

func LoadLastRecord(jaFile string) *EdgeRecord {

	data, err := os.ReadFile(jaFile)
	if err != nil {
		log.Fatalln(err)
	}
	js := string(data)

	start := strings.LastIndex(js, `"Name":`) // indicator for searching
	end := strings.LastIndex(js, "}")
	block := "{" + js[start:end] + "}"

	pts := &EdgeRecord{}
	json.Unmarshal([]byte(block), pts)
	return pts
}

func (r *EdgeRecord) AddPtInfo(x, y int, grayPeaks, rPeaks, gPeaks, bPeaks []byte) {
	r.Pts = append(
		r.Pts,
		struct {
			X          int
			Y          int
			GrayPeaks  []byte
			RedPeaks   []byte
			GreenPeaks []byte
			BluePeaks  []byte
		}{
			X:          x,
			Y:          y,
			GrayPeaks:  grayPeaks,
			RedPeaks:   rPeaks,
			GreenPeaks: gPeaks,
			BluePeaks:  bPeaks,
		})
}

func (r *EdgeRecord) Log(jafile string) {

	newFile := false
	if !fd.FileExists(jafile) {
		if os.WriteFile(jafile, []byte{'[', ']'}, os.ModePerm) != nil {
			log.Fatalln("PtsRecord Log Error @ Creating first json array file")
		}
		newFile = true
	}

	prevData, err := os.ReadFile(jafile)
	if err != nil {
		log.Fatalln("PtsRecord Log Error @ ReadFile")
	}
	end := bytes.LastIndex(prevData, []byte{']'})
	prevData = prevData[:end]

	if !newFile {
		prevData = append(prevData, ',')
	}

	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		log.Fatalln("PtsRecord Log Error @ Marshal")
	}

	added := bytes.Join([][]byte{prevData, data}, []byte{})
	added = append(added, ']')

	os.WriteFile(jafile, added, os.ModePerm)
}
