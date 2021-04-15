package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request SNSMessage
		err     error
	}{
		{
			request: SNSMessage{},
			err:     ErrInvalidRecords,
		},
		{
			request: SNSMessage{
				Records: []Record{
					{
						SNS: SNS{
							Type: "Notification",
							MessageId: "dc1e94d9-56c5-5e96-808d-cc7f68faa162",
							TopicArn: "arn:aws:sns:us-east-2:111122223333:ExampleTopic1",
							Subject: "TestSubject",
							Message: "{\"eventTriggerName\":\"test-noti\",\"status\":\"FAILED\",\"applicationName\":\"code-deploy-notify-slack\",\"deploymentGroupName\":\"code-deploy-notify\",\"region\":\"ap-northeast-2\",\"deploymentId\":\"d-testdeploy\",\"createTime\":\"2018-05-16 08:23:44\",\"completeTime\":\"\"}",
							Timestamp: time.Now(),
							SignatureVersion: "1",
							Signature: "FMG5tlZhJNHLHUXvZgtZzlk24FzVa7oX0T4P03neeXw8ZEXZx6z35j2FOTuNYShn2h0bKNC/zLTnMyIxEzmi2X1shOBWsJHkrW2xkR58ABZF+4uWHEE73yDVR4SyYAikP9jstZzDRm+bcVs8+T0yaLiEGLrIIIL4esi1llhIkgErCuy5btPcWXBdio2fpCRD5x9oR6gmE/rd5O7lX1c1uvnv4r1Lkk4pqP2/iUfxFZva1xLSRvgyfm6D9hNklVyPfy+7TalMD0lzmJuOrExtnSIbZew3foxgx8GT+lbZkLd0ZdtdRJlIyPRP44eyq78sU0Eo/LsDr0Iak4ZDpg8dXg==",
							SigningCertURL: "https://sns.us-east-2.amazonaws.com/SimpleNotificationService-010a507c1833636cd94bdb98bd93083a.pem",
							UnsubscribeURL: "https://sns.us-east-2.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:us-east-2:111122223333:ExampleTopic1:e1039402-24e7-40a3-a0d4-797da162b297",
						}},
				},
			},
			err: nil,
		},
	}
	for _, test := range tests {
		err := Handler(test.request)
		assert.IsType(t, test.err, err)
	}
}
