package audio2face

type StandardResp struct {
	Status  string      `json:"status"`
	Result  interface{} `json:"result"`
	Message string      `json:"message"`
	Details []Detail    `json:"detail"`
}

type Detail struct {
	Loc  []string `json:"loc"`
	Msg  string   `json:"msg"`
	Type string   `json:"type"`
}

type Instances struct {
	StandardResp
	Result struct {
		FullfaceInstances []string `json:"fullface_instances"`
		RegularInstances  []string `json:"regular_instances"`
	} `json:"result"`
}

type PlayerInstances struct {
	StandardResp
	Result []string `json:"result"`
}

type GETRootPath struct {
	StandardResp
	Result string `json:"result"`
}
