package controllers

import "github.com/VMETA3/sui-hackathon/utils/nvidia/audio2face"

var Con *Controller

type Controller struct {
	A2F *audio2face.A2F
}

func NewController() *Controller {
	if Con == nil {
		Con = &Controller{
			A2F: audio2face.InitA2F(),
		}
	}
	return Con
}
