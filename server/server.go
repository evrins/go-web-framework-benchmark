package server

import (
	"fmt"
)

type BenchmarkFunction func(int)

var BFMap = map[string]BenchmarkFunction{
	"default":            startDefaultMux,
	"atreugo":            startAtreugo,
	"baa":                startBaa,
	"beego":              startBeego,
	"bone":               startBone,
	"bxog":               startBxog,
	"chi":                startChi,
	"denco":              startDenco,
	"don":                startDon,
	"echo":               startEcho,
	"fasthttp":           startFasthttp,
	"fasthttp-router":    startFastHTTPRouter,
	"fasthttp-routing":   startFastHTTPRouting,
	"fiber":              startFiber,
	"gear":               startGear,
	"gearbox":            startGearbox,
	"gin":                startGin,
	"go-craft-web":       startGocraftWeb,
	"go-frame":           startGoframe,
	"goji":               startGoji,
	"golf":               startGolf,
	"go-restful":         startGoRestful,
	"gorilla":            startGorilla,
	"go-router":          startGorouter,
	"go-router-fasthttp": startGorouterFastHTTP,
	"go-ozzo":            startGoozzo,
	"http-router":        startHTTPRouter,
	"http-tree-mux":      starthttpTreeMux,
	"httpz":              startHTTPZ,
	"indigo":             startIndigo,
	"lars":               startLars,
	"lion":               startLion,
	"muxie":              startMuxie,
	"negroni":            startNegroni,
	"pat":                startPat,
	"pulse":              startPulse,
	"pure":               startPure,
	"r2router":           startR2router,
	"tango":              startTango,
	"tiger":              startTigerTonic,
	"tiny-router":        startTinyRouter,
	"violetear":          startVioletear,
	"vulcan":             startVulcan,
	"webgo":              startWebgo,
	"goyave":             startGoyave,
}

func GetWebFramework(frameworkName string) (BenchmarkFunction, error) {
	fn, exists := BFMap[frameworkName]
	if !exists {
		return nil, fmt.Errorf("framework %s not found", frameworkName)
	}
	return fn, nil
}
