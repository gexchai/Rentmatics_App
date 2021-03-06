package Model

type User struct {
	Id        int
	Username  string
	Password  string
	Loginid   string
	Logintype string
}

type RentUser struct {
	Username string
	Password string
	Role     string
}

type AdminResponse struct {
	Username string
	Role     string
}

type AdminLoginResponse struct {
	Userid   int
	Username string
	Password string
	Role     string
}

type UserResponse struct {
	Username string
	Loginid  string
	Status   string
}

type LoginUser struct {
	Username string
	Password string
}

type StatusResponse struct {
	Status string
}

type Getlogin struct {
	Username string
	Loginid  string
	Password string
}

type Changepass struct {
	Loginid     string
	OldPassword string
	Newpassword string
}

type Subscribeuser struct {
	Sub_City  string
	Sub_Email string
}
type Feedback struct {
	Name    string
	Emailid string
	Rating  int
}

type Referandearn struct {
	Refername    string
	Refernumber  string
	Ownername    string
	Ownerphone   string
	Owneraddress string
}

type Refer struct {
	Loginid string
}
