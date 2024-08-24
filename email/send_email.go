package emailSend

func SendEmail(title string, content string, to []string, cc, bcc, attachFile []string) error {
	sender := NewGmailSender()
	err := sender.SendEmail(title, content, to, cc, bcc, attachFile)
	return err
}
