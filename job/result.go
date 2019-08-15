package json

func GetOperationStatus(targetTrain string, delayTrains DelayTrains) string {
	operationStatus := "正常に運行しています"
	for _, train := range delayTrains {
		if targetTrain == train.Name {
			operationStatus = "遅延しています。"
		}
	}
	return operationStatus
}
