package magic_link

import "github.com/pecet3/quizex/pkg/utils"

func (*MagicLink) SendEmail(to, code, userName string) error {
	subject := "🔒Auth - Magic Link🔒 pecet.it (no reply)"
	body := `
    <html>
    	<body>
    		<h1>Hello ` + userName + `,</h1>
    			<p>This is code:</p>
				<h2>
					<i>` + code + `</i>
				</h2>
    	</body>
    </html>
    `
	if err := utils.SendEmail(to, subject, body); err != nil {
		return err
	}
	return nil
}
