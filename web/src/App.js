import React, {useEffect, useState} from 'react';
import axios from 'axios';
import AuthForm from './AuthForm';
import Rooms from './Rooms';
import ConnectFour from './ConnectFour';
function App ()  {
    const[isAuth, setIsAuth] = useState({status: false, id: null, username: null})
    const [game, setGame] = useState({id: null})

    useEffect( () => {
        // console.log("game", game)
    }, [isAuth, game])

    const deleteToken = () => {
        // Remove the token from localStorage
        localStorage.removeItem('token');
        setGame({id: null})
        setIsAuth({status: false, id: null, username: null})
    }

    return (
        <div>
            {isAuth.status ? <div> <button onClick={deleteToken}> LogOut</button> <br/> UserId : {isAuth.id} <br/> Name : {isAuth.username} <br/> </div> :<AuthForm onAuth={(auth)=> setIsAuth(auth)} />}
            <br/>- <br/>
            {isAuth.status ?
                (game.id==null? <Rooms onGame={(status)=> setGame(status)} isAuth={isAuth}/> : <ConnectFour room={game} onGame={()=>setGame({id: null})} isAuth={isAuth}/>) : null
            }
        </div>
    );
};

export default App;