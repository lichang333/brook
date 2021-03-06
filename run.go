package brook

// RunBKServer used to make a new BKServer and start to listen
func RunBKServer(address, password string, timeout, deadline int, m string) error {
	s, err := NewBKServer(address, password, timeout, deadline, m)
	if err != nil {
		return err
	}
	return s.ListenAndServe()
}

// RunBKClient used to make a new BKClient and start a socks5 proxy to listen
func RunBKClient(address, server, password string, timeout, deadline int, m string) error {
	c, err := NewBKClient(address, server, password, timeout, deadline, m, nil)
	if err != nil {
		return err
	}
	return c.ListenAndServe(nil)
}

// RunBKHTTPClient used to make a new BKClient and start a http proxy to listen
func RunBKHTTPClient(address, server, password string, timeout, deadline int, m string) error {
	c, err := NewBKClient(address, server, password, timeout, deadline, m, nil)
	if err != nil {
		return err
	}
	return c.ListenAndServeHTTP(nil)
}

// RunS5Server used to make a new S5Server and start to listen
func RunS5Server(address, password string, timeout, deadline int) error {
	s := NewS5Server(address, password, timeout, deadline)
	return s.ListenAndServe()
}

// RunS5Client used to make a new S5Client and start a socks5 proxy to listen
func RunS5Client(address, server, password string, timeout, deadline int) error {
	c := NewS5Client(address, server, password, timeout, deadline, nil)
	return c.ListenAndServe()
}

// RunSSServer used to make a new SSServer and start to listen
func RunSSServer(address, password string, timeout, deadline int) error {
	s := NewSSServer(address, password, timeout, deadline)
	return s.ListenAndServe()
}

// RunSSClient used to make a new SSClient and start a socks5 proxy to listen
func RunSSClient(address, server, password string, timeout, deadline int) error {
	c := NewSSClient(address, server, password, timeout, deadline, nil)
	return c.ListenAndServe(nil)
}

// RunSSHTTPClient used to make a new SSClient and start a http proxy to listen
func RunSSHTTPClient(address, server, password string, timeout, deadline int) error {
	c := NewSSClient(address, server, password, timeout, deadline, nil)
	return c.ListenAndServeHTTP(nil)
}

// RunRelay used to make a new Relay and start to listen
func RunRelay(address, server string, timeout, deadline int) error {
	r := NewRelay(address, server, timeout, deadline)
	return r.ListenAndServe()
}
