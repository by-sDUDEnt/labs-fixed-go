import React, {useEffect, useState} from 'react';
import axios from 'axios';

function Rooms({onGame, isAuth}) {
    const [rooms, setRooms] = useState([]);
    useEffect(() => {
        getRooms()
    }, [])

    useEffect(() => {
    }, [rooms])
    const getRooms = () => {
        axios.get('http://16.171.22.36:8080/api/v1/room', {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('token')}`,
            },
        })
            .then((response) => {
                if (response.status === 200) {
                    setRooms(response.data)
                } else {
                    // Handle error
                }
            })
            .catch((error) => {
                // Handle network errors
            });
    }


    const deleteRoom = (id) => {
        axios.delete(`http://16.171.22.36:8080/api/v1/room/${id}`, {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('token')}`,
            },
        })
            .then((response) => {
                if (response.status === 200) {
                    getRooms()
                } else {
                    // Handle error
                }
            })
            .catch((error) => {
                // Handle network errors
            });
    }
    const JoinRoom = (id) => {
        axios.post(`http://16.171.22.36:8080/api/v1/room/${id}`, {}, {
            headers: {
                'Content-Type': 'application/json',
                'accept': 'application/json',
                Authorization: `Bearer ${localStorage.getItem('token')}`,
            },
        })
            .then((response) => {
                if (response.status === 201) {
                    onGame({id: response.data.id})
                } else {
                    console.log(response)
                }
            })
            .catch((error) => {
                console.log(error)
            });
    }


    const CreateRoom = () => {
        axios.post('http://16.171.22.36:8080/api/v1/room/create', {}, {
            headers: {
                'Content-Type': 'application/json',
                'accept': 'application/json',
                Authorization: `Bearer ${localStorage.getItem('token')}`,
            },
        })
            .then((response) => {
                if (response.status === 201) {
                    getRooms()
                    onGame({id: response.data.id})

                } else {
                    console.log(response)
                }
            })
            .catch((error) => {
                console.log(error)
            });
    }

    const RejoinRoom = (id) => {
        onGame({id: id})
    }


    return (
        <div> Rooms:
            {rooms.map((room) => <ul  key={room.id}> <li>Id: {room.id} <br/></li>
            {((room.players[0]==isAuth.id) || (room.players[1]==isAuth.id) )
                ?
                <div style={{margin: "5px"}}>
                    <button onClick={()=>{RejoinRoom(room.id)}}>Rejoin Room</button>  <button style={{margin: "5px"}} onClick={() => deleteRoom(room.id)}>Leave Room</button>
                </div>
                : (room.players.length==2) ? <div>Status: <b>Full</b></div> :
                    <ul> Players: <li>{room.players[0]}</li><li>{(room.players[1])? <div>room.players[1]</div> :<div>Free</div> }</li>
                         <button onClick={() => JoinRoom(room.id)}>Join</button> </ul>
                }




        </ul>)}
            <br/>
            <button onClick={CreateRoom}>Create Room</button>
        </div>
    );
}

export default Rooms;