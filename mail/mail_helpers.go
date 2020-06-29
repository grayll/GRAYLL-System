package mail

import (
	"fmt"
	// "log"
	// "time"
	//"bitbucket.org/grayll/grayll.io-grz-arkady/api"
)

// func GenOpenPositionGRY(position AlgorithmPosition) (string, string, string, []string) {

// 	title := fmt.Sprintf(`GRAYLL | %s Algo Position Opened`, position.AlgorithmType)
// 	openFee := GetManagementFeePercent(position.AlgorithmType)
// 	contents := []string{
// 		time.Unix(position.OpenPositionTimestamp, 0).Format(`15:04 | 02-01-2006`),

// 		fmt.Sprintf(`You have opened a new %s algorithmic position.`, GetAlgorithmString(position.AlgorithmType)),

// 		fmt.Sprintf(`Open Algo Position Value: $%.5f`, position.OpenPositionValueUSD),

// 		fmt.Sprintf(`%s Rate | $%.5f`, position.AlgorithmType[:3], position.OpenValueGRY),

// 		fmt.Sprintf(`%s | Algo Management Fee @ %.1f%% | $%.5f`, position.AlgorithmType, openFee, position.OpenPositionFeeUSD),

// 		fmt.Sprintf(`GRX Rate | $%.5f`, position.OpenValueGRX),

// 		fmt.Sprintf(`GRX | Total Algo Position Value | %.7f GRX`, position.OpenPositionTotalGRX),

// 		fmt.Sprintf(`%s | Algo Management Fee @ %.1f%% | %.7f GRX`, position.AlgorithmType, openFee, position.OpenPositionFeeGRX),

// 		fmt.Sprintf(`GRAYLL | Transaction ID | %s`, position.GrayllTxId),
// 	}
// 	content := ""
// 	for i, sent := range contents {
// 		if i > 0 && i != 2 {
// 			content = content + ". " + sent
// 		} else {
// 			content = content + sent
// 		}
// 	}

// 	url := "Stellar | Transaction ID | https://stellar.expert/explorer/public/search?term=" + position.StellarTxId
// 	log.Println(content)
// 	contents = append(contents, url)

// 	return title, content, url, contents
// }

// func GetManagementFeePercent(algorithm string) float64 {
// 	switch algorithm {
// 	case "GRZ":
// 		return 0.3
// 	case "GRY 1":
// 		return 1.8
// 	case "GRY 2":
// 		return 1.3
// 	case "GRY 3":
// 		return 0.9
// 	}
// 	return 0
// }
// func GenPerfChange(positionMap map[string]interface{}, durationString, algorithm string) (string, string, string, []string) {

// 	title := fmt.Sprintf(`GRAYLL | %s Algo Position Performance`, algorithm)
// 	openTime := int64(positionMap["open_position_timestamp"].(float64))
// 	last12hROIPercent := positionMap["current_position_ROI_percent"].(float64)
// 	if val, ok := positionMap["last_12h_ROI_percent"]; ok {
// 		last12hROIPercent = val.(float64)
// 	}
// 	contents := []string{
// 		time.Unix(openTime, 0).Format(`15:04 | 02-01-2006`),

// 		fmt.Sprintf(`Your scheduled %s algorithmic position performance notification for: `, algorithm),

// 		fmt.Sprintf(`GRAYLL | Transaction ID | %s`, positionMap["grayll_transaction_id"].(string)),

// 		fmt.Sprintf(`%s Algorithmic Position | 24 hour ROI: %.5f%%`, algorithm, last12hROIPercent),

// 		fmt.Sprintf(`Algo Position Duration: %v`, durationString),

// 		fmt.Sprintf(`Current Algo Position Value: $%.5f`, positionMap["current_position_value_$"].(float64)),

// 		fmt.Sprintf(`Current GRX Rate | $%.7f`, positionMap["current_value_GRX"].(float64)),

// 		fmt.Sprintf(`Current Algo Position Value: %.7f GRX`, positionMap["current_position_value_GRX"].(float64)),

// 		fmt.Sprintf(`Total %s Algorithmic Position ROI: %.5f%%`, algorithm, positionMap["current_position_ROI_percent"].(float64)),
// 	}
// 	content := ""
// 	for i, sent := range contents {
// 		if i > 0 {
// 			content = content + ". " + sent
// 		} else {
// 			content = content + sent
// 		}
// 	}
// 	return title, content, "", contents
// }

