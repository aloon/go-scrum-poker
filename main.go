package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/gosimple/slug"
)

type Action struct {
	Action string `json:"action"`
	Value  string `json:"value"`
}

type Participant struct {
	UserName string `json:"username"`
	Vote     string `json:"vote"`
	TempVote string `json:"-"`
}

type Room struct {
	ID           string        `json:"id"`
	CreatedAt    time.Time     `json:"-"`
	UpdatedAt    time.Time     `json:"-"`
	Participants []Participant `json:"participants"`
	YouAre       string        `json:"you_are,omitempty"`
}

var rooms = make(map[string]Room)
var roomClients = make(map[string]map[*websocket.Conn]string)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func createRoom(roomID, userName string) {
	room := Room{
		ID:           roomID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Participants: []Participant{{UserName: userName, Vote: ""}},
	}

	rooms[roomID] = room
}

func joinRoom(roomID, userName string) {
	room, exists := rooms[roomID]
	if !exists {
		return
	}

	room.Participants = append(room.Participants, Participant{UserName: userName, Vote: ""})
	room.UpdatedAt = time.Now()

	rooms[roomID] = room
}

func leaveRoom(roomID, userName string) {
	room, exists := rooms[roomID]
	if !exists {
		return
	}

	for i, participant := range room.Participants {
		if participant.UserName == userName {
			room.Participants = append(room.Participants[:i], room.Participants[i+1:]...)
			break
		}
	}

	room.UpdatedAt = time.Now()

	rooms[roomID] = room
}

func sendRoomToClients(roomID string, userName string) {
	connections, exists := roomClients[roomID]
	if !exists {
		return
	}

	room, exists := rooms[roomID]
	if !exists {
		return
	}

	if userName != "" {
		room.YouAre = userName
	}

	roomJSON, err := json.Marshal(room)
	if err != nil {
		fmt.Println("Error al convertir la sala a JSON:", err)
		return
	}

	for conn := range connections {
		err := conn.WriteMessage(websocket.TextMessage, roomJSON)
		if err != nil {
			fmt.Println("Error al enviar la sala al cliente:", err)
			continue
		}
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.LoadHTMLGlob(filepath.Join("templates", "*.html"))
	router.Static("/static", "./static")

	router.GET("/:room/ws", func(c *gin.Context) {
		roomID := c.Param("room")

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("Error al actualizar a WebSocket:", err)
			return
		}

		xUsername := c.Query("username")
		userName := ""
		if xUsername != "" {
			userName = xUsername
		} else {
			userName = generateFunnyUsername()
		}

		connections, exists := roomClients[roomID]
		if !exists {
			connections = make(map[*websocket.Conn]string)
			roomClients[roomID] = connections
		}

		connections[conn] = userName

		if !exists {
			createRoom(roomID, userName)
		} else {
			joinRoom(roomID, userName)
		}

		sendRoomToClients(roomID, userName)

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Error al leer el mensaje:", err)
				break
			}

			var action Action
			err = json.Unmarshal(msg, &action)
			if err != nil {
				fmt.Println("Error al decodificar el mensaje:", err)
				break
			}

			switch action.Action {
			case "vote":
				voteValue := action.Value

				room := rooms[roomID]
				room.UpdatedAt = time.Now()
				for i, participant := range room.Participants {
					if participant.UserName == userName {
						room.Participants[i].TempVote = voteValue
						room.Participants[i].Vote = "✋"
						break
					}
				}

				sendRoomToClients(roomID, "")

			case "showVotes":
				room := rooms[roomID]
				room.UpdatedAt = time.Now()
				for i := range room.Participants {
					room.Participants[i].Vote = room.Participants[i].TempVote
				}
				sendRoomToClients(roomID, "")

			case "cleanVotes":
				room := rooms[roomID]
				room.UpdatedAt = time.Now()
				for i := range room.Participants {
					room.Participants[i].Vote = ""
					room.Participants[i].TempVote = ""
				}
				sendRoomToClients(roomID, "")
			case "editUsername":
				room := rooms[roomID]
				room.UpdatedAt = time.Now()
				for i := range room.Participants {
					if room.Participants[i].UserName == userName {
						room.Participants[i].UserName = action.Value
						userName = action.Value
						sendRoomToClients(roomID, action.Value)
						break
					}
				}
			}
		}

		delete(connections, conn)
		leaveRoom(roomID, userName)

		sendRoomToClients(roomID, "")

		err = conn.Close()
		if err != nil {
			fmt.Println("Error al cerrar la conexión WebSocket:", err)
		}
	})

	router.GET("/:room", func(c *gin.Context) {
		roomId := c.Params.ByName("room")
		room, ok := rooms[roomId]
		if !ok {
			now := time.Now()
			rooms[roomId] = Room{CreatedAt: now, UpdatedAt: now, ID: roomId}
			room = rooms[roomId]
		}

		c.HTML(http.StatusOK, "room.html", gin.H{
			"room": room,
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.POST("/create-room", func(c *gin.Context) {
		type RoomName struct {
			RoomName string `json:"roomName"`
		}
		var roomName RoomName
		if err := c.BindJSON(&roomName); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data := map[string]interface{}{
			"slug": slug.Make(roomName.RoomName),
		}

		c.JSON(http.StatusOK, data)
	})

	router.Run(":8080")
}
