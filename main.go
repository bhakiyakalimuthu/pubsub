import(
	"context"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func main(){
	pubsub := gochannel.NewGoChannel(gochannel.Config{},logger)
	publishMessages(pubsub)
}

func publishMessages(publisher message.Publisher){
		data := struct{vehicleID,userID,sourceID, source string}{"dummyvehicle","dummyuser",
		"dummysourceid","dummysource"}
		payload, _ := json.Marshal(&data)
		 msg := message.NewMessage(watermill.NewUUID(),payload)
		middleware.SetCorrelationID(watermill.NewUUID(), msg)
		log.Printf("sending message %s, correlation id: %s\n", msg.UUID, middleware.MessageCorrelationID(msg))
		if err := publisher.Publish("vehicle-terminator-svc",msg);err!=nil	{
			panic(err)
		}
		 	
}
