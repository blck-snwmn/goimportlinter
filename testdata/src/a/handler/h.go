package handler

import "a/domain" // want `this file can't import \"a/domain\"`

func handler() {
	domain.Do()
}
