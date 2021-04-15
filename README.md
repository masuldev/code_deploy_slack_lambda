# Example golang slack
SNS, CodeDeploy, Lambda Example

![Screenshot of the slack](docs/example.png?raw=true "Screenshot")

## Quick Start
1) git clone
2) slack webhook url import
3) GOOS=linux GOARCH=amd64 go build -o main .
4) zip function.zip main
5) upload this app lambda
6) test!
7) profit!

###test json
```json
{
  "Records": [
    {
      "EventVersion": "0.0",
      "EventSubscriptionArn": "arn:aws:sns:EXAMPLE",
      "EventSource": "aws:sns",
      "Sns": {
        "Subject": "test notification",
        "Message": "{\"eventTriggerName\":\"test\",\"status\":\"FAILED\",\"applicationName\":\"code-deploy-sns-slack\",\"deploymentGroupName\":\"code-deploy-sns\",\"region\":\"ap-northeast-2\",\"deploymentId\":\"test-deploy\",\"createTime\":\"Thu Apr 15 08:24:51 UTC 2021\",\"completeTime\":\"Thu Apr 15 08:28:51 UTC 2021\"}"
      }
    }
  ]
}
```

###! Have a nice day !
