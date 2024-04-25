package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var resource Resource

	// open the resource; be persistant if error is transient
	for resource, err = opener(); err != nil; resource, err = opener() {
		if _, ok := err.(TransientError); !ok {
			return
		}
	}

	// ensure we close the resource
	defer resource.Close()

	// ensure we call defrob if Frob panics
	defer func() {
		if e := recover(); e != nil {
			if fe, ok := e.(FrobError); ok {
				resource.Defrob(fe.defrobTag)
				err = fe
			} else {
				err = e.(error)
			}
		}
	}()

	// call frob
	resource.Frob(input)
	return
}