// func GenPerformanceChange(positionMap map[string]interface{}, durationString, algorithm string) (string, string, string, []string) {

// 	title := fmt.Sprintf(`GRAYLL | %s Algo Position Performance`, algorithm)
// 	openTime := int64(positionMap["open_position_timestamp"].(float64))
// 	last12hROIPercent := positionMap["current_position_ROI_percent"].(float64)
// 	if val, ok := positionMap["last_12h_ROI_percent"]; ok {
// 		last12hROIPercent = val.(float64)
// 	}
// 	contents := []string{
// 		time.Unix(openTime, 0).Format(`15:04 | 02-01-2006`),

// 		fmt.Sprintf(`Your scheduled %s algorithmic position performance notification for: `, algorithm),

// 		fmt.Sprintf(`GRAYLL | Transaction ID | %s`, positionMap["grayll_transaction_id"].(string)),

// 		fmt.Sprintf(`%s Algorithmic Position | 24 hour ROI: %.5f%%`, algorithm, last12hROIPercent),

// 		fmt.Sprintf(`Algo Position Duration: %v`, durationString),

// 		fmt.Sprintf(`Current Algo Position Value: $%.5f`, positionMap["current_position_value_$"].(float64)),

// 		fmt.Sprintf(`Current GRX Rate | $%.7f`, positionMap["current_value_GRX"].(float64)),

// 		fmt.Sprintf(`Current Algo Position Value: %.7f GRX`, positionMap["current_position_value_GRX"].(float64)),

// 		fmt.Sprintf(`Total %s Algorithmic Position ROI: %.5f%%`, algorithm, positionMap["current_position_ROI_percent"].(float64)),
// 	}
// 	content := ""
// 	for i, sent := range contents {
// 		if i > 0 {
// 			content = content + ". " + sent
// 		} else {
// 			content = content + sent
// 		}
// 	}
// 	return title, content, "", contents
// }

// func GenPositionValueWarning(positionMap map[string]interface{}, algorithm, grayllTxId string) (string, string, string, []string) {

// 	title := fmt.Sprintf(`GRAYLL | %s Algo Position Value Warning`, algorithm)
// 	openTime := int64(positionMap["open_position_timestamp"].(float64))
// 	// last12hROIPercent := positionMap["current_position_ROI_percent"].(float64)
// 	// if val, ok := positionMap["last_12h_ROI_percent"]; ok {
// 	// 	last12hROIPercent = val.(float64)
// 	// }

// 	// GRAYLL | GRY 1 Balthazar Algo Position Value Warning (*Subject)
// 	// 08:23 | 22/02/2020 (*Time | Date)

// 	// Your GRY 1 | Balthazar algorithmic position currently has a lower value than your principal.
// 	// You may want to consider reviewing your performance targets to avoid potentially further decline in value.

// 	// GRAYLL | Transaction ID | 000000000000000645

// 	// GRY 1 | Balthazar Algorithmic Position | 24 hour ROI: -1.80125%

// 	// Total GRY 1 | Balthazar Algorithmic Position ROI: -2.80125%

// 	// Current Algo Position Value: $10.58608

// 	// Current Algo Position Value: 74.6908687 GRX

// 	contents := []string{
// 		time.Unix(openTime, 0).Format(`15:04 | 02-01-2006`),

// 		fmt.Sprintf(`Your %s algorithmic position currently has a lower value than your principal.`, algorithm),

// 		fmt.Sprintf(`You may want to consider reviewing your performance targets to avoid potentially further decline in value.`),

// 		fmt.Sprintf(`GRAYLL | Transaction ID | %s`, grayllTxId),

// 		//fmt.Sprintf(`%s Algorithmic Position | 24 hour ROI: %.5f%%`, algorithm, last12hROIPercent),

// 		fmt.Sprintf(`Total %s Algorithmic Position ROI: %.5f%%`, algorithm, positionMap["current_position_ROI_percent"].(float64)),

// 		fmt.Sprintf(`Current Algo Position Value: $%.5f`, positionMap["current_position_value_$"].(float64)),

// 		fmt.Sprintf(`Current Algo Position Value: %.7f GRX`, positionMap["current_position_value_GRX"].(float64)),
// 	}
// 	content := ""
// 	for i, sent := range contents {
// 		if i > 0 {
// 			content = content + ". " + sent
// 		} else {
// 			content = content + sent
// 		}
// 	}
// 	return title, content, "", contents
// }

