# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test ./... -coverprofile=coverage.txt
```

## Files:

README.md: Describes the project, requirements, and gives instructions on how to run the code.

/grade-calculator/go.mod: Defines the module’s path and Go version.

/grade-calculator/grade_calculator.go: Defines the structure of the grades, grade types, and grade calculators. Also allows for the adding of grades and calculates the numerical and final grades in each category.

/grade-calculator/grade_calculator_test.go: Unit tests calling the functions defined in grade_calculator.go and checking if they are working correctly.


## Unit tests (/grade-calculator/grade_calculator_test.go):

TestGetGradeA() creates a grade calculator and fills it with grades that are 100% for each category, ensuring that the final grade must be an A. If the outcome is not the expected grade (A) then an error is sent.

TestGetGradeB() creates a grade calculator and fills it with grades between 80% and 85%, which should result in a B grade. Sends an error message if the resulting grade is not a B.

TestGetGradeC() creates a grade calculator and fills it with grades between 70% and 79%, which should result in a C grade. Sends an error message if the resulting grade is not a C.

TestGetGradeD() creates a grade calculator and fills it with grades between 58% and 69%, which should result in a D grade. Sends an error message if the resulting grade is not a D.

TestGetGradeF() creates a grade calculator and fills it with a single grade of 59% for assignments. Sends an error message if the resulting grade is not a D.

TestGetGradeTypeAssignment() creates a grade calculator with one grade of type assignment and sends an error message if the type is not "assignment".

TestPass() creates & fills the grade calculator with a single assignment at 60%. Sets PassOrFail boolean to true. Returns an error if the final grade is not "Pass".

TestFail() creates & fills the grade calculator with a single assignment at 59%. Sets PassOrFail boolean to true. Returns an error if the final grade is not "Fail".

All the tests should pass.

## Resources:

https://www.w3schools.com/go/index.php
