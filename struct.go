package main

// Calendar 자료구조
type Calendar struct {
	Layers []Layer
}

//Layer 자료구조
type Layer struct {
	Title     string
	Color     string //#FF3366
	Greyscale bool
	Hidden    bool
	Schedules []Schedule
}

// Schedule 자료구조
type Schedule struct {
	Title string
	Start string
	End   string
}
