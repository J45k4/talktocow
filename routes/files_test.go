package routes

import (
	"mime/multipart"
	"net/textproto"
	"testing"
)

func TestNormalizedContentTypeRemovesParametersAndLowercases(t *testing.T) {
	got := normalizedContentType("Audio/WebM;codecs=opus")

	if got != "audio/webm" {
		t.Fatalf("expected normalized content type audio/webm, got %q", got)
	}
}

func TestRecordingUploadContentTypeOptionsAllowAudioAndWebM(t *testing.T) {
	options := storeUploadedFileOptions{
		AllowedContentTypePrefixes: []string{"audio/"},
		AllowedContentTypes:        []string{"video/webm"},
	}

	if !contentTypeAllowed("audio/webm", options) {
		t.Fatal("expected audio/webm to be allowed")
	}

	if !contentTypeAllowed("video/webm", options) {
		t.Fatal("expected video/webm to be allowed as a browser WebM fallback")
	}

	if contentTypeAllowed("image/png", options) {
		t.Fatal("expected image/png to be rejected")
	}
}

func TestTrustedHeaderPrefixOnlyTrustsConfiguredPrefixes(t *testing.T) {
	if !contentTypeAllowedByPrefixes("audio/webm", []string{"audio/"}) {
		t.Fatal("expected audio/webm to match audio prefix")
	}

	if contentTypeAllowedByPrefixes("image/png", []string{"audio/"}) {
		t.Fatal("expected image/png not to match audio prefix")
	}
}

func TestDetectedUploadedContentTypePreservesWebMAudioHeader(t *testing.T) {
	fileHeader := &multipart.FileHeader{
		Header: textproto.MIMEHeader{
			"Content-Type": []string{"audio/webm;codecs=opus"},
		},
	}

	got := detectedUploadedContentType([]byte{0x1a, 0x45, 0xdf, 0xa3, 0x00}, fileHeader, storeUploadedFileOptions{})

	if got != "audio/webm" {
		t.Fatalf("expected audio/webm, got %q", got)
	}
}

func TestDetectedUploadedContentTypeDoesNotTrustAudioHeaderForTextData(t *testing.T) {
	fileHeader := &multipart.FileHeader{
		Header: textproto.MIMEHeader{
			"Content-Type": []string{"audio/webm"},
		},
	}

	got := detectedUploadedContentType([]byte("this is not an audio container"), fileHeader, storeUploadedFileOptions{})

	if got == "audio/webm" {
		t.Fatal("expected text data not to be classified as audio/webm")
	}
}
