package utils

import "geektrust/core/program"

// MapStringToProgramCategory - Maps the program category to enum.
// The enums will helps the code maintainable with minimal amount of memory allocation.
func MapStringToProgramCategory(input string) program.Category {
	switch input {
	case program.CategoryCertification.String():
		return program.CategoryCertification
	case program.CategoryDegree.String():
		return program.CategoryDegree
	case program.CategoryDiploma.String():
		return program.CategoryDiploma
	}
	return program.CategoryUnknown
}
