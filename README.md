# GOLANG
func TestPasswd(t *testing.T) {
	var s = new(slapd.Slapd)
	s.Config = &slapd.DefaultConfig
	err := s.StartAndInitialize()
	defer s.Stop()
	if err != nil {
		t.Error(err)
	}

	lc := ldap.NewConnection("localhost:9999")
	err = lc.Connect()
	if err != nil {
		t.Error(err)
	}

	err = lc.Bind(slapd.DefaultConfig.Rootdn.Dn, slapd.DefaultConfig.Rootdn.Password)
	if err != nil {
		t.Error(err)
	}

	c := New(lc, "dc=example,dc=com")

	// create test person
	err = c.Create(&fritzFoobarPerson)
	if err != nil {
		t.Error(err)
	}

	// set password of test person to "foobaz"
	err = c.Passwd(&fritzFoobarPerson, "foobaz")
	if err != nil {
		t.Error(err)
	}

	c.Close()

	lc = ldap.NewConnection("localhost:9999")
	err = lc.Connect()
	if err != nil {
		t.Error(err)
	}

	// try to login as the test person with password "foobaz"
	err = lc.Bind(fritzFoobarPerson.Dn()+","+slapd.DefaultConfig.Suffix.Dn, "foobaz")
	if err != nil {
		t.Error(err)
	}

	c = New(lc, "dc=example,dc=com")

	// let the test person change its own password to "foobar"
	// this needs these acls set in slapd.conf:
	// access to attrs=userPassword
	//	by self write
	//	by anonymous auth
	//	by users none
	// access to * by * read

	err = c.Passwd(nil, "foobar")
	if err != nil {
		t.Error(err)
	}

	c.Close()

	lc = ldap.NewConnection("localhost:9999")
	err = lc.Connect()
	if err != nil {
		t.Error(err)
	}

	// try to login as the test person with password "foobar"
	err = lc.Bind(fritzFoobarPerson.Dn()+","+slapd.DefaultConfig.Suffix.Dn, "foobar")
	if err != nil {
		t.Error(err)
	}

}
