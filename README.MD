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
go test -coverprofile=coverage.txt ./...
```

## Files:

README.md: Describes the project, requirements, and gives instructions on how to run the code.

/grade-calculator/go.mod: Defines the module’s path and Go version.

/grade-calculator/grade_calculator.go: Defines the structure of the grades, grade types, and grade calculators. Also allows for the adding of grades and calculates the numerical and final grades in each category.

/grade-calculator/grade_calculator_test.go: Unit tests calling the functions defined in grade_calculator.go and checking if they are working correctly.


## Unit tests (/grade-calculator/grade_calculator_test.go):

TestGetGradeA() creates a grade calculator and fills it with grades that are 100% for each category, ensuring that the final grade must be an A. If the outcome is not the expected grade (A) then an error is sent. This test should pass.

TestGetGradeB() creates a grade calculator and fills it with grades between 80% and 85%, which should result in a B grade. Sends an error message if the resulting grade is not a B. This test should pass.

TestGetGradeF() creates a grade calculator and fills it with grades between 91% and 100% in each category, expecting an F grade despite the inputs averaging out to an A. Sends an error if the resulting grade is not an F. This test should fail.

## Resources:

https://www.w3schools.com/go/index.php
