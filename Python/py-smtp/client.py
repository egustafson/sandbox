from smtplib import SMTP
from email.utils import formataddr
from email.mime.text import MIMEText

msg = MIMEText('Email message body.')
msg['To'] = formataddr(('Recipient', 'rc@stub.net'))
msg['From'] = formataddr(('Sender', 'sender@stub.net'))
msg['Subject'] = 'Message Subject'

client = SMTP('127.0.0.1', 1025)
client.set_debuglevel(True)
try:
    client.sendmail('sender@stub.net', ['rc@stub.net'], msg.as_string())
finally:
    client.quit()

print('done.')
