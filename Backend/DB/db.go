package DB
import "go.mongodb.org/mongo-driver/bson/primitive"


type Todo struct {
    Title       string             `bson:"title" validate:"required"`
    Description string             `bson:"description" validate:"required"`
    Completed   bool               `bson:"completed"`
    ID          primitive.ObjectID `bson:"_id,omitempty"`
}

