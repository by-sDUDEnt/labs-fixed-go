import React, {useEffect, useState} from 'react';
import axios from 'axios';
function ConnectFour ({room, onGame, isAuth}) {
    const [socket, setSocket] = useState(null);
    const [receivedMessage, setReceivedMessage] = useState({});
    const [enemy, setEnemy] = useState(null);
    const [start, setStart] = useState(false);
    const [board, setBoard] = useState(null);
    const [turn, setTurn] = useState(null);
    const [winner, setWinner] = useState(null);
    const [color, setColor] = useState(null);
    useEffect(() => {
        console.log("board : ", board)
        console.log("color : ", color)
    },[board, turn, winner, start, enemy, receivedMessage])


    useEffect(() => {
        // Create a new WebSocket connection when the component mounts
            if (room.id != null) {
                const newSocket = new WebSocket('ws://16.171.22.36:8080/api/v1/game/'+room.id+'/ws?token='+localStorage.getItem('token')
                );
                newSocket.onopen = () => {
                    console.log('WebSocket connection established.');
                };
                // Set the socket state
                newSocket.onmessage = (event) => {
                    const receivedMessage = event.data;
                    setReceivedMessage(JSON.parse(JSON.parse(receivedMessage)));

                };
                setSocket(newSocket);
                // create timer to send ping to server every 5 seconds

                const interval = setInterval(() => {
                    axios.get(`http://16.171.22.36:8080/api/v1/room/${room.id}`, {
                        headers: {
                            Authorization: `Bearer ${localStorage.getItem('token')}`,
                        }
                    }).then((response) => {
                        console.log("response : ", response.data.status)
                        if(response.data.players.length === 2){
                            setEnemy(response.data.players[0] === isAuth.id ? response.data.players[1] : response.data.players[0])
                            setColor(response.data.players[0] === isAuth.id ? 1 : 2)
                            // alert("Enemy arrived!")
                        }
                        if(response.data.status === "in-progress"){
                            setStart(true)
                            setTurn(response.data.current_move_player_id)
                            setBoard(response.data.table)

                            // setWinner(response.data.winner)
                            // alert("Game started!")
                            clearInterval(interval)

                        }
                        if(response.data.status === "finished"){
                            clearInterval(interval)
                            leaveGame()
                            onGame()

                        }
                    })
                }, 1000);
                // Clean up the socket connection when the component unmounts
                return () => {
                    newSocket.close();
                };

            }
    }, []);


    useEffect(() => {
        console.log("receivedMessage : ", receivedMessage)
        switch (receivedMessage.type) {
            case 0:
                setStart(true)
                setTurn(receivedMessage.first_move_player_id)
                setBoard(receivedMessage.table)
                break;
            case 3:
                if (receivedMessage.player_id === isAuth.id){
                    setTurn(enemy)
                }else{
                    setTurn(isAuth.id)
                }
                ColorTable(receivedMessage.position, receivedMessage.player_id)
                break;
            case 4:
                setWinner(receivedMessage.win_player_id)
                if(receivedMessage.win_player_id === isAuth.id){
                    alert("You won!")
                    leaveGame()

                }else{
                    alert("You lost!")
                   leaveGame()
                    onGame()

                }
                break;

        }
    }, [receivedMessage])



    const startGame = () => {
        socket.send(JSON.stringify({"type": 1}))
    }
    const makeMove = (col) => {
        // make sure position is a number
        let position = setPositions(parseInt(col))

        if(turn === isAuth.id){
            socket.send(JSON.stringify({"type": 2, position: position}))
            console.log("makeMove", position)
        }else{
            console.log("Not your turn")
        }


    }


    const leaveGame = () => {
        axios.delete(`http://16.171.22.36:8080/api/v1/room/${room.id}`, {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('token')}`,
            },
        })
            .then((response) => {
                if (response.status === 200) {
                    // getRooms()
                    onGame()
                } else {
                    // Handle error
                }
            })
            .catch((error) => {
                // Handle network errors
            });
    }

    const checkRoom = () => {
        axios.get(`http://16.171.22.36:8080/api/v1/room/${room.id}`, {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('token')}`,
            }
        }).then((response) => {
            console.log(response.data)
        })
    }

    const ColorTable=(pos, id)=>{
        let col = pos%7
        console.log("col : ", col)
        let row = Math.floor(pos/7)
        console.log("row : ", row)
        let newBoard = board.slice()
        if(id === isAuth.id){
            newBoard[row][col] = color
            setBoard(newBoard)
        }
        else{
            newBoard[row][col] = color===1 ? 2 : 1
            setBoard(newBoard)
        }

    }

    const setPositions = (col) => {
        if(turn === isAuth.id){
            let position = parseInt(col)
            let newBoard = board.slice()
            let pos;
            for(let i = 5; i > -1; i--){

                if(newBoard[i][position] === 0){
                        newBoard[i][position] = color
                        pos=i*7+position
                        setBoard(newBoard)

                    console.log("NewBoard : ", board)
                    return i*7+position
                }
            }
        } else {
            alert("It's not your turn!")
        }

    }


    return (
        <div>
            <h1>You are in game! Room id : {room.id}</h1>
            <h3>Enemy: <b>{enemy}</b></h3>
            <h4>Turn: <b>{turn===isAuth.id? <span style={{backgroundColor:"green", color: "white"}}>Your turn!</span>: <span>Enemy</span>}</b></h4>
            {color=== 1 ? <h5 style={{backgroundColor: "red", width: "100px"}} >Color: <b>Red</b></h5> :  <h5 style={{backgroundColor: "yellow",  width: "100px"}} >Color: <b>Yellow</b></h5>}

            Game id: <b>{room.id}</b>
            <button onClick={leaveGame}>Leave game</button>
            { enemy ? ( start? <button onClick={makeMove}>Make Move</button> : <button onClick={startGame}>Start game</button>) : null
            }

            <div>
                <table style={{border: "3px black solid"}}>
                    <thead>
                        <tr>
                            <td><button onClick={()=>makeMove(0)}>Make Move 1</button></td>
                            <td><button onClick={()=>makeMove(1)}>Make Move 2</button></td>
                            <td><button onClick={()=>makeMove(2)}>Make Move 3</button></td>
                            <td><button onClick={()=>makeMove(3)}>Make Move 4</button></td>
                            <td><button onClick={()=>makeMove(4)}>Make Move 5</button></td>
                            <td><button onClick={()=>makeMove(5)}>Make Move 6</button></td>
                            <td><button onClick={()=>makeMove(6)}>Make Move 7</button></td>
                            </tr>
                        </thead>
                    <tbody>
                    {board ? board.map((value, row) => <tr key={row} style={{height : "10px", width: "10px"}}>{value.map((v, col) =>{
                        if(v === 0){
                            return <td key={(row*7)+col}  style={{backgroundColor: "white", border: "1px black solid"}}>Row: {(row*7)+col} " Pos : {col}</td>
                        }
                        if(v === 1){
                            return <td key={(row*7)+col}  style={{backgroundColor: "red", border: "1px black solid"}}>Row: {(row*7)+col} " Pos : {col} WOW 1</td>
                        }
                        if(v === 2){
                            return <td key={(row*7)+col}  style={{backgroundColor: "yellow", border: "1px black solid"}}>Row: {(row*7)+col} " Pos : {col}</td>
                        }
                    })}</tr>): null}
                    </tbody>


                </table>
            </div>
            <div>

                <button onClick={checkRoom}>Check room in console</button>
            </div>
    </div>)
}
export default ConnectFour;