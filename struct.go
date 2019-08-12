package main

type Calendar struct {
	Layers []Layer
}

type Layer struct {
	Title     string
	Color     string //#FF3366
	Greyscale bool
	Hidden    bool
	Schedules []Schedule
}

type Schedule struct {
	Title   string
	Timein  string
	Timeout string
}
