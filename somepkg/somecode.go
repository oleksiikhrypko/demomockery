package somepkg

type DataProvider interface {
	GetData(idx int) (int, error)
}

func SomeLogic(data DataProvider) (int, error) {
	first, err := data.GetData(1)
	if err != nil {
		return 0, err
	}
	second, err := data.GetData(first)
	if err != nil {
		return 0, err
	}
	return first + second, nil
}
