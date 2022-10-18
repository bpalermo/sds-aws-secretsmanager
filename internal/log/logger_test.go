package log

import "testing"

func TestLogger_Debugf(t *testing.T) {
	l := Logger{
		Debug: true,
	}
	l.Debugf("works")
}

func TestLogger_Errorf(t *testing.T) {
	l := Logger{
		Debug: false,
	}
	l.Errorf("works %s", "yeah")
}

func TestLogger_Infof(t *testing.T) {
	l := Logger{
		Debug: true,
	}
	l.Infof("works %s", "yeah")
}

func TestLogger_Warnf(t *testing.T) {
	l := Logger{
		Debug: false,
	}
	l.Warnf("works %s", "yeah")
}
