package esepunittests

type GradeCalculator struct {
	studentGrades []Grade
	PassOrFail    bool
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		studentGrades: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	} else if numericalGrade <= 59 && numericalGrade >= 0 {
		return "F"
	}

	return "N/A"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	switch gradeType {
	case Assignment:
		gc.studentGrades = append(gc.studentGrades, Grade{
			Name:  name,
			Grade: grade,
			Type:  Assignment,
		})

	case Exam:
		gc.studentGrades = append(gc.studentGrades, Grade{
			Name:  name,
			Grade: grade,
			Type:  Exam,
		})

	case Essay:
		gc.studentGrades = append(gc.studentGrades, Grade{
			Name:  name,
			Grade: grade,
			Type:  Essay,
		})

	}
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	var assignments []Grade
	var exams []Grade
	var essays []Grade
	assignment_weight := 0.5
	exam_weight := 0.35
	essay_weight := 0.15

	for _, grade := range gc.studentGrades {
		currType := grade.Type
		switch currType {
		case Assignment:
			assignments = append(assignments, grade)
		case Exam:
			exams = append(exams, grade)
		case Essay:
			essays = append(essays, grade)
		}
	}
	assignment_average := computeAverage(assignments)
	exam_average := computeAverage(exams)
	essay_average := computeAverage(essays)

	sum_weights := 0.0

	if len(assignments) != 0 {
		sum_weights += assignment_weight
	}

	if len(exams) != 0 {
		sum_weights += exam_weight
	}

	if len(essays) != 0 {
		sum_weights += essay_weight
	}

	weighted_grade := float64(assignment_average)*assignment_weight + float64(exam_average)*exam_weight + float64(essay_average)*essay_weight

	if sum_weights == 0 {
		return -1
	}

	weighted_grade = weighted_grade / sum_weights

	return int(weighted_grade)
}

func computeAverage(grades []Grade) int {
	if len(grades) == 0 {
		return 0
	}

	sum := 0

	for _, grade := range grades {
		sum += grade.Grade
	}

	return sum / len(grades)
}
