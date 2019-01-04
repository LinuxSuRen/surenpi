package main

import (
	"testing"

	"./mock_main"

	"github.com/golang/mock/gomock"
)

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	expect := 2

	m := mock_main.NewMockFoo(ctrl)
	m.EXPECT().Max(1, 2).Return(expect)

	if m.Max(1, 2) != expect {
		t.Errorf("max result is wrong, should be %d\n", expect)
	}
}
