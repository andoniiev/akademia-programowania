package academy

import "math"

type Student struct {
	Name       string
	Grades     []int
	Project    int
	Attendance []bool
}

// AverageGrade returns an average grade given a
// slice containing all grades received during a
// semester, rounded to the nearest integer.
func AverageGrade(grades []int) int {
	if len(grades) == 0 {
		return 0
	}
	var sum float64 = 0.0
	for _, grade := range grades {
		sum += float64(grade)
	}

	return int(math.Round(sum / float64(len(grades))))
}

// AttendancePercentage returns a percentage of class
// attendance, given a slice containing information
// whether a student was present (true) of absent (false).
//
// The percentage of attendance is represented as a
// floating-point number ranging from  0 to 1,
// with 2 digits of precision.
func AttendancePercentage(attendance []bool) float64 {
	attendedClasses := 0
	totalClasses := len(attendance)

	for _, attend := range attendance {
		if attend {
			attendedClasses++
		}
	}

	return float64(attendedClasses) / float64(totalClasses)
}

// FinalGrade returns a final grade achieved by a student,
// ranging from 1 to 5.
//
// The final grade is calculated as the average of a project grade
// and an average grade from the semester, with adjustments based
// on the student's attendance. The final grade is rounded
// to the nearest integer.

// If the student's attendance is below 80%, the final grade is
// decreased by 1. If the student's attendance is below 60%, average
// grade is 1 or project grade is 1, the final grade is 1.
func FinalGrade(s Student) int {
	projectGrade := s.Project
	semesterGrade := AverageGrade(s.Grades)

	attendancePercentage := AttendancePercentage(s.Attendance)

	if attendancePercentage < 0.6 || semesterGrade == 1 || projectGrade == 1 {
		return 1
	} else if attendancePercentage < 0.8 {
		finalGrade := (projectGrade + semesterGrade) / 2
		finalGrade -= 1
		return int(math.Round(float64(finalGrade)))
	} else {
		finalGrade := int(math.Round((float64(projectGrade + semesterGrade)) / 2))
		return finalGrade
	}
}

// GradeStudents returns a map of final grades for a given slice of
// Student structs. The key is a student's name and the value is a
// final grade.
func GradeStudents(students []Student) map[string]uint8 {
	grades := make(map[string]uint8)

	for _, student := range students {
		finalGrade := FinalGrade(student)
		grades[student.Name] = uint8(finalGrade)
	}

	return grades
}
