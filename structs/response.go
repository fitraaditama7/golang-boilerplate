package structs

type Response struct {
	Error 		bool 		`json:"error"`
	Code 		int			`json:"code"`
	Message 	string 		`json:"message"`
	Results 	[]Person 	`json:"results"`
	Result 		Person 		`json:"result"`
}
