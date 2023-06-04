package audio2face

type A2FFunctional string

const (
	// Default
	Status            A2FFunctional = "/status"
	ControlportStatus A2FFunctional = "/controlport/status"

	// General A2F
	GetInstances A2FFunctional = "/A2F/GetInstances" // Returns the existing instances as dictionary {'fullface_instances':[paths],'regular_instances':[paths] }
	USDLoad      A2FFunctional = "/A2F/USD/Load"     // Load a USD file into the stage
	SetFrame     A2FFunctional = "/A2F/SetFrame"     // Sets the current frame
	GetFrame     A2FFunctional = "/A2F/GetFrame"     // Returns the current frame

	// Player
	PlayGetInstances A2FFunctional = "/A2F/Player/GetInstances" // Lists the existing player instances in the scene
	GetRootPath      A2FFunctional = "/A2F/Player/GetRootPath"
	SetRootPath      A2FFunctional = "/A2F/Player/SetRootPath"
	GetTracks        A2FFunctional = "/A2F/Player/GetTracks"
	GetCurrentTrack  A2FFunctional = "/A2F/Player/GetCurrentTrack"
	SetTrack A2FFunctional = "/A2F/Player/SetTrack"

	// Audio2Emotion
	GenerateKeys   A2FFunctional = "/A2F/A2E/GenerateKeys"
	A2EGetSettings A2FFunctional = "/A2F/A2E/GetSettings"
	A2ESetSettings A2FFunctional = "/A2F/A2E/SetSettings"

	// Export
	ExportBlendshapes A2FFunctional = "/A2F/Exporter/ExportBlendshapes"

	// Special return
	REPORT_SUCCESS = "\"OK\""
)
