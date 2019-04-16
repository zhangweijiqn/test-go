package main

import (
	"container/list"
	"encoding/json"
)

type PushTask struct {
	Id                   int64   `json:"id"`
	ExperimentId         string  `json:"experimentId"`
	MemberId             int64   `json:"memberId"`
	ContentId            int64   `json:"contentId"`
	ContentType          string  `json:"contentType"`
	ContentTitle         string  `json:"contentTitle"`
	PushTitle            string  `json:"pushTitle"`
	PushTime             int64   `json:"pushTime"`
	PushDeadLine         string  `json:"pushDeadLine"`
	ParentId             string  `json:"parentId"`
	Grade                int     `json:"grade"`
	IosMinAppVersion     string  `json:"iosMinAppVersion"`
	AndroidMinAppVersion string  `json:"androidMinAppVersion"`
	Rating               float32 `json:"rating"`
	PushType             string  `json:"pushType"`
	PushIcon             string  `json:"pushIcon"`
	IsPuFilter           int     `json:"isPuFilter"`
}

type PushTaskAlg struct {
	PushTask  PushTask  `json:"pushTask"`
	MemberIds list.List `json:"memberIds"`
}

func (s *PushTaskAlg) Encode() ([]byte, error) {
	return json.Marshal(*s)
}

func (s *PushTaskAlg) Length() int {
	v, _ := s.Encode()
	return len(v)
}

func (s *PushTask) Encode() ([]byte, error) {
	return json.Marshal(*s)
}

func (s *PushTask) Length() int {
	v, _ := s.Encode()
	return len(v)
}

func main() {
	msg := "{\"pushTask\":{\"id\":0,\"experimentId\":\"alswr_8pm_2018_12_04\",\"memberId\":0,\"contentId\":\"472904679\",\"contentType\":\"answer\",\"contentTitle\":\"航空航天领域对动力的巨大需求是别的工程领域很难见到的，几百吨的客机要以接近音速飞行。或者几十吨的战斗机要以超过2倍音速飞行。再或者上千吨的……\",\"pushTitle\":\"有哪些航空航天上的事实，没有一定航空航天知识的人不会相信？\",\"pushTime\":1543924800000,\"pushDeadLine\":null,\"parentId\":null,\"grade\":0,\"iosMinAppVersion\":null,\"androidMinAppVersion\":null,\"rating\":0.0,\"pushType\":\"algorithm\",\"pushIcon\":null,\"isPuFilter\":0},\"memberIds\":[196514455]}"
	pushInfo := PushTaskAlg{}
	if err := json.Unmarshal(msg, &pushInfo); err != nil {
		return
	}
}
