package utils

func GetCacheTxId(algoType, grayll_tx_id string) (string, string) {
	cacheTxId := ""
	unreadPath := ""
	switch algoType {
	case "GRZ":
		unreadPath = "UrGRZ"
		cacheTxId = "z1" + grayll_tx_id
	case "GRY 1":
		unreadPath = "UrGRY1"
		cacheTxId = "y1" + grayll_tx_id
	case "GRY 2":
		unreadPath = "UrGRY2"
		cacheTxId = "y2" + grayll_tx_id
	case "GRY 3":
		unreadPath = "UrGRY3"
		cacheTxId = "y3" + grayll_tx_id
	}
	return cacheTxId, unreadPath
}
