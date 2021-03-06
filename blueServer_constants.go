package main

// Paths on the server
const (
	moeVocabPath   = "/moe-vocab/"
	vocabPath      = "/vocab/"
	htmlLookupPath = "/html-lookup/"
	convertPath    = "/convert/"
	sentencesPath  = "/sentences/"
	mcdsPath       = "/mcds/"

	moeVocabLookupPath  = "/moe-vocab/lookup/"
	vocabLookupPath     = "/vocab/lookup/"
	convertLookupPath   = "/convert/lookup/"
	sentencesLookupPath = "/sentences/lookup/"
	mcdsLookupPath      = "/mcds/lookup/"

	settingsPath  = "/settings/"
	helpAboutPath = "/help-about/"

	assetsPath = "/assets/"
)

const layoutFile = "layout.html"

// mcd related constants
const (
	clozeBegin = `<span style="font-weight:600;color:#ff12c7;">`
	clozeEnd   = `</span>`
)
