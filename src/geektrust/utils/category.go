package utils

import "geektrust/domain/program"

func MapStringToProgramCategory(input string) program.ProgramCategory {
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
