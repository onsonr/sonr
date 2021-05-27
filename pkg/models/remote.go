package models

func (r *RemoteJoinRequest) Display() string {

}

func (r *RemoteJoinRequest) Words() []string {
	r.GetTopicWords()
}

func (r *RemoteJoinRequest) Topic() string {
	
}
