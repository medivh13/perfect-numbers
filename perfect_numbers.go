package perfect

import "errors"

type Classification string

const ClassificationDeficient Classification = "ClassificationDeficient"

const ClassificationPerfect Classification = "ClassificationPerfect"

const ClassificationAbundant Classification = "ClassificationAbundant"

var ErrOnlyPositive = errors.New("Not a positve number")

func Classify(n int64) (class Classification, err error) {
	if n <= 0 {
		return "", ErrOnlyPositive
	}
	factors := findFacotrs(n)
	var sum int64
	for _, factor := range factors {
		sum += int64(factor)
	}

	if sum < n {
		return ClassificationDeficient, nil
	}

	if sum > n {
		return ClassificationAbundant, nil
	}

	return ClassificationPerfect, nil
}

func findFacotrs(n int64) (factors []int) {
	for i := 1; int64(i) <= n/2; i++ {
		if int(n)%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}
