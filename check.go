package main

import (
	"regexp"
)

var (
	regexTime     = regexp.MustCompile(`^\d{4}-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])$`) //2019-0909
	regexWebColor = regexp.MustCompile(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`)                 //#FF0011, #ff0011, #F01, #f01
)
