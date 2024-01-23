package behavioral_patterns

import "fmt"

type iOtp interface {
	genRandomOtp(int) string
	saveOtpCache(string)
	getMessage(string) string
	sendNotification(string) error
	publishMetric()
}

type otp struct {
	iOtp iOtp
}

func (o *otp) genAndSendMessage(otpLen int) error {
	otp := o.iOtp.genRandomOtp(otpLen)
	o.iOtp.saveOtpCache(otp)
	msg := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(msg)

	if err != nil {
		return err
	}

	o.iOtp.publishMetric()
	return nil
}

type sms struct {
	otp
}

func (s *sms) genRandomOtp(len int) string {
	fmt.Println("SMS: generation random OTP.")
	return fmt.Sprintf("smsOtpLen%v", len) 
}

func (s *sms) saveOtpCache(otp string) {
	fmt.Printf("Saving SMS OTP: %v to cache\n", otp)
}

func (s *sms) getMessage(otp string) string {
	return fmt.Sprintf("SMS: otp for login is %v", otp)
}

func (s *sms) sendNotification(msg string) error {
	fmt.Printf("Sending SMS => %v\n", msg)
	return nil
}

func (s *sms) publishMetric() {
	fmt.Println("SMS: Publishing metrics")
}

type email struct {
	otp
}

func (s *email) genRandomOtp(len int) string {
	fmt.Println("Email: generation random OTP.")
	return fmt.Sprintf("emailOtpLen%v", len) 
}

func (s *email) saveOtpCache(otp string) {
	fmt.Printf("Saving Email OTP: %v to cache\n", otp)
}

func (s *email) getMessage(otp string) string {
	return fmt.Sprintf("Email: otp for login is %v", otp)
}

func (s *email) sendNotification(msg string) error {
	fmt.Printf("Sending Email => %v\n", msg)
	return nil
}

func (s *email) publishMetric() {
	fmt.Println("Email: Publishing metrics")
}

func TemplatePattern() {
	smsOtp := &sms{}
	smsOtp.iOtp = smsOtp
	smsOtp.genAndSendMessage(4)

	emailOpt := &email{}
	oEmail := otp {
		iOtp: emailOpt,
	}
	oEmail.genAndSendMessage(5)
}



