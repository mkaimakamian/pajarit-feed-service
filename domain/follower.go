package domain

// Representa una relación muy básica de seguidor - seguido
// Idealmente debería contar con un atributo que de idea de la fecha de alta;
// otros datos de auditoría podríuan ser útiles también a menos que se audite por otro medio
type Follower struct {
	FollowerId string
	FollowedId string
}
