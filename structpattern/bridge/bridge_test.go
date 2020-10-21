package bridge

import "testing"

func TestSendCommonSMS(t *testing.T)  {
	NewCommonMessage(ViaSMS()).SendMessage("#common_SMS msg#", "lizeyuan")
}

func TestSendCommonEmail(t *testing.T)  {
	NewCommonMessage(ViaEmail()).SendMessage("#common_Email msg#", "lizeyuan")
}

func TestSendUrgencySMS(t *testing.T)  {
	NewUrgencyMessage(ViaSMS()).SendMessage("#urgency_SMS msg#", "lizeyuan")
}

func TestSendUrgencyEmail(t *testing.T)  {
	NewUrgencyMessage(ViaSMS()).SendMessage("#urgency_Email msg#", "lizeyuan")
}