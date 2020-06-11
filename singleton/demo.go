package singleton

import "sync"

// 非线程安全
type notSafe struct {
}

var iNotSafe *notSafe

func GetINotSafe() *notSafe {
	if iNotSafe == nil {
		iNotSafe = &notSafe{}
	}
	return iNotSafe
}

// 线程安全
type easySafe struct {
}

var iEasySafe *easySafe = &easySafe{}

func GetIEasySafe() *easySafe {
	return iEasySafe
}

// 加锁线程安全
type lockSafe struct {
}

var (
	iLockSafe *lockSafe
	iOnceLock sync.Mutex
)

func GetILockSafe() *lockSafe {
	iOnceLock.Lock()
	defer iOnceLock.Unlock()
	if iLockSafe == nil {
		iLockSafe = &lockSafe{}
	}
	return iLockSafe
}

// 加锁两次检测线程安全
type checkLockSafe struct {
}

var (
	iCheckLockSafe *checkLockSafe
	iCheckLock     sync.Mutex
)

func GetCheckLockSafe() *checkLockSafe {
	if iCheckLockSafe == nil {
		iCheckLock.Lock()
		defer iCheckLock.Unlock()
		if iCheckLockSafe == nil {
			iCheckLockSafe = &checkLockSafe{}
		}
	}
	return iCheckLockSafe
}

// 最佳方式-once
type onceSafe struct {
}

var (
	iOnceSafe *onceSafe
	once      sync.Once
)

func GetOnceSafe() *onceSafe {
	once.Do(func() {
		iOnceSafe = &onceSafe{}
	})
	return iOnceSafe
}
