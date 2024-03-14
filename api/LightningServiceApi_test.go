package api

import (
	"reflect"
	"testing"
)

func TestOpenChannel(t *testing.T) {
	want := "*lnrpc.Lightning_OpenChannelClient"
	got := reflect.TypeOf(OpenChannel()).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestGetInfo(t *testing.T) {
	want := "*lnrpc.GetInfoResponse"
	got := reflect.TypeOf(GetInfo()).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