// func GenExpiredPosition(grayllTxId, algorithm string) (string, string, []string) {

// 	title := fmt.Sprintf(`GRAYLL | %s Algo Position Expired`, algorithm)
// 	openTime := time.Now().Unix()

// 	contents := []string{
// 		time.Unix(openTime, 0).Format(`15:04 | 02-01-2006`),

// 		fmt.Sprintf(`Your %s algorithmic position has expired: `, algorithm),

// 		fmt.Sprintf(`GRAYLL | Transaction ID | %s`, grayllTxId),

// 		fmt.Sprintf(`Expired Algo Position Value: $0.00000`),

// 		fmt.Sprintf(`GRX | Expired Algo Position Value | 0.00000 GRX`),

// 		fmt.Sprintf(`Total %s Algorithmic Position ROI: -100.00000%% `, algorithm),
// 	}
// 	content := ""
// 	for i, sent := range contents {
// 		if i > 0 {
// 			content = content + ". " + sent
// 		} else {
// 			content = content + sent
// 		}
// 	}
// 	return title, content, contents
// }
// func GenOptimalAlgo(positionMap map[string]interface{}, algorithm string) (string, string, []string) {

// 	title := fmt.Sprintf(`GRAYLL | %s Algo Position Optimal Performance Range`, algorithm)
// 	openTime := int64(positionMap["open_position_timestamp"].(float64))
// 	last12hROIPercent := positionMap["current_position_ROI_percent"].(float64)
// 	if val, ok := positionMap["last_12h_ROI_percent"]; ok {
// 		last12hROIPercent = val.(float64)
// 	}
// 	contents := []string{
// 		time.Unix(openTime, 0).Format(`15:04 | 02-01-2006`),

// 		fmt.Sprintf(`Your %s algorithmic position is within the optimal performance range.`, algorithm),

// 		fmt.Sprintf(`GRAYLL | Transaction ID | %s`, positionMap["grayll_transaction_id"].(string)),

// 		fmt.Sprintf(`%s Algorithmic Position | 24 hour ROI: %.5f%%`, algorithm, last12hROIPercent),

// 		fmt.Sprintf(`Total %s Algorithmic Position ROI: %.5f%%`, algorithm, positionMap["current_position_ROI_percent"].(float64)),

// 		fmt.Sprintf(`Current Algo Position Value: $%.5f`, positionMap["current_position_value_$"].(float64)),

// 		fmt.Sprintf(`Current Algo Position Value: %.7f GRX`, positionMap["current_position_value_GRX"].(float64)),
// 	}
// 	content := ""
// 	for i, sent := range contents {
// 		if i > 0 {
// 			content = content + ". " + sent
// 		} else {
// 			content = content + sent
// 		}
// 	}
// 	return title, content, contents
// }

// func GenClosePositionGRY(position AlgorithmPosition) (string, string, string, []string) {

// 	title := fmt.Sprintf(`GRAYLL | %s Algo Position Closed`, position.AlgorithmType)
// 	openFee := GetManagementFeePercent(position.AlgorithmType)
// 	contents := []string{
// 		time.Unix(position.ClosePositionTimestamp, 0).Format(`15:04 | 02-01-2006`),

// 		fmt.Sprintf(`You have closed a %s position`, GetAlgorithmString(position.AlgorithmType)),
// 		fmt.Sprintf(`Algo Position Duration: %s`, position.DurationString),

// 		fmt.Sprintf(`Close Algo Position Value: $%.5f`, position.ClosePositionValueUSD),
// 		fmt.Sprintf(`%s Rate | $%v`, position.AlgorithmType[:3], position.CloseValueGRY),
// 		fmt.Sprintf(`%s | Algo Management Fee @ %.1f%% | $%.5f`, position.AlgorithmType, openFee, position.ClosePositionFeeUSD),
// 		fmt.Sprintf(`%s | Algo Performance Fee @ 18%% | $%.5f`, position.AlgorithmType, position.ClosePerformanceFeeUSD),

// 		fmt.Sprintf(`GRX Rate | $%.5f`, position.CloseValueGRX),
// 		fmt.Sprintf(`%s | Algo Management Fee @ %.1f%% | %.7f GRX`, position.AlgorithmType, openFee, position.ClosePositionFeeGRX),
// 		fmt.Sprintf(`%s | Algo Performance Fee @ 18%% | %.7f GRX`, position.AlgorithmType, position.ClosePerformanceFeeGRX),

