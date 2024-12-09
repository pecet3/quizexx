package magic_link

import "github.com/pecet3/quizex/pkg/utils"

func (*MagicLink) SendEmailLogin(to, code, userName string) error {
	subject := "🎲 Quizex 🎲 Magic Code (noreply)"
	body := `
    <html>
    	<body>
    		<h2>Hello, ` + userName + `,</h2>
    			<p>This is a magic code:</p>
				<h1>
					<i>` + code + `</i>
				</h1>
				<i>Please, copy and then paste it to the Quizex App.</i>
    	</body>
    </html>
    `
	if err := utils.SendEmail(to, subject, body); err != nil {
		return err
	}
	return nil
}
func (*MagicLink) SendEmailRegister(to, code, userName string) error {
	subject := "🎲 Quizex 🎲  Welcome the first time! (noreply)"
	body := `
    <html>
    	<body>
    		<h2>Hello, ` + userName + `,</h2>
				<p>We are happy You joined us!</p>
				<p>We wish you good luck and many good games</p>
				</br>
    			<p>This is a magic code:</p>
				<h1>
					<i>` + code + `</i>
				</h1>
				<i>Please, copy and then paste it to the Quizex App.</i>
    	</body>
    </html>
    `
	if err := utils.SendEmail(to, subject, body); err != nil {
		return err
	}
	return nil
}
