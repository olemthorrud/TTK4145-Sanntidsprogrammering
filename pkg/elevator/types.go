package elevator

const (
	N_FLOORS  = 4
	N_BUTTONS = 3
)

type ElevatorBehaviour int

const (
	EB_Idle ElevatorBehaviour = iota
	EB_DoorOpen
	EB_Moving
	EB_Stopped
)

type Elevator struct {
	Id                  string
	OrderClearedCounter int
	OrderCounter        int
	Floor               int
	Dirn                MotorDirection
	Requests            [N_FLOORS][N_BUTTONS]bool
	Lights              [N_FLOORS][N_BUTTONS]bool
	Behaviour           ElevatorBehaviour
	similarity			int
}