// 		fmt.Sprintf(`GRX | Close Algo Position Value | %.7f GRX`, position.ClosePositionValueGRX),

// 		fmt.Sprintf(`GRAYLL | Transaction ID | %s`, position.GrayllTxId),
// 	}

// 	content := ""
// 	for i, sent := range contents {
// 		if i > 0 {
// 			content = content + ". " + sent
// 		} else {
// 			content = content + sent
// 		}
// 	}

// 	url := "Stellar | Transaction ID | https://stellar.expert/explorer/public/search?term=" + position.CloseStellarTxId
// 	contents = append(contents, url)
// 	return title, content, url, contents
// }

// func GenCloseAllPositionGRY(position AlgorithmCloseAllPosition) (string, string, string, []string) {

// 	title := fmt.Sprintf(`GRAYLL | %s Algo Position Closed`, position.AlgorithmType)
// 	openFee := GetManagementFeePercent(position.AlgorithmType)
// 	contents := []string{
// 		time.Unix(position.ClosePositionTimestamp, 0).Format(`15:04 | 02-01-2006`),

// 		fmt.Sprintf(`You have closed %s position. Algo Position Duration: %s `, GetAlgorithmString(position.AlgorithmType), position.DurationString),

// 		fmt.Sprintf(`Close Algo Position Value: $%.5f`, position.ClosePositionValueUSD),
// 		fmt.Sprintf(`%s Rate | $%v.`, position.AlgorithmType[:3], position.CloseValueGRY),
// 		fmt.Sprintf(`%s | Algo Management Fee @ %.1f%% | $%.5f`, position.AlgorithmType, openFee, position.ClosePositionFeeUSD),
// 		fmt.Sprintf(`%s | Algo Performance Fee @ 18%% | $%.5f`, position.AlgorithmType, position.ClosePerformanceFeeUSD),

// 		fmt.Sprintf(`GRX Rate | $%.5f`, position.CloseValueGRX),
// 		fmt.Sprintf(`%s | Algo Management Fee @ %.1f%% | %.7f GRX`, position.AlgorithmType, openFee, position.ClosePositionFeeUSD),
// 		fmt.Sprintf(`%s | Algo Performance Fee @ 18%% | %.7f GRX`, position.AlgorithmType, position.ClosePerformanceFeeUSD),

// 		fmt.Sprintf(`GRX | Close Algo Position Value | %.7f GRX`, position.ClosePositionValueGRX),

// 		fmt.Sprintf(`GRAYLL | Transaction ID | %s. `, position.GrayllTxId),
// 	}

// 	content := ""
// 	for i, sent := range contents {
// 		if i > 0 {
// 			content = content + ". " + sent
// 		} else {
// 			content = content + sent
// 		}
// 	}

// 	url := "Stellar | Transaction ID | https://stellar.expert/explorer/public/search?term=" + position.StellarTxId
// 	//log.Println(content)
// 	contents = append(contents, url)

// 	return title, content, url, contents
// }

func GenRefererProfit(name, lname, stellarTx, grayllTx string, feeGrx, grxusd float64) (string, string, string, []string) {
	// GRAYLL | Referral Fee — received from — {First Name Referral} {Last Name Referral}
	// Dear {First Name Referrer},
	// You have received {....} GRX in Referral Fees from your Referral Contact {First Name Referrer} {Last Name Referrer}.

	// GRX Rate | {$0.12795}

	// Referral Fee Value: {$8.99999)

	// Stellar Transaction ID | https://stellar.expert/explorer/public/search?term=128358172838031361

	title := fmt.Sprintf(`GRAYLL | Referral Fee — received from — %s %s`, name, lname)
	contents := []string{

		fmt.Sprintf(`You have received %.7f GRX Referral Fee from your Referral Contact %s %s`, feeGrx, name, lname),

		fmt.Sprintf(`GRX Rate | %.5f`, grxusd),

		fmt.Sprintf(`Referral Fee Value: $%.5f`, feeGrx*grxusd),

		fmt.Sprintf(`Grayll | Transaction ID | %s. `, grayllTx),
	}
	content := ""
	for i, sent := range contents {
		if i > 0 {
			content = content + ". " + sent
		} else {
			content = content + sent
		}
	}
	url := "Stellar | Transaction ID | https://stellar.expert/explorer/public/search?term=" + stellarTx
	contents = append(contents, url)
	return title, content, url, contents

}
