import React, {useEffect, useState} from 'react';
import axios from 'axios';

function AuthForm({onAuth}) {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [isSignUp, setIsSignUp] = useState(false);
    useEffect(() => {
        let token = getToken()
        if (token) {
            axios.get('http://16.171.22.36:8080/api/v1/user', {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            })
                .then((response) => {
                    if (response.status === 200) {
                        getUserId()
                    } else {
                        onAuth({status: false, id: null, username: null})
                    }
                })
                .catch((error) => {
                    // Handle network errors
                });
        }
        const getUserId = () => {
            axios.get('http://16.171.22.36:8080/api/v1/user',
                {
                    headers: {
                        Authorization: `Bearer ${getToken()}`,
                    }
                })
                .then((response) => {
                    if (response.status === 200) {
                        onAuth({status: true, id: response.data.id, username: response.data.username})
                    } else {
                        onAuth({status: false, id: null, username: null})
                    }

                })
        }
    }, [])


    const handleSubmit = (e) => {
        e.preventDefault();
        const authEndpoint = isSignUp ? 'http://localhost:8080/api/v1/register' : 'http://localhost:8080/api/v1/login';
        let auth = {status: false, id: null, username: null}
        fetch(authEndpoint, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${getToken()}`,
            },
            body: JSON.stringify({username, password}),
        })
            .then((response) => {
                if (response.ok) {
                    response.json().then(r => {
                        saveToken(r.token)
                        auth.status = true
                        getUserId()
                    })
                } else {
                }
            })
            .catch((error) => {
                // Handle network errors
            });

        const getUserId = () => {
            axios.get('http://16.171.22.36:8080/api/v1/user',
                {
                    headers: {
                        Authorization: `Bearer ${getToken()}`,
                    }
                })
                .then((response) => {
                    if (response.status === 200) {
                        auth.id = response.data.id
                        auth.username = response.data.username
                        onAuth(auth)
                    } else {
                        auth = {status: false, id: null}
                    }

                })
        }


    };

    const saveToken = (token) => {
        // Save the token in localStorage
        localStorage.setItem('token', token);
    };

    const getToken = () => {
        // Retrieve the token from localStorage
        return localStorage.getItem('token');
    };


    return (
        <div>
            <div>
                <h2>{isSignUp ? 'Sign Up' : 'Login'}</h2>
                <form onSubmit={handleSubmit}>
                    <div>
                        <label htmlFor="username">Username:</label>
                        <input
                            type="text"
                            id="username"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                        />
                    </div>
                    <div>
                        <label htmlFor="password">Password:</label>
                        <input
                            type="password"
                            id="password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                        />
                    </div>
                    <button type="submit">{isSignUp ? 'Sign Up' : 'Login'}</button>
                </form>
                <p>
                    {isSignUp ? 'Already have an account?' : 'Don\'t have an account?'}
                    <button onClick={() => setIsSignUp(!isSignUp)}>
                        {isSignUp ? 'Login' : 'Sign Up'}
                    </button>
                </p>
            </div>
        </div>
    );
};

export default AuthForm;