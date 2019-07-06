# OhMyPushBot

Send a telegram message to you with simple HTTP POST request!

[OhMyPushBot](https://t.me/ohmypushbot) in Telegram messager.

Useful in unattended script to notification script manager.

## How to use

1. Add this bot ([OhMyPushBot](https://t.me/ohmypushbot)) to you chat,click "start" button.
1. You can click `/url` in promot message,or type `/url` in textbox and send.
1. Bot will send you a url,like `https://ohmypushbot.zjyl1994.com/send?chatid=1234567890&sign=xxxxxxxxxxxxxxxxxxx`.
1. Send a POST request with you message in body,you will recvice you message
in bot chat.

## Q&A

### How it works?

I write this telegram bot,when you POST message as request body,my server will call telegram bot api 
and send you message to you chat.

### Is it safe?

Sure! You can check my code,i just pass you message to telegram bot api.If you not trust me,you can deploy
a self-hosted instance with you own telegram bot account.