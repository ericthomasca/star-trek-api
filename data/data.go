package data

import "github.com/ericthomasca/star-trek-api/models"

var SeriesData = []models.Series{
	TOSData,
	TASData,
	TNGData,
	DS9Data,
	VOYData,
	ENTData,
}

func strPtr(s string) *string {
	return &s
}

func floatPtr(f float64) *float64 {
	return &f
}