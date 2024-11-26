package bson_test

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func getSample() bson.Raw {
	cmd := "{\"find\": \"ForbidWithdraw\",\"filter\": {\"PlayerId\": {\"$numberInt\":\"1171498\"},\"Forbid\": true},\"limit\": {\"$numberLong\":\"1\"},\"singleBatch\": true,\"lsid\": {\"id\": {\"$binary\":{\"base64\":\"QEh8Z5b4RGOzFtK9wUgwQQ==\",\"subType\":\"04\"}}},\"$clusterTime\": {\"clusterTime\": {\"$timestamp\":{\"t\":1688696773,\"i\":1}},\"signature\": {\"hash\": {\"$binary\":{\"base64\":\"cfQmn/XA32JIegi4/AXsBNG4uYM=\",\"subType\":\"00\"}},\"keyId\": {\"$numberLong\":\"7205238149381881858\"}}},\"$db\": \"ConfigDB\",\"$readPreference\": {\"mode\": \"primary\"}}"
	//var j map[string]interface{}
	//if err := json.Unmarshal([]byte(cmd), &j); err != nil {
	//	panic(err)
	//}
	//
	//b, err := bson.Marshal(j)
	//if err != nil {
	//	panic(err)
	//}

	var r bson.Raw
	if err := bson.Unmarshal([]byte(cmd), &r); err != nil {
		panic(err)
	}

	return r

}

func Test_decode(t *testing.T) {

	raw := getSample()
	fmt.Print(raw)

}
