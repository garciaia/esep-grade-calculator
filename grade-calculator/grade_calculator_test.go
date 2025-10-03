package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.PassOrFail = false

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.PassOrFail = false

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.PassOrFail = false

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 79, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 78, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.PassOrFail = false

	gradeCalculator.AddGrade("open source assignment", 58, Assignment)
	gradeCalculator.AddGrade("exam 1", 67, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 69, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.PassOrFail = false

	gradeCalculator.AddGrade("open source assignment", 59, Assignment)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeTypeAssignment(t *testing.T) {
	expected_value := "assignment"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.PassOrFail = false

	gradeCalculator.AddGrade("open source assignment", 59, Assignment)

	actual_value := gradeCalculator.studentGrades[0].Type.String()

	if expected_value != actual_value {
		t.Errorf("Expected String() to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestPass(t *testing.T) {
	expected_value := "Pass"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.PassOrFail = true

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected String() to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestFail(t *testing.T) {
	expected_value := "Fail"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.PassOrFail = true

	gradeCalculator.AddGrade("open source assignment", 59, Assignment)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected String() to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestNoGrades(t *testing.T) {
	expected_value := "N/A"

	gradeCalculator := NewGradeCalculator()
	gradeCalculator.PassOrFail = true

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected String() to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
