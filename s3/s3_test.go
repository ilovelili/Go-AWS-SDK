package s3

import (
	"testing"
)

// TestUpload test object upload
func TestUpload(t *testing.T) {
	svc := NewService("<<access key>>", "<<secret key>>")
	svc.SetRegion("ap-northeast-1")
	svc.SetBucket("monappy-logs")

	opts := &UploadOptions{
		FileName:       "./test.png",
		Public:         true,
		Encryption:     KMS,
		EncrptionKeyID: "<<KMS key id>>",
	}

	resp := svc.Upload(opts)
	if resp.Error != nil {
		t.Error(resp.Error)
	}
}
