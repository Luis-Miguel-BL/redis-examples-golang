package contract

type Semaphore interface {
	Lock() (err error)
}
