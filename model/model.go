// /home/krylon/go/src/github.com/blicero/uruk/model/model.go
// -*- mode: go; coding: utf-8; -*-
// Created on 16. 12. 2024 by Benjamin Walkenhorst
// (c) 2024 Benjamin Walkenhorst
// Time-stamp: <2024-12-17 19:27:22 krylon>

// Package model provides data types that model the game.
package model

// City is a city state
type City struct {
	Name        string
	Age         uint
	Population  uint
	SoilQuality float32
}
